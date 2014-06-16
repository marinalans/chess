package chess

//Minimax-Decision(state) returns an action
func miniMaxDecision(state ChessNode) Move {
	//v = Max-Value(state)
	//return the action in successors(state) with value v
	return Move{}
}

//Max-Value(state) returns a utility value
func maxValue(state ChessNode) int {
	//If Terminal-Test(state) then return Utility(state)
	//v <= -infinity
	//for a, s in Successors(state) do
	//  v <= Max(v, Min-Value(s))
	//return v
	return 0
}

//Min-Value(state) returns a utility value
func minValue(state ChessNode) int {
	//If Terminal-Test(state) then return Utility(state)
	//v <= +infinity
	//for a, s in Successors(state) do
	//  v <= Min(v, Max-Value(s))
	//return v
	return 0
}

func terminalTest(state ChessNode) bool {
	return false
}

func successors(state ChessNode) []Move {
	var moves []Move
	return moves
}

func utility(state ChessNode) int {
	return 0
}
