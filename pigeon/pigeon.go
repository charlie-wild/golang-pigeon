package pigeon

import "time"

//Pigeon is a collection of the pigeon attributes
type Pigeon struct {
	Name            string
	Colour          string
	Wingspan        int
	HomeLocation    string
	CurrentLocation string
}

//this function is a method on a pointer receiver.
func (p *Pigeon) flyToLocation(newLocation string) {
	time.Sleep(2)
	p.CurrentLocation = newLocation
}

func (p Pigeon) getLocation() string {
	return p.CurrentLocation
}

func (p *Pigeon) returnHome() {
	p.CurrentLocation = p.HomeLocation
}
