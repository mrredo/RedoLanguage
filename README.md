# RedoLanguage
Interpreter for redo language
## ADDING SEMICOLONS AFTER EXPRESSIONS WILL
for example
```
var hello = "Hello " + "World" + "!";
hello += "10";
println(hello)//outputs: Hello World!10
```
if u dont put `;` at end of expression(not required in functions, after functions u dont need semicolons)
```
var hello = "Hello " + "World" + "!"
hello += "10"
```
it would say `hello is not defined` because it doesnt stop reading the expression, and stops reading when semicolon is added
NOTE: This is required in this current situation, in further updates hopefully ill figure it out



#BETA features
### IF statements
Note: Only supports hard coded values
```go
if true {
print(true)
}
```

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
