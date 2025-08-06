package numbers

type Formatter interface {
	Format(num float64, opt ...Option) string
}
type Option struct {
	MinimumFractionDigits *int
	MaximumFractionDigits *int
	MinimumIntegerDigits  int
}

func MinimumFractionDigits(min int) Option {
	if min >= 1<<8 {
		min = (1 << 8) - 1
	}
	return Option{MinimumFractionDigits: &min}
}

func MaximumFractionDigits(max int) Option {
	if max >= 1<<15 {
		max = (1 << 15) - 1
	}
	return Option{MaximumFractionDigits: &max}
}

func MinimumIntegerDigits(n int) Option {
	return Option{MinimumIntegerDigits: n}
}
