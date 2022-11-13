package main

import "fmt"

type Command interface {
	execute()
}

type Restaurant struct {
	AllDishes   int
	CleanDishes int
}

func NewRestaurant() *Restaurant {
	const totalDishes = 10
	return &Restaurant{
		AllDishes:   totalDishes,
		CleanDishes: totalDishes,
	}
}

type MakePizzaCommand struct {
	n          int
	restaurant *Restaurant
}

func (c *MakePizzaCommand) execute() {
	c.restaurant.CleanDishes -= c.n
	fmt.Println("Made:", c.n, "pizzas")
}

type MakeSaladCommand struct {
	n          int
	restaurant *Restaurant
}

func (c *MakeSaladCommand) execute() {
	c.restaurant.CleanDishes -= c.n
	fmt.Println("Made:", c.n, "salads")
}

type CleanDishesCommand struct {
	restaurant *Restaurant
}

func (c *CleanDishesCommand) execute() {
	c.restaurant.CleanDishes = c.restaurant.AllDishes
	fmt.Println("All dishes cleaned")
}

func (r *Restaurant) MakePizza(count int) Command {
	return &MakePizzaCommand{
		restaurant: r,
		n:          count,
	}
}

func (r *Restaurant) MakeSalad(count int) Command {
	return &MakeSaladCommand{
		restaurant: r,
		n:          count,
	}
}

func (r *Restaurant) CleanAllDishes() Command {
	return &CleanDishesCommand{
		restaurant: r,
	}
}

type Cook struct {
	Commands []Command
}

func (c *Cook) executeCommands() {
	for _, c := range c.Commands {
		c.execute()
	}
}

func main() {
	rest := NewRestaurant()

	tasks := []Command{
		rest.MakePizza(2),
		rest.MakeSalad(5),
		rest.MakePizza(1),
		rest.CleanAllDishes(),
		rest.MakePizza(10),
		rest.CleanAllDishes(),
	}

	cooks := []*Cook {
		&Cook{},
		&Cook{},
	}

	for i, task := range tasks {
		cook := cooks[i % len(cooks)]
		cook.Commands = append(cook.Commands, task)
	}

	for i, c := range cooks {
		fmt.Println("cook", i, ":")
		c.executeCommands()
	}
}

// Шаблон команды
// Шаблон команды полезен, когда вам нужно выполнить задачи,
// но вы хотите отделить управление задачами от выполнения самой задачи.

