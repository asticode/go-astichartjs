package astichartjs

// Chart types
const (
	ChartTypeLine = "line"
	ChartTypePie  = "pie"
)

// Chart axis positions
const (
	ChartAxisPositionsBottom = "bottom"
)

// Chart axis type
const (
	ChartAxisTypesLinear = "linear"
)

// Chart background colors
const (
	ChartBackgroundColorBlue   = "rgba(54, 162, 235, 0.2)"
	ChartBackgroundColorGreen  = "rgba(75, 192, 192, 0.2)"
	ChartBackgroundColorOrange = "rgba(255, 159, 64, 0.2)"
	ChartBackgroundColorPurple = "rgba(153, 102, 255, 0.2)"
	ChartBackgroundColorRed    = "rgba(255, 99, 132, 0.2)"
	ChartBackgroundColorYellow = "rgba(255, 206, 86, 0.2)"
)

// Chart border colors
const (
	ChartBorderColorBlue   = "rgba(54, 162, 235, 1)"
	ChartBorderColorGreen  = "rgba(75, 192, 192, 1)"
	ChartBorderColorOrange = "rgba(255, 159, 64, 1)"
	ChartBorderColorPurple = "rgba(153, 102, 255, 1)"
	ChartBorderColorRed    = "rgba(255,99,132,1)"
	ChartBorderColorYellow = "rgba(255, 206, 86, 1)"
)

// Chart represents a chart
type Chart struct {
	Data    *Data    `json:"data,omitempty"`
	Options *Options `json:"options,omitempty"`
	Type    string   `json:"type"`
}

// Data represents data
type Data struct {
	Datasets []Dataset `json:"datasets,omitempty"`
	Labels   []string  `json:"labels,omitempty"`
}

// Dataset represents a dataset
type Dataset struct {
	BackgroundColor interface{}   `json:"backgroundColor,omitempty"`
	BorderColor     interface{}   `json:"borderColor,omitempty"`
	Data            []interface{} `json:"data,omitempty"`
	Label           string        `json:"label,omitempty"`
}

// DataPoint represents a data point
type DataPoint struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Options represents options
type Options struct {
	Responsive *bool   `json:"responsive,omitempty"`
	Scales     *Scales `json:"scales,omitempty"`
	Title      *Title  `json:"title,omitempty"`
}

// Scales represents scales options
type Scales struct {
	XAxes []Axis `json:"xAxes,omitempty"`
	YAxes []Axis `json:"yAxes,omitempty"`
}

// Axis represents an axis options
type Axis struct {
	Position   string      `json:"position,omitempty,omitempty"`
	ScaleLabel *ScaleLabel `json:"scaleLabel,omitempty"`
	Type       string      `json:"type,omitempty,omitempty"`
}

// ScaleLabel represents a scale label options
type ScaleLabel struct {
	Display     *bool  `json:"display,omitempty"`
	LabelString string `json:"labelString,omitempty"`
}

// Title represents a title options
type Title struct {
	Display *bool  `json:"display,omitempty"`
	Text    string `json:"text,omitempty"`
}
