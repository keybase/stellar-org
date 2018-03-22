package amount_test

import (
	"testing"

	"github.com/keybase/stellar-org/amount"
	"github.com/keybase/stellar-org/xdr"
)

var Tests = []struct {
	S     string
	I     xdr.Int64
	valid bool
}{
	{"100.0000000", 1000000000, true},
	{"-100.0000000", -1000000000, true},
	{"100.0000001", 1000000001, true},
	{"123.0000001", 1230000001, true},
	{"123.00000001", 0, false},
	{"922337203685.4775807", 9223372036854775807, true},
	{"922337203685.4775808", 0, false},
	{"922337203686", 0, false},
	{"-922337203685.4775808", -9223372036854775808, true},
	{"-922337203685.4775809", 0, false},
	{"-922337203686", 0, false},
	{"1000000000000.0000000", 0, false},
	{"1000000000000", 0, false},
}

func TestParse(t *testing.T) {
	for _, v := range Tests {
		o, err := amount.Parse(v.S)
		if !v.valid && err == nil {
			t.Errorf("expected err for input %s", v.S)
			continue
		}
		if v.valid && err != nil {
			t.Errorf("couldn't parse %s: %v", v.S, err)
			continue
		}

		if o != v.I {
			t.Errorf("%s parsed to %d, not %d", v.S, o, v.I)
		}
	}
}

func TestString(t *testing.T) {
	for _, v := range Tests {
		if !v.valid {
			continue
		}

		o := amount.String(v.I)

		if o != v.S {
			t.Errorf("%d stringified to %s, not %s", v.I, o, v.S)
		}
	}
}
