class Solution {
public:
    vector<int> twoSum(vector<int>& numbers, int target) {
        int left = 0;
        int right = numbers.size() - 1;
        int summed = 0;
        while (left < right) {
            summed = numbers[left] + numbers[right];
            if (summed == target) return vector<int>{ left + 1, right + 1};
            if (summed < target) {
                ++left;
            } else if (summed > target) {
                --right;
            }
        }
        return vector<int>{ -1 }; // Indicate failure
    }
};