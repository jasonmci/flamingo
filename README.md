# flamingo
A word with Go in it

## Starting up

- [ ] `main.go` is a convention we typically use and it is good practice
- [ ] `package main` is also idiomatic
- [ ] each line of each go file must be a package declaration
- [ ] we can return more than one thing from a function, which is incredibly useful

## Pointers

- [ ] instead of returning something from a function and then using that as an arg to change the function
      we can use pointers. 
- [ ] lowercase func is internal to package
- [ ] uppercase func is external and visible outside the package

## Structs

- [ ] There are a lot of thigns to like about Go and structs are one of them
- [ ] And usiing them with pointers is how we can use a struct as the rceiver for a function

## Maps 

- [ ] Maps are immutable, but what does that really mean?
    We dont need to provide a pointer to a maps
    Order is not preserved. 

## Keyboard shortcuts
- [ ] To format Golang  SHIFT+OPTION+F

## Packages

- [ ] go mod init github.com/jasonmci/packagename
- [ ] this makes a package we can reuse. 
- [ ] then we can add the helpers folder and a file inside of that

```golang
"github.com/jasonmci/myniceprogram/helpers"
```

but the go.mod file has 

```
module github.com/jasonmci/myniceprogram

go 1.19
```

and we can add the subdirectory helpers under there and we can use it like

```go
var myVar helpers.SomeType
```
which is a type that is declared in the helpers file


## Channels. 

- [ ] A means of sending info from one part of the program toa nother. 
