public class Solution {
    public int CountComponents(int n, int[][] edges) {
        Dictionary<int, List<int>> adj = new();

        foreach(var edge in edges) {
            if (!adj.ContainsKey(edge[0])) {
                adj[edge[0]] = new List<int>();
            }

            if (!adj.ContainsKey(edge[1])) {
                adj[edge[1]] = new List<int>();
            }
            
            adj[edge[0]].Add(edge[1]);
            adj[edge[1]].Add(edge[0]);
        }


        var components = 0;
        var visited = new bool[n];

        void dfs(int node) {
            if (visited[node]) {
                return;
            }

           visited[node] = true;
            
            if (adj.TryGetValue(node, out var list)) {
                foreach(var neighbor in list) {
                    dfs(neighbor);
                }
            }
        }

        for (int i = 0; i < n; i++) {
            if (visited[i]) {
                continue;
            }
            dfs(i);
            components++;
        }

        return components;
    }
}
