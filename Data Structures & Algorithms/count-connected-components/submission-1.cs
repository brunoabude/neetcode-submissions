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
        var remaining = new HashSet<int>();

        for(int i = 0; i < n; i++){
            remaining.Add(i);
        }

        void dfs(int node) {
            if (!remaining.Contains(node)) {
                return;
            }

            remaining.Remove(node);
            
            if (adj.TryGetValue(node, out var list)) {
                foreach(var neighbor in list) {
                    dfs(neighbor);
                }
            }
        }

        while (remaining.Count > 0) {
            var node = remaining.First();
            dfs(node);
            components++;
        }

        return components;
    }
}
