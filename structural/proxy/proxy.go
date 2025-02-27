package proxy
import "fmt"
// Image Interface
type Image interface {
  Display()
}

// Real Image (Expensive to Load)
type RealImage struct {
  filename string
}

func (r *RealImage) LoadImage() {
  fmt.Println("Loading image:", r.filename)
}

func (r *RealImage) Display() {
  fmt.Println("Displaying image:", r.filename)
}

// Proxy Image (Delays Loading)
type ProxyImage struct {
  filename string
  realImage *RealImage
}


func (p *ProxyImage) Display() {
  if p.realImage == nil {
    p.realImage = &RealImage{filename: p.filename}
    p.realImage.LoadImage()
  }
  p.realImage.Display()
}

func Run() {
  // Using Proxy to delay Loading
  image1 := &ProxyImage{filename: "photo1.jpg"}
  image2 := &ProxyImage{filename: "photo2.jpg"}

  // Image is loaded only when Display() is called
  image1.Display() // Loads and 
  image1.Display()

  image2.Display() // Loads and displays new images
}
