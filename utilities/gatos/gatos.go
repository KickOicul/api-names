package gatos

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	randomGenerator          = rand.New(rand.NewSource(time.Now().UnixNano()))
	suggestedCatNamesByColor = make(map[string][]string)
	mutex                    sync.Mutex
	nombresPorColor          = map[string][]string{
		"naranja": {"Garfield", "Flame", "Rusty", "Cinnamon", "Mars"},
		"negro":   {"Shadow", "Onyx", "Panther", "Salem", "Midnight"},
	}
)

func GetCatNameByColor(color string) (string, error) {

	Aleatorio := randomGenerator.Intn(10)

	nombres, ok := nombresPorColor[color]
	if !ok {
		return "", fmt.Errorf("No hay nombres disponibles para el gato color (%v)", color)
	}

	// Incluye nombres sugeridos en la lista
	nombres = append(nombres, suggestedCatNamesByColor[color]...)

	Aleatorio = Aleatorio % len(nombres)

	// Devuelve el nombre aleatorio
	return nombres[Aleatorio], nil
}

func AddSuggestedCatName(color, suggestedName string) error {
	mutex.Lock()
	defer mutex.Unlock()

	// Verifica si el color está en el mapa nombresPorColor
	if _, ok := nombresPorColor[color]; !ok {
		return fmt.Errorf("El color (%v) no está permitido para sugerencias de nombres", color)
	}

	// Inicializa la lista de nombres sugeridos si aún no existe
	if suggestedCatNamesByColor[color] == nil {
		suggestedCatNamesByColor[color] = make([]string, 0)
	}

	// Agrega el nombre sugerido
	suggestedCatNamesByColor[color] = append(suggestedCatNamesByColor[color], suggestedName)

	return nil
}
