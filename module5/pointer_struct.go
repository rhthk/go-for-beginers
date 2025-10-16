package module5

type Person struct {
	name string
	age  int
}

func updateAge(p *Person, age int) {
	p.age = age
}
