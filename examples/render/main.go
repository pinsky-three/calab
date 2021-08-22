package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"time"

	"github.com/gen2brain/x264-go"
	lksdk "github.com/livekit/server-sdk-go"
	"github.com/minskylab/calab/experiments"
	"github.com/minskylab/calab/experiments/petridish"
	"github.com/minskylab/calab/spaces/board"
	"github.com/minskylab/calab/systems/lifelike"
	"github.com/pion/webrtc/v3"
	"github.com/pion/webrtc/v3/pkg/media"
)

// type SProvider struct{}

// func (p *SProvider) NextSample() (media.Sample, error) {
// 	media.Sample{}
// }

func basicLifeLike(w, h int, lifeRule *lifelike.Rule) *petridish.PetriDish {
	dynamic := lifelike.MustNew(lifeRule, lifelike.ToroidBounded, lifelike.MooreNeighborhood(1, false))
	space := board.MustNew(w, h).Fill(board.UniformNoise, dynamic)

	return petridish.NewFromSpaceAndDynamic(space, dynamic, petridish.WithTPSMonitor)
}

func main() {
	gameOfLife := basicLifeLike(256, 256, lifelike.GameOfLifeRule)

	experiment := experiments.New()

	experiment.AddPetriDish(gameOfLife)

	fmt.Printf("gameOfLife id: %s\n", gameOfLife.ID)

	frames, err := experiment.Observe(gameOfLife.ID)
	if err != nil {
		panic(err)
	}

	go gameOfLife.Run(30 * time.Minute)

	host := "ws://127.0.0.1:7880"
	apiKey := "APIwLeah7g4fuLYDYAJeaKsSE"
	apiSecret := "8nTlwISkb-63DPP7OH4e.nw.J44JjicvZDiz8J59EoQ+"
	roomName := "myroom"
	identity := "botuser"

	room, err := lksdk.ConnectToRoom(host, lksdk.ConnectInfo{
		APIKey:              apiKey,
		APISecret:           apiSecret,
		RoomName:            roomName,
		ParticipantIdentity: identity,
	})
	if err != nil {
		panic(err)
	}

	time.Sleep(5 * time.Second)

	// sampler, err := lksdk.NewLoadTestProvider(1920)
	// if err != nil {
	// 	panic(err)
	// }

	buf := bytes.NewBuffer(make([]byte, 0))

	opts := &x264.Options{
		Width:     256,
		Height:    256,
		FrameRate: 30,
		Tune:      "animation",
		Preset:    "fast",
		Profile:   "baseline",
		// LogLevel:  x264.LogDebug,
	}

	enc, err := x264.NewEncoder(buf, opts)
	if err != nil {
		panic(err)
	}

	// webrtc.NewTra

	track, err := webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeH264, Channels: 1}, "video", "pion")
	if err != nil {
		panic(err)
	}

	// track, err := lksdk.NewLocalSampleTrack(webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeVP8}, sampler)
	// if err != nil {
	// 	panic(err)
	// }

	if _, err = room.LocalParticipant.PublishTrack(track, "track test"); err != nil {
		panic(err)
	}

	// local.SetMuted(false)

	for frame := range frames {
		img := x264.NewYCbCr(image.Rect(0, 0, opts.Width, opts.Height))
		draw.Draw(img, img.Bounds(), image.White, image.ZP, draw.Src)

		img.Set(0, opts.Height/2, color.RGBA{255, 0, 0, 255})

		err = enc.Encode(img)
		if err != nil {
			panic(err)
		}

		fmt.Println(frame.ColorModel())
		// img := x264.NewYCbCr(frame.Bounds())
		// img.ToYCbCr(frame)
		// // img := x264.NewYCbCr(image.Rect(0, 0, opts.Width, opts.Height))
		// if err := enc.Encode(img); err != nil {
		// 	panic(err)
		// }

		if err = track.WriteSample(media.Sample{
			Data:      buf.Bytes(),
			Duration:  time.Second,
			Timestamp: time.Now(),
		}); err != nil {
			panic(err)
		}

		buf.Reset()
		// fmt.Printf("mean tps: %.2f\n", gameOfLife.GetMeanTPS())
	}

	enc.Flush()
	// time.Sleep(30 * time.Second)
}
