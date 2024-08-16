package main

import (
	"github.com/ignite-laboratories/JanOS"
	"github.com/ignite-laboratories/arwen"
	"github.com/ignite-laboratories/arwen/AI_Music"
	"time"
)

var waveformSys = Arwen.NewWaveformSystem()
var aiMusicSys = AI_Music.NewAI_MusicSystem()

type window int

var ecsWorld = JanOS.NewECSWorld("Logic", waveformSys, aiMusicSys)

func main() {
	JanOS.Universe.Start(PreFlight, Loop, ecsWorld)
}

var performance AI_Music.Performance
var binaryData AI_Music.BinaryData

func PreFlight() {
	performance, _ = aiMusicSys.LookupPerformance(AI_Music.FamilyBrass, AI_Music.NameTrumpetInC, AI_Music.PitchA5, AI_Music.DynamicFortissimo)
	binaryData, _ = aiMusicSys.GetBinaryData(performance.Entity)
}

func Loop(delta time.Duration) {
}
