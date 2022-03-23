/* © 2022 Ben C. Forsberg <benfrsbrg@gmail.com> */

package state

import (
	"gotron/action"
)

type State interface {
	// Reduce may alter the State. It returns a new State, and a bool
	// which indicates whether the returned State is different than
	// the original. If the dirty bool is true, the caller should replace
	// their state with the returned State.
	Reduce(action action.ActionName, payload interface{}) (State, bool)
}

type States map[string]State

// Reduce implements the State interface, and ranges over all the States
// in the collection, returning the updated map and an OR'd dirty value.
func (ss States) Reduce(action action.ActionName, payload interface{}) (State, bool) {
	var bigDirty bool
	for key, state := range ss {
		newState, dirty := state.Reduce(action, payload)
		if dirty {
			ss[key] = newState
			bigDirty = true
		}
	}
	return ss, bigDirty
}

type SessionState struct {
  Username string `json:"username"`
  NavPath  string `json:"navPath"`
}

func (ss *SessionState) Reduce(actionName action.ActionName, payload interface{}) (State, bool) {
	switch actionName {
	case action.SessionUpdateUsername:
		ss.Username = payload.(string)
		return ss, true
	case action.SessionNavigate:
		ss.NavPath = payload.(string)
		return ss, true
	default:
		return ss, false
	}
}

type AppState struct {
	Session *SessionState `json:"session"`
	Data    States        `json:"data"`
	Views   States        `json:"views"`
}

func (as *AppState) Reduce(action action.ActionName, payload interface{}) (State, bool) {
	var (
		state    State
		dirty    bool
		bigDirty bool
	)
	state, dirty = as.Session.Reduce(action, payload)
	if dirty {
		bigDirty = true
		as.Session = state.(*SessionState)
	}
	state, dirty = as.Data.Reduce(action, payload)
	if dirty {
		bigDirty = true
		as.Data = state.(States)
	}
	state, dirty = as.Views.Reduce(action, payload)
	if dirty {
		bigDirty = true
		as.Views = state.(States)
	}
	return as, bigDirty
}

func NewAppState() *AppState {
	return &AppState{
    Session: &SessionState{ Username: `Trevor`, NavPath: `/` },
		Data:    make(States),
    Views:   States{`fortune`: FortuneState(``)},
	}
}
