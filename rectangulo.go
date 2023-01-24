package figuras

type Rectangulo struct {
	Ancho  float64
	Altura float64
}

func (rectangulo *Rectangulo) perimetro() float64 {
	return 2*rectangulo.Altura + 2*rectangulo.Ancho
}

func (rectangulo *Rectangulo) area() float64 {
	return rectangulo.Altura * rectangulo.Ancho
}
