package strategy

import "fmt"

// Strategy interface
type PaymentStrategy interface {
  Pay(amount float64)
}


// Concrete Strategy
type CreditCardPayment struct {}

func (c *CreditCardPayment) Pay (amount float64) {
  fmt.Printf("Paid %.2f using Credit Card. \n", amount)
}

type PayPalPayment struct {}

func (p *PayPalPayment) Pay (amount float64) {
  fmt.Printf("Paid %.2f using PayPal. \n", amount)
}

type BitcoinPayment struct {}

func (b *BitcoinPayment) Pay (amount float64) {
  fmt.Printf("Paid %.2f using PayPal. \n", amount)
}

// Context: Uses a payment strategy
type PaymentContext struct {
  strategy PaymentStrategy
}

func (p *PaymentContext) SetStrategy(strategy PaymentStrategy) {
  p.strategy = strategy
}

func (p *PaymentContext) ProcessPayment(amount float64) {
  if p.strategy == nil {
    fmt.Println("No payment strategy selected.")
    return
  }
  p.strategy.Pay(amount)
}

func Run() {
  context :=  &PaymentContext{}
  // Pay with credit Card
  context.SetStrategy(&CreditCardPayment{})
  context.ProcessPayment(100.69)

  context.SetStrategy(&PayPalPayment{})
  context.ProcessPayment(85.75)

  context.SetStrategy(&BitcoinPayment{})
  context.ProcessPayment(420.010)
}
