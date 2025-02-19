package factory

import "fmt"

// Product
type Pizza struct {
  Crust string
  Sauce string
  Toppings []string
}

// Builder interface
type PizzaBuilder interface {
  SetCrust(string)
  SetSauce(string)
  AddTopping(string)
  Build() Pizza 
}

// Concrete Builder: Margherita Pizza
type MargheritaPizzaBuilder struct {
  pizza Pizza
}

func (b * MargheritaPizzaBuilder)SetCrust(crust string) {
  b.pizza.Crust = crust
}

func (b *MargheritaPizzaBuilder) SetSauce(sauce string) {
  b.pizza.Sauce = sauce
}

func (b *MargheritaPizzaBuilder) AddTopping(topping string) {
	b.pizza.Toppings = append(b.pizza.Toppings, topping)
}

func (b *MargheritaPizzaBuilder) Build() Pizza {
  return b.pizza
}

// Concrete Builder: Vegetable Pizza
type VegetablePizzaBuilder struct {
  pizza Pizza
}

func (b *VegetablePizzaBuilder) SetCrust(crust string) {
  b.pizza.Crust = crust
}

func (b *VegetablePizzaBuilder) SetSauce(sauce string) {
  b.pizza.Sauce = sauce
}

func (b *VegetablePizzaBuilder) AddTopping(topping string) {
  b.pizza.Toppings = append(b.pizza.Toppings, topping)
}

func (b *VegetablePizzaBuilder) Build() Pizza {
  return b.pizza
}
// Concrete Builder: Pepperoni Pizza
type PepperoniPizzaBuilder struct {
  pizza Pizza
}

func (b *PepperoniPizzaBuilder) SetCrust(crust string) {
  b.pizza.Crust = crust
}

func (b *PepperoniPizzaBuilder) SetSauce(sauce string) {
  b.pizza.Sauce = sauce
}

func (b *PepperoniPizzaBuilder) AddTopping(toppping string) {
  b.pizza.Toppings = append(b.pizza.Toppings, toppping)
}

func (b *PepperoniPizzaBuilder) Build() Pizza {
  return b.pizza
}

// Director
type PizzaDirector struct {
  builder PizzaBuilder
}

func (d *PizzaDirector) ConstructMargheritaPizza() {
  d.builder.SetCrust("Thin Crust")
  d.builder.SetSauce("Tomato")
  d.builder.AddTopping("Mozzarella")
  d.builder.AddTopping("Basil")
}

func (d *PizzaDirector) ConstructVegetablePizza() {
  d.builder.SetSauce("Whole Wheat Crust")
  d.builder.SetCrust("Pesto")
  d.builder.AddTopping("Olives")
  d.builder.AddTopping("Bell Peppers")
  d.builder.AddTopping("Mushrooms")
}

func (d *PizzaDirector) ConstructPepperoniPizza() {
  d.builder.SetCrust("Thic Crust")
  d.builder.SetSauce("Tomato")
  d.builder.AddTopping("Pepperoni")
  d.builder.AddTopping("Chedar")
}


func RunBuilder() {
  // Client code using the builder pattern
// Margherita Pizza
  margheritaBuilder := &MargheritaPizzaBuilder{}
	margheritaDirector := PizzaDirector{builder: margheritaBuilder}
	margheritaDirector.ConstructMargheritaPizza()
	margheritaPizza := margheritaBuilder.Build()
	fmt.Printf("Margherita Pizza: %s with %s sauce and toppings %v\n", margheritaPizza.Crust, margheritaPizza.Sauce, margheritaPizza.Toppings)

	// Vegetable Pizza
	vegetableBuilder := &VegetablePizzaBuilder{}
	vegetableDirector := PizzaDirector{builder: vegetableBuilder}
	vegetableDirector.ConstructVegetablePizza()
	vegetablePizza := vegetableBuilder.Build()
	fmt.Printf("Vegetable Pizza: %s with %s sauce and toppings %v\n", vegetablePizza.Crust, vegetablePizza.Sauce, vegetablePizza.Toppings)

	// Pepperoni Pizza
	pepperoniBuilder := &PepperoniPizzaBuilder{}
	pepperoniDirector := PizzaDirector{builder: pepperoniBuilder}
	pepperoniDirector.ConstructPepperoniPizza()
	pepperoniPizza := pepperoniBuilder.Build()
	fmt.Printf("Pepperoni Pizza: %s with %s sauce and toppings %v\n", pepperoniPizza.Crust, pepperoniPizza.Sauce, pepperoniPizza.Toppings)
}

