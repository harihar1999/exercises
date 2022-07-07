package main

import "fmt"

type salary interface {
	calculateSalary() int
}
type fulltime struct {
	basic int
	day int
}

type contractor struct {
	basic int
	day int
}

type freelancer struct {
	basic int
	hour int
}

func (emp1 fulltime) calculateSalary() int{
	return emp1.basic*emp1.day
}

func (emp2 contractor) calculateSalary() int {
	return emp2.basic*emp2.day
}

func (emp3 freelancer) calculateSalary() int {
	return emp3.basic*emp3.hour
}




func main() {

	var sal salary

	emp1 := fulltime{500,30}
	emp2 := contractor{100,30}
	emp3 := freelancer{10,20}

	sal=emp1
	fmt.Println(sal.calculateSalary())

	sal=emp2
	fmt.Println(sal.calculateSalary())

	sal=emp3
	fmt.Println(sal.calculateSalary())




}
