//基礎から分かるGO言語
//実行したいところのコメントを外せば実行できる

package main

import "fmt"

func main() {
	//hello world表示
	//fmt.Printf("Hello World\n")

	//P45　型変換
	//変換の書式は　変換先の型(変換する値)
	var i int = 1234

	var u uint32 = uint32(i)

	var f float32 = float32(u)

	var s string = string(i)

	var b []byte = []byte("abc")

	fmt.Println(u)
	fmt.Println(f)
	fmt.Println(s)
	fmt.Println(b)
}
