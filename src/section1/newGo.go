package main

import "fmt"

type ByteSize uint64

var (
	i int
	b bool
	s string
)

var name, title, num1, num2 = "jaeik", "golang", 1, 2

const NICKNAME = "laidback"

const (
	GO = iota // 여러 상수의 값을 0부터 1씩 값을 증가시킨다.
	JAVA
	PYTHON
	C
)
const (
	_           = iota             // 초기값이 0이기 때문에 버린다.
	KB ByteSize = 1 << (10 * iota) // << 연산자는 비트를 이동시킨다. 본 문법에서는 1을 왼쪽으로 10번 이동하므로 2의 10승은 1024가 됩니다.
	MB
	GB
	TB
	PT
	EB
)

func main() {
	i, b, s = 1, true, "example"

	fmt.Println(i, b, s)
	fmt.Println("nickname : ", NICKNAME)
	fmt.Println(name, title, num1, num2)

	fmt.Println(GO, JAVA, PYTHON, C)

	fmt.Println(KB, MB, GB, TB, PT, EB)
}
