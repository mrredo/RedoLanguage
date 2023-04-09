# RedoLanguage
Interpreter for redo langugae
# URGENT: use 1.6 version if u want to use string arguments, in 1.7v they are broken
# Features
### Methods:

- print(args...)
- println(args...)
- printf(format string, args...)
- free(variableKey string)

### Variable declaration
```ts
var key = value
var key = false
var key = "hello"
var key = 10
```
### Comments
supports
one line comments //

multi line comments

/*

*/

value can be a string, boolean or a number
# DONE
- =, +=, -=, *=, /=, %= operators for variables
- +, -, *, /, ==, <, >, <=, >=, % general operators
# TODO

- if/else if/else statements
- for/while loops
- function declaration
