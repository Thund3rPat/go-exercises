package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func playSong(f io.ReadCloser) {
	streamer, format, err := mp3.Decode(f)
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
	song := os.Args[1]
	f, err := os.Open(song)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Currently Playing: ", song)
	playSong(f)
	fmt.Println("Finished Playing")

}
