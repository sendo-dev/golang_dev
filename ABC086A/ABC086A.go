package main

// 必要なパッケージのインポート
import "fmt"

// 戻り値の型は引数()の後ろに記載する。なければ省略
func main() {

	/*
		変数宣言はvarのみ
		型宣言は変数名の後ろ
	*/
	var a, b int

	/*
		標準入力
		パッケージからの関数の利用は先頭が大文字（fmt.Scan）
	*/
	fmt.Scan(&a, &b)

	// 条件分岐。判定式を()で括らない
	if (a*b)%2 == 0 {
		/*
			標準出力
			パッケージからの関数の利用は先頭が大文字（fmt.Println）
		*/
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
