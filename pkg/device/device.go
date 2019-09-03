package device

import (
	"github.com/stianeikeland/go-rpio"
)

// On set raspberry PIN to `HIGH`
func On(pin int) {
	p, status := pinStatus(pin)

	if status != rpio.High {
		p.Output()
		p.High()
	}
}

// Off set raspberry PIN to `LOW`
func Off(pin int) {
	p, status := pinStatus(pin)

	if status != rpio.Low {
		p.Output()
		p.Low()
	}
}

func pinStatus(pin int) (rpio.Pin, rpio.State) {
	p := rpio.Pin(pin)

	p.Input()

	return p, p.Read()
}

// Status returns a Human Readble status
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
