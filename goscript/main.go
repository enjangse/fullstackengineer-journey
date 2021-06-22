package main

import (
  "fmt"
  "bufio"
  "os"
  //"goscript/calculation"
  //"goscript/multiplication"
)
//Function and Package Tutorial
/*func main() {
    fmt.Println("Hello, Belajar Golang")

    result := multiplication.Add(1, 9)

    fmt.Println(result)
}*/


//Variable Tutorial
/*
func main() {
  var name string = "Halo"
  var secondvar string
  secondvar = "Second Var"
  age := 20

   fmt.Println(name)
   fmt.Println(secondvar)
   fmt.Println(age)
}
*/

//If else
/*func main() {
  age := 10

  if age > 17 {
    fmt.Println("Boleh bermain game PUBG mobile")
  }else {
    fmt.Println("Hanya boleh bermain Ninja Fruit")
  }
}*/

//if elseif
/*func main() {
  nilai := 75
  var grade string

  if nilai < 50 {
    grade = "D"
  } else if nilai < 60 {
    grade = "C"
  } else if nilai < 70 {
    grade = "B"
  } else {
    grade = "A"
  }

    fmt.Println(grade)
}*/

/*func main() {
  nilai := 3

  switch nilai {
  case 1: fmt.Println("Satu")
  case 2: fmt.Println("Dua")
  case 3: fmt.Println("Tiga")
  default: fmt.Println("Default")
  }

}*/
//looping
/*func main() {

  for i :=1; i <= 10; i++ {
    fmt.Println("saya sedang belajar loop %n", i)
  }
}*/

//input output from users
func main() {
  //var name string
  consoleReader := bufio.NewReader(os.Stdin)
  fmt.Println("Whats your name? ")
  //fmt.Scanln(&name)
  name, _ := consoleReader.ReadString('\n' )
  fmt.Println("Hello ", name)
}
