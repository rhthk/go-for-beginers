package module4

type Emp struct {
	name       string
	baseSalary float64
	bonus      float64
}

func (emp Emp) totalSalary() float64 {
	return emp.baseSalary + emp.bonus
}

func employeeCalculator() float64 {
	john := Emp{name: "John Doe", baseSalary: 200, bonus: 10}
	return john.totalSalary()
}
