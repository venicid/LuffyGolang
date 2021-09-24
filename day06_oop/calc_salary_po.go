package main
/*面向过程编程
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

*/

import "fmt"

const (
	factorA       = 4
	factorC       = 1
	salaryARatio = 0.06
	salaryCRatio = 0.03
)

func calcSalary(ratio float32, factor int, info map[string]interface{}) (
	username string, bonus int, nextSalary float32) {
	// interface 需要转换类型
	salary := info["salary"].(int)
	username = info["name"].(string)
	bonus = salary * factor
	nextSalary = float32(salary)*ratio + float32(salary)
	return username, bonus, nextSalary
}

// todo 有新的员工P6入职
func calcSalary1(ratio float32, factor int, info map[string]interface{}) (
	username string, bonus int, nextSalary float32) {
	// interface 需要转换类型
	salary := info["salary"].(int)
	username = info["name"].(string)
	bonus = salary * factor
	nextSalary = float32(salary)*ratio + float32(salary)
	return username, bonus, nextSalary
}

// TODO 有新的员工P5入职
func calcSalary2(ratio float32, factor int, info map[string]interface{}) (
	username string, bonus int, nextSalary float32) {
	// interface 需要转换类型
	salary := info["salary"].(int)
	username = info["name"].(string)
	bonus = salary * factor
	nextSalary = float32(salary)*ratio + float32(salary)
	return username, bonus, nextSalary
}

func epibolyCalcSalary(ratio float32, factor int, info map[string]interface{}) (
	username string, bonus int, nextSalary float32) {
	// interface 需要转换类型
	salary := info["salary"].(int)
	username = info["name"].(string)
	bonus = salary * factor
	return username, bonus, float32(salary)
}

func main() {
	var employeeInfo []map[string]interface{}
	var companyKPI string = "A"
	var ratio float32
	var factor int

	employee1 := map[string]interface{}{
		"name":   "Alex",
		"level":  8,
		"salary": 40000,
	}
	employee2 := map[string]interface{}{
		"name":   "Boyle",
		"level":  7,
		"salary": 30000,
	}
	epiboly1 := map[string]interface{}{  // 代表外包
		"name":   "张三",
		"level":  0,   //代表外包
		"salary": 10000,
	}
	employeeInfo = append(employeeInfo, employee1, employee2)
	fmt.Println("员工当前薪资信息\n", employeeInfo)
	for _, info := range employeeInfo {
		if companyKPI == "A" {
			ratio = salaryARatio
			factor = factorA +1

		} else if companyKPI == "C" {
			ratio = salaryCRatio
			factor = factorC +1
		} else {
			panic("Error!")
		}
		username,bonus, nextSalary := calcSalary(ratio,factor, info)
		fmt.Printf("%s在本年度的奖金是%d,下一年度工资是%.1f\n", username, bonus, nextSalary)
	}
	username,bonus, nextSalary := calcSalary(0,2, epiboly1)
	fmt.Printf("外包:%s在本年度的奖金是%d,下一年度工资是%.1f\n", username, bonus, nextSalary)
}