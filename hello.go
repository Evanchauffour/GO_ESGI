package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// message et commentaire
	fmt.Println("Hello, World!")
	// calculer 64 exposant 8
	calculate(64, 8)
	// afficher l'âge
	timeNow(time.Date(2003, 05, 13, 0, 0, 0, 0, time.UTC))
}

func calculate(number int, power int) {
	fmt.Printf("Result: %.2f\n", math.Pow(float64(number), float64(power)))
}

func timeNow(birthDate time.Time) {
	//  Afficher le jour, l'heure et l'âge
	fmt.Println("Day: ", time.Now().Day())
	// Afficher l'heure
	fmt.Println("Hour: ", time.Now().Hour())
	// Calculer l'âge et le stocker dans une variable
	age := int(time.Since(birthDate).Hours() / (24 * 365))
	// Afficher l'âge
	fmt.Printf("J'ai %d ans\n", age)
}
