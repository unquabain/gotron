/* © 2022 Ben C. Forsberg <benfrsbrg@gmail.com> */

package state

import (
  "os/exec"
  "gotron/action"
  "fmt"
)

type FortuneState string
func (fs FortuneState) Reduce(actionName action.ActionName, payload interface{}) (State, bool) {
  switch actionName {
  case action.FortuneGenerate:
    cmd := exec.Command(`fortune`, `-l`)
    stdout, err := cmd.Output()
    if err != nil {
      return FortuneState(fmt.Sprintf(`error running fortune: %v`, err)), true
    }
    return FortuneState(stdout), true
  default:
    return fs, false
  }
}
