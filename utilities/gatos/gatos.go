package gatos

import (
	"math/rand"
	"sync"
	"time"
)

var (
	randomGenerator          = rand.New(rand.NewSource(time.Now().UnixNano()))
	suggestedCatNamesByColor = make(map[string][]string)
	mutex                    sync.Mutex
)

func GetCatNameByColor(color string) (string, error) {
	nombresPorColor := map[string][]string{
		"naranja": {"Garfield", "Flame", "Rusty", "Cinnamon", "Mars"},
		"negro":   {"Shadow", "Onyx", "Panther", "Salem", "Midnight"},
	}

	Aleatorio := randomGenerator.Intn(10)

	mutex.Lock()
	defer mutex.Unlock()

	nombres, ok := nombresPorColor[color]
	if !ok {
		nombres = make([]string, 0)
	}

	nombres = append(nombres, suggestedCatNamesByColor[color]...)

	Aleatorio = Aleatorio % len(nombres)

	// Devuelve el nombre aleatorio
	return nombres[Aleatorio], nil
}

func AddSuggestedCatName(color, suggestedName string) {
	mutex.Lock()
	defer mutex.Unlock()

	if suggestedCatNamesByColor[color] == nil {
		suggestedCatNamesByColor[color] = make([]string, 0)
	}

	// Agrega el nombre sugerido
	suggestedCatNamesByColor[color] = append(suggestedCatNamesByColor[color], suggestedName)
}
