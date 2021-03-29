package clockface

import (
	"math"
	"time"
)

// Point represents a two dimensional Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

const (
	MinuteHandLength = 80
	SecondHandLength = 90
	HourHandLength   = 50

	clockCenterX = 150
	clockCenterY = 150
)

// HourHandPoint returns the end point of a hour hand of an analogue clock at time `t`
func HourHandPoint(t time.Time) Point {
	return handPoint(-hoursInRadians(t)+math.Pi/2, HourHandLength)
}

// MinuteHandPoint returns the end point of a minute hand of an analogue clock at time `t`
func MinuteHandPoint(t time.Time) Point {
	return handPoint(-minutesInRadians(t)+math.Pi/2, MinuteHandLength)
}

// SecondHandPoint returns the end point of a second hand of an analogue clock at time `t`
func SecondHandPoint(t time.Time) Point {
	return handPoint(-secondsInRadians(t)+math.Pi/2, SecondHandLength)
}

func handPoint(angle float64, handLength float64) Point {
	return Point{clockCenterX + handLength*math.Cos(angle),
		clockCenterY - handLength*math.Sin(angle)}
}

func hoursInRadians(t time.Time) float64 {
	h := t.Hour() % 12
	return math.Pi * (3600*float64(h) + 60*float64(t.Minute()) + float64(t.Second())) / (6 * 3600)
}

func minutesInRadians(t time.Time) float64 {
	return math.Pi * (60*float64(t.Minute()) + float64(t.Second())) / 1800
}

func secondsInRadians(t time.Time) float64 {
	return math.Pi * float64(t.Second()) / 30
}
