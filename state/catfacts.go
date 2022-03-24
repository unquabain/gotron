/* © 2022 Ben C. Forsberg <benfrsbrg@gmail.com */

package state

import (
  "net/http"
  "gotron/action"
  "encoding/json"
  "bytes"
  "fmt"
)

type factResponse struct {
  Fact string `json:"fact"`
}

func getCatFact(dispatcher action.Dispatcher) {
  var response factResponse
  buff := new(bytes.Buffer)
  resp, err := http.Get(`https://catfact.ninja/fact`)
  if err != nil {
    dispatcher.Dispatch(action.Action{
      Name: action.CatFactError,
      Payload: fmt.Errorf(`could not get a cat fact: %w`, err),
    })
    return
  }
  if _, err := buff.ReadFrom(resp.Body); err != nil {
    dispatcher.Dispatch(action.Action{
      Name: action.CatFactError,
      Payload: fmt.Errorf(`could not read cat fact response: %w`, err),
    })
    return
  }
  resp.Body.Close()
  if err := json.Unmarshal(buff.Bytes(), &response); err != nil {
    dispatcher.Dispatch(action.Action{
      Name: action.CatFactError,
      Payload: fmt.Errorf(`could not understand cat fact response: %w`, err),
    })
    return
  }
  dispatcher.Dispatch(action.Action{
    Name: action.CatFactSuccess,
    Payload: response.Fact,
  })
}

type CatFact string
func (cf CatFact)Reduce(actionName action.ActionName, payload interface{}, dispatcher action.Dispatcher) (State, bool) {
  switch actionName {
  case action.CatFactFetch:
    go getCatFact(dispatcher)
    return CatFact(`loading...`), true
  case action.CatFactError:
    return CatFact(fmt.Sprintf(`ERROR: %v`, payload)), true
  case action.CatFactSuccess:
    return CatFact(payload.(string)), true
  default:
    return cf, false
  }
}

