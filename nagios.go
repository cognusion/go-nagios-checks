package gonagioschecks

import (
	"fmt"
	"github.com/spf13/cast"
	"os"
	"strings"
)

const OK int = 0
const WARNING int = 1
const CRITICAL int = 2
const UNKNOWN int = 3

type Nagios struct {
	Code    int
	Message string
	Metrics []string
}

// Escalate the Nagios status code, if a more severe one
// is passed
func (n *Nagios) EscalateIf(code int) {
	switch code {
	case OK:
		if n.Code == UNKNOWN {
			n.Code = code
		}
	case WARNING:
		if n.Code != CRITICAL {
			n.Code = code
		}
	case CRITICAL:
		n.Code = code
	}
}

// Prepend a message to the message
func (n *Nagios) PrependMessage(message string) {
	n.Message = message + n.Message
}

// Append a message to the message
func (n *Nagios) AddMessage(message string) {
	n.Message = n.Message + message
}

// Append a message to the message if the condition isn't empty
func (n *Nagios) AddMessageIf(message, cond string) {
	n.AddMessageIfBool(message, cond != "")
}

// Append a message to the message if the condition boolean isn't false
func (n *Nagios) AddMessageIfBool(message string, cond bool) {
	if cond {
		n.Message = n.Message + message
	}
}

// Add metrics to the output, from values
func (n *Nagios) AddMetricNumbers(name string, value, warn, crit, min, max interface{}) {
	n.AddMetrics(fmt.Sprintf("%s=%s;%s;%s;%s;%s", name,
		cast.ToString(value), cast.ToString(warn),
		cast.ToString(crit), cast.ToString(min), cast.ToString(max)))
}

// Add metrics to the output from a well-formed Nagios-compatible string
func (n *Nagios) AddMetrics(metrics string) {
	n.Metrics = append(n.Metrics, metrics)
}

// Return the full Nagios-compatible output
func (n *Nagios) FullMessage() (message string) {
	message = n.Message

	// if there are any metrics, make them
	if len(n.Metrics) > 0 {
		message = message + "| " + strings.Join(n.Metrics, " ")
	}

	return
}

// Correctly Nagios-compatibly exit by analyzing the status code, and displaying
// the message appropriately.
func (n *Nagios) Exit() {

	switch n.Code {
	case OK:
		ExitOk(n.FullMessage())
	case WARNING:
		ExitWarning(n.FullMessage())
	case CRITICAL:
		ExitCritical(n.FullMessage())
	case UNKNOWN:
		fallthrough
	default:
		ExitUnknown(n.FullMessage())
	}
}

// Nagios-compatible exit with the OK status
func ExitOk(message string) {
	fmt.Printf("OK: %s\n", message)
	os.Exit(OK)
}

// Nagios-compatible exit with the WARNING status
func ExitWarning(message string) {
	fmt.Printf("WARNING: %s\n", message)
	os.Exit(WARNING)
}

// Nagios-compatible exit with the CRITICAL status
func ExitCritical(message string) {
	fmt.Printf("CRITICAL: %s\n", message)
	os.Exit(CRITICAL)
}

// Nagios-compatible exit with the UNKNOWN status
func ExitUnknown(message string) {
	fmt.Printf("UNKNOWN: %s\n", message)
	os.Exit(UNKNOWN)
}
