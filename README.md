# go-nagios-checks
Go library to ease the writing a Nagios-compatible checks

[![GoDoc](https://godoc.org/github.com/cognusion/go-nagios-checks?status.svg)](https://godoc.org/github.com/cognusion/go-nagios-checks)

```bash
go get github.com/cognusion/go-nagios-checks
```

```go
package main

import (
	nagios "github.com/cognusion/go-nagios-checks"
	"time"
	"fmt"
)

func main() {
	// Make a new check object, and set the default state
	n := nagios.Nagios{Code: nagios.UNKNOWN}
	
	
	// fill this in.
	
	
	// Graceful exit, with the right code, 
	// the right message
	// and the right metrics
	n.Exit()
}
```