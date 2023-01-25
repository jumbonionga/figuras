package figuras

import "fmt"

type figura interface {
	area() float64
	perimetro() float64
}

func Medidas(figura figura) {
	fmt.Println("Medidas:", figura)
	fmt.Println("Area:", figura.area())
	fmt.Println("Perimetro:", figura.perimetro())
}
