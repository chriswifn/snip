package snip

import _ "embed"

//go:embed text/en/get.md
var _get string

//go:embed text/en/init.md
var _init string

//go:embed text/en/snip.md
var _snip string
