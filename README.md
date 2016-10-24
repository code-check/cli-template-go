# Command line application template for Go

Implement CLI application by editing [app.go](app.go).

## How to get input parameters
app.go file has a function `run`.  

``` go
func run(args []string) {
  // code to run
}
```

`args` is simply came from `os.Args`, passed by `main` function in [main.go](main.go). [main.go](main.go) is also an editable file; if you want to do option parse, you may customize this file as well.

## How to output result
You can use `Print` methods in `fmt` package.

``` go
fmt.Println(args)
```
