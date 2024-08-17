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
	integralObserver := Observers.NewIntegralObserver(OnTrigger)
	theta.Sample(10, time.Duration(time.Second), integralObserver)

	for {
		time.Sleep(5 * time.Second)
	}
}

func OnTrigger(signal *JanOS.Signal, instant time.Time, value float64) {
	JanOS.Universe.Printf(signal, "%s Integral: %v", signal.Symbol, value)
}
