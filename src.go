// Click here and start typing.
package main

import (
	"fmt"
	"math"
)

// Exibir display de escolha para o usuário
func display() string {
	fmt.Println("Calculadora de Geometria Plana e Espacial")
	fmt.Println("(1) Triângulo equilátero")
	fmt.Println("(2) Retângulo")
	fmt.Println("(3) Quadrado")
	fmt.Println("(4) Círculo")
	fmt.Println("(5) Pirâmide com base quadrangular")
	fmt.Println("(6) Cubo")
	fmt.Println("(7) Paralelepîpedo	")
	fmt.Println("(8) Esfera")
	fmt.Println("(0) Sair")
	fmt.Print("Digite a sua opção: ")
	//fmt.Scanln(&opcao)
	return ""
}

func triangulo(base, altura float32) (float32, float32) {
	// No menu apresentado trata-se de um triângulo equilátero
	area := (base * altura) / 2
	perimetro := base + altura + base
	return area, perimetro
}

func retangulo(base, altura float32) (float32, float32) {
	area := base * altura
	perimetro := 2 * (base + altura)
	return area, perimetro
}

func quadrado(lado float32) (float32, float32) {
	area := lado * lado
	perimetro := 4 * lado
	return area, perimetro
}

func circulo(raio float32) (float32, float32) {
	area := (math.Pi) * raio * raio
	perimetro := (math.Pi) * 2 * raio
	return area, perimetro
}

func piramide(areabase, arealateral, altura float32) (float32, float32) {
	area := areabase + arealateral
	volume := 1 / 3 * areabase * altura
	return area, volume
}

func cubo(aresta float32) (float32, float32) {
	area := 6 * aresta * aresta
	volume := aresta * aresta * aresta
	return area, volume
}

func paralelepipedo(aresta1, aresta2, aresta3 float32) (float32, float32) {
	area := (2 * aresta1 * aresta2) + (2 * aresta1 * aresta3) + (2 * aresta2 * aresta3)
	volume := aresta1 * aresta2 * aresta3
	return area, volume
}

func esfera(raio float32) (float32, float32) {
	area := 4 / 3 * (math.Pi) * raio * raio
	volume := 4 / 3 * (math.Pi) * raio * raio * raio
	return area, volume
}

func main() {
	fmt.Println(math.Pi)
	var op int

	//fmt.Scanln(&op)
	fmt.Println(display())
	fmt.Scanln(&op)
	switch op {
	case 1:
		var base, altura float32
		// No menu apresentado trata-se de um triângulo equilátero
		fmt.Print("Digite a base do triângulo: ")
		fmt.Scanln(&base)
		fmt.Print("Digite a altura do triângulo: ")
		fmt.Scanln(&altura)
		area, perimetro := triangulo(base, altura)
		fmt.Println("A área do Triângulo é ", area)
		fmt.Println("O perímetro do Triângulo é ", perimetro)
	case 2:
		var base, altura float32
		fmt.Print("Digite a base do retângulo: ")
		fmt.Scanln(&base)
		fmt.Print("Digite a altura do retângulo: ")
		fmt.Scanln(&altura)
		area, perimetro := retangulo(base, altura)
		fmt.Println("A área do Retângulo é ", area)
		fmt.Println("O perímetro do Retângulo é ", perimetro)
	case 3:
		var lado float32
		fmt.Print("Digite o lado do quadrado: ")
		fmt.Scanln(&lado)
		area, perimetro := quadrado(lado)
		fmt.Println("A área do Quadrado é ", area)
		fmt.Println("O perímetro do Quadrado é ", perimetro)
	case 4:
		var raio float32
		fmt.Print("Digite o raio do círculo: ")
		fmt.Scanln(&raio)
		area, perimetro := circulo(raio)
		fmt.Println("A área do Círculo é ", area)
		fmt.Println("O perímetro do Círculo é ", perimetro)
	case 5:
		var areabase, arealateral, altura float32
		fmt.Print("Digite a área da base da pirâmide: ")
		fmt.Scanln(&areabase)
		fmt.Print("Digite a área lateral da pirâmide: ")
		fmt.Scanln(&arealateral)
		fmt.Print("Digite a altura da pirâmide: ")
		fmt.Scanln(&altura)
		area, volume := piramide(areabase, arealateral, altura)
		fmt.Println("A área da Pirâmide com base quadrangular é ", area)
		fmt.Println("O volume da Pirâmide com base quadrangular é ", volume)
	case 6:
		var raio float32
		fmt.Print("Digite o raio do cubo: ")
		fmt.Scanln(&raio)
		area, volume := cubo(raio)
		fmt.Println("A área do Cubo é ", area)
		fmt.Println("O volume do Cubo é ", volume)
	case 7:
		var aresta1, aresta2, aresta3 float32
		fmt.Print("Digite o valor da aresta1 do paralelepípedo: ")
		fmt.Scanln(&aresta1)
		fmt.Print("Digite o valor da aresta2 do paralelepípedo: ")
		fmt.Scanln(&aresta2)
		fmt.Print("Digite o valor da aresta3 do paralelepípedo: ")
		fmt.Scanln(&aresta3)
		area, volume := paralelepipedo(aresta1, aresta2, aresta3)
		fmt.Println("A área do Paralelepípedo é ", area)
		fmt.Println("O volume do Paralelepîpedo é ", volume)
	case 8:
		var raio float32
		fmt.Print("Digite o raio da esfera: ")
		fmt.Scanln(&raio)
		area, volume := esfera(raio)
		fmt.Println("A área da Esfera é ", area)
		fmt.Println("O volume da Esfera é ", volume)
	case 0:
		fmt.Println("Finalizando programa!")
		break
	default:
		fmt.Println("Digite uma opção válida")
	}

}
