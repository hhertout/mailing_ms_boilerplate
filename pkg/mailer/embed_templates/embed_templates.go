package embed_templates

import (
	"embed"
)

//go:embed defaults.html
var DefaultsTemplate embed.FS
