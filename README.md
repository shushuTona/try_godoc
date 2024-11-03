# try_godoc

## `godoc` とは

[https://pkg.go.dev/golang.org/x/tools/cmd/godoc](https://pkg.go.dev/golang.org/x/tools/cmd/godoc)

決められたフォーマットでコードにコメントを記載することで、コードのドキュメントをブラウザで閲覧することができるパッケージ。

（コマンドを実行したディレクトリに記載された `go.mod` を確認して、依存パッケージと標準パッケージの一覧を表示する）

## `go doc` との違い

`go doc` は、CLI上でパッケージのドキュメントを閲覧することができる `go` コマンドのサブコマンド。

`godoc` はパッケージ内容をブラウザで閲覧することができる。

## インストール方法

```bash
go install golang.org/x/tools/cmd/godoc@latest
```

## ドキュメントの書き方

[https://go.dev/doc/comment](https://go.dev/doc/comment)

### パッケージの概要

`package` 前に記載したコメントはドキュメントの `Overview` として表示される。

```go
// calcパッケージは、対象の数値に対して実行される下記計算処理の過程と結果を構造体で管理します。
//   - 加算（[Calc.Add]）
//   - 減算（[Calc.Sub]）
package calc
```

### 構造体

構造体定義の前の行に記載したコメントが構造体自体の説明として表示される。

フィールドの前の行のコメントも構造体のコードと一緒に表示される。

フィールドはパブリック定義の場合はドキュメントに表示され、プライベート定義のフィールドは `contains filtered or unexported fields` というコメントアウトで省略される。

```go
// Action
// 計算処理毎の記録
type Action struct {
	// 計算時の1つ目の項
	term1 int

	// 計算時の2つ目の項
	term2 int

	symbol string
}
```

### Exampleコード

`Example` テストコードを記載することで、メソッドのexampleとして表示される。

```go
func ExampleCalc_Add() {
	num := calc.NewCalc(0)

	num.Add(10)

	fmt.Println(num.GetNum())
	// Output:
	// 10
}
```

## 起動方法

```bash
godoc -http=:6060
```
