package metrics

// Type data type for metric const type
type Type string

const (
	TypeCounter       Type = "counter"
	TypeUpDownCounter Type = "up_down_counter"
	TypeHistogram     Type = "histogram"
)

const (
	NameSeparator = "_"
)
