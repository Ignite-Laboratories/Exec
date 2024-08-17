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

// The reference that ebiten drives
type window int

var amplitude = JanOS.Universe.Signals.NewSignalWithValue("Amplitude", Symbols.Alpha, 100)
var frequency = JanOS.Universe.Signals.NewSignalWithValue("Frequency", Symbols.Omega, 1)
var theta = JanOS.Universe.Signals.NewSignal("Theta", Symbols.Theta)

func main() {
	theta.SineWave(amplitude, frequency)
	theta.Sample(10, time.Duration(time.Second*2), func(ts JanOS.TimeSlice) {
		JanOS.Universe.Printf(theta, "Observation: %v", ts.Data)
	})

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
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", theta.Name, theta.GetValue(now)), 0, 0)
}

// Layout is required by ebiten
func (w *window) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
