package bridge

import "fmt"

type Notifier interface {
  SendNotification(message string)
}

type EmailNotifier struct{}

func (e *EmailNotifier) SendNotification(message string) {
  fmt.Println("Sending Email: " + message)
}


type SMSNotifier struct{}

func (s *SMSNotifier) SendNotification(message string) {
  fmt.Println("Sending SMS: " + message)
}

// Abstraction
type Notification struct {
  notifier Notifier
}

func (n *Notification) Notify(message string) {
  n.notifier.SendNotification(message)
}

func Run() {
  // Send an email Notificatione
  email := &EmailNotifier{}
  notification := Notification{notifier: email}
  notification.Notify("This is an email alert!")

  // Send an SMS notificaton
  sms := &SMSNotifier{}
  notification = Notification{notifier: sms}
  notification.Notify("This us ab SMS alert!")
}
