package timestamp

var timeOfDayString = [...]string{
	"Night",
	"Morning",
	"Afternoon",
	"Evening",
}

type TimeOfDay int

const (
	// Night is the time between 00:00 and 05:59
	Night TimeOfDay = 0
	// Morning is the time between 06:00 and 11:59
	Morning TimeOfDay = 1
	// Afternoon is the time between 12:00 and 17:59
	Afternoon TimeOfDay = 2
	// Evening is the time between 18:00 and 23:59
	Evening TimeOfDay = 3
)

func (t TimeOfDay) String() string {
	return timeOfDayString[t]
}
