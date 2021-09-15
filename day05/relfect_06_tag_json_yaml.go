package main

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"reflect"
)

/* a.yaml
yaml_name: 李师师
yaml_age: 90
yaml_rich: false
http:
  ips:
    - 1.1
    - 2.2
  port: 88
*/

// - 代表忽略
type Person struct {
	Name string     `json:"name" yaml:"yaml_name"  xiaoyi:"name" `
	Age  int        `json:"age" yaml:"yaml_age"  xiaoyi:"age" `
	Rich bool       `json:"rich" yaml:"-"  xiaoyi:"-" `
	Hc   HttpConfig `yaml:"http"`
}

type HttpConfig struct {
	Ip   []string `yaml:"ips"  `
	Port int      `yaml:"port"  `
}

/*
# 结构体标签和反射
- json的标签解析json
- yaml的标签解析yaml
- 自定义xiaoyi标签
- 原理是t.Field.Tag.Lookup("标签名")
*/

// json解析
// json.Marshal
// json.Unmarshal
func jsonWork() {
	p := Person{
		Name: "小乙",
		Age:  19,
		Rich: true,
	}
	// 1. 先将对象解析成字符串
	data, err := json.Marshal(p)
	if err != nil {
		log.Printf("[json.Marshal.error][err:%v]", err)
		return
	}
	log.Printf("[person.json.Marshal.res :%v]", string(data))

	// 2. 从json字符串解析回这个对象
	p2Str := `{
		"name":"李逵",
		"age":28,
		"rich":true
	}`
	var p2 Person
	err = json.Unmarshal([]byte(p2Str), &p2)
	if err != nil {
		log.Printf("[json.Unmarshal.error][err:%v]", err)
		return
	}
	log.Printf("[person.json.Unmarshal.res :%v]", p2)

}

// yaml解析
// yaml.Unmarshal
// yaml.Marshal
func ymlWork() {

	// yaml解析为对象
	fileName := "a.yaml"
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("[ioutil.ReadFile.error][err:%v]", err)
		return
	}

	var p Person
	err = yaml.Unmarshal(content, &p)
	if err != nil {
		log.Printf("[yaml.Unmarshal.error][err:%v]", err)
		return
	}
	log.Printf("[person.yaml.Unmarshal.res :%v]", p)


	// 对象解析为yaml
	p1 := Person{
		Name: "abc",
		Age:  20,
		Rich: false,
	}
	data, err := yaml.Marshal(p1)
	if err != nil {
		log.Printf("[yaml.Marshal.error][err:%v]", err)
		return
	}
	err = ioutil.WriteFile("b.yaml", data, 0644)
	if err != nil {
		log.Printf("[ioutil.WriteFile.error][err:%v]", err)
		return
	}
}


// 自定义xiaoyi标签
func myTagWork() {
	p := Person{
		Name: "abc",
		Age:  10,
		Rich: false,
		Hc:   HttpConfig{},
	}
	serializeStructTag(p)
}

func serializeStructTag(s interface{}) {
	t := reflect.TypeOf(s)
	//value := reflect.ValueOf(s)
	for i := 0; i < t.NumField(); i++ {

		field := t.Field(i)
		key := field.Name
		jsonV := field.Tag.Get("json")
		yamlV := field.Tag.Get("yaml")

		if tag, ok := field.Tag.Lookup("xiaoyi"); ok {
			log.Printf("[找到了xiaoyi标签 :key：%s xiaoyi=%s]", key, tag)
		}
		log.Printf("[key=%s json=%s yaml=%s]", key, jsonV, yamlV)

	}
}

func main() {
	jsonWork()
	ymlWork()
	myTagWork()
}


/*

2021/09/16 00:55:19 [person.json.Marshal.res :{"name":"小乙","age":19,"rich":true,"Hc":{"Ip":null,"Port":0}}]
2021/09/16 00:55:19 [person.json.Unmarshal.res :{李逵 28 true {[] 0}}]

2021/09/16 00:55:19 [person.yaml.Unmarshal.res :{小姨 23 false {[] 0}}]

2021/09/16 00:55:19 [找到了xiaoyi标签 :key：Name xiaoyi=name]
2021/09/16 00:55:19 [key=Name json=name yaml=yaml_name]
2021/09/16 00:55:19 [找到了xiaoyi标签 :key：Age xiaoyi=age]
2021/09/16 00:55:19 [key=Age json=age yaml=yaml_age]
2021/09/16 00:55:19 [找到了xiaoyi标签 :key：Rich xiaoyi=-]
2021/09/16 00:55:19 [key=Rich json=rich yaml=-]
2021/09/16 00:55:19 [key=Hc json= yaml=http]
*/