package main

import (
	"github.com/ignite-laboratories/JanOS"
	"github.com/ignite-laboratories/JanOS/Symbols"
	"log"
	"time"
)

// ---------------------------------------------------------------------------------------------------------
// This is an example that sets up a signal and observes it at a lower frequency and up-samples the result.
// ---------------------------------------------------------------------------------------------------------

type observer struct{}

func main() {
	JanOS.Universe.StdResolution = 60

	amplitude := JanOS.Universe.Signals.NewSignalWithValue("Amplitude", Symbols.Alpha, 100)
	frequency := JanOS.Universe.Signals.NewSignalWithValue("Frequency", Symbols.Omega, 1)
	theta := JanOS.Universe.Signals.NewSignal("Theta", Symbols.Theta)
	theta.SineWave(amplitude, frequency)

	theta.Sample(50, time.Duration(time.Second), &observer{})

	for {
		time.Sleep(time.Second)
	}
}

func (o *observer) GetName() string {
	return "Observer"
}

func (o *observer) OnSample(signal *JanOS.Signal, ts JanOS.TimeSlice) {
	result := ts.UpSample(500)
	log.Println(GetValues(ts))
	log.Println(GetValues(result))
}

func GetValues(ts JanOS.TimeSlice) []int {
	return JanOS.Select(ts.Data, func(val JanOS.PointValue) int {
		return int(val.Value)
	})
}
