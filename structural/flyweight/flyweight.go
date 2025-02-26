package flyweight

import "fmt"


// Flyweight: Character Formating
type CharacterFormat struct {
  font string
  size int 
  color string
}

// Flyweight: Factory: Manage shared objects
type FormatFactory struct {
  formats map[string]*CharacterFormat
}

func NewFormatFactory() *FormatFactory {
  return &FormatFactory{
    formats: make(map[string]*CharacterFormat),
  }
}

func (f *FormatFactory) GetFormat(font string, size int, color string) *CharacterFormat {
  key := fmt.Sprintf("%s-%d-%s", font, size, color)
  if format, exists := f.formats[key]; exists {
    return format
  }
  newFormat := &CharacterFormat{font, size, color}
  f.formats[key] = newFormat
  return newFormat
}


// Character: Uses shared format
type Character struct {
  letter string
  format *CharacterFormat
}

func (c *Character) Display() {
	fmt.Printf("Character: %s, Font: %s, Size: %d, Color: %s\n",
		c.letter, c.format.font, c.format.size, c.format.color)
}


func Run() {
  factory := NewFormatFactory()

  // Shared format objects
  format1 := factory.GetFormat("Arial", 12, "Black")
  format2 := factory.GetFormat("Arial", 12, "Black")
  format3 := factory.GetFormat("Time New Roman", 14, "Blue")

  // Creating characters
  charA := &Character{letter: "A", format: format1}
  charB := &Character{letter: "B", format: format2}
  charC := &Character{letter: "C", format: format3}
  
  charA.Display()
  charB.Display()
  charC.Display()
  
  // Displaying memory addresses tp show shared objects
  fmt.Println("Format of A:", format1)
  fmt.Println("Format of B (shared with A):", format2)
  fmt.Println("Format of C (unique)", format3)

}


