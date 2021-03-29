package clockface

import (
	"math"
	"testing"
	"time"

	"github.com/brighteyed/learn-go-with-tests/math/util"
)

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{time: simpleTime(0, 0, 0), angle: 0},
		{time: simpleTime(0, 0, 15), angle: math.Pi / 2},
		{time: simpleTime(0, 0, 30), angle: math.Pi},
		{time: simpleTime(0, 0, 45), angle: 3 * math.Pi / 2},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := secondsInRadians(c.time)

			if !util.RoughlyEqualFloat64(got, c.angle, util.Float64EqualityThreshold) {
				t.Errorf("want angle %v but got %v", c.angle, got)
			}
		})
	}
}

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{time: simpleTime(0, 0, 5), angle: math.Pi / 360},
		{time: simpleTime(0, 15, 30), angle: math.Pi/2 + math.Pi/60},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := minutesInRadians(c.time)

			if !util.RoughlyEqualFloat64(got, c.angle, util.Float64EqualityThreshold) {
				t.Errorf("want angle %v but got %v", c.angle, got)
			}
		})
	}
}

func TestHoursInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{time: simpleTime(1, 10, 30), angle: math.Pi/6 + math.Pi*105/3600},
		{time: simpleTime(13, 30, 30), angle: math.Pi/6 + math.Pi*305/3600},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := hoursInRadians(c.time)

			if !util.RoughlyEqualFloat64(got, c.angle, util.Float64EqualityThreshold) {
				t.Errorf("want angle %v but got %v", c.angle, got)
			}
		})
	}
}

func TestHourHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{time: simpleTime(0, 0, 0), point: Point{X: 150, Y: 150 - 50}},
		{time: simpleTime(3, 0, 0), point: Point{X: 150 + 50, Y: 150}},
		{time: simpleTime(6, 0, 0), point: Point{X: 150, Y: 150 + 50}},
		{time: simpleTime(9, 0, 0), point: Point{X: 150 - 50, Y: 150}},
		{time: simpleTime(12, 0, 0), point: Point{X: 150, Y: 150 - 50}},
		{time: simpleTime(15, 0, 0), point: Point{X: 150 + 50, Y: 150}},
		{time: simpleTime(18, 0, 0), point: Point{X: 150, Y: 150 + 50}},
		{time: simpleTime(21, 0, 0), point: Point{X: 150 - 50, Y: 150}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := HourHandPoint(c.time)

			if !roughlyEqualPoints(got, c.point) {
				t.Errorf("want point %v but got %v", c.point, got)
			}
		})
	}
}

func TestMinuteHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{time: simpleTime(0, 0, 0), point: Point{X: 150, Y: 150 - 80}},
		{time: simpleTime(0, 15, 0), point: Point{X: 150 + 80, Y: 150}},
		{time: simpleTime(0, 30, 0), point: Point{X: 150, Y: 150 + 80}},
		{time: simpleTime(0, 45, 0), point: Point{X: 150 - 80, Y: 150}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := MinuteHandPoint(c.time)

			if !roughlyEqualPoints(got, c.point) {
				t.Errorf("want point %v but got %v", c.point, got)
			}
		})
	}
}

func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{time: simpleTime(0, 0, 0), point: Point{X: 150, Y: 150 - 90}},
		{time: simpleTime(0, 0, 15), point: Point{X: 150 + 90, Y: 150}},
		{time: simpleTime(0, 0, 30), point: Point{X: 150, Y: 150 + 90}},
		{time: simpleTime(0, 0, 45), point: Point{X: 150 - 90, Y: 150}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := SecondHandPoint(c.time)

			if !roughlyEqualPoints(got, c.point) {
				t.Errorf("want point %v but got %v", c.point, got)
			}
		})
	}
}

func roughlyEqualPoints(p1, p2 Point) bool {
	return util.RoughlyEqualFloat64(p1.X, p2.X, util.Float64EqualityThreshold) &&
		util.RoughlyEqualFloat64(p1.Y, p2.Y, util.Float64EqualityThreshold)
}

func simpleTime(h, m, s int) time.Time {
	return time.Date(1337, time.January, 1, h, m, s, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}
