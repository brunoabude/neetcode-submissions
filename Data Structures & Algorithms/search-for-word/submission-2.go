/*
["A","B","C","E"]
["S","F","C","S"]
["A","D","E","E"]
*/
const VisitedFlag byte = '#'

func exist(board [][]byte, word string) bool {
    wordLen := len(word)
    m, n := len(board), len(board[0])

    var dfs func(i, j, charIdx int) bool
    dfs = func(i, j, charIdx int) bool {
        if charIdx >= wordLen {
            return true
        }

        // bound check
        if i < 0 || i >= m || j < 0 || j >= n {
            return false
        }

        if board[i][j] == VisitedFlag {
            return false
        }

        if word[charIdx] != board[i][j] {
            return false
        }

        originalChar := board[i][j]

        board[i][j] = VisitedFlag
        defer func(){ board[i][j] = originalChar }()

        return dfs(i+1, j, charIdx+1) ||
        dfs(i-1, j, charIdx+1) ||
        dfs(i, j+1, charIdx+1) ||
        dfs(i, j-1, charIdx+1) 
    }

    for i := range m {
        for j := range n {
            if dfs(i, j, 0) {
                return true
            }
        }
    }

    return false
}
