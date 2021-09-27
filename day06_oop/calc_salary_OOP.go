package main

import "fmt"

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
	factorA      = 4    // KPI为A工资系数
	factorC      = 1    // KPI为C工资系数
	salaryARatio = 0.06 // KPI为A工资涨幅率
	salaryCRatio = 0.03 //// KPI为C工资涨幅率
)

type BaseInfo struct {
	Name string
	Age  int
}

type Employee struct {
	BaseInfo
	Level  int
	Salary int
}

type CompanyAdmin struct {
	AllYearKPI   string
	EmployeeList []*Employee
}

func (c *CompanyAdmin) EnterEmployeeInfo(
	employee *Employee) []*Employee {
	newEmployeeList := append(c.EmployeeList, employee)
	return newEmployeeList
}

func (c *CompanyAdmin) GetEmployeeInfo(
	employees []*Employee) []*Employee {
	return employees
}

func (c *CompanyAdmin) CalcSalary(
	ratio float32, factor int, info Employee) (
	username string, bonus int, nextSalary float32) {
	salary := info.Salary
	username = info.Name
	bonus = salary * factor
	nextSalary = float32(salary)*ratio + float32(salary)
	return username, bonus, nextSalary

}

func (c *CompanyAdmin) GetSalary(employees []*Employee) string {
	var ratio float32
	var factor int
	for _, info := range employees {
		if c.AllYearKPI == "A" {
			ratio = salaryARatio
			factor = factorA + 1
		} else if c.AllYearKPI == "C" {
			ratio = salaryCRatio
			factor = factorC + 1
		} else {
			panic("Error!")
		}
		username, bonus, nextSalary := c.CalcSalary(ratio, factor, *info)
		fmt.Printf("%s在本年度的奖金是%d,下一年度工资是%.1f\n", username, bonus, nextSalary)
	}
	return fmt.Sprintln(ratio, factor)
}

type EpibolyAdmin struct {
	EpiboleList []*Employee
	CompanyAdmin
}

func (e *EpibolyAdmin) EnterEmployeeInfo(
	epiboly *Employee) []*Employee {
	newEpibolyList := append(e.EpiboleList, epiboly)
	return newEpibolyList
}

func (e *EpibolyAdmin) GetSalary(epiboles []*Employee) string {
	var ratio float32
	var factor int
	for _, info := range epiboles {
		username, bonus, nextSalary := e.CalcSalary(0, 2, *info)
		fmt.Printf("%s在本年度的奖金是%d,下一年度工资是%.1f\n", username, bonus, nextSalary)
	}
	return fmt.Sprintln(ratio, factor)
}

type PaySalary interface {
	GetSalary([]*Employee)(string)
}

func PaySomeSalary(paySalary PaySalary, user []*Employee){
	paySalary.GetSalary(user)
}

func main() {
	hr := &CompanyAdmin{AllYearKPI: "A"}
	EmployeeInfo1 := hr.EnterEmployeeInfo(&Employee{
		BaseInfo{
			"Alex", 30,
		},
		8,
		40000,
	})
	EmployeeInfo2 := hr.EnterEmployeeInfo(&Employee{
		BaseInfo{
			"Boyle", 18,
		},
		7,
		30000,
	})

	hr2 := &EpibolyAdmin{}
	epibole1 := hr2.EnterEmployeeInfo(
		&Employee{
			BaseInfo{
				"张三", 18,
			},
			0,
			10000,
		},
	)
	fmt.Println(hr.EmployeeList)
	hr.GetSalary(EmployeeInfo1)
	// hr.GetSalary(EmployeeInfo2)
	// hr2.GetSalary(epibole1)
	PaySomeSalary(hr, EmployeeInfo1)
	PaySomeSalary(hr, EmployeeInfo2)
	PaySomeSalary(hr2, epibole1)

}
