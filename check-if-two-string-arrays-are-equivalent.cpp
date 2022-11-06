class Solution {
public:
    bool arrayStringsAreEqual(vector<string>& word1, vector<string>& word2) {
        string a = "";
        for(auto x : word1) a = a.append(x);
        string b = "";
        for(auto x : word2) b = b.append(x);
        return a == b;
    }
};