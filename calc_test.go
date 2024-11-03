package calc_test

import (
	"calc"
	"fmt"
	"testing"
)

func ExampleNewCalc() {
	num := calc.NewCalc(10)

	fmt.Println(num)
	// Output:
	// &{10 []}
}

func TestCalc_GetNum(t *testing.T) {
	type fields struct {
		num int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "初期化した数値を取得する",
			fields: fields{
				num: 10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := calc.NewCalc(tt.fields.num)
			if got := c.GetNum(); got != tt.want {
				t.Errorf("Calc.GetNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleCalc_GetNum() {
	num := calc.NewCalc(10)

	num.Add(10)

	fmt.Println(num.GetNum())
	// Output:
	// 20
}

func TestCalc_Add(t *testing.T) {
	type fields struct {
		num int
	}
	type args struct {
		num int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "加算処理",
			fields: fields{
				num: 10,
			},
			args: args{
				num: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := calc.NewCalc(tt.fields.num)
			c.Add(tt.args.num)
		})
	}
}

func ExampleCalc_Add() {
	num := calc.NewCalc(0)

	num.Add(10)

	fmt.Println(num.GetNum())
	// Output:
	// 10
}

func TestCalc_Sub(t *testing.T) {
	type fields struct {
		num int
	}
	type args struct {
		num int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "減算処理",
			fields: fields{
				num: 10,
			},
			args: args{
				num: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := calc.NewCalc(tt.fields.num)
			c.Sub(tt.args.num)
		})
	}
}

func ExampleCalc_Sub() {
	num := calc.NewCalc(100)

	num.Sub(10)

	fmt.Println(num.GetNum())
	// Output:
	// 90
}

func TestCalc_GetActions(t *testing.T) {
	type fields struct {
		num int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "",
			fields: fields{},
			want: `0+10
10+50
60-10`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := calc.NewCalc(tt.fields.num)
			c.Add(10)
			c.Add(50)
			c.Sub(10)
			if got := c.GetActions(); got != tt.want {
				t.Errorf("Calc.GetActions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleCalc_GetActions() {
	num := calc.NewCalc(0)

	num.Add(10)
	num.Add(50)
	num.Sub(10)

	fmt.Printf("%s", num.GetActions())
	// Output:
	// 0+10
	// 10+50
	// 60-10
}
