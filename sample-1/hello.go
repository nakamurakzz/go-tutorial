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
	// defer fmt.Println("defer")
	// defer fmt.Println("defer2")
	// defer fmt.Println("defer3")

	nn := 1
	// defer fmt.Println(nn) // 処理が遅れるだけなので 1 が出力される
	nn = 2
	fmt.Println(nn) 

	fmt.Println("main")


	// *int: 整数型のポインタ
	// ポインタ変数：メモリ上に保存されている値のアドレスを保持する変数
	var p *int
	i := 42
	// &i: 整数型のポインタ
	p = &i
	fmt.Println(*p)
	fmt.Println(p)

	// pointer
	xxx := 100
	fmt.Println("x address:\t", &xxx) // -> xxx変数のアドレス
	var yyy *int
	fmt.Println("y value:\t", yyy) // -> yyy変数には値が入っていないためnil
	fmt.Println("y address:\t", &yyy) // -> yyy変数のアドレス
	yyy = &xxx // yyy変数にxxx変数のアドレスを代入
	fmt.Println("y value:\t", yyy) // -> yyy変数にはxxx変数のアドレスが入っているため、xxx変数のアドレスが出力される
	fmt.Println("y address:\t", &yyy) // -> yyy変数のアドレス

	var z **int
	fmt.Println("z value:\t", z) // -> z変数には値が入っていないためnil
	fmt.Println("z address:\t", &z) // -> z変数のアドレス
	z = &yyy // z変数にyyy変数のアドレスを代入
	fmt.Println("z value:\t", z) // -> z変数にはyyy変数のアドレスが入っているため、yyy変数のアドレスが出力される
	fmt.Println("z address:\t", &z) // -> z変数のアドレス

	// &変数名: 変数が確保しているアドレス
	// *変数名： ポインタ変数の指すアドレスの実データ
	// *型名: 指定した型を保持するために確保したアドレス => ポインタ変数

	// Goはすべて値渡し
	// 参照渡ししたい時は、&変数名を渡す

	var xxxx = "alice"
	fmt.Println(&xxxx) // xxxx のアドレスを出力
	yyyy := &xxxx
	fmt.Println(*yyyy) // xxxx の値を出力
	fmt.Println(yyyy) // xxxx のアドレスを出力
	fmt.Println(&yyyy) // yyyy のアドレスを出力
	show(yyyy)


	// 構造体
	type Person struct {
		name string
		age int
	}
	fmt.Println(Person{"alice", 20})
	person := Person{"alice", 20}
	fmt.Println(person.name)
	personPointer := &person
	personPointer.name = "bob"
	fmt.Println(*personPointer)
	fmt.Println(person)

	personPassValue := person
	personPassValue.name = "bob2"
	fmt.Println(person)
	fmt.Println(personPassValue)

	type Site struct {
		siteName, url string
		from int
	}

	// 初期値を指定しない場合はnil相当の値が入る
	site := Site{}
	fmt.Println(site)

	site2 := Site{siteName: "golang", url: "https://golang.org/"}
	fmt.Println(site2)

	// 配列
	var names [2]string
	names[0] = "alice"
	names[1] = "bob"
	fmt.Println(names)
	// names[2] = "carol" // エラー

	// スライス
	var names2 []string
	names2 = names[1:2]

	fmt.Println(names2)

	// スライスのリテラル
	names3 := []string{"alice", "bob", "carol"}
	names3 = append(names3, "dave")
	fmt.Println(names3)

	fmt.Println(len(names3))
	fmt.Println(cap(names3))

	var ns []int
	fmt.Println(ns, len(ns), cap(ns))
	if ns == nil {
		fmt.Println("nil")
	}

	if len(ns) == 0 {
		fmt.Println("nil")
	}

	var pow = []int{1,2,4,6,18,32,64,128}
	// range
	// array.forEach((v,i)=>console.log(i,v))
	for i, v:= range pow {
		fmt.Println(i,v)
	}

	// map
	var m map[string]Vertex
	m = make(map[string]Vertex) // 指定され型のmapを初期化
	fmt.Println(m)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	ver := Vertex{40.68433, -74.39967}
	fmt.Println(ver)

	mm := map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google": {37.42202, -122.08408},
	}
	fmt.Println(mm)
	delete(mm, "Google")
	fmt.Println(mm)

	elm, ok := mm["Google"]
	fmt.Println(elm, ok)
	elm1, ok1 := mm["Bell Labs"]
	fmt.Println(elm1, ok1)

	// mapはstructとkeyを紐付けて管理
	// map
	// {
	//    key: Vertex
	// }[]
}

type Vertex struct {
	Lat, Long float64
}

func show(y *string) {
	fmt.Println("val:", *y)
	fmt.Println("address:", &y)
	fmt.Println("address:", y)
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