func isValidSudoku(board [][]byte) bool {
	for i := range 9 {
		digits := 0

		for j := range 9 {
			d := board[i][j]

			if d == '.' {
				continue
			}

			d = d - '1'
			bit := 1 << d

			if digits&bit == bit {
				return false
			}

			digits = digits | bit
		}
	}

	for j := range 9 {
		digits := 0

		for i := range 9 {
			d := board[i][j]

			if d == '.' {
				continue
			}

			d = d - '1'
			bit := 1 << d

			if digits&bit == bit {
				return false
			}

			digits = digits | bit
		}
	}

	for groupI := range 3 {
		for groupJ := range 3 {
			digits := 0

			for j := range 3 {
				for i := range 3 {
					d := board[groupI*3+i][groupJ*3+j]

					if d == '.' {
						continue
					}

					d = d - '1'
					bit := 1 << d

					if digits&bit == bit {
						return false
					}

					digits = digits | bit
				}
			}
		}
	}

	return true
}