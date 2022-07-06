package main

import "fmt"

type dog struct {
	name   string
	weight float64
}

func (a dog) getFoodPerMonth() float64 {
	return a.weight / 5 * 10
}
func (a dog) String() string {
	return "Собака"
}
func (a dog) getAnimalName() string {
	return a.name
}

type cat struct {
	name   string
	weight float64
}

func (a cat) getFoodPerMonth() float64 {
	return a.weight * 7
}
func (a cat) String() string {
	return "Кіт"
}
func (a cat) getAnimalName() string {
	return a.name
}

type cow struct {
	name   string
	weight float64
}

func (a cow) getFoodPerMonth() float64 {
	return a.weight * 25
}
func (a cow) String() string {
	return "Корова"
}
func (a cow) getAnimalName() string {
	return a.name
}

type animal interface {
	watsFoodForMonthGetter
	animalNameGetter
	fmt.Stringer
}

type animalNameGetter interface {
	getAnimalName() string
}

type watsFoodForMonthGetter interface {
	getFoodPerMonth() float64
}

func getSumFoodPerMonthForAnymalsFarm(animals []animal) float64 {
	var SumFood float64
	for _, v := range animals {
		fmt.Printf("%v: %v - потрібно на місяць %v кг.корму\n", v.String(), v.getAnimalName(), v.getFoodPerMonth())
		SumFood += v.getFoodPerMonth()
	}
	return SumFood

}

func main() {

	var myFarm = []animal{
		dog{"Рекс", 25}, dog{"Луна", 15.},
		cat{"Мурчик", 5.}, cat{"Том", 3.},
		cow{"Зірка", 350.}, cow{"Дуня", 300.},
	}

	sFood := getSumFoodPerMonthForAnymalsFarm(myFarm)
	fmt.Printf("-----------------------------------------------------------------------\n")
	fmt.Printf("Щоб прогодувати тварин на потрібно %v кг корму", sFood)

}
