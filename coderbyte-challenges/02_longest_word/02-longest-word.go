package main
import ( 
    "fmt"
    "regexp"
)

func LongestWord(sen string) string {
  // code goes here
  r, _ := regexp.Compile("\\w+")
  words := r.FindAll([]byte(sen), -1)
  longest_word := "a"

  for i := 0; i < len(words); i++ {
    if len(words[i]) > len(longest_word) {
      longest_word = string(words[i])
    }
  }

  return longest_word
}

func main() {

  // do not modify below here, readline is our function
  // that properly reads in the input for you
  fmt.Println(LongestWord("fun!!time"))

}