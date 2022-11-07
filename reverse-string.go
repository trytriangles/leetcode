func reverseString(s []byte) {
    max := len(s) - 1
    for i := 0; i < len(s) / 2; i++ {
        o := max - i
        s[i], s[o] = s[o], s[i]
    }
}

