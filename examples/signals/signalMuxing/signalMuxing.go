package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/ignite-laboratories/JanOS"
	"github.com/ignite-laboratories/JanOS/Formula"
	"github.com/ignite-laboratories/JanOS/Symbols"
	"log"
	"time"
)

// ---------------------------------------------------------------------------------------------------------
// This is an example that sets up two sine waves and muxes them together into a third signal
// ---------------------------------------------------------------------------------------------------------

// The reference that ebiten drives
type window int

var amplitude = JanOS.Universe.Signals.NewSignalWithValue("Amplitude", Symbols.Alpha, 100, nil)
var frequency = JanOS.Universe.Signals.NewSignalWithValue("Frequency", Symbols.Omega, 0.25, nil)
var signalA = JanOS.Universe.Signals.NewSignal("Signal A", Symbols.Theta, nil)
var signalB = JanOS.Universe.Signals.NewSignal("Signal B", Symbols.Psi, nil)
var signalC *JanOS.Signal

func main() {
	signalA.SineWave(amplitude, frequency)
	signalB.SineWave(amplitude, frequency)

	signalC = JanOS.Universe.Signals.Mux("Signal C", Symbols.Lambda, Formula.Additive, signalA, signalB)

	// Launch ebiten...
	var w *window
	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("Spark")
	if err := ebiten.RunGame(w); err != nil {
		log.Panic(err)
	}
}

func (w *window) Update() error {
	return nil
}

func (w *window) Draw(screen *ebiten.Image) {
	now := time.Now()
	// This is used in lieu of a newline in graphics-land
	offset := 15

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", signalA.Name, signalA.GetInstantValue(now).Point), 0, 0*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", signalB.Name, signalB.GetInstantValue(now).Point), 0, 1*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", signalC.Name, signalC.GetInstantValue(now).Point), 0, 2*offset)
}

// Layout is required by ebiten
func (w *window) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
