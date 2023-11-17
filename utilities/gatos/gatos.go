package gatos

import (
	"fmt"
	"math/rand"
	"time"
)

var randomGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))

func GetCatNameByColor(color string) (string, error) {
	nombresPorColor := map[string][]string{
		"naranja": {"Garfield", "Flame", "Rusty", "Cinnamon", "Mars"},
		"negro":   {"Shadow", "Onyx", "Panther", "Salem", "Midnight"},
	}

	Aleatorio := randomGenerator.Intn(10)

	nombres, ok := nombresPorColor[fmt.Sprintf(color)]
	if !ok {
		return "", fmt.Errorf("No hay nombres disponibles para el gato color (%v)", color)
	}

	Aleatorio = Aleatorio % len(nombres)

	// Devuelve el nombre aleatorio
	return nombres[Aleatorio], nil
}
