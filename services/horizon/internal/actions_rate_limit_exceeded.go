package horizon

import (
	"net/http"

	"github.com/zenazn/goji/web"

	hProblem "github.com/keybase/stellar-org/services/horizon/internal/render/problem"
	"github.com/keybase/stellar-org/support/render/problem"
)

// RateLimitExceededAction renders a 429 response
type RateLimitExceededAction struct {
	Action
	App *App
}

// ServeHTTPC is a method for web.Handler
func (action RateLimitExceededAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ap := &action.Action
	c := web.C{
		Env: map[interface{}]interface{}{
			"app": action.App,
		},
	}
	ap.Prepare(c, w, r)
	ap.App = action.App
	problem.Render(action.Ctx, action.W, hProblem.RateLimitExceeded)
}
