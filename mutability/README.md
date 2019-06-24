# Mutability

- only constants(const) are immutable.
- a function mutating an argument won't change the original

```go
package main

import "fmt"

type Artist struct {
  Name, Genre string
  Songs int
}

func newRelease(a Artist) int {
  a.Songs++
  return a.Songs
}

func main() {
  me := Artist{Name: "Matt", Genre: "Electro", Songs: 42}
  fmt.Printf("%s released their %dth song\n", me.Name, NewRelease(me))
  fmt.Printf("%s has a total of %d songs", me.Name, me.Songs)
}
```

### Example with pointer

```go
package main
import "fmt"

type Artist struct {
  Name, Genre string
  Songs int
}

func newRelease(a *Artist) int {//Now takes in a pointer
  a.Songs++
  return a.Songs
}

func main(){
  me := &Artist{Name: "Matt", Genre: "Electro", Sogns: 42} //Now assigns me to a poitner
  fmt.Printf("%s released their %dth song\n", me.Name, newRelease(me))
  fmt.Printf("%s has a total of %d songs". me.Name, me.Songs)
}
```
