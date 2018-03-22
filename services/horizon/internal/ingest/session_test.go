package ingest

import (
	"testing"

	"github.com/keybase/stellar-org/services/horizon/internal/db2/history"
	"github.com/keybase/stellar-org/services/horizon/internal/test"
)

func Test_ingestSignerEffects(t *testing.T) {
	tt := test.Start(t).ScenarioWithoutHorizon("set_options")
	defer tt.Finish()

	s := ingest(tt)
	tt.Require.NoError(s.Err)

	q := &history.Q{Session: tt.HorizonSession()}

	// Regression: https://github.com/stellar/horizon/issues/390 doesn't produce a signer effect when
	// inflation has changed
	var effects []history.Effect
	err := q.Effects().ForLedger(3).Select(&effects)
	tt.Require.NoError(err)

	if tt.Assert.Len(effects, 1) {
		tt.Assert.NotEqual(history.EffectSignerUpdated, effects[0].Type)
	}
}

func Test_ingestOperationEffects(t *testing.T) {
	tt := test.Start(t).ScenarioWithoutHorizon("set_options")
	defer tt.Finish()

	s := ingest(tt)
	tt.Require.NoError(s.Err)

	// ensure inflation destination change is correctly recorded
	q := &history.Q{Session: tt.HorizonSession()}
	var effects []history.Effect
	err := q.Effects().ForLedger(3).Select(&effects)
	tt.Require.NoError(err)

	if tt.Assert.Len(effects, 1) {
		tt.Assert.Equal(history.EffectAccountInflationDestinationUpdated, effects[0].Type)
	}
}
