package main 
import "fmt" 
func prime(n int) bool { 
 if n <= 1 { return false } 
 for i := 2; i <= n; i++ {
  if n%i == 0 && i != n { 
   return false 
} 
} 
return true 
} 

func main() {
 var n int 
 fmt.Println("Enter the number:") 
 fmt.Scan(&n) 
 if prime(n) { 
  fmt.Println(n, "is a prime") 
} else { fmt.Println(n, "not prime") } 
}
