package main 
import "fmt" 
func main() { 
var num1, num2 int 
fmt.Println("Enter number 1:") 
fmt.Scan(&num1) 
fmt.Println("Enter number 2:") 
fmt.Scan(&num2) 
sum := num1 + num2 
fmt.Println("Sum of both numbers:", sum) 
}
