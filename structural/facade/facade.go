package facade

import "fmt"

// Subsystyem: Amplifier
type Amplifier struct{}

func (a *Amplifier) On() {
  fmt.Println("Amilifier is On")
}

func (a *Amplifier) Off() {
  fmt.Println("Amilifier is Off")
}

// Subsystem: DVD Player
type DVDPlayer struct{}

func (d *DVDPlayer) Play() {
  fmt.Println("DVD Player is playing")
}

func (d *DVDPlayer) Stop() {
  fmt.Println("DVD Player stopped")
}

// Subsystem: Lights
type Lights struct {}

func (l *Lights) Dim() {
  fmt.Println("Lights are dimmed")
}

func (l *Lights) Brighten() {
  fmt.Println("Lights are brightened")
}

// Facade: Home Theater
type HomeTheater struct {
  amp *Amplifier
  dvd *DVDPlayer
  lights *Lights
}

func NewHomeTheater() *HomeTheater {
  return &HomeTheater{
    amp: &Amplifier{},
    dvd: &DVDPlayer{},
    lights: &Lights{},
  }
}

func (h *HomeTheater) WatchMovie() {
  fmt.Println("Starting Movie...")
  h.lights.Dim()
  h.amp.On()
  h.dvd.Play()
}

func (h *HomeTheater) EndMovie() {
  fmt.Println("Stopping Movie...")
  h.dvd.Stop()
  h.amp.Off()
  h.lights.Brighten()
}


func Run() {
  homeTheater := NewHomeTheater()
  homeTheater.WatchMovie()
  homeTheater.EndMovie()
}
