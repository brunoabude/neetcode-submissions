public class Solution {
    public int NumDecodings(string s) {
        Dictionary<int, int> memo = [];

        int dfs(int i) {
            if (i == s.Length) {
                return 1;
            }

            if (i > s.Length || s[i] == '0' || s[i] > '9') {
                return 0;
            }

            if (memo.TryGetValue(i, out int v)) {
                return v;
            }

            var aux = 0;

            if (i+1 < s.Length && (
                (s[i] == '1' && s[i+1] >= '0' && s[i+1] <= '9') ||
                (s[i] == '2' && s[i+1] >= '0' && s[i+1] <= '6')
            )) {
                aux += dfs(i+2);
            }

            aux += dfs(i+1);

            memo[i] = aux;
            return aux;
        }
        // 26 1 10 5 5 9 7 17 5 6 5 6 2
        // 2 6 1 10 5 5 9 7 17 5 6 5 6 2
        // 26 1 10 5 5 9 7 1 7 5 6 5 6 2
        // 2 6 1 10 5 5 9 7 1 7 5 6 5 6 2

        return dfs(0);
    }
}
