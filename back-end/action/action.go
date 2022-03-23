/* © 2022 Ben C. Forsberg <benfrsbrg@gmail.com> */

package action

// ActionName is for enumerating string constants for actions to avoid
// misspellings.
type ActionName string

const (
	SessionUpdateUsername ActionName = `SESSION.UPDATE_USERNAME`
	SessionNavigate                  = `SESSION.NAVIGATE`
)

type Action struct {
	Name    ActionName  `json:"action"`
	Payload interface{} `json:"payload"`
}
