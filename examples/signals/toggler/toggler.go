package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/ignite-laboratories/JanOS"
	"github.com/ignite-laboratories/JanOS/Symbols"
	"log"
	"time"
)

// ---------------------------------------------------------------------------------------------------------
// This is an example that sets up two sine waves and muxes them together into a third signal
// ---------------------------------------------------------------------------------------------------------

// The reference that ebiten drives
type window int

var signalA = JanOS.Universe.Signals.NewToggleSignal("Signal A", Symbols.Theta, 1)

func main() {
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

	ebitenutil.DebugPrint(screen, fmt.Sprintf("%s - %f", signalA.Name, signalA.GetInstantValue(now).Point))
}

// Layout is required by ebiten
func (w *window) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
