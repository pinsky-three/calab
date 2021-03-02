package main

import (
	"fmt"
	"image"
	"image/color/palette"
	"image/gif"
	"os"
	"sync"
	"time"

	"github.com/minskylab/calab"
	"github.com/minskylab/calab/gol"
)

func main() {
	images := make(chan image.Image)

	model := gol.NewGoLDynamicalSystem(512, 512, time.Now().Unix())

	renderer := gol.NewImageRenderer(images, 512, 512)

	vm := calab.NewVM(model, renderer)

	wg := &sync.WaitGroup{}

	animationFile, err := os.OpenFile("animation.gif", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	imagesArr := []*image.Paletted{}

	go func(wg *sync.WaitGroup, imagesArr *[]*image.Paletted) {
		mu := &sync.Mutex{}

		for img := range images {
			mu.Lock()
			wg.Add(1)

			gImg := image.NewPaletted(img.Bounds(), palette.Plan9)
			gImg.Pix = img.(*image.Gray).Pix
			*imagesArr = append(*imagesArr, gImg)

			mu.Unlock()
			wg.Done()
		}

	}(wg, &imagesArr)

	vm.Run(1 * time.Minute)

	wg.Wait()

	fmt.Println(len(imagesArr))
	delays := []int{}
	for range imagesArr {
		delays = append(delays, 0)
	}

	if err := gif.EncodeAll(animationFile, &gif.GIF{
		Image: imagesArr,
		Delay: delays,
	}); err != nil {
		panic(err)
	}

	animationFile.Close()
}
