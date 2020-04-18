//基礎から分かるGO言語
//実行したいところのコメントを外せば実行できる

package main

import "fmt"

func main() {
	//hello world表示
	//fmt.Printf("Hello World\n")

	//P45　型変換
	//変換の書式は　変換先の型(変換する値)
	/*
		var i int = 1234

		var u uint32 = uint32(i)

		var f float32 = float32(u)

		var s string = string(i)

		var b []byte = []byte("abc")

		fmt.Println(u)
		fmt.Println(f)
		fmt.Println(s)
		fmt.Println(b)
	*/
	//P69 for文
	/*
		i := 0

		for i < 5 {
			fmt.Println(i)
			i++
		}
	*/

	//Javaでいう拡張for文

	/*
		arr := [...]int{0, 1, 2, 3, 4}

		for i := range arr {
			fmt.Println(i)
		}
	*/

	//for break文

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)

	// 	if i == 2 {
	// 		break
	// 	}
	// }

	//continue
	//偶数・奇数判定
	// for i := 0; i < 5; i++ {
	// 	if i%2 != 0 {
	// 		//iの値が奇数のときスキップ（for文に戻る）
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	//P81 ポインタについて

	// //ポインタ変数宣言
	// var ptr *int
	// //int型の変数を宣言
	// var i int = 12345

	// ptr = &i

	// fmt.Println("iのアドレス", &i)
	// fmt.Println("ptrの値(変数iのアドレス):", ptr)

	// fmt.Println("iの値", i)
	// fmt.Println("ポインタ経由のiの値", *ptr)

	// *ptr = 999
	// fmt.Println("ポインタ経由のiの値", *ptr)

	// a, b := 2, 2
	// //関数に変数を渡す。aはそのまま渡す。値渡し
	// //bはアドレス演算子を使ってポインタとして渡す。ポインタ渡し
	// double(a, &b)
	// fmt.Println("値渡し", a)
	// fmt.Println("ポインタ渡し", b)\

	//P91
	// answer := plus(1, 2) //answerの変数にplus関数の戻り値を入れている。変数の型宣言がないのは関数で型宣言があり、型は分かっているため
	// fmt.Println(answer)

	//P91　多数の値を返す関数の呼び出し
	// add, sub, mul, div := calc(1, 2)
	// fmt.Println(add, sub, mul, div)

	//P93
	// holiday(1, "元旦", "成人の日")
	// holiday(2, "建国記念日の日")
	// holiday(3, "春分の日")

	//関数リテラルを使用して、関数外の変数valにアクセス
	// val := 123
	// func(i int) {
	// 	fmt.Println(i * val)
	// }(2)

	// f := func(s string) {
	// 	fmt.Println(s)
	// }
	// f("hoge")

	// var z myType = 1234
	// z.println()

	//fizzbuzz
	// 	1~100までの数字で、
	// 3で割り切れれば、「Fizz!」を表示する
	// 5で割り切れれば、「Buzz!」を表示する
	// 3と5で割り切れれば、「Fizz Buzz!」を表示する

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizz Buzz!")
		} else if i%5 == 0 {
			fmt.Println("Buzz!")
		} else if i%3 == 0 {
			fmt.Println("Fizz!")
		} else {
			fmt.Println(i)
		}
	}
}

// func double(x int, y *int) {
// 	x = x * 2
// 	*y = *y * 2
// }

//}

//P91多数の値を返す関数の宣言
// func calc(a int, b int) (int, int, int, float32) {
// 	return a + b, a - b, a * b, float32(a) / float32(b)
// }

//P93 可変長パラメーターdaysをもつ関数
// func holiday(month int, days ...string) {
// 	fmt.Printf("%d月の祝日は%d日あります。\n", month, len(days))
// 	for _, day := range days {
// 		fmt.Println(day)
// 	}
// 	fmt.Println()
// }

// type myType int

// func (varue myType) println() {
// 	fmt.Println(varue)
// }
