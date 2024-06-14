package bootstrapper

type Option int

const (
	WithUnknownOption Option = iota
	WithDebug
)

func HasOption(options []Option, option Option) bool {
	if len(options) == 0 {
		return false
	}

	for i := range options {
		if options[i] == option {
			return true
		}
	}

	return false
}
