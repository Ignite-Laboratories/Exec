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

var amplitude = JanOS.Universe.Signals.NewSignalWithValue("Amplitude", Symbols.Alpha, 100)
var frequency = JanOS.Universe.Signals.NewSignalWithValue("Frequency", Symbols.Omega, 0.25)
var signalA = JanOS.Universe.Signals.NewSignal("Signal A", Symbols.Theta)
var signalB = JanOS.Universe.Signals.NewSignal("Signal B", Symbols.Psi)
var signalC *JanOS.Signal

func main() {
	signalA.SineWave(amplitude, frequency)
	signalB.SineWave(amplitude, frequency)

	signalC = signalA.Mux("Signal C", Symbols.Lambda, Formula.Additive, signalB)

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

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", signalA.Name, signalA.GetValue(now)), 0, 0*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", signalB.Name, signalB.GetValue(now)), 0, 1*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", signalC.Name, signalC.GetValue(now)), 0, 2*offset)
}

// Layout is required by ebiten
func (w *window) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
