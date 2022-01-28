package state

var state = "[kylling rev korn hs VEST---\\_____/______________/---ØST]"

func ViewState() string {
	return state
}

func PutInBoat() string {
	state = "[kylling rev korn VEST---\\_kylling_hs_/______________/---ØST]"
	return state
}

func CrossRiver() string {
	state = "[kylling rev korn VEST---\\______________\\_kylling_hs_//---ØST]"
	return state
}
