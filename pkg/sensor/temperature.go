package sensor

import (
	"errors"
	"fmt"
	"strings"

	"github.com/d2r2/go-dht"
)

func dht11Read(pin, retryTimes int) (float32, float32, int, error) {
	return dht.ReadDHTxxWithRetry(dht.DHT11, 4, false, retryTimes)
}

func dht12Read(pin, retryTimes int) (float32, float32, int, error) {
	return dht.ReadDHTxxWithRetry(dht.DHT12, 4, false, retryTimes)
}

func dht22Read(pin, retryTimes int) (float32, float32, int, error) {
	return dht.ReadDHTxxWithRetry(dht.DHT22, pin, false, retryTimes)
}

func ReadTemperature(dhtType string, pin, retryTimes int) (float32, float32, int, error) {
	switch strings.ToLower(dhtType) {
	case "dht11":
		return dht11Read(pin, retryTimes)
	case "dht12":
		return dht12Read(pin, retryTimes)
	case "dht22":
		return dht22Read(pin, retryTimes)
	default:
		return 0.0, 0.0, 0, errors.New(fmt.Sprintf("Unsupported dhttype %s", dhtType))
	}
}
