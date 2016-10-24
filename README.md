# Command line application template for Go

Implement CLI application by editing [app.go](app.go).  
You may add new files to keep your code clean, if it is allowed in your challenge.

## How to get input parameters
You can get arguments as `args` in [app.go](app.go) file where the `run` method is defined.  

``` go
func run(args []string) {
  // code to run
}
```

`args` is simply came from `os.Args`, passed by `main` function in [main.go](main.go). [main.go](main.go) is also an editable file; if you want to parse options, you may customize this file as well.

## How to output result
You can use `fmt.Println` methods to output your results.

``` go
fmt.Println(args)
```
