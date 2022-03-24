/* © 2022 Ben C. Forsberg <benfrsbrg@gmail.com> */

package action

type Dispatcher interface {
	Dispatch(action Action) error
}

// ActionName is for enumerating string constants for actions to avoid
// misspellings.
type ActionName string

const (
	SessionUpdateUsername ActionName = `SESSION.UPDATE_USERNAME`
	SessionNavigate                  = `SESSION.NAVIGATE`
	FortuneGenerate                  = `VIEWS.FORTUNE.GENERATE`
	CatFactFetch                     = `VIEWS.CATFACT.FETCH`
	CatFactError                     = `VIEWS.CATFACT.ERROR`
	CatFactSuccess                   = `VIEWS.CATFACT.SUCCESS`
)

type Action struct {
	Name    ActionName  `json:"action"`
	Payload interface{} `json:"payload"`
}
