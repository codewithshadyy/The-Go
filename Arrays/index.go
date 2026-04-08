

// arrays are used to store elements of the same data types instea of having multiple variables
// variables are declare in this ways 


// 1. With the var keyword:
// Syntax
// var array_name = [length]datatype{values} // here length is defined
//  var names  = [5] string{"superdad", "layla"}

// or

// var array_name = [...]datatype{values} // here length is inferred 
// var names = [...]string{"superdada", "layla"}



// 2. With the := sign:
// Syntax
// array_name := [length]datatype{values} // here length is defined

// or

// array_name := [...]datatype{values} // here length is inferred

package main
import "fmt"

func main(){

	var names  = [5] string{"superdad", "layla"}
	students := [4] string {"superdad", "layla"}



	names[0] = "job"
	fmt.Println(names[0])
	
	fmt.Println(students[1])
}

// we acces an element by passing an index e.g students[0] which return students at index position 0
// we change an element by parsing an index into the variable the a new value e.g students[0] = "james"
// to access the length of an array with call the len(array_name)