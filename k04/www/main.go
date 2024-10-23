package main

import (
	"fmt"
	"net/http"
)

// ハンドラー関数
// ハンドラーで使う関数は引数は決まっている
// 第1引数：レスポンス用
// 第1引数：リクエスト用
// 同じパッケージで関数を使う場合は関数の頭文字を小文字にする(大文字でも動いたのはたまたま…？)
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "テスト")
}

func main() {
	//ルート設定
	http.HandleFunc("/", Handler) //第1引数にエンドポイント、第2引数に実行関数
	http.HandleFunc("/test", Test)

	//webサーバー起動
	http.ListenAndServe(":8080", nil)
}
