package main

import (
	"fmt"
)

func main() {
	fmt.Println("Initialisiere Heilungs-Protokoll...")

	var heilungsKristall *int

	if heilungsKristall != nil {
		fmt.Printf("Energie-Level des Kristalls: %d\n", *heilungsKristall)
	} else {
		fmt.Println("⚠️ Warnung: Kristall ist noch nicht aufgeladen (nil). Überspringe Messung.")
	}

	energieFluss := 40
	istHeilungErfolgreich := false

	if energieFluss >= 50 {
		istHeilungErfolgreich = true
	}

	if istHeilungErfolgreich {
		fmt.Println("✅ Die Aura breitet sich aus! Die Highlands erblühen.")
	} else {
		fmt.Println("❌ Die Energie reicht nicht aus (Fluss unter 50). Die Highlands bleiben grau.")
	}

	baeumeGepflanzt := 0
	fmt.Println("Beginne Aufforstung...")

	for baeumeGepflanzt < 5 {
		baeumeGepflanzt++
		fmt.Printf("Baum Nr. %d gepflanzt.\n", baeumeGepflanzt)
	}

	fmt.Println("Mission abgeschlossen: Der Weltencode ist rein!")
}
