package framerr

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"

	"golang.org/x/xerrors"
)

// Frame represents a single error frame, including its message, and the stack
// frame location information, if available.
//
// If no stack frame information is available, Source will be nil. If the
// information is available, all fields on Source will be set.
type Frame struct {
	Text string

	Source *struct {
		Package  string
		Function string
		File     string
		Line     int
	}
}

// Extract the Frames that are contained in the error. Errors must implement the
// xerrors.Formatter interface to be able to extract error chains and stack
// frame source details.
func Extract(err error) []Frame {
	frames := []Frame{}
	if err == nil {
		return frames
	}

	// We can only access this information by printing it using the Formatter
	// interface. Iterate over it ourself, so we know for sure where each error
	// ends.
	for next := err; next != nil; {
		if f, ok := next.(xerrors.Formatter); ok {
			pr := &printRecorder{}
			next = f.FormatError(pr)
			frames = append(frames, pr.Parse())
		} else {
			frames = append(frames, Frame{Text: next.Error()})
			next = nil
		}
	}

	return frames
}

// printRecorder is an xerrors.Printer that always requests detail (the frame),
// and parses the provideed info to create a Frame
type printRecorder struct {
	text     bytes.Buffer
	detail   bytes.Buffer
	inDetail bool
}

func (p *printRecorder) Print(args ...interface{}) {
	buf := &p.text
	if p.inDetail {
		buf = &p.detail
	}

	fmt.Fprint(buf, args...)
}

func (p *printRecorder) Printf(format string, args ...interface{}) {
	buf := &p.text
	if p.inDetail {
		buf = &p.detail
	}

	fmt.Fprintf(buf, format, args...)
}

func (p *printRecorder) Detail() bool {
	// The standard formatter code expects Detail to be called right before
	// detail is printed, so we'll do the same
	p.inDetail = true
	return p.inDetail
}

var detailsRegexp = regexp.MustCompile(`(?m)^(.*)\.(.*)\n  (.*):([0-9])+$`)

func (p *printRecorder) Parse() Frame {
	var source *struct {
		Package  string
		Function string
		File     string
		Line     int
	}

	matches := detailsRegexp.FindStringSubmatch(p.detail.String())
	if len(matches) > 1 {
		line, err := strconv.Atoi(matches[4])
		if err == nil {
			source = &struct {
				Package  string
				Function string
				File     string
				Line     int
			}{
				Package:  matches[1],
				Function: matches[2],
				File:     matches[3],
				Line:     line,
			}
		}

	}

	return Frame{
		Text:   p.text.String(),
		Source: source,
	}
}
