package observer

import "fmt"

// Observer interface
type Observer interface {
  Update(stockPrice float64)
}


// Subject Interface 
type Stock interface {
  Register(observer Observer)
  Deregister(observer Observer)
  Notify()
}

// Concrete Subject
type TeslaStock struct {
  observers []Observer
  price float64
}

func (t *TeslaStock) Register(o Observer) {
  t.observers = append(t.observers, o)
}


func (t *TeslaStock) Deregister(o Observer) {
  for i, obs := range t.observers {
    if obs == o {
      t.observers = append(t.observers[:i], t.observers[i+1:]...)
      break
    }
  }
}

func (t *TeslaStock) Notify() {
  for _, observer := range t.observers {
    observer.Update(t.price)
  }
}

func (t *TeslaStock) SetPrice(price  float64) {
  t.price = price
  t.Notify()
}

// Concrete Observer
type Investor struct {
  name string
}

func (i *Investor) Update(stockPrice float64) {
  fmt.Printf("%s notified: New Tesla stock price is $%.2f\n", i.name, stockPrice)
}


func Run() {
  // Create stock
  tesla := &TeslaStock{}
  // Create investors
  investor1 := &Investor{name: "Pegah"}
  investor2 := &Investor{name: "Fredrick"}

  // Register investors
  tesla.Register(investor1)
  tesla.Register(investor2)

  // Stock price update
  tesla.SetPrice(420.14)
  tesla.SetPrice(420.69)
  
  // Fredrick unsubscribe
  
  tesla.Deregister(investor2)
  

  tesla.SetPrice(420.85)

}
