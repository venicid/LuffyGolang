package main

import "testing"

func TestAdd(t *testing.T) {
	a := 10
	b := 20
	want := 40
	actual := Add(a, b)
	if want != actual {
		t.Errorf("[Add 参数a:%v b:%v][期望：%v 实际:%v]", a, b, want, actual)
	}
}

func TestMul1(t *testing.T) {
	a := 10
	b := 20
	want := 200
	actual := Mul(a, b)
	if want != actual {
		t.Errorf("[Mul 参数a:%v b:%v][期望：%v 实际:%v]", a, b, want, actual)
	}
}


func TestMul2(t *testing.T)  {
	t.Run("zhengshu+", func(t *testing.T) {
		if Mul(4,5) != 20{
			t.Fatal("mult.zhengshu.error")
		}
	})

	t.Run("fushu-", func(t *testing.T) {
		if Mul(4,-2) != -8{
			t.Fatal("mult.fushu.error")
		}
	})
}

func TestMul3(t *testing.T)  {

	type tt struct {
		a int
		b int
		want int
		name string
	}

	cases := []tt{
		{4, 5, 21, "zhengshu"},
		{2, -3, -4, "fushu"},
		{0, -3, 5, "ling"},
	}

	for _, c := range cases{
		actual := Mul(c.a,c.b)
		if c.want != actual{
			t.Errorf("mul.%s.error", c.name)
		}
	}

}


func TestDiv(t *testing.T) {
	a := 10
	b := 20
	want := 2
	actual := Div(b, a)
	if want != actual {
		t.Errorf("[Add 参数a:%v b:%v][期望：%v 实际:%v]", a, b, want, actual)
	}
}


/*
E:\golang\HelloGolang\day05\test>go test -cover -run=TestMul -v .
=== RUN   TestMul1
--- PASS: TestMul1 (0.00s)
=== RUN   TestMul2
=== RUN   TestMul2/zhengshu+
=== RUN   TestMul2/fushu-
--- PASS: TestMul2 (0.00s)
    --- PASS: TestMul2/zhengshu+ (0.00s)
    --- PASS: TestMul2/fushu- (0.00s)
=== RUN   TestMul3
    compute_test.go:58: mul.zhengshu.error
    compute_test.go:58: mul.fushu.error
    compute_test.go:58: mul.ling.error
--- FAIL: TestMul3 (0.00s)
FAIL
coverage: 33.3% of statements
FAIL    day05/test      0.476s
FAIL

*/