package static

import (
	_ "embed"
)

//go:embed templates/defaults.html
var DefaultsTemplate string
