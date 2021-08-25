package main

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"time"

	lksdk "github.com/livekit/server-sdk-go"
	"github.com/minskylab/calab"
	"github.com/minskylab/calab/experiments"
	"github.com/minskylab/calab/experiments/petridish"
	"github.com/minskylab/calab/spaces/board"
	"github.com/minskylab/calab/systems/lifelike"
	"github.com/pion/webrtc/v3"
	"github.com/pion/webrtc/v3/pkg/media"
)

func basicLifeLike(w, h int, lifeRule *lifelike.Rule) *petridish.PetriDish {
	dynamic := lifelike.MustNew(lifeRule, lifelike.ToroidBounded, lifelike.MooreNeighborhood(1, false))
	space := board.MustNew(w, h).Fill(board.UniformNoise, dynamic)

	return petridish.NewDefault(calab.BulkDynamicalSystem(space, dynamic))
}

// func fastCyclicAutomata(w, h int, radius, states, threshold, stochastic int) *petridish.PetriDish {
// 	nh := cyclic.MooreNeighborhood(radius, false)
// 	dynamic := cyclic.MustNewRockPaperScissor(cyclic.ToroidBounded, nh, states, threshold, stochastic)

// 	space := board.MustNew(w, h).Fill(board.UniformNoise, dynamic)

// 	return petridish.NewDefault(calab.BulkDynamicalSystem(space, dynamic))
// }

func main() {
	gameOfLife := basicLifeLike(256, 256, lifelike.GameOfLifeRule)
	// rockPaperSicsors := fastCyclicAutomata(256, 256, 2, 6, 2, 1)

	experiment := experiments.New()

	// experiment.AddPetriDish(classicLifeLike)
	experiment.AddPetriDish(gameOfLife)

	// fmt.Printf("classicLifeLike id: %s\n", classicLifeLike.ID)
	fmt.Printf("rockPaperSicsors id: %s\n", gameOfLife.ID)

	// server.ServeExperiment(experiment, 8080)

	frames, err := experiment.Observe(gameOfLife.ID)
	if err != nil {
		panic(err)
	}

	go gameOfLife.Run(30 * time.Minute)

	host := "ws://143.244.182.101:7880"
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

	buf := bytes.NewBuffer([]byte{})

	// allParticipants := []string{}
	// for _, participant := range room.GetParticipants() {
	// 	allParticipants = append(allParticipants, participant.SID())
	// }
	// prov := lksdk.NewNullSampleProvider(256)
	// // webrtc.

	// reader, err := h264writer.New("nil")
	// if err != nil {
	// 	panic(err)
	// }

	track, err := webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeVP8}, "video", "pion")
	if err != nil {
		panic(err)
	}

	local, err := room.LocalParticipant.PublishTrack(track, "track test")
	if err != nil {
		panic(err)
	}

	local.SetMuted(true)

	// track.WriteSample()
	for frame := range frames {
		if err := jpeg.Encode(buf, frame, &jpeg.Options{Quality: 90}); err != nil {
			panic(err)
		}

		// if err := room.LocalParticipant.PublishData(buf.Bytes(), livekit.DataPacket_RELIABLE, allParticipants); err != nil {
		// 	panic(err)
		// }
		track.WriteSample(media.Sample{
			Data:     buf.Bytes(),
			Duration: time.Second,
		})

		buf.Reset()
		fmt.Printf("mean tps: %.2f\n", gameOfLife.GetMeanTPS())
	}
}
