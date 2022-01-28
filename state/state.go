package state

var state = "[kylling rev korn hs VEST---\\_____/______________/---ØST]"

func ViewState() string {
	return state
}

func PutInBoat() {
	state = "[kylling rev korn VEST---\\_kylling_hs_/______________/---ØST]"
}

func CrossRiver() {
	state = "[kylling rev korn VEST---\\______________\\_kylling_hs_//---ØST]"
}
