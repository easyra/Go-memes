#Const

- Can only be character, string, boolean, or nummeric value
- Can't be declared using := syntax
- untyped constant takes type needed by its context

### Examples

```go
const Pi = 3.14
const (
  StatusOk = 200
  StatusCreated = 201
  StatusAccepted = 202
  StatusNonAuthoritativeInfo = 203
  StatusNoContent = 204
  StatusResetContent = 205
  StatusPartialContent = 206
)
```

```go
package main

import "fmt"

const (
  Pi = 3.14
  Truth= false
  Big = 1<< 62
  small = Big >> 61
)

func main(){
  const Greeting = "Hi"
  fmt.PrintLn(Greeting)
  fmt.PrintLn(Pi)
  fmt.PrintLn(Truth)
  fmt.PrintLn(Big)
}
```
