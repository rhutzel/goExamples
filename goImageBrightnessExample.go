package main

import (
	"flag"
	"log"
	"time"
	"github.com/rhutzel/goImageBrightness"
)

func main() {
	flagSamplingPtr := flag.Bool("sampling", false, "Linear brightness averaging of small image fixture sampling.")
	flagLinearPtr := flag.String("linear", "", "Linear brightness averaging of large image.")
	flagThreadedPtr := flag.String("threaded", "", "Threaded brightness averaging of large image.")
	flag.Parse()

	switch {
	case *flagSamplingPtr:
		doSampling()
		return
	case len(*flagLinearPtr) > 0:
		doLargeLinear(flagLinearPtr)
		return
	case len(*flagThreadedPtr) > 0:
		doLargeThreaded(flagThreadedPtr)
		return
	}

	flag.PrintDefaults()
}

func doSampling() {
	img, imgType, err := goImageBrightness.ImageFromFile("sampleImageAllWhite.png")
	if err != nil {
		log.Fatalf("Failed to decode [%s] image.", imgType)
	}
	brightnessPercentage := goImageBrightness.AnalyseImage(img)
	log.Printf("Total relative luminance is %d%%.", brightnessPercentage)

	img, imgType, err = goImageBrightness.ImageFromFile("sampleImageMostlyBrightGreen.png")
	if err != nil {
		log.Fatalf("Failed to decode [%s] image.", imgType)
	}
	brightnessPercentage = goImageBrightness.AnalyseImage(img)
	log.Printf("Total relative luminance is %d%%.", brightnessPercentage)

	img, imgType, err = goImageBrightness.ImageFromFile("sampleImageSlightlyBrightGreen.png")
	if err != nil {
		log.Fatalf("Failed to decode [%s] image.", imgType)
	}
	brightnessPercentage = goImageBrightness.AnalyseImage(img)
	log.Printf("Total relative luminance is %d%%.", brightnessPercentage)

	img, imgType, err = goImageBrightness.ImageFromFile("sampleImageAllBlack.png")
	if err != nil {
		log.Fatalf("Failed to decode [%s] image.", imgType)
	}
	brightnessPercentage = goImageBrightness.AnalyseImage(img)
	log.Printf("Total relative luminance is %d%%.", brightnessPercentage)
}

func doLargeLinear(pngPath *string) {
	start := time.Now()
	img, imgType, err := goImageBrightness.ImageFromFile(*pngPath)
	if err != nil {
		log.Fatalf("Failed to decode [%s] image.", imgType)
	}
	brightnessPercentage := goImageBrightness.AnalyseImage(img)
	log.Printf("Linear total relative luminance is %d%%. (%d ms)", brightnessPercentage, time.Since(start) / 1000000)
}

func doLargeThreaded(pngPath *string) {
	start := time.Now()
	img, imgType, err := goImageBrightness.ImageFromFile(*pngPath)
	if err != nil {
		log.Fatalf("Failed to decode [%s] image.", imgType)
	}
	brightnessPercentage := goImageBrightness.ParallelAnalyseImage(img, 4)
	log.Printf("Threaded total relative luminance is %d%%. (%d ms)", brightnessPercentage, time.Since(start) / 1000000)
}
