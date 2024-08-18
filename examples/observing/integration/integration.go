package main

import (
	"github.com/ignite-laboratories/JanOS"
	"github.com/ignite-laboratories/JanOS/Symbols"
	"time"
)

// ---------------------------------------------------------------------------------------------------------
// This is an example that simply sets up a sine wave and then samples it.
// ---------------------------------------------------------------------------------------------------------

var amplitude = JanOS.Universe.Signals.NewSignalWithValue("Amplitude", Symbols.Alpha, 100)
var frequency = JanOS.Universe.Signals.NewSignalWithValue("Frequency", Symbols.Omega, 1)
var theta = JanOS.Universe.Signals.NewSignal("Theta", Symbols.Theta)

type observer struct{}

func main() {
	theta.SineWave(amplitude, frequency)
	theta.Sample(10, time.Duration(time.Second), &observer{})

	for {
		time.Sleep(5 * time.Second)
	}
}

func OnTrigger(signal *JanOS.Signal, instant time.Time, value float64) {
	JanOS.Universe.Printf(signal, "%s Integral: %v", signal.Symbol, value)
}

func (o *observer) GetName() string {
	return "Observer"
}
func (o *observer) OnSample(signal *JanOS.Signal, ts JanOS.TimeSlice) {
	JanOS.Universe.Printf(signal, "%s Integral: %v", signal.Symbol, ts.Integrate())
}
