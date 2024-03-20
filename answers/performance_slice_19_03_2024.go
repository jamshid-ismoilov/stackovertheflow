package answers

// QUESTION

/*
I am moving from C++ to Golang and have a clarification question about slice in Golang.

I am trying to understand/confirm if it is an expensive operation to remove an element located at the front of the slice.

From the below example, it looks like the slice (frontRemoved) is still pointing to the original slice even after the removal of the first element.
So, I think it is not an expensive operation to remove an element from the front of the slice (unlike vector in C++).
*/

import "fmt"

func main() {
	original := []int{1, 2, 3, 4}
	frontRemoved := original[1:]
	frontRemoved[0] = 10
	fmt.Println("original slice:", original)
	fmt.Println("frontRemoved slice:", frontRemoved)
}

/*
**Output**:
original slice: [1 10 3 4]
frontRemoved slice: [10 3 4]

Can you please let me know if it is an expensive operation to remove the first element of the slice? Also, should I use a container/list to implement a queue?
Or slice is an acceptable solution where performance is important?

Thank you

I have experimented with the above code to understand the slice implementation.
*/

// ANSWER

/*
In Golang, removing an element from the front of the slice is not an expensive operation.
In your example, Go doesn't create a new underlying array. Instead, it creates a new slice header that points to the same array but with different starting index.

So you're not copying elements around memory, you're just adjusting slice header's starting point and length.

For your second question, if your queue's size is fixed and you don't do much insertions or deletions in the middle, then using slice is fine.

But if your queue dynamically resize with much elements and you do a lot of insertions/deletions then you can use container/list as you mention.
*/
