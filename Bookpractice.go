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

	//ポインタ変数宣言
	var ptr *int
	//int型の変数を宣言
	var i int = 12345

	ptr = &i

	fmt.Println("iのアドレス", &i)
	fmt.Println("ptrの値(変数iのアドレス):", ptr)

	fmt.Println("iの値", i)
	fmt.Println("ポインタ経由のiの値", *ptr)

	*ptr = 999
	fmt.Println("ポインタ経由のiの値", *ptr)

	a, b := 2, 2
	//関数に変数を渡す。aはそのまま渡す。値渡し
	//bはアドレス演算子を使ってポインタとして渡す。ポインタ渡し
	double(a, &b)
	fmt.Println("値渡し", a)
	fmt.Println("ポインタ渡し", b)

}

func double(x int, y *int) {
	x = x * 2
	*y = *y * 2
}
