package main

import "fmt"

func main() {

	/*
	 * Fragment 1: Die kontrollierte Bewässerung
	 * Problem: wasser wurde nie erhöht.
	 * Lösung: wasser++ einbauen.
	 */
	wasser := 0
	limit := 10

	fmt.Println("--- Start: Bewässerung ---")
	for wasser < limit {
		wasser++
		fmt.Println("Gieße magisches Wasser...")
		fmt.Println("Wasserstand:", wasser, "/", limit)
	}
	fmt.Println("Bewässerung abgeschlossen!\n")

	/*
	 * Fragment 2: Die korrekte Licht-Zufuhr
	 * Problem: lichtEnergie wurde erhöht statt verringert.
	 * Lösung: lichtEnergie-- verwenden.
	 */
	lichtEnergie := 100

	fmt.Println("--- Start: Licht-Zufuhr ---")
	for lichtEnergie > 0 {
		fmt.Println("Lichtstrahl wird fokussiert...")
		fmt.Println("Lichtenergie:", lichtEnergie)
		lichtEnergie--
	}
	fmt.Println("Lichtenergie verbraucht!\n")

	/*
	 * Fragment 3: Der erfolgreiche Wächter-Check
	 * Problem: istGezähmt wurde nie auf true gesetzt.
	 * Lösung: nach einer bestimmten Anzahl Versuche true setzen.
	 */
	istGezähmt := false
	versuche := 0
	maxVersuche := 10

	fmt.Println("--- Start: Wächter-Check ---")
	for !istGezähmt {
		versuche++
		fmt.Printf("Versuch %d: Besänftige die Ranken...\n", versuche)

		if versuche >= maxVersuche {
			istGezähmt = true
		}
	}

	fmt.Println("Mission abgeschlossen: Die Ranken sind gezähmt!")
}
