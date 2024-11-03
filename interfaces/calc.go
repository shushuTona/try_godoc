// calcパッケージのinterface定義
package interfaces

type ICalc interface {
	GetNum() int
	GetActions() string
	Add(num int)
	Sub(num int)
}
