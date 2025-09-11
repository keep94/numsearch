package numsearch_test

import (
	"fmt"

	"github.com/keep94/numsearch"
	"github.com/keep94/sqrt"
)

func ExampleAll() {

	// sqrt(2) = 0.14142135... * 10^1
	n := sqrt.Sqrt(2)

	// '14' matches at index 0, 2, 144, ...
	count := 0
	for index := range numsearch.All(n, numsearch.String("14")) {
		fmt.Println(index)
		count++
		if count == 3 {
			break
		}
	}
	// Output:
	// 0
	// 2
	// 144
}

func ExampleFirst() {

	// sqrt(3) = 0.1732050807... * 10^1
	n := sqrt.Sqrt(3)

	fmt.Println(numsearch.First(n, numsearch.String("0508")))
	// Output:
	// 4
}

func ExampleLast() {
	n := sqrt.Sqrt(2)
	fmt.Println(numsearch.Last(n.WithEnd(1000), numsearch.String("14")))
	// Output:
	// 945
}

func ExampleBackward() {
	n := sqrt.Sqrt(2)
	count := 0
	iterator := numsearch.Backward(n.WithEnd(1000), numsearch.Ints(1, 4))
	for index := range iterator {
		fmt.Println(index)
		count++
		if count == 3 {
			break
		}
	}
	// Output:
	// 945
	// 916
	// 631
}
