package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
)

const (
	Duration   = 2
	SampleRate = 44100
	Frequency  = 440
)

var (
	tau = math.Pi * 2
)

func sineWaver() {
	var (
		start float64 = 1.0
		end   float64 = 1.0e-4
	)
	numero_de_samples := Duration * SampleRate
	var angle float64 = tau / float64(numero_de_samples)
	archivo_audio := "sine.bin"
	f, _ := os.Create(archivo_audio)
	decayfac := math.Pow(end/start, 1.0/float64(numero_de_samples))

	for i := 0; i < numero_de_samples; i++ {
		sample := math.Sin(angle * Frequency * float64(i))
		sample *= start
		start *= decayfac
		var buf [8]byte
		binary.LittleEndian.PutUint32(buf[:], math.Float32bits(float32(sample)))
		bw, err := f.Write(buf[:])
		if err != nil {
			panic(err)
		}
		fmt.Printf("bw %v, file %v", bw, archivo_audio)
	}
}

func main() {
	sineWaver()
}
