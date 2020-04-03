package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

func play(song string) {
	var streamer beep.StreamSeekCloser
	var format beep.Format
	var err error

	file, err := os.Open(song)
	filetype := path.Ext(song)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	switch filetype {
	case ".mp3":
		streamer, format, err = mp3.Decode(file)
	case ".wav":
		streamer, format, err = wav.Decode(file)
	default:
		log.Fatal("File not supported!")
	}
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No file specified")
	}

	var songs []string = os.Args[1:]

	for _, song := range songs {
		fmt.Println("Currently Playing: ", song)
		play(song)
		fmt.Println("Finished Playing")
	}
}
