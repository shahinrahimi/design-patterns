package decorator
import "fmt"

type Coffee interface {
  GetCost() int
  GetDescription() string
}


type SimpleCoffee struct{}

func (s *SimpleCoffee) GetCost() int {
  return 5
}

func (s *SimpleCoffee) GetDescription() string {
  return "Simple Coffee"
}

// Decorator Milk
type MilkDecorator struct {
  coffee Coffee
}


func (m *MilkDecorator) GetCost() int {
  return m.coffee.GetCost() + 2
}

func (m *MilkDecorator) GetDescription() string {
  return m.coffee.GetDescription() + ", Milk"
}

// Decorator Sugar
type SugarDecorator struct {
  coffee Coffee
}

func (s *SugarDecorator) GetCost() int {
  return s.coffee.GetCost() + 1
}

func (s *SugarDecorator) GetDescription() string {
  return s.coffee.GetDescription() + ", Sugar"
}

func Run() {
  coffee := &SimpleCoffee{}
  fmt.Println(coffee.GetDescription(), "costs", coffee.GetCost())

  milkCoffee := &MilkDecorator{coffee}
  fmt.Println(milkCoffee.GetDescription(), "costs", milkCoffee.GetCost())

  sugarCoffee := &SugarDecorator{coffee}
  fmt.Println(sugarCoffee.GetDescription(), "costs", sugarCoffee.GetCost())

  sugarMilkCoffee := &SugarDecorator{milkCoffee}
  fmt.Println(sugarMilkCoffee.GetDescription(), "costs", sugarMilkCoffee.GetCost())
}
