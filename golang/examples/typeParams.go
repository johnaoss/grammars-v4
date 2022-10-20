package samples

import "fmt"

type Complex interface {
	~complex64 | ~complex128
}

type Stack[T any] struct {
	arr []T
}

type SameLineExtra interface{ ~[]*Stack; String() string }

func (p *Stack[T]) Push(a T) {
	{
		p.arr = append(p.arr, a)
	}
}

func semiUnion[K, V interface{ ~uint8 | ~int8 }](x K, y V) {
	fmt.Println(x, y)
}

func byteString[S interface{ ~[]byte | string }](a S) {
	fmt.Println(string(a))
}
