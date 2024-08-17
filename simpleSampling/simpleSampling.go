package main

import (
	"github.com/ignite-laboratories/JanOS"
	"github.com/ignite-laboratories/JanOS/Observers"
	"github.com/ignite-laboratories/JanOS/Symbols"
	"time"
)

// ---------------------------------------------------------------------------------------------------------
// This is an example that simply sets up a sine wave and then samples it.
// ---------------------------------------------------------------------------------------------------------

var amplitude = JanOS.Universe.Signals.NewSignalWithValue("Amplitude", Symbols.Alpha, 100)
var frequency = JanOS.Universe.Signals.NewSignalWithValue("Frequency", Symbols.Omega, 1)
var theta = JanOS.Universe.Signals.NewSignal("Theta", Symbols.Theta)

func main() {
	JanOS.Universe.StdResolution = 60
	theta.SineWave(amplitude, frequency)
	theta.Sample(10, time.Duration(time.Second), &Observers.LoggingObserver{})

	for {
		time.Sleep(5 * time.Second)
	}
}
