# Packages

- Name same as last element of import path
- Ex: The name of import "fmt" is fmt. import "math/rand" === rand

```go
package main

func main() {
  print("Hello, World!\n")
}
```

# Imports

```go
import "fmt"
import "math/rand"
```

###Grouped Imports

```go
import (
  "fmt"
  "math/rand"
)
```
