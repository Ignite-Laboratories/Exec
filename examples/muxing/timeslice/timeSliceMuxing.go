package main

import (
	"github.com/ignite-laboratories/JanOS"
	"github.com/ignite-laboratories/JanOS/Formula"
	"github.com/ignite-laboratories/JanOS/Symbols"
	"log"
	"time"
)

// ---------------------------------------------------------------------------------------------------------
// This is an example that samples a sine wave and then fires a trigger function whenever the sine wave
// changes value at an extreme rate.
// ---------------------------------------------------------------------------------------------------------

type observer struct {
	Name string
}

func main() {
	JanOS.Universe.StdResolution = 60

	amplitude := JanOS.Universe.Signals.NewSignalWithValue("Amplitude", Symbols.Alpha, 100)
	frequency := JanOS.Universe.Signals.NewSignalWithValue("Frequency", Symbols.Omega, 1)
	theta := JanOS.Universe.Signals.NewSignal("Theta", Symbols.Theta)
	theta.SineWave(amplitude, frequency)

	o := &observer{Name: "Observer"}
	theta.Sample(10, time.Duration(time.Second), o)

	for {
		time.Sleep(time.Second)
	}
}

// GetNamedValue returns the assigned name to this instance.
func (o *observer) GetNamedValue() string {
	return o.Name
}

var lastSlice JanOS.TimeSlice

func (o *observer) OnSample(signal *JanOS.Signal, ts JanOS.TimeSlice) {
	log.Println(lastSlice.Data)
	log.Println(ts.Data)
	result := ts.Mux(Formula.Additive, lastSlice)
	log.Println(result.Data)
	log.Println()

	lastSlice = ts
}
