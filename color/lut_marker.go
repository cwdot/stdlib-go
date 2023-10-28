package color

type LutOpts struct {
	Lut          map[string]Color
	DefaultColor Color
}

func Lut(lut map[string]Color) func(opts *LutOpts) {
	return func(opts *LutOpts) {
		opts.Lut = lut
	}
}

func DefaultColor(defColor Color) func(opts *LutOpts) {
	return func(opts *LutOpts) {
		opts.DefaultColor = defColor
	}
}

func NewLutMarker(opts ...func(opt *LutOpts)) *LutMarker {
	lut := make(map[string]Color)
	var defaultColor Color
	o := &LutOpts{
		Lut:          lut,
		DefaultColor: defaultColor,
	}
	for _, opt := range opts {
		opt(o)
	}
	return &LutMarker{
		lut:          o.Lut,
		defaultColor: o.DefaultColor,
	}
}

type LutMarker struct {
	lut          map[string]Color
	defaultColor Color
}

func (m *LutMarker) Set(name string, c Color) {
	m.lut[name] = c
}

func (m *LutMarker) Mark(value string) string {
	if c, ok := m.lut[value]; ok {
		return c.It(value)
	}
	return m.defaultColor.It(value)
}
