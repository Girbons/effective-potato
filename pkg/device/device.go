package device

import (
	"github.com/stianeikeland/go-rpio"
)

func On(pin int) {
	p, status := pinStatus(pin)

	if status != rpio.High {
		p.High()
	}
}

func Off(pin int) {
	p, status := pinStatus(pin)

	if status != rpio.Low {
		p.Low()
	}
}

func pinStatus(pin int) (rpio.Pin, rpio.State) {
	p := rpio.Pin(pin)

	return p, p.Read()
}

func Status(pin int) string {
	var humanStatus string
	_, status := pinStatus(pin)

	switch status {
	case rpio.Low:
		humanStatus = "Low"
	case rpio.High:
		humanStatus = "High"
	default:
		humanStatus = "Unkwown"
	}

	return humanStatus
}
