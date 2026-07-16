package render

const (
	ansiReset = "\x1b[0m"
	ansiBold  = "\x1b[1m"
	ansiDim   = "\x1b[2m"
	ansiCyan  = "\x1b[36m"
)

// Options controls whether output is machine-friendly or decorated for a terminal.
type Options struct {
	Plain bool
	Color bool
}

func styled(value, codes string, enabled bool) string {
	if !enabled {
		return value
	}
	return codes + value + ansiReset
}
