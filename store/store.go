/* © 2022 Ben C. Forsberg <benfrsbrg@gmail.com> */

package store

import (
	"encoding/json"
	"fmt"
	"github.com/webview/webview"
	"gotron/action"
	"gotron/state"
  "log"
  "os"
  "sync"
)

type Store struct {
	wv webview.WebView
	State *state.AppState
  lock *sync.Mutex
}

func NewStore(wv webview.WebView) *Store {
	store := &Store{
		State: state.NewAppState(),
		wv:    wv,
    lock: new(sync.Mutex),
	}
  jsLogger := log.New(os.Stderr, `From JS: `, 0)
	wv.Bind(`gotronDispatch`, store.Dispatch)
	wv.Bind(`gotronLog`, jsLogger.Println)
	wv.Bind(`gotronPushState`, store.pushState)
	return store
}

func (s *Store) pushState() error {
	stateJson, err := json.Marshal(s.State)
	if err != nil {
		return fmt.Errorf(`could not marshal state: %w`, err)
	}
	command := fmt.Sprintf(`window.gotronSetState(%s)`, stateJson)
	s.wv.Eval(command)
	return nil
}

func (s *Store) Dispatch(action action.Action) error {
  s.lock.Lock()
  defer s.lock.Unlock()
	st, dirty := s.State.Reduce(action.Name, action.Payload, s)
	if dirty {
		s.State = st.(*state.AppState)
		return s.pushState()
	}
	return nil
}
