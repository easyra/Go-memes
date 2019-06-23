# Printing Cosntants and Variables

- fmt package === best practice

##### fmt.Println

- prints passed variable's value/ appends to newline

```go
func main() {
  sylonModel := 6
  fmt.Println(clyonModel)
}
```

##### fmt.Printf

- for one or more values using a defined format.

```go
func main() {
  name := "Caperica-Six"
  aka := fmt.Sprintf("Number %d", 6)
  fmt.Printf("%s is also known as %s",
    name,aka)
}
```
