package main
import "fmt"
func palindrome(s string) bool {
    n: = len(s) for i: = 0;i < n / 2;i++{
        if s[i] != s[n - i - 1] {
            return false
        }
    }
    return true
}
func main() {
    var s string fmt.Println("Enter a string:") fmt.scan( & s) if palindrome(s) {
        fmt.Println(s, "is a palindrome")
    } else {
        fmt.Println(s, "not a palindrome")
    }
}
