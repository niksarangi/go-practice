package main
import "fmt"
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n - 1) + fibonacci(n - 2)
}
func main() {
    var n int 
    fmt.Println("Enter number of terms for fibonacci:") 
    fmt.Scan( & n) 
    fmt.Println("Fibonnacci series:") 
    for i: = 0;i < n;i++{
        fmt.Print(fibonacci(i), " ")
    }
}
