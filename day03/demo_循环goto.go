package main

import "log"

func main() {

	i := 0
	sum:
		log.Println(i)
		i++

		// i小于100，无条件循环
		if i < 100 {
			goto sum
		}

}
