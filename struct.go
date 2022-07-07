package main

type cat struct {
	weight float64
}

func (cat cat) howWeigh() float64 {
	return cat.weight
}
func (cat cat) foodWeigh() float64 {
	foodPerMonth := 7.0
	return cat.weight * foodPerMonth
}
func (cat cat) String() string {
	return "Кот"
}

type dog struct {
	weight float64
}

func (dog dog) howWeigh() float64 {
	return dog.weight
}
func (dog dog) foodWeigh() float64 {
	foodPerMonth := 2.0
	return dog.weight / foodPerMonth
}
func (dog dog) String() string {
	return "Собака"
}

type cow struct {
	weight float64
}

func (cow cow) howWeigh() float64 {
	return cow.weight
}
func (cow cow) foodWeigh() float64 {
	foodPerMonth := 25.0
	return cow.weight * foodPerMonth
}
func (cow cow) String() string {
	return "Корова"
}
