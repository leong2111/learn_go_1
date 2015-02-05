package main
import "fmt"

type mySentence string


type Contact struct {
name string
greeting string
}

func main() {
message := "Hello World!"
var greeting *string = &message
fmt.Println(message, greeting)
// the * makes a pointer of a certain type
// the above code makes greeting a pointer to a string
// the & before a variable says, "Give me the memory location of that variable"


var message mySentence = "Hello World!"
// var message string = "Hello World!"
fmt.Println(message)

// go is not an OOP language
// the type system that exists in go
// -- makes it so you don't need classes
// -- gives you more flexibility b/c you're not constrained by class requirements
// instead of using classes, we have user defined types

var c = Contact{"Marcus", "Hello!"}
fmt.Println(c.name)
fmt.Println(c.greeting)
}
// we can use a STRUCT like a class
// go is not an OOP language
// the type sustem that exists in go
// -- makes it so you don't need classes
// -- gives you more flexibility b/c you're not constrained by class requirements
// instead of using classes, we have user defined types
