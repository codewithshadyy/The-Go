
// Like arrays, slices are also used to store multiple values of the same type in a single variable.

// However, unlike arrays, the length of a slice can grow and shrink as you see fit.

// In Go, there are several ways to create a slice:

// Using the []datatype{values} format
// Create a slice from an array
// Using the make() function

package main
import "fmt"

func main(){

	// // Using the []datatype{values} format
	my_slice := []int{1, 2, 4}

	// len() function - returns the length of the slice (the number of elements in the slice)
    //  cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)
	fmt.Println(len(my_slice))
	fmt.Println(cap(my_slice))


	// creating a slice from an array
	var numbers = [4]int{4, 5, 6, 78}
	slice_me := numbers[0:2]
	fmt.Printf("slice: %v\n", slice_me)
	fmt.Printf("length:%d\n", len(slice_me))
	fmt.Printf("capacity:%d\n", cap(slice_me))


	// appending elelments to a slice
	slice_me = append(slice_me, 2, 3, 6)
	fmt.Println(slice_me)


	// appending one slice to another slice
	new_slice :=[]int{45, 67, 6, 2}
	final_slice :=append(slice_me, new_slice...)

	fmt.Println(final_slice)


	// creating slice with the make() function

	// var course  = make([] string, 4, 6)


}