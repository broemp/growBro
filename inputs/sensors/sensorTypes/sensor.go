package sensortypes

type Sense interface{}

type ConnectionType int

const (
	USB ConnectionType = iota
	GPIO
	Http
)

type Sensor struct {
	Model      string
	Connection ConnectionType
	GPIOPin    int
	Path       string
}
