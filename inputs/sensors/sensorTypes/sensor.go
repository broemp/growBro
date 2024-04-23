package sensortypes

type SensorInterface interface {
	Read() (string, error)
}

const (
	USB ConnectionType = iota
	GPIO
	Http
)
