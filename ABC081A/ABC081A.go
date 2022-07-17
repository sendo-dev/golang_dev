package main

// 必要なパッケージのインポート
import "fmt"

// 戻り値の型は引数()の後ろに記載する。なければ省略
func main() {

	/*
		変数宣言はvarのみ
		型宣言は変数名の後ろ
	*/
	var a string
	var result int = 0
	const hit = "1"

	/*
		標準入力
		パッケージからの関数の利用は先頭が大文字（fmt.Scan）
	*/
	fmt.Scan(&a)

	/*
		文字列の長さを取得
		バイト数：len()
		文字数　：utf8.RuneCountInString()
	*/
	for i := 0; i < len(a); i++ {
		// ヨーダ記法は注意される（定数を左に書く記法）
		if a[i:i+1] == hit {
			result++
		}
	}

	fmt.Print(result)

}
