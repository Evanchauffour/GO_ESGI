package main

import (
	"fmt"
	"math"
	"strings"
)

// 1. Interface Payeur
type Payeur interface {
	Payer(montant float64) (string, error)
}

// Indice sujet : vérification à la compilation
var _ Payeur = &CarteCredit{}
var _ Payeur = &PayPal{}
var _ Payeur = &Crypto{}

// 2. CarteCredit
type CarteCredit struct {
	Numero    string
	Titulaire string
	Solde     float64
}

// Payer déduit du solde et retourne "Transaction CB #XXXX confirmée".
// Erreur si solde insuffisant.
func (c *CarteCredit) Payer(montant float64) (string, error) {
	if montant > c.Solde {
		return "", fmt.Errorf("solde insuffisant")
	}
	c.Solde -= montant

	// Indice : strings.HasPrefix(cc.Numero, ...) — XXXX = 4 derniers chiffres
	xxxx := c.Numero[len(c.Numero)-4:]
	prefix := c.Numero[:len(c.Numero)-4]
	if !strings.HasPrefix(c.Numero, prefix) {
		return "", fmt.Errorf("numéro de carte invalide")
	}

	return fmt.Sprintf("Transaction CB #%s confirmée", xxxx), nil
}

// 3. PayPal
type PayPal struct {
	Email string
	Solde float64
}

// Payer déduit du solde et retourne "Paiement PayPal de X€ vers email".
func (p *PayPal) Payer(montant float64) (string, error) {
	if montant > p.Solde {
		return "", fmt.Errorf("solde insuffisant")
	}
	p.Solde -= montant
	return fmt.Sprintf("Paiement PayPal de %.2f€ vers %s", montant, p.Email), nil
}

// 4. Crypto
type Crypto struct {
	Adresse string
	Solde   float64
	Monnaie string
}

// Payer convertit le montant en crypto (1 BTC = 50000€) et retourne le détail.
func (c *Crypto) Payer(montant float64) (string, error) {
	// Indice : math.Round(montant/50000*1000)/1000
	quantite := math.Round(montant/50000*1000) / 1000

	if quantite > c.Solde {
		return "", fmt.Errorf("solde insuffisant")
	}
	c.Solde -= quantite

	return fmt.Sprintf("%.3f %s envoyés à %s (%.2f€)", quantite, c.Monnaie, c.Adresse, montant), nil
}

// 5. ProcesserPanier affiche le total, le mode (type switch), puis appelle Payer.
func ProcesserPanier(payeur Payeur, articles []float64) {
	total := 0.0
	for _, prix := range articles {
		total += prix
	}
	fmt.Printf("Total : %.2f€\n", total)

	// Type switch pour afficher le mode utilisé
	switch payeur.(type) {
	case *CarteCredit:
		fmt.Println("Mode utilisé : CarteCredit")
	case *PayPal:
		fmt.Println("Mode utilisé : PayPal")
	case *Crypto:
		fmt.Println("Mode utilisé : Crypto")
	}

	msg, err := payeur.Payer(total)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}
	fmt.Println(msg)
}

func main() {
	articles := []float64{29.99, 49.50, 15.00}

	fmt.Println("--- CarteCredit ---")
	cb := &CarteCredit{Numero: "4532123456784242", Titulaire: "Alice", Solde: 200}
	ProcesserPanier(cb, articles)

	fmt.Println("\n--- PayPal ---")
	pp := &PayPal{Email: "bob@mail.com", Solde: 150}
	ProcesserPanier(pp, articles)

	fmt.Println("\n--- Crypto ---")
	crypto := &Crypto{Adresse: "1ABC...", Solde: 0.01, Monnaie: "BTC"}
	ProcesserPanier(crypto, articles)
}
