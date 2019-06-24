# Pointers

- no pointer arithmetic
- struct fields acced through struct pointer
- Need to pass pointers to pass arguments by reference.(As opposed to passing by value aka. copying the arguments)
- get value of pointer with & symbol in front of value
- "\*" symbol gets pointer's value
- Methods often defined on pointers and not values
- pointers often stored in variables

```go
client := &http.Client{}
resp, err := client.Get("http://gobootcamp.com")
```
