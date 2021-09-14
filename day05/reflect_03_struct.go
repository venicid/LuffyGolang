package main

import (
	"log"
	"reflect"
)

type Person struct {
	Name string
	Age  int
	abc  string
}

type Student struct {
	Person
	StudentId  int
	SchoolName string
	IsBaoSOng  bool
	Hobbies    []string
	//hobbies    []string
	Labels     map[string]string
}

// 非指针型的方法
func (s Student) GoHome() {
	log.Printf("[回家了][%v]", s.Name)
}

// 指针型的方法
func (s *Student) GoToSchool() {
	log.Printf("[上学了][%v]", s.Name)
}

// 小写方法
func (s *Student) baoSong() {
	log.Printf("[参加竞赛，报送][%v]", s.Name)
}

func main() {
	p := Person{
		Name: "alex",
		Age:  0,
		abc:  "test",
	}

	s := Student{
		Person:     p,
		StudentId:  1234,
		SchoolName: "五道口男子技术学院",
		IsBaoSOng:  true,
		Hobbies:    []string{"唱歌", "跳舞", "Rap"},
		//hobbies:    []string{"唱歌", "跳舞", "Rap"},
		Labels: map[string]string{"k1":"v1", "k2":"v2"},
	}

	// TypeOf 获取模板对象的类型
	t := reflect.TypeOf(s)
	// ValueOf 获取模板对象的值类型
	v := reflect.ValueOf(s)

	// .Name  获取目标的名称
	log.Printf("[对象类型的名称：%v]", t.Name())

	// NumField返回type's field count
	for i := 0; i < t.NumField(); i++ {

		// Field代表对象的字段名对象
		key := t.Field(i)
		// 通过v.Field.Interface获取字段的值
		value := v.Field(i).Interface()

		// Anonymous代表字段是否是匿名字段
		anonymous := "非匿名"
		if key.Anonymous {
			anonymous = "匿名"
		}

		log.Printf("[%s 字段][第:%d个字段][字段的名称:%s][字段的类型:%v][字段的值：%v]",
			anonymous,
			i+1,
			key.Name,
			key.Type,
			value,
		)
	}


	// 通过NumMethod, 获取对象绑定的所有方法的counts
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		log.Printf("[第:%d个方法][方法的名称:%s][方法的类型:%v]", i+1, m.Name, m.Type)
	}
}

/*
2021/09/15 01:05:53 [对象类型的名称：Student]
2021/09/15 01:05:53 [匿名 字段][第:1个字段][字段的名称:Person][字段的类型:main.Person][字段的值：{alex 0 test}]
2021/09/15 01:05:53 [非匿名 字段][第:2个字段][字段的名称:StudentId][字段的类型:int][字段的值：1234]
2021/09/15 01:05:53 [非匿名 字段][第:3个字段][字段的名称:SchoolName][字段的类型:string][字段的值：五道口男子技术学院]
2021/09/15 01:05:53 [非匿名 字段][第:4个字段][字段的名称:IsBaoSOng][字段的类型:bool][字段的值：true]
2021/09/15 01:05:53 [非匿名 字段][第:5个字段][字段的名称:Hobbies][字段的类型:[]string][字段的值：[唱歌 跳舞 Rap]]
2021/09/15 01:05:53 [非匿名 字段][第:6个字段][字段的名称:Labels][字段的类型:map[string]string][字段的值：map[k1:v1 k2:v2]]
2021/09/15 01:05:53 [第:1个方法][方法的名称:GoHome][方法的类型:func(main.Student)]

Process finished with the exit code 0

*/