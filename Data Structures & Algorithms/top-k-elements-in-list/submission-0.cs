public class Solution {
    public int[] TopKFrequent(int[] nums, int k) {
        var freq = new Dictionary<int, int>();

        foreach(var i in nums) {
            if (freq.TryGetValue(i, out var v)) {
                freq[i] = v + 1;
            } else {
                freq[i] = 1;
            }
        }

        var q = new PriorityQueue<int, int>();
        
        foreach (var (val,f) in freq) {
            q.Enqueue(val,-f);
        }


        var result = new int[k];

        for (int i = 0; i < k; i++) {
            result[i] = q.Dequeue();
        }

        return result;        
    }
}
