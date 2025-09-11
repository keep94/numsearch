numsearch
=========

A package to find a pattern within a sequence of digits.

While this package is meant to work with the data structures
github.com/keep94/sqrt, it will work with anything that implements the
required interfaces.

## Examples

```golang
package main

import (
    "fmt"

    "github.com/keep94/numsearch"
    "github.com/keep94/sqrt"
)

func main() {

    // Find the 0-based position of the first occurrence of "07" within
    // Sqrt(3). This prints 8.
    fmt.Println(numsearch.FindFirst(sqrt.Sqrt(3), numsearch.String("07")))
}
```

More documentation and examples can be found [here](https://pkg.go.dev/github.com/keep94/numsearch).
