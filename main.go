package main

import (
	"flag"
	"fmt"
	"time"
)

const (
	// distance to reach max ground speed
	// http://aviation.stackexchange.com/a/14359
	// http://flightaware.com/live/flight/QFA18/history/20150424/0655Z/KLAX/YSSY/tracklog
	ascentDescentDistance        = 100 // in nm - source, the internet
	ascentDescentSpeedMultiplier = 0.64
)

func validateParams(distance, airspeed, headwinds float64, startTimeInHST time.Time) {
	if headwinds > airspeed {
		fmt.Errorf("plane is moving backwards if trueairspeed is %v and headwinds are %v", airspeed, headwinds)
	}
	if distance <= 0 {
		fmt.Errorf("already in hawai'i, Mahalo!")
	}

}

func calculateTimeAtHalfwayPoint(distance, airspeed, headwinds float64, startTimeInHST time.Time) (time.Time, error) {
	trueGroundSpeed := airspeed - headwinds

	// assuming that ascent and descent are about 50nm each
	distanceCoveredAtMaxGroundSpeed := distance - ascentDescentDistance*2

	// assuming that speed descreases linearly during ascent and descent from 0 to trueGroundSpeed and trueGroundSpeed to 0
	ascentDescentGroundSpeed := trueGroundSpeed * ascentDescentSpeedMultiplier

	durationAscentDescent := (ascentDescentDistance * 2) / ascentDescentGroundSpeed
	durationCruising := distanceCoveredAtMaxGroundSpeed / trueGroundSpeed

	// duration to halfway point is half of ascent + descent + cruising
	durationToHalfWayPoint, err := time.ParseDuration(fmt.Sprintf("%fh", (durationAscentDescent+durationCruising)/2))
	if err != nil {
		return startTimeInHST, err
	}

	return startTimeInHST.Add(durationToHalfWayPoint), nil
}

func main() {
	distance := flag.Float64("distance", 0, "distance in nautical miles")
	airSpeed := flag.Float64("airspeed", 0, "airspeed in nautical miles")
	headWinds := flag.Float64("headwinds", 0, "headwinds in nautical miles")
	startTimeStr := flag.String("startTime", "", "Start time in HH:MM:SS")
	flag.Parse()

	startTime, err := time.Parse("15:04:05", *startTimeStr)
	if err != nil {
		flag.Usage()
		return
	}
	validateParams(*distance, *airSpeed, *headWinds, startTime)

	half, err := calculateTimeAtHalfwayPoint(*distance, *airSpeed, *headWinds, startTime)

	if err != nil {
		flag.Usage()
		fmt.Errorf("failed to run %v", err)
	} else {
		fmt.Printf("Approximate time at halfway point: %s\n", half.Format("15:04:05"))
	}
}
