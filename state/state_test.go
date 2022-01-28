package state

import "testing"

func TestViewState(t *testing.T) {
	wanted := "[kylling rev korn hs VEST---\\_____/______________/---ØST]"
	state := ViewState()
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}

func TestPutInBoat(t *testing.T) {
	wanted := state
	state := ViewState()
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}

func TestCrossRiver(t *testing.T) {
	wanted := state
	state := ViewState()
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}
