package main

// 必要なパッケージのインポート
import "fmt"

// 戻り値の型は引数()の後ろに記載する。なければ省略
func main() {

	/*
		Goでは
		固定長 の配列が Array（宣言時、[]内に大きさを指定）
		可変長 の配列が Slice（宣言時、[]内は未指定）
		いずれも宣言は var
		※また、配列の大きさにはconst値しか利用できない（sliceLinを利用してsliceの大きさを定義できない）
	*/
	var slice []int
	var sliceLen int

	/*
		標準入力
		パッケージからの関数の利用は先頭が大文字（fmt.Scan）
	*/
	fmt.Scan(&sliceLen)

	for i := 0; i < sliceLen; i++ {
		var tmp int
		fmt.Scan(&tmp)
		// slice配列にtmpを追加
		slice = append(slice, tmp)
	}

	/*
		Goにwhile文はない
		loopやgotoを利用する
	*/
	var check bool = true
	var loopCnt int = 0

loop:

	for i := 0; i < sliceLen; i++ {
		if slice[i]%2 == 0 {
			slice[i] = slice[i] / 2
		} else {
			check = false
		}
	}

	if check {
		loopCnt++
		goto loop
	}

	fmt.Print(loopCnt)

}
