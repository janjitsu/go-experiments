package main
import "fmt"

func FirstFactorial(num int) int {
  
  // code goes here
  total := num
    
  for i := num - 1; i > 1; i-- {
    total = total * i
  }
  return total;

}

func main() {

  // do not modify below here, readline is our function
  // that properly reads in the input for you
  fmt.Println(FirstFactorial(4))

}