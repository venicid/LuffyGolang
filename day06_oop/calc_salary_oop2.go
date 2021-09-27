package main

import (
	"fmt"
	"log"
)

/*面向对象编程
为“北京露菲有限公司”开发一款“年底工资&奖金结算生成器”程序，需求如下：
公司目前除了老板之外，只有两名正式员工，一名外包。
两名正式员工分别职级P7和职级P8，工资为3w/月、4w/月，外包1w/月。
年底需要根据公司业绩A(盈利)、C(亏损)进行员工年底奖金核算。
奖金发放规则：
      （已知这两名正式员工年底的个人KPI绩效都已达标）
      a. 如果公司业绩评为A，那么这两名员工年底多拿4个月工资+普调6%
      b. 如果公司业绩评为C,  年底多拿1个月工资+普调3%
      ========
      c. 外包无须受公司业绩限制，奖金统一年底多拿1个月，无普调


// UML建模
员工基类
   员工类(包含外包员工)

企业HR管理类
    外包HR管理类
    -- 员工信息录入
    -- 员工信息查询
    -- 工资计算
    -- 工资结果查看
*/

const (
	factorA       = 4
	factorC       = 1
	salaryARatio = 0.06
	salaryCRatio = 0.03
)

type Person struct {
	Name string
	Age int
}

type Employee struct {
	Person
	Salary int
	Level int
}

type CompanyAdmin struct {
	AllYearKPI string
	EmployeeList []*Employee
}

// EnterEmployeeInfo 录入员工信息/**
func (c *CompanyAdmin) EnterEmployeeInfo(emp *Employee) []*Employee {
	c.EmployeeList = append(c.EmployeeList, emp)
	return c.EmployeeList
}

func (c *CompanyAdmin) GetEmployeeInfo() []*Employee  {
	return c.EmployeeList
}

func (c *CompanyAdmin) CalcSalary(ratio float32, factor int,emp Employee)(
	userName string, bonus int, newSalary float32)  {
	salary := emp.Salary
	userName = emp.Name
	bonus = salary * factor
	newSalary = float32(salary) * (ratio + float32(1))
	return userName, bonus, newSalary
}

func (c *CompanyAdmin)  GetSalary(emp []*Employee) string{
	var ratio float32
	var factor int
	for _, emp := range emp{
		if c.AllYearKPI == "A"{
			ratio = salaryARatio
			factor = factorA + 1
		} else if c.AllYearKPI =="C"{
			ratio = salaryCRatio
			factor = factorC + 1
		} else{
			log.Println("不支持")
		}

		username, bonus, nextSalary := c.CalcSalary(ratio, factor, *emp)
		fmt.Printf("%s在本年度的奖金是%d,下一年度工资是%.1f\n", username, bonus, nextSalary)
	}
	return fmt.Sprintln(ratio, factor)
}

/**
外包人员
 */
type EpibolyAdmin struct {
	EpiboleList []*Employee
	CompanyAdmin
}

func (e *EpibolyAdmin) EnterEmployeeInfo(epiboly *Employee) []*Employee {
	newEpibolyList := append(e.EpiboleList, epiboly)
	return newEpibolyList
}

func (e *EpibolyAdmin) GetSalary(epiboles []*Employee) string {
	var ratio float32
	var factor int
	for _, info := range epiboles {
		username, bonus, nextSalary := e.CalcSalary(0, 1, *info)
		fmt.Printf("%s在本年度的奖金是%d,下一年度工资是%.1f\n", username, bonus, nextSalary)
	}
	return fmt.Sprintln(ratio, factor)
}

/**
统一支付接口
 */
type PaySalary interface {
	GetSalary([]*Employee)(string)
}

func PaySomeSalary(paySalary PaySalary, user []*Employee)  {
	paySalary.GetSalary(user)
}

func main(){
	hr := &CompanyAdmin{
		AllYearKPI: "A",
	}

	emp1 := Employee{
		Person{
			"jack",
			 33,
		},
		30000,
		7,
	}

	emp2 := Employee{
		 Person: Person{
			Name: "alex",
			Age: 33,
		},
		Salary: 40000,
		Level: 8,
	}

	//fmt.Println(emp1)
	empList1 := hr.EnterEmployeeInfo(&emp1)
	//fmt.Println(empList1)
	hr.GetSalary(empList1)

	//fmt.Println(emp2)
	empList2 := hr.EnterEmployeeInfo(&emp2)
	//fmt.Println(empList2)
	hr.GetSalary(empList2)

	//fmt.Println(hr.GetEmployeeInfo())

	hr2 := &EpibolyAdmin{}
	epibole1 := hr2.EnterEmployeeInfo(
		&Employee{
			Person{
				"张三", 18,
			},
			10000,
			0,
		},
	)
	hr2.GetSalary(epibole1)

	// 多态
	fmt.Println("######################")
	PaySomeSalary(hr, empList2)
	PaySomeSalary(hr2, epibole1)
}

/*
jack在本年度的奖金是150000,下一年度工资是31800.0
jack在本年度的奖金是150000,下一年度工资是31800.0
alex在本年度的奖金是200000,下一年度工资是42400.0
张三在本年度的奖金是10000,下一年度工资是10000.0
######################
jack在本年度的奖金是150000,下一年度工资是31800.0
alex在本年度的奖金是200000,下一年度工资是42400.0
张三在本年度的奖金是10000,下一年度工资是10000.0

Process finished with the exit code 0
*/