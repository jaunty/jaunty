package version

var version string

// Version returns a version string set during compile.
func Version() string {
	if version == "" {
		return "(indev)"
	}
	return version
}
