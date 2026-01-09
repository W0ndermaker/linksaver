package events

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type Processor interface {
	Process(e Event) error
}

type Type int // кастомный тип

const (
	Unknown Type = iota
	Message
)

type Event struct {
	Type Type
	Text string
}
