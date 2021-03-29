package svg

import (
	"bytes"
	"encoding/xml"
	"math"
	"testing"
	"time"

	clockface "github.com/brighteyed/learn-go-with-tests/math"
	"github.com/brighteyed/learn-go-with-tests/math/util"
)

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

// Length returns the length of a line
func (l Line) Length() float64 {
	return math.Sqrt((l.X1-l.X2)*(l.X1-l.X2) + (l.Y1-l.Y2)*(l.Y1-l.Y2))
}

func TestSVGWriter(t *testing.T) {
	b := bytes.Buffer{}
	Write(&b, time.Date(1111, time.January, 1, 0, 0, 0, 0, time.UTC))

	svg := SVG{}
	xml.Unmarshal(b.Bytes(), &svg)

	if len(svg.Line) != 3 {
		t.Fatalf("must be 3 hands, but got %d", len(svg.Line))
	}

	t.Run("all hands should have the same center", func(t *testing.T) {
		if !(svg.Line[0].X1 == svg.Line[1].X1 && svg.Line[1].X1 == svg.Line[2].X1 &&
			svg.Line[0].Y1 == svg.Line[1].Y1 && svg.Line[1].Y1 == svg.Line[2].Y1) {
			t.Errorf("start points of hands %+v differ", svg.Line)
		}
	})

	t.Run("length of hands", func(t *testing.T) {
		for _, line := range svg.Line {
			length := line.Length()
			if !util.RoughlyEqualFloat64(length, clockface.HourHandLength, util.Float64EqualityThreshold) &&
				!util.RoughlyEqualFloat64(length, clockface.MinuteHandLength, util.Float64EqualityThreshold) &&
				!util.RoughlyEqualFloat64(length, clockface.SecondHandLength, util.Float64EqualityThreshold) {
				t.Errorf("hand length for %v is %f that is not expected", line, length)
			}
		}
	})
}
