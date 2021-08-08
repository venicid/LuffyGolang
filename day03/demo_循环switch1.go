package main

import "log"

func calcEnv(env string) string {
	var res string

	switch env {
	case "fat":
		res = "FAT"
	case "pre":
		res =  "PRE"
	case "pro":
		res = "PRO"
	}

	return res
}

func main()  {

	res := calcEnv("fat")
	log.Println(res)

	res1 := calcEnv("pre")
	log.Println(res1)
}


