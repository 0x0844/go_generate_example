package types


type Event int


const (
	Add Event = iota
	Sub
	Update
	Delete
)

//go:generate stringer -type=Event
