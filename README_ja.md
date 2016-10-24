# コマンドラインアプリケーション(CLI アプリ)作成用テンプレート(Go)

[app.go](app.go)を編集して、CLIアプリを実装してください。

## コマンドライン引数の取得方法
app.goファイルには`run`というメソッドが定義されています。

``` go
func run(args []string) {
  // code to run
}
```

ここでの `args` は、[main.go](main.go) の `main` 関数内で渡された `os.Args` です。[main.go](main.go) も編集可能なファイルとして用意されているので、CLI アプリケーションとして、オプション解析を行いたい場合は [main.go](main.go) を編集すると良いでしょう。

## コマンド実行結果の標準出力への出力
標準出力への出力は `fmt` パッケージ内の `Print` 系のメソッドで可能です。

``` go
fmt.Println(args)
```
