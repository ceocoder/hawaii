package main

import (
	"testing"
	"time"
)

func TestCalculateTimeAtHalfWayPoint(t *testing.T) {
	layout := "15:04:05"
	startTime, _ := time.Parse(layout, "03:04:05")
	want, _ := time.Parse(layout, "05:34:58")
	halfWayTime, err := calculateTimeAtHalfwayPoint(2000, 450, 30, startTime)

	if err != nil {
		t.Fatal(err)
	} else if halfWayTime.Format(layout) != want.Format(layout) {
		t.Fatalf("wanted %v, got %v", want.Format(layout), halfWayTime.Format(layout))
	}
}
