## sys
This is an unnecessary "wrapper" package over some built-in functions found in the ["builtin" package](https://pkg.go.dev/builtin) for Go.

I like being explicit when it comes to functions and methods, and would want to know where they come from.

I am in the process of designing a programming language, and I would love to make it necessary for whatever being used (except for some keywords maybe) to be packaged and explicitly imported into a source code file.

`system.out.println` anyone?

### Usage
Simply import the package, and call the necessary function.

```go
package main

import (
	"github.com/bsljth/sys"
)

func main() {
	sys.Log("Hello, world!")
}
```