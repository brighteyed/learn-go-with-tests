package svg

import (
	"fmt"
	"io"
	"time"

	clockface "github.com/brighteyed/learn-go-with-tests/math"
)

// HourHand writes an SVG node for a hour hand
func HourHand(w io.Writer, t time.Time) {
	point := clockface.HourHandPoint(t)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#000;stroke-width:3px;"/>`,
		point.X, point.Y)
}

// MinuteHand writes an SVG node for a minute hand
func MinuteHand(w io.Writer, t time.Time) {
	point := clockface.MinuteHandPoint(t)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#000;stroke-width:3px;"/>`,
		point.X, point.Y)
}

// SecondHand writes an SVG node for a second hand
func SecondHand(w io.Writer, t time.Time) {
	point := clockface.SecondHandPoint(t)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`,
		point.X, point.Y)
}

// Write writes an SVG representation of an analogue clock for given time `t` to the writer `w`
func Write(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)

	HourHand(w, t)
	MinuteHand(w, t)
	SecondHand(w, t)

	io.WriteString(w, svgEnd)
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`
