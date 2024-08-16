package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/ignite-laboratories/JanOS"
	"github.com/ignite-laboratories/JanOS/Symbol"
	"log"
	"time"
)

// ---------------------------------------------------------------------------------------------------------
// This is an example of leveraging dimensions by themselves.  JanOS does not require setting up the universe
// unless you want to utilize its features.  In this example, we create 15 dimensions - seven oscillators,
// seven frequencies, and a singular amplitude.  The dimensions can be adjusted at runtime using 'A' and 'S'
// on your keyboard, which multiplies all the frequencies by 0.9 or 1.1, respectively.  The oscillating
// dimensions are then queried whenever we perform a draw call to grab the currently relevant timeline
// information, relative to time.Now, and output their incoming data points.

// Each dimension has its own loop to maintain itself while remaining thread-safe for querying.

// The output window utilized in this example is using ebiten, but you can bring your own visualization engine.
// ---------------------------------------------------------------------------------------------------------

// The reference that ebiten drives
type window int

// Our dimensions
var alpha *JanOS.Dimension
var omega *JanOS.Dimension
var theta *JanOS.Dimension
var sigma *JanOS.Dimension
var tau *JanOS.Dimension
var upsilon *JanOS.Dimension
var phi *JanOS.Dimension
var chi *JanOS.Dimension
var psi *JanOS.Dimension
var mu *JanOS.Dimension
var nu *JanOS.Dimension
var xi *JanOS.Dimension
var omicron *JanOS.Dimension
var lambda *JanOS.Dimension
var rho *JanOS.Dimension

func main() {
	// Set up the dimensions:

	// Amplitude dimension
	alpha = JanOS.Universe.Dimensions.NewDimension("Alpha", Symbol.Alpha, 100)

	// Theta's frequency dimension - Omega
	omega = JanOS.Universe.Dimensions.NewDimension("Omega", Symbol.Omega, 1)
	// Oscillator Theta
	theta = JanOS.Universe.Dimensions.NewOscillatingDimension("Theta", Symbol.Theta, alpha, omega)

	// Sigma's frequency dimension - Mu
	mu = JanOS.Universe.Dimensions.NewDimension("Mu", Symbol.Mu, 1.1)
	// Oscillator Sigma
	sigma = JanOS.Universe.Dimensions.NewOscillatingDimension("Sigma", Symbol.Sigma, alpha, mu)

	// Tau's frequency dimension - Nu
	nu = JanOS.Universe.Dimensions.NewDimension("Nu", Symbol.Nu, 1.2)
	// Oscillator Tau
	tau = JanOS.Universe.Dimensions.NewOscillatingDimension("Tau", Symbol.Tau, alpha, nu)

	// Upsilon's frequency dimension - Xi
	xi = JanOS.Universe.Dimensions.NewDimension("Xi", Symbol.Xi, 1.3)
	// Oscillator Upsilon
	upsilon = JanOS.Universe.Dimensions.NewOscillatingDimension("Upsilon", Symbol.Upsilon, alpha, xi)

	// Phi's frequency dimension - Omicron
	omicron = JanOS.Universe.Dimensions.NewDimension("Omicron", Symbol.Omicron, 1.4)
	// Oscillator Phi
	phi = JanOS.Universe.Dimensions.NewOscillatingDimension("Phi", Symbol.Phi, alpha, omicron)

	// Chi's frequency dimension - Pi
	lambda = JanOS.Universe.Dimensions.NewDimension("Lambda", Symbol.Lambda, 1.5)
	// Oscillator Chi
	chi = JanOS.Universe.Dimensions.NewOscillatingDimension("Chi", Symbol.Chi, alpha, lambda)

	// Psi's frequency dimension - Rho
	rho = JanOS.Universe.Dimensions.NewDimension("Rho", Symbol.Rho, 1.6)
	// Oscillator Psi
	psi = JanOS.Universe.Dimensions.NewOscillatingDimension("Psi", Symbol.Psi, alpha, rho)

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
		omega.SetValue(now, omega.GetValue(now)*0.9)
		mu.SetValue(now, mu.GetValue(now)*0.9)
		nu.SetValue(now, nu.GetValue(now)*0.9)
		xi.SetValue(now, xi.GetValue(now)*0.9)
		omicron.SetValue(now, omicron.GetValue(now)*0.9)
		lambda.SetValue(now, lambda.GetValue(now)*0.9)
		rho.SetValue(now, rho.GetValue(now)*0.9)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		// On 'S' we speed them up
		omega.SetValue(now, omega.GetValue(now)*1.1)
		mu.SetValue(now, mu.GetValue(now)*1.1)
		nu.SetValue(now, nu.GetValue(now)*1.1)
		xi.SetValue(now, xi.GetValue(now)*1.1)
		omicron.SetValue(now, omicron.GetValue(now)*1.1)
		lambda.SetValue(now, lambda.GetValue(now)*1.1)
		rho.SetValue(now, rho.GetValue(now)*1.1)
	}

	return nil
}

// Layout is required by ebiten
func (w *window) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (w *window) Draw(screen *ebiten.Image) {
	now := time.Now()
	// This is used in lieu of a newline in graphics-land
	offset := 15

	// In the drawing loop, we just debug print out the dimensional values

	// For the point dimensions, we just always print their values
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", alpha.Name, alpha.GetValue(now)), 0, 0*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", omega.Name, omega.GetValue(now)), 0, 1*offset)
	// Theta is an oscillator, but can be observed as a point
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", theta.Name, theta.GetValue(now)), 0, 2*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", mu.Name, mu.GetValue(now)), 0, 3*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", nu.Name, nu.GetValue(now)), 0, 4*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", xi.Name, xi.GetValue(now)), 0, 5*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", omicron.Name, omicron.GetValue(now)), 0, 6*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", lambda.Name, lambda.GetValue(now)), 0, 7*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", rho.Name, rho.GetValue(now)), 0, 8*offset)

	// For the oscillating dimensions, we capture 10 indices into the future and output their values
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", theta.Name, theta.Timeline.SliceFutureIndices(now, 10).Data), 0, 9*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", sigma.Name, sigma.Timeline.SliceFutureIndices(now, 10).Data), 0, 10*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", tau.Name, tau.Timeline.SliceFutureIndices(now, 10).Data), 0, 11*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", upsilon.Name, upsilon.Timeline.SliceFutureIndices(now, 10).Data), 0, 12*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", phi.Name, phi.Timeline.SliceFutureIndices(now, 10).Data), 0, 13*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", chi.Name, chi.Timeline.SliceFutureIndices(now, 10).Data), 0, 14*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", psi.Name, psi.Timeline.SliceFutureIndices(now, 10).Data), 0, 15*offset)
}
