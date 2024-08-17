package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/ignite-laboratories/JanOS"
	"github.com/ignite-laboratories/JanOS/Symbols"
	"log"
	"time"
)

// ---------------------------------------------------------------------------------------------------------
// This is an example of leveraging signals by themselves.  JanOS does not require setting up the universe
// unless you want to utilize its features.  In this example, we create 15 signals - seven oscillators,
// seven frequencies, and a singular amplitude.  The signals can be adjusted at runtime using 'A' and 'S'
// on your keyboard, which multiplies all the frequencies by 0.9 or 1.1, respectively.  The oscillating
// signals are then queried whenever we perform a draw call to grab the currently relevant timeline
// information, relative to time.Now, and output their incoming data points.

// Each signal has its own loop to maintain itself while remaining thread-safe for querying.

// The output window utilized in this example is using ebiten, but you can bring your own visualization engine.
// ---------------------------------------------------------------------------------------------------------

// The reference that ebiten drives
type window int

// Amplitude signal
var alpha = JanOS.Universe.Signals.NewSignalWithValue("Alpha", Symbols.Alpha, 100)

// Theta's frequency signal - Omega
var omega = JanOS.Universe.Signals.NewSignalWithValue("Omega", Symbols.Omega, 1)

// Oscillator Theta
var theta = JanOS.Universe.Signals.NewSignal("Theta", Symbols.Theta)

// Sigma's frequency signal - Mu
var mu = JanOS.Universe.Signals.NewSignalWithValue("Mu", Symbols.Mu, 1.1)

// Oscillator Sigma
var sigma = JanOS.Universe.Signals.NewSignal("Sigma", Symbols.Sigma)

// Tau's frequency signal - Nu
var nu = JanOS.Universe.Signals.NewSignalWithValue("Nu", Symbols.Nu, 1.2)

// Oscillator Tau
var tau = JanOS.Universe.Signals.NewSignal("Tau", Symbols.Tau)

// Upsilon's frequency signal - Xi
var xi = JanOS.Universe.Signals.NewSignalWithValue("Xi", Symbols.Xi, 1.3)

// Oscillator Upsilon
var upsilon = JanOS.Universe.Signals.NewSignal("Upsilon", Symbols.Upsilon)

// Phi's frequency signal - Omicron
var omicron = JanOS.Universe.Signals.NewSignalWithValue("Omicron", Symbols.Omicron, 1.4)

// Oscillator Phi
var phi = JanOS.Universe.Signals.NewSignal("Phi", Symbols.Phi)

// Chi's frequency signal - Lambda
var lambda = JanOS.Universe.Signals.NewSignalWithValue("Lambda", Symbols.Lambda, 1.5)

// Oscillator Chi
var chi = JanOS.Universe.Signals.NewSignal("Chi", Symbols.Chi)

// Psi's frequency signal - Rho
var rho = JanOS.Universe.Signals.NewSignalWithValue("Rho", Symbols.Rho, 1.6)

// Oscillator Psi
var psi = JanOS.Universe.Signals.NewSignal("Psi", Symbols.Psi)

func main() {
	// Start the oscillating loops on the appropriate signals
	theta.SineWave(alpha, omega)
	sigma.SineWave(alpha, mu)
	tau.SineWave(alpha, nu)
	upsilon.SineWave(alpha, xi)
	phi.SineWave(alpha, omicron)
	chi.SineWave(alpha, lambda)
	psi.SineWave(alpha, rho)

	theta.Timeline.SetResolution(500)

	// Launch ebiten...
	var w *window
	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("Spark")
	if err := ebiten.RunGame(w); err != nil {
		log.Panic(err)
	}
}

func (w *window) Update() error {
	now := time.Now()

	// In the logic loop, we just check for key input
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		// On 'A' we slow the oscillators
		omega.SetValue(now, omega.GetInstantValue(now).Value*0.9)
		mu.SetValue(now, mu.GetInstantValue(now).Value*0.9)
		nu.SetValue(now, nu.GetInstantValue(now).Value*0.9)
		xi.SetValue(now, xi.GetInstantValue(now).Value*0.9)
		omicron.SetValue(now, omicron.GetInstantValue(now).Value*0.9)
		lambda.SetValue(now, lambda.GetInstantValue(now).Value*0.9)
		rho.SetValue(now, rho.GetInstantValue(now).Value*0.9)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		// On 'S' we speed them up
		omega.SetValue(now, omega.GetInstantValue(now).Value*1.1)
		mu.SetValue(now, mu.GetInstantValue(now).Value*1.1)
		nu.SetValue(now, nu.GetInstantValue(now).Value*1.1)
		xi.SetValue(now, xi.GetInstantValue(now).Value*1.1)
		omicron.SetValue(now, omicron.GetInstantValue(now).Value*1.1)
		lambda.SetValue(now, lambda.GetInstantValue(now).Value*1.1)
		rho.SetValue(now, rho.GetInstantValue(now).Value*1.1)
	}

	return nil
}

func (w *window) Draw(screen *ebiten.Image) {
	now := time.Now()
	// This is used in lieu of a newline in graphics-land
	offset := 15

	// In the drawing loop, we just debug print out the signalal values

	// For the point signals, we just always print their values
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", alpha.Name, alpha.GetInstantValue(now)), 0, 0*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", omega.Name, omega.GetInstantValue(now)), 0, 1*offset)
	// Theta is an oscillator, but can be observed as a point
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", theta.Name, theta.GetInstantValue(now)), 0, 2*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", mu.Name, mu.GetInstantValue(now)), 0, 3*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", nu.Name, nu.GetInstantValue(now)), 0, 4*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", xi.Name, xi.GetInstantValue(now)), 0, 5*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", omicron.Name, omicron.GetInstantValue(now)), 0, 6*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", lambda.Name, lambda.GetInstantValue(now)), 0, 7*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", rho.Name, rho.GetInstantValue(now)), 0, 8*offset)

	// For the oscillating signals, we capture 10 indices from the past and output their values
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", theta.Name, theta.Timeline.SlicePastIndices(now, 10).Data), 0, 9*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", sigma.Name, sigma.Timeline.SlicePastIndices(now, 10).Data), 0, 10*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", tau.Name, tau.Timeline.SlicePastIndices(now, 10).Data), 0, 11*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", upsilon.Name, upsilon.Timeline.SlicePastIndices(now, 10).Data), 0, 12*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", phi.Name, phi.Timeline.SlicePastIndices(now, 10).Data), 0, 13*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", chi.Name, chi.Timeline.SlicePastIndices(now, 10).Data), 0, 14*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", psi.Name, psi.Timeline.SlicePastIndices(now, 10).Data), 0, 15*offset)
}

// Layout is required by ebiten
func (w *window) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
