package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File 


// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	var squares = 0
	for _, v := range cb[file] {
		if v {
			squares++
		}
	}
	return squares
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	var squares = 0

	if rank < 1 || rank > 8 {
		return 0
	}

	for _, v := range cb {
		var x = v[rank - 1]
		if x {
			squares++
		}
	}

	return squares
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var squares = 0

	for _, v := range cb {
		squares += len(v)
	}	
	
	return squares
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var squares = 0

	for _, v := range cb {
		for _, square := range v {
			if square {
				squares++
			}	
		}	
	}

	return squares
}
