func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    charMap := make(map[rune]int)

    for _, char := range s {
        charMap[char]++
    }

    for _, char := range t {
        if _, ok := charMap[char]; !ok || charMap[char] == 0 {
            return false
        }

        charMap[char]--
    }

    return true
}
