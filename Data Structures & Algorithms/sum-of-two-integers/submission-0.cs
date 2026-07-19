public class Solution {
    public int GetSum(int a, int b) {
        int carry = 0;
        int sum = 0;

        for(int i = 0; i < 32; i++) {
            int mask = 1 << i;
            int bitA = a & mask;
            int bitB = b & mask;
            int bitRes = 0;

            if (bitA == mask && bitB == mask) {
                if (carry == 1) { bitRes = mask; }
                carry = 1;
            } else if (bitA == mask || bitB == mask) {
                if (carry == 0) {
                    bitRes = mask;
                } else {
                    bitRes = 0;
                    carry = 1;
                }
            } else {
                if (carry == 1) { bitRes = mask; }
                carry = 0;
            }

            sum |= bitRes;
        }

        return sum;
    }
}
