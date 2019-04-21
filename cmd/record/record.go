package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gordonklaus/portaudio"
	wave "github.com/zenwerk/go-wave"
	"github.com/zjinys/serpent/utils"
)

func main() {
	audioFileName := "record.wav"

	waveFile, err := os.Create(audioFileName)

	utils.Chk(err)

	inputChannels := 1
	outputChannels := 0
	sampleRate := 44100
	framesPerBuffer := make([]byte, 64)

	portaudio.Initialize()
	defer portaudio.Terminate()
	stream, err := portaudio.OpenDefaultStream(inputChannels, outputChannels, float64(sampleRate), len(framesPerBuffer), framesPerBuffer)
	defer stream.Close()
	utils.Chk(err)

	param := wave.WriterParam{
		Out:           waveFile,
		Channel:       inputChannels,
		SampleRate:    sampleRate,
		BitsPerSample: 8, // if 16, change to WriteSample16()
	}

	waveWriter, err := wave.NewWriter(param)
	defer waveWriter.Close()
	utils.Chk(err)

	utils.Chk(stream.Start())

	start := time.Now().UnixNano()
	fmt.Printf("START: %d\n", start)
	for {
		utils.Chk(stream.Read())
		_, err := waveWriter.Write([]byte(framesPerBuffer)) // WriteSample16 for 16 bits
		utils.Chk(err)

		now := time.Now().UnixNano()

		//log.Printf()
		if now-start > 1020000000 {
			break
		}
	}

	utils.Chk(stream.Stop())
}
