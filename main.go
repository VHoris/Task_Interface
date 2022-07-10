package main

import (
	"errors"
	"fmt"
)

type dog struct {
	name   string
	weight float64
}

func (a dog) foodPerMonth() float64 {
	return a.weight / 5 * 10
}
func (a dog) String() string {
	return "Собака"
}
func (a dog) getAnimalName() string {
	return a.name
}
func (a dog) getWeight() float64 {
	return a.weight
}

type cat struct {
	name   string
	weight float64
}

func (a cat) foodPerMonth() float64 {
	return a.weight * 7
}
func (a cat) String() string {
	return "Кіт"
}
func (a cat) getAnimalName() string {
	return a.name
}
func (a cat) getWeight() float64 {
	return a.weight
}

type cow struct {
	name   string
	weight float64
}

func (a cow) foodPerMonth() float64 {
	return a.weight * 25
}
func (a cow) String() string {
	return "Корова"
}
func (a cow) getAnimalName() string {
	return a.name
}
func (a cow) getWeight() float64 {
	return a.weight
}

type animal interface {
	foodConsumer
	animalNameGetter
	fmt.Stringer
	animalWeightGetter
}

type animalNameGetter interface {
	getAnimalName() string
}

type foodConsumer interface {
	foodPerMonth() float64
}

type animalWeightGetter interface {
	getWeight() float64
}

func sumFoodPerMonth(animals []animal) (float64, error) {
	var SumFood float64
	for _, v := range animals {

		err := validateAnimalType(v)
		if errors.Is(err, animalTypeError) {
			fmt.Printf("failed to calculate sum food %s - %s: %v\n", v.String(), v.getAnimalName(), err)
			continue
		}

		err = validateAnimalWeight(v)
		if errors.Is(err, animalWeightError) {
			return 0, fmt.Errorf("failed to calculate sum food %s - %s: %v\n", v.String(), v.getAnimalName(), err)
		}

		err = validateCanEat(v)
		if errors.Is(err, animalCanEatError) {
			fmt.Printf("failed to calculate sum food %s - %s: %v\n", v.String(), v.getAnimalName(), err)
			continue
		}

		fmt.Printf("%v: %v - потрібно на місяць %v кг.корму\n", v.String(), v.getAnimalName(), v.foodPerMonth())
		SumFood += v.foodPerMonth()
	}
	return SumFood, nil

}

var animalTypeError = errors.New("wrong animal's type:")

func validateAnimalType(a animal) error {
	animalsName := map[string]string{"Рекс": "Собака", "Луна": "Собака", "Мурчик": "Кіт", "Том": "Кіт", "Зірка": "Корова", "Дуня": "Корова"}
	if anName := a.getAnimalName(); animalsName[anName] != a.String() {
		return animalTypeError
	}
	return nil
}

var animalWeightError = errors.New("weight is more than normal:")

func validateAnimalWeight(a animal) error {
	animalsMaxWeight := map[string]float64{"Собака": 50, "Кіт": 20, "Корова": 300}
	if weight := a.getWeight(); weight > animalsMaxWeight[a.String()] {
		return animalWeightError
	}
	return nil
}

var animalCanEatError = errors.New("can't eat")

func validateCanEat(a animal) error {
	animalsCanEat := map[string]bool{"Собака": true, "Кіт": false, "Корова": true}
	if !animalsCanEat[a.String()] {
		return animalCanEatError
	}
	return nil
}

func main() {

	var myFarm = []animal{
		dog{"Мурчик", 25}, dog{"Луна", 15.},
		cat{"Мурчик", 5.}, cat{"Том", 3.},
		cow{"Зірка", 300.}, cow{"Дуня", 350.},
	}

	sFood, err := sumFoodPerMonth(myFarm)
	if err != nil {
		fmt.Printf("All is bed: %v", err)
		return
	}
	fmt.Printf("-----------------------------------------------------------------------\n")
	fmt.Printf("Щоб прогодувати тварин на потрібно %v кг корму", sFood)

}
