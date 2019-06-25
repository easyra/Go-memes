# Types

#### Non-Numeric Types

- bool
- string

#### Numeric Types

##### unit

- either32 or 64 bits

##### int

- same size as uint

##### uintptr

- unsigned integer large enough to store the uninterpreted bits of a pointer value

##### uint(x)

- 8
- 16
- 32
- 64

##### int(x)

- 8
- 16
- 32
- 64

##### float

- float32
- float64

##### complex

- complex64
- complex128

##### byte

- alais for unit8

##### rune

- alias for int32

#### Examples

```go
package main

import (
  "fmt"
  "math/cmplx"
)

var (
  goIsFun bool = true
  maxInt uint64 = 1<<64 - 1
  complex complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
  const f = "%T(%v)\n"
  fmt.Printf(f,goIsFun, goIsFun)
  fmt.Printf(f, maxInt, maxInt)
  fmt.Printf(f,complex,complex)
}
```

# Type Conversion

- types need to be manually converted if varible is being passed a to function expecting a different type

#### T(v)

- converts value v into type T

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

```go
i := 42
f = float64(i)
u := uint(f)
```

# Type Assertion

- takes a value an copies it with a different type

```go
import (
  "fmt"
  "time"
)
func timeMap(y interface{}){
  z, ok := y.(map[string]interface{})
  if ok {
    z["updated_at"] = time.Now()
  }
}

func main() {
  foo := map[string]interface{} {
    "Matt": 42,
  }
  timeMap(foo)
  fmt.Println(foo)
}
```
