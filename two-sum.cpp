class Solution {
public:
    vector<int> twoSum(vector<int> nums, int target) {
        map<int, int> seen;
        for (unsigned int i = 0; i < nums.size(); i++) {
            int number = nums[i];
            int complement = target - number;
            auto it = seen.find(complement);
            if (it != seen.end()) return vector<int>{ it->second, (int)i };
            seen.insert({ number, i });
        }
        return vector<int>{ -1 }; // Unreachable under challenge parameters.
    }
};