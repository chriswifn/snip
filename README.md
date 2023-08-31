# 🌳 Go Bonzai Snippet Branch

Snippet tool.

## Installation

This snip command can be installed as a standalone program or composed into a Bonzai command tree.

Standalone
```
go install github.com/chriswifn/snip/cmd/snip@latest
```

Composed

```go
package z

import (
    Z "github.com/rwxrob/bonzai/z"
    "github.com/chriswifn/snip"
)

var Cmd = &Z.Cmd{
    Name: `z`,
    Commands: []*Z.Cmd{help.Cmd, snip.Cmd},
}
```

## Tab Completion
To activate bash completion just use the `complete -C` option from your
`.bashrc` or command line. There is no messy sourcing required. All the
completion is done by the program itself.

```
complete -C snip snip
```

