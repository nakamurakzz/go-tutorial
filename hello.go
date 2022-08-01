package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"rsc.io/quote"
)

// 後から変更できる変数
var c bool = true
// d := "hello" // 関数の外側では暗黙的な宣言ができない

// 定数
const xx = "hello"
const yy = "world"

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	z int
)

// 変換
var i int = 42
var f float64 = float64(i)

func main() {
	// xx = "world" // const で定義された変数は変更できない
	fmt.Println(f)
	fmt.Println(add(z,2))
	d := "hello"
	c = false
	fmt.Println(d)
	d = "world"
	fmt.Println(d)
	fmt.Println(c)
	fmt.Println(quote.Go())
	fmt.Println(time.Now())
	fmt.Println(rand.Intn(100))
	fmt.Println(math.Pi)
	fmt.Println(add(1, 2))
	fmt.Println(returnOrgStr("1", "2"))
	fmt.Println(split(10))

	// for文、カッコは不要
	for i = 0; i < 10; i ++ {
		fmt.Println(i)
	}

	// 初期化と後処理ステートメントは不要
	for ; i < 20; {
		fmt.Println(i)
		i++
	}

	sum := 1
	for sum < 10 {
		fmt.Println(sum)
		sum += sum
	}

	// 無限ループ
	// for {
	// }

	// if文
	aa := 1
	if aa == 1 {
		fmt.Println("aa:1")
	}

	bb := 1
	if bb == 1 {
		fmt.Println("bb:1")
	} else {
		fmt.Println("bb:0")
	}

	// switch文
	// breakは不要
	switch bb {
	case 1:
		fmt.Println("bb:1 ok")
	case 2:
		fmt.Println("bb:2")
	default:
		fmt.Println("bb:0")
	}

	switch {
	case aa == 1:
		fmt.Println("aa:1")
	case aa == 2:
		fmt.Println("aa:2")
	default:
		fmt.Println("aa:0")
	}

	// deferは後処理ステートメント
	// deferはスタックに積まれていく
	// deferは、関数が終了するときに実行される
	// defer3 -> defer2 -> defer の順にコンソール出力される
	defer fmt.Println("defer")
	defer fmt.Println("defer2")
	defer fmt.Println("defer3")

	nn := 1
	defer fmt.Println(nn) // 処理が遅れるだけなので 1 が出力される
	nn = 2
	fmt.Println(nn) 

	fmt.Println("main")
}


// 引数の型が同じ場合は、最後以外の型を省略できる
// func add(x int, y int) int {
func add(x, y int) int {
	return x + y
}

func returnOrgStr(x, y string) (string, string) {
	return x , y
}

// naked return
func split (sum int)(x,y int){
	x = sum * 4 / 9
	y = sum - x
	return
}