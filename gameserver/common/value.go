package common

// Value represents a constant value.
type Value float64

func NewValue(v float64) Value {
	return Value(v)
}

func (v Value) Add(other Value) Value {
	return Value(float64(v) + float64(other))
}

func (v Value) Mult(other Percentage) Value {
	return Value(float64(v) * (1 + float64(other)))
}

func (v Value) Sub(other Value) Value {
	return Value(float64(v) - float64(other))
}

func (v Value) Div(other Percentage) Value {
	return Value(float64(v) / (1 + float64(other)))
}

type PercentageUnit float64

type Percentage float64

func NewPercentage(v float64) Percentage {
	return Percentage(v)
}

func (p Percentage) Add(other PercentageUnit) Percentage {
	return Percentage(float64(p) + float64(other))
}

func (p Percentage) Sub(other PercentageUnit) Percentage {
	return Percentage(float64(p) - float64(other))
}

func (p Percentage) Mult(other Percentage) Percentage {
	return Percentage(1 - (1-float64(p))*(1-float64(other)))
}

type ModifierKind int

const (
	ModifierKindStatic ModifierKind = iota
	ModifierKindPercentage
)

type Modifier struct {
	Kind  ModifierKind
	Value float64
}

func (m Modifier) Apply(v Value) Value {
	switch m.Kind {
	case ModifierKindStatic:
		return v.Add(Value(m.Value))
		// case ModifierKindPercentage:
		// 	return v.Mult(Value(m.Value))
	}
	panic("unknown modifier kind")
}
