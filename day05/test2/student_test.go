package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewStudent(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Start test new", t, func() {
		stu, err := NewStudent("")

		Convey("空name初始化错误", func() {
			So(err, ShouldBeError)
		})
		Convey("stu应该是个空指针", func() {
			So(stu, ShouldBeNil)
		})
	})
}

func TestStudent_GetAvgScore(t *testing.T) {
	stu, _ := NewStudent("xiaoyi")

	Convey("不设置分数会出错", t, func() {
		_,err := stu.GetAvgScore()

		Convey("获取平均分出错", func() {
			So(err, ShouldBeError)
		})
	})

	Convey("正常情况", t, func() {
		stu.ChiScore = 99
		stu.MathScore = 98
		stu.EngScore = 100

		score, err := stu.GetAvgScore()
		Convey("获取平均分错误",  func() {
			So(err, ShouldBeNil)
		})
		Convey("平均分大于60",  func() {
			So(score, ShouldBeGreaterThan, 60)
		})

	})

}

/*
E:\golang\HelloGolang\day05\test2>go test -v .
=== RUN   TestNewStudent

  Start test new
    空name初始化错误 .
    stu应该是个空指针 .


2 total assertions

--- PASS: TestNewStudent (0.00s)
=== RUN   TestStudent_GetAvgScore

  不设置分数会出错
    获取平均分出错 .


3 total assertions


  正常情况
    获取平均分错误 .
    平均分大于60 .


5 total assertions

--- PASS: TestStudent_GetAvgScore (0.00s)
PASS
ok      day05/test2     0.298s
*/
