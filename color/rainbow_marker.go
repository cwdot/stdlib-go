package color

type RainbowOpts struct {
	Colors []Color
}

func Colors(newColors []Color) func(opts *RainbowOpts) {
	return func(opts *RainbowOpts) {
		opts.Colors = newColors
	}
}

func NewRainbowMarker(opts ...func(opt *RainbowOpts)) *RainbowMarker {
	colors := []Color{
		Red,
		Green,
		Yellow,
		Blue,
		Magenta,
		Cyan,
	}
	o := &RainbowOpts{Colors: colors}
	for _, opt := range opts {
		opt(o)
	}
	return &RainbowMarker{
		colors: o.Colors,
	}
}

type RainbowMarker struct {
	colors   []Color
	position int
}

func (m *RainbowMarker) Mark(value string) string {
	if m.position >= len(m.colors) {
		m.position = 0
	}
	c := m.colors[m.position]
	m.position++
	return c.It(value)
}
