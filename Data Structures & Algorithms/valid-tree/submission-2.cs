public class Solution {
    public bool ValidTree(int n, int[][] edges) {
        var adj = new List<List<int>>();

        for (int i = 0; i < n; i++){ adj.Add(new List<int>()); }

        for (int i = 0; i < edges.Length; i++) {
            var v0 = edges[i][0];
            var v1 = edges[i][1];
            
            adj[v0].Add(v1);
            adj[v1].Add(v0);
        }

        var visited = new bool[n];

        bool dfs(int prev, int v) {
            if (visited[v]) { return false; }

            visited[v] = true;
            
            for (int i = 0; i < adj[v].Count; i++) {
                if (adj[v][i] == prev) { continue; }
                if (adj[v][i] == v) { return false; }
                if (!dfs(v, adj[v][i])) { return false; }
            }

            return true;
        }

        return dfs(-1, 0) && visited.All(v => v == true);
    }
}
