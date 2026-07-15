package buildinfo

// These values are replaced with release metadata at build time.
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// Info describes the application build.
type Info struct {
	Version string
	Commit  string
	Date    string
}

// Current returns the metadata compiled into the running binary.
func Current() Info {
	return Info{
		Version: version,
		Commit:  commit,
		Date:    date,
	}
}
