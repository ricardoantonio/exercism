package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	f, ok := cb[file]
	if !ok {
		return 0
	}
	squaresOccupied := 0
	for _, occupied := range f {
		if occupied {
			squaresOccupied++
		}
	}
	return squaresOccupied
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}
	squaresOccupied := 0
	for _, file := range cb {
		if file[rank-1] {
			squaresOccupied++
		}
	}
	return squaresOccupied
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	for _, file := range cb {
		return len(cb) * len(file)
	}
	return 0
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	totalSquaresOccupied := 0
	for file := range cb {
		totalSquaresOccupied += CountInFile(cb, file)
	}
	return totalSquaresOccupied
}
