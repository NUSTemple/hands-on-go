package leetcode

func unorderedEqualString(first, second []string) bool {
    if len(first) != len(second) {
        return false
    }
    exists := make(map[string]bool)
    for _, value := range first {
        exists[value] = true
    }
    for _, value := range second {
        if !exists[value] {
            return false
        }
    }
    return true
}

func unorderedEqualInt(first, second []int) bool {
    if len(first) != len(second) {
        return false
    }
    exists := make(map[int]bool)
    for _, value := range first {
        exists[value] = true
    }
    for _, value := range second {
        if !exists[value] {
            return false
        }
    }
    return true
}