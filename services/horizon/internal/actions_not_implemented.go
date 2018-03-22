package horizon

import (
	hProblem "github.com/keybase/stellar-org/services/horizon/internal/render/problem"
	"github.com/keybase/stellar-org/support/render/problem"
)

// NotImplementedAction renders a NotImplemented prblem
type NotImplementedAction struct {
	Action
}

// JSON is a method for actions.JSON
func (action *NotImplementedAction) JSON() {
	problem.Render(action.Ctx, action.W, hProblem.NotImplemented)
}
