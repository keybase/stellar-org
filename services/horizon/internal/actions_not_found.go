package horizon

import (
	"github.com/keybase/stellar-org/support/render/problem"
)

// NotFoundAction renders a 404 response
type NotFoundAction struct {
	Action
}

// JSON is a method for actions.JSON
func (action *NotFoundAction) JSON() {
	problem.Render(action.Ctx, action.W, problem.NotFound)
}
