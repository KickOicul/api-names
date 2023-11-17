package barcos

import (
	"fmt"
	"math/rand"
	"time"
)

var randomGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))

var (
	adjectives = []string{"cursed", "Majestic", "Sapphire", "Golden", "Daring", "Brilliant", "Eternal", "Radiant", "Swift", "Infinite", "Vibrant"}
	nouns      = []string{"Voyager", "Phoenix", "Harmony", "Aurora", "Serenity", "Horizon", "Odyssey", "Trident", "Eclipse", "Cascade", "SeaDogs"}
)

func GeneraNombreBarcoAleatorio() string {

	adjective := adjectives[randomGenerator.Intn(len(adjectives))]
	noun := nouns[randomGenerator.Intn(len(nouns))]

	shipName := fmt.Sprintf("%v %v", adjective, noun)

	//devuelve nombre barco
	return shipName
}
