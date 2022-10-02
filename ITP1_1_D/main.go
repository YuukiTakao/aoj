/**
時計

秒単位の時間 S が与えられるので、h:m:s の形式へ変換して出力してください。ここで、h は時間、m は 60 未満の分、s は 60 未満の秒とします。

Input
S が１行に与えられます。

Output
h、m、s を :（コロン）区切りで１行に出力してください。数値が１桁の場合、0 を付けて２桁表示をする必要はありません。

Constraints

*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var S int
	fmt.Fscan(reader, &S)
	//fmt.Printf("%d\n", S)

	//fmt.Printf("%d\n", S/3600)
	h := S / 3600
	m := (S - h*3600) / 60
	s := S - (h*3600 + m*60)
	fmt.Printf("%d:%d:%d\n", h, m, s)
}

var reader = bufio.NewReader(os.Stdin)
