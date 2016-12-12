package main

/*

#include <stdio.h>
#include <stdlib.h>

int Hoge(){
	puts("hoge");
	fflush(stdout);
	return 777;
}

*/
import "C"
import "fmt"

func main() {
	fmt.Println("hello")

	ret := C.Hoge()
	fmt.Printf("ret: %T, %v\n", ret, ret)

	v := C.random()
	fmt.Println("v:", v)

	fmt.Println("bye!")
}
