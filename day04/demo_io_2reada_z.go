package main

import "fmt"

//  举例 a-z A-Z

func alpha(r byte) byte{
	if (r>'A' && r<'Z') || (r>='a' && r<='z'){
		return r
	}
	return 0
}

func main() {

	fmt.Println(alpha('3'))
	fmt.Println(alpha('a'))
	fmt.Println(alpha('z'))
	fmt.Println(alpha('$'))

}


/*
0
97
122
0
*/