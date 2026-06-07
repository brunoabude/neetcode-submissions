/*
["A","B","C","E"]
["S","F","C","S"]
["A","D","E","E"]
*/

func exist(board [][]byte, word string) bool {
    wordLen := len(word)
    m, n := len(board), len(board[0])
    visited := make([][]bool, m)

    for i := range m {
        visited[i] = make([]bool, n)
    }

    var dfs func(i, j, charIdx int) bool
    dfs = func(i, j, charIdx int) bool {
        if charIdx >= wordLen {
            return true
        }

        if i < 0 || i >= m {
            return false
        }

        if j < 0 || j >= n {
            return false
        }

        if visited[i][j] {
            return false
        }

        if word[charIdx] != board[i][j] {
            return false
        }

        visited[i][j] = true
        defer func(){ visited[i][j] = false }()

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
