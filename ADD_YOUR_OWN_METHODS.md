# How to make your own methods for RedoLanguage in Go
### Now you can only use Go since RedoLanguage doesn't support function declaration

In `std/functions.go` you can find the `Functions` map, there are 2 ways to add to the map
And commands must follow this to be stable `func(args ...interface{}) interface{}`
```go
func LoadCommands() {
    Functions["input"] = Input
}
func Input(args ...interface{}) interface{}
  fmt.Println(args...)  
}
```
or 
```GO
var Functions = map[string]func(args ...interface{}) interface{}{
    //OTHER METHODS
    "input": Input
    
}
```
```GO
var Functions = map[string]func(args ...interface{}) interface{}{
    //OTHER METHODS
    "input": func(args ...interface{}) interface{} {
         fmt.Println(args...)  
    },
    
}
```
