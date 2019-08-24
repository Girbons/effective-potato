package device

import (
	"github.com/stianeikeland/go-rpio"
)

func On(pin int) {
	p, status := pinStatus(pin)

	if status != rpio.High {
		p.High()
	}

	rpio.Close()
}

func Off(pin int) {
	p, status := pinStatus(pin)

	if status != rpio.Low {
		p.Low()
	}

	rpio.Close()
}

func pinStatus(pin int) (rpio.Pin, rpio.State) {
	p := rpio.Pin(pin)

	return p, p.Read()
}

func Status(pin int) string {
	_, status := pinStatus(pin)

	return string(status)
}
