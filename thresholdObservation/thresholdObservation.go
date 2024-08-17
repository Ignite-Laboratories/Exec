package main

import (
	"github.com/ignite-laboratories/JanOS"
	"github.com/ignite-laboratories/JanOS/Observers"
	"github.com/ignite-laboratories/JanOS/Symbols"
	"time"
)

// ---------------------------------------------------------------------------------------------------------
// This is an example that samples a sine wave and then fires a trigger function whenever the sine wave
// changes value at an extreme rate.
// ---------------------------------------------------------------------------------------------------------

func main() {
	JanOS.Universe.StdResolution = 60

	amplitude := JanOS.Universe.Signals.NewSignalWithValue("Amplitude", Symbols.Alpha, 100)
	frequency := JanOS.Universe.Signals.NewSignalWithValue("Frequency", Symbols.Omega, 1)
	theta := JanOS.Universe.Signals.NewSignal("Theta", Symbols.Theta)
	observer := Observers.NewThresholdObserver("Observer", 0.4, OnTrigger)
	theta.SineWave(amplitude, frequency)
	theta.Sample(10, time.Duration(time.Second), observer)

	for {
	}
}

func OnTrigger(observation JanOS.Observation) {
	JanOS.Universe.Printf(observation.Observer, "Found %d trigger points on %s", len(observation.Values), string(observation.Signal.Symbol))
}
