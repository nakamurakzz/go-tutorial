package main

import (
	"fmt"
	"math"
)

// インターフェース
type Abser interface {
	Abs() float64
}


// 構造体
type Vertex struct {
	X, Y float64
}

type MyFloat float64

// レシーバを伴う関数がメソッド
// v（Vertex型）のレシーバを持つ
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// ポインタレシーバとすることで、レシーバで受けたインスタンスの値を変更できる
// 新しく変数を作成しないため、メモリの節約
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// TypeScriptで書くとこうなる？
// class Vertex {
// 	X: number,
// 	Y: number,

// 	constructor(x:number, y: number) {
// 		this.X = x
// 		this.Y = y
// 	}

// 	scale(f: number){
// 		this.X = this.X * f
// 		this.Y = this.Y * f
// 	}

// 	abs():number {
// 		return math.sqrt(this.X * this.X + this.Y * this.Y)
// 	}
// }

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// 通常の関数
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println(t.S)
}

func main(){
	// empty interface
	// 任意の型
	// var nilT: any
	var nilI interface{}
	nilI = 42
	fmt.Println(nilI)
	nilI = "hello"
	fmt.Println(nilI)

	var typeI interface{} = "hello"

	// 型アサーション
	s := typeI.(string)
	fmt.Println(s)

	s, isOk := typeI.(string)
	fmt.Println(s, isOk)

	// f, isOk := typeI.(float64)
	// fmt.Println(f, ok)

	// f = typeI.(float64) // panic
	// fmt.Println(f)

	var i I

	var t *T

	i = t
	i.M()

	i = &T{"hello"}
	i.M()

	v := Vertex{3,4}
	fmt.Println((v.Abs()))
	fmt.Println(Abs(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v.Scale(10)
	fmt.Println(v.Abs())

	var a Abser

	a = f // MyFloat型はAbsメソッドがある
	a = &v // *Vertex型はAbsメソッドがある

	// a = v // Vertex型はAbsメソッドがない

	fmt.Println(a.Abs())

	do("test")
	do(1)
	do(0.1)

	var p = Person{"Bob", 20}
	fmt.Println(p)

	var ip = IpAddress{1, 2, 3, 4}
	fmt.Println(ip)
}

// 型のチェック
func do(i interface{}){
	// switch on type of i
	switch v := i.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("unknown", v)
	}
}

type Person struct  {
	name string
	age int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %v", p.name, p.age)
}

type IpAddress struct {
	first, second, third, fourth int
}

func (ip IpAddress) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip.first, ip.second, ip.third, ip.fourth)
}