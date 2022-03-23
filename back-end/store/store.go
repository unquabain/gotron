/* © 2022 Ben C. Forsberg <benfrsbrg@gmail.com> */

package store

import (
	"encoding/json"
	"fmt"
	"github.com/webview/webview"
	"gotron/back-end/action"
	"gotron/back-end/state"
  "log"
  "os"
)

type Store struct {
	wv webview.WebView
	state.State
}

func NewStore(wv webview.WebView) *Store {
	store := &Store{
		State: state.NewAppState(),
		wv:    wv,
	}
  jsLogger := log.New(os.Stderr, `From JS: `, 0)
	wv.Bind(`gotronDispatch`, store.dispatch)
	wv.Bind(`gotronLog`, jsLogger.Println)
	wv.Bind(`gotronPushState`, store.setState)
	return store
}

func (s *Store) setState() error {
	stateJson, err := json.Marshal(s.State)
	if err != nil {
		return fmt.Errorf(`could not marshal state: %w`, err)
	}
	command := fmt.Sprintf(`window.gotronSetState(%s)`, stateJson)
	s.wv.Eval(command)
	return nil
}

func (s *Store) dispatch(action action.Action) error {
	state, dirty := s.State.Reduce(action.Name, action.Payload)
	if dirty {
		s.State = state
		return s.setState()
	}
	return nil
}
