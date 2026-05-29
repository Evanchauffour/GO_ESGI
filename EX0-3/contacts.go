package main

import (
	"fmt"
)

type Personne struct {
	Prenom string
	Nom    string
	Age    int
	Email  string
}

func (p Personne) NomComplet() string {
	return fmt.Sprintf("%s %s", p.Prenom, p.Nom)
}

func (p Personne) Presentation() string {
	return fmt.Sprintf("Je suis %s et j'ai %d ans (email: %s)", p.NomComplet(), p.Age, p.Email)
}

type Adresse struct {
	Rue        string
	Ville      string
	CodePostal string
}

func (a Adresse) Format() string {
	return fmt.Sprintf("%s, %s %s", a.Rue, a.CodePostal, a.Ville)
}

type Employe struct {
	Personne
	Adresse
	Poste   string
	Salaire float64
}

func (e Employe) FicheEmploye() string {
	return fmt.Sprintf(
		"Employe: %s\nAge: %d\nEmail: %s\nAdresse: %s\nPoste: %s\nSalaire: %.2f",
		e.NomComplet(),
		e.Age,
		e.Email,
		e.Adresse.Format(),
		e.Poste,
		e.Salaire,
	)
}

func (e *Employe) AugmenterSalaire(pct float64) {
	e.Salaire = e.Salaire * (1 + pct/100)
}

type Etudiant struct {
	Personne
	Promo   string
	Moyenne float64
}

func (e Etudiant) MentionObtenue() string {
	switch {
	case e.Moyenne >= 16:
		return "TB"
	case e.Moyenne >= 14:
		return "B"
	case e.Moyenne >= 12:
		return "AB"
	default:
		return "P"
	}
}

func main() {
	employe1 := Employe{
		Personne: Personne{
			Nom:    "Doe",
			Prenom: "John",
			Age:    30,
			Email:  "john.doe@example.com",
		},
		Adresse: Adresse{
			Rue:        "10 rue de Paris",
			Ville:      "Paris",
			CodePostal: "75001",
		},
		Poste:   "Développeur Go",
		Salaire: 42000,
	}

	employe2 := Employe{
		Personne: Personne{
			Nom:    "Doe",
			Prenom: "Jane",
			Age:    25,
			Email:  "jane.doe@example.com",
		},
		Adresse: Adresse{
			Rue:        "3 avenue Victor Hugo",
			Ville:      "Lyon",
			CodePostal: "69002",
		},
		Poste:   "Cheffe de projet",
		Salaire: 48000,
	}

	etudiant1 := Etudiant{
		Personne: Personne{
			Nom:    "Martin",
			Prenom: "Lucas",
			Age:    20,
			Email:  "lucas.martin@example.com",
		},
		Promo:   "M1",
		Moyenne: 15.2,
	}

	etudiant2 := Etudiant{
		Personne: Personne{
			Nom:    "Durand",
			Prenom: "Emma",
			Age:    21,
			Email:  "emma.durand@example.com",
		},
		Promo:   "M1",
		Moyenne: 12.4,
	}

	employe1.AugmenterSalaire(10)

	fmt.Println(employe1.FicheEmploye())
	fmt.Println()
	fmt.Println(employe2.FicheEmploye())
	fmt.Println()

	fmt.Printf("Etudiant: %s | Promo: %s | Moyenne: %.2f | Mention: %s\n",
		etudiant1.NomComplet(), etudiant1.Promo, etudiant1.Moyenne, etudiant1.MentionObtenue())
	fmt.Printf("Etudiant: %s | Promo: %s | Moyenne: %.2f | Mention: %s\n",
		etudiant2.NomComplet(), etudiant2.Promo, etudiant2.Moyenne, etudiant2.MentionObtenue())
}
