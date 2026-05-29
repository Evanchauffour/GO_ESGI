package main

import (
	"fmt"
)

var weight float64
var height float64

const IMCMaigreur = 18.5
const IMCNormal = 25.0
const IMCSurPoids = 30.0

func main() {
	calculateIMC()
}

func calculateIMC() {
	fmt.Print("Poids (kg) : ")
	fmt.Scan(&weight)
	fmt.Print("Taille (m) : ")
	fmt.Scan(&height)

	if height <= 0 {
		fmt.Println("La taille doit être > 0")
		return
	}

	imc := weight / (height * height)
	fmt.Printf("Votre IMC est de %.2f\n", imc)

	if imc < IMCMaigreur {
		fmt.Println("Vous êtes en sous-poids")
	} else if imc < IMCNormal {
		fmt.Println("Vous êtes en poids normal")
	} else if imc < IMCSurPoids {
		fmt.Println("Vous êtes en surpoids")
	} else if imc >= IMCSurPoids {
		fmt.Println("Vous êtes en obésité")
	}
}
