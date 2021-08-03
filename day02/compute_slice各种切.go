package main

import "fmt"

func main(){

	s1 := []int{1,2,3,4,5,6,7,8,9}
	fmt.Printf("s1[值：%v][新切片的长度=%d 容量=%d]\n", s1,len(s1), cap(s1))
	s2 := s1[2:6]
	fmt.Printf("s1[2:6][从索引为2，第3个元素开始，往后切4个][值：%v][新切片的长度=%d 容量=%d]\n", s2,len(s2), cap(s2))
	s3 :=s1[5:]
	fmt.Printf("s1[5:][从索引为5，第6个元素开始，切到最后][值：%v][新切片的长度=%d 容量=%d]\n", s3,len(s3), cap(s3))
	s4 :=s1[:4]
	fmt.Printf("s1[:4][从开始，切到索引为4，第5个元素结束][值：%v][新切片的长度=%d 容量=%d]\n", s4,len(s4), cap(s4))
	s5 :=s1[:]
	fmt.Printf("s1[:][从开始，切到最后][值：%v][新切片的长度=%d 容量=%d]\n", s5,len(s5), cap(s5))
	s6 :=s1[2:6:6]
	fmt.Printf("s1[2:6:6][从索引为2，第3个元素开始，，往后切4个元素][值：%v][新切片的长度=%d 容量=%d]\n", s6,len(s6), cap(s6))
	s7 :=s1[2:6:9]
	fmt.Printf("s1[2:6:9][从索引为2，第3个元素开始，往后切4个元素][值：%v][新切片的长度=%d 容量=%d]\n", s7,len(s7), cap(s7))

	fmt.Println("切片容量 = cap/end - start")

	/*

	s1[值：[1 2 3 4 5 6 7 8 9]][新切片的长度=9 容量=9]
	切片容量 = cap/end - start
	s1[2:6][从索引为2，第3个元素开始，往后切4个][值：[3 4 5 6]][新切片的长度=4 容量=7]
	s1[5:][从索引为5，第6个元素开始，切到最后][值：[6 7 8 9]][新切片的长度=4 容量=4]
	s1[:4][从开始，切到索引为4，第5个元素结束][值：[1 2 3 4]][新切片的长度=4 容量=9]
	s1[:][从开始，切到最后][值：[1 2 3 4 5 6 7 8 9]][新切片的长度=9 容量=9]
	s1[2:6:6][从索引为2，第3个元素开始，切到索引为6，第7个元素][值：[3 4 5 6]][新切片的长度=4 容量=4]
	s1[2:6:9][从索引为2，第3个元素开始，切到索引为6，第7个元素][值：[3 4 5 6]][新切片的长度=4 容量=7]
	*/

	s1[5] = 999
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println(s6)
	fmt.Println(s7)

	/*
	[1 2 3 4 5 999 7 8 9]
	[3 4 5 999]
	[999 7 8 9]
	[1 2 3 4]
	[1 2 3 4 5 999 7 8 9]
	[3 4 5 999]
	[3 4 5 999]
	*/

}