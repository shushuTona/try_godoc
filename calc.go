// calcパッケージは、対象の数値に対して実行される下記計算処理の過程と結果を構造体で管理します。
//   - 加算（[Calc.Add]）
//   - 減算（[Calc.Sub]）
package calc

import "fmt"

// Calc
// 計算内容を管理する構造体
type Calc struct {
	// 計算結果を格納するフィールド
	num int

	// 計算処理の一覧を格納するスライス
	actionStack []*Action
}

// Action
// 計算処理毎の記録
type Action struct {
	// 計算時の1つ目の項
	term1 int

	// 計算時の2つ目の項
	term2 int

	symbol string
}

func NewCalc(basenum int) *Calc {
	return &Calc{
		num:         basenum,
		actionStack: []*Action{},
	}
}

// 実行時点での計算結果を取得する
func (c *Calc) GetNum() int {
	return c.num
}

// 実行時点での計算過程を文字列として取得する
func (c *Calc) GetActions() string {
	s := ""
	for _, action := range c.actionStack {
		if s != "" {
			s += "\n"
		}
		s += fmt.Sprintf("%d%s%d", action.term1, action.symbol, action.term2)
	}

	return s
}

// 引数に指定した数値を使用して加算処理を行う
func (c *Calc) Add(num int) {
	c.actionStack = append(c.actionStack, &Action{term1: c.num, term2: num, symbol: "+"})
	c.num = c.num + num
}

// 引数に指定した数値を使用して減算処理を行う
func (c *Calc) Sub(num int) {
	c.actionStack = append(c.actionStack, &Action{term1: c.num, term2: num, symbol: "-"})
	c.num = c.num - num
}
