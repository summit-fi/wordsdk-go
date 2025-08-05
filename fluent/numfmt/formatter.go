package numfmt

type Formatter interface {
	Format(num float64, opt ...Option) string
}
type Option struct {
	MinimumFractionDigits int
	MaximumFractionDigits int
	MinimumIntegerDigits  int
}

func MinimumFractionDigits(n int) Option {
	return Option{MinimumFractionDigits: n}
}

func MaximumFractionDigits(n int) Option {
	return Option{MaximumFractionDigits: n}
}

func MinimumIntegerDigits(n int) Option {
	return Option{MinimumIntegerDigits: n}
}
