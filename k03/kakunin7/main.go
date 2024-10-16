package main

import "fmt"

func SafeRecover() {
	defer func() {
		r := recover() //panicが発生したかを取得
		if r != nil {
			//panic発生
			fmt.Println("エラーの内容：", r)
		}
	}()
	panic("エラー")
	//上と同じ処理
	// if r := recover(); r!=nil{

	// }
}

func main() {
	fmt.Println("start")
	SafeRecover()
	fmt.Println("end")
}
