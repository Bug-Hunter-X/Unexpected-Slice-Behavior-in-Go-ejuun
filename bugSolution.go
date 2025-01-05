func myFunc(a []int) []int {
    b := make([]int, len(a))
    for i, v := range a {
        b[i] = v * 2
    }
    return b
}

func main() {
    a := []int{1, 2, 3}
    b := myFunc(a)
    fmt.Println(b) // Output: [2 4 6]
    a[0] = 10
    fmt.Println(b) // Output: [2 4 6] 
    fmt.Println(a) // Output: [10 2 3]
} 