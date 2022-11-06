class Solution {
public:
    vector<int> twoSum(vector<int>& numbers, int target) {
        int left = 0;
        int right = numbers.size() - 1;
        while (left < right) {
            int summed = numbers[left] + numbers[right];
            if (summed == target) {
                return vector<int>{ left + 1, right + 1};
            }
            else if (summed < target) {
                ++left;
            }
            else if (summed > target) {
                --right;
            }
        }
        return vector<int>{ -1 }; // Indicate failure
    }
};