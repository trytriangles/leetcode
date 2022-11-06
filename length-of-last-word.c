int lengthOfLastWord(char *s) {
    int wordLength = 0;
    int sLen = (int)strlen(s) - 1;
    for (int i = sLen; i > -1; i--) {
        if (s[i] == ' ') {
            if (wordLength > 0)
                return wordLength;
            continue;
        }
        wordLength++;
        if (i == 0)
            return wordLength;
    }
    return -1;
}
