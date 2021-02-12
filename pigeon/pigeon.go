package pigeon

import (
	"math/rand"
	"time"
)

//Pigeon is a collection of the pigeon attributes
type Pigeon struct {
	Name            string
	Colour          string
	wingspan        int
	HomeLocation    string
	currentLocation string
	message         string
}

//this function is a method on a pointer receiver.
func (p *Pigeon) FlyToLocation(newLocation string) {
	r := rand.Intn(10)
	time.Sleep(time.Duration(r) * time.Microsecond)
	p.currentLocation = newLocation
}

func (p Pigeon) GetLocation() string {
	return p.currentLocation
}

func (p *Pigeon) ReturnHome() {
	r := rand.Intn(10)
	time.Sleep(time.Duration(r) * time.Microsecond)
	p.currentLocation = p.HomeLocation
}

func CreateNewPigeon(name, colour, homelocation string) Pigeon {
	rand.Seed(time.Now().Unix())
	pg := Pigeon{
		Name:            name,
		Colour:          colour,
		HomeLocation:    homelocation,
		currentLocation: homelocation,
		wingspan:        rand.Intn(10),
	}

	return pg
}
