package calc

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // go.mod依存確認用のパッケージ
)

const (
	CONST_1 = 100 // 定数1
	CONST_2 = 200 // 定数2
	CONST_3 = 300 // 定数3
)

// グローバル変数
var GlobalVariable = 100

// グローバル関数
func GlobalFunc() {
	fmt.Println("GlobalFunc")
}

// CommentStruct
// コメント記載用の構造体
//
// # 見出しを記載する際は前後に改行が必要
//
//   - リスト1
//   - リスト2
//   - リスト3
//
// コードブロック
//
//	func HogeFunc(num int) {
//		fmt.Println(num)
//	}
//
// 通常リンク
//
// https://github.com/shushuTona/try_godoc
//
// ドキュメントリンク
//
// [Calc.Add]
type CommentStruct struct {
	// 公開フィールド
	PublicField string

	// 非公開フィールド
	privateField string
}
