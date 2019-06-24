# Functions

- can return multiple values

### Examples

```go
func add(x int, y int) int {
  return x + y
}

func main(){
  fmt.Println(add(42, 13))
}
```

```go
func location(city string) (string,string){
  var region string
  var continent string

  switch city {
    case "Los Angeles", "LA", "Santa Monica":
      region, continent = "California", "North America"
    case "New York", "NYC":
      region, continent = "New York", "North America"
    default:
      region, continent = "Unknown", "UnKnown"
  }
  return region, continent // Returning multiple values
}

func main () {
  region, continent := location("Santa Monica")
  fmt.Printf("Matt lives in %s, %s", region, continent)
}
```

### Example with result parameters

- pretty confusing. probably shouldn't use

```go
func locatin(name, city string) (region, continent string){
  switch city {
    case "New York", "LA", "Chicago":
      continent = "North America"
    default:
      continent = "Unknown"
  }
  return
}

func main() {
  region, continent := location("Matt", "LA")
  fmt.Printf("%s lives in %s", region, continent)
}
```
