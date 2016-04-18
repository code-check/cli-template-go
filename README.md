# CLI template for Go

This is template app for CLI test.  
You can make console application by editing [src/main/app.go](src/main/app.go)

It uses cli.go. See detail in [GitHub repo of cli.go](https://github.com/codegangsta/cli)

## How to get input parameters
app.go file has a function `doMain`.  
It takes a parameter `c *cli.Context`.

``` go
func doMain(c *cli.Context) {
}
```

All parameters are passed as `c.Args()` array

If you want to use option parameter, you can customize cli  in [src/main/main.go](src/main/main.rb)

## How to output result
You can use `fmt.Println` method

``` go
  fmt.Println(c.Argus[0])
```