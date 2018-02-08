package main

import (
	"image"
	"image/png"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/sync", mandelbrotHand(genMandelbrotSync))
	http.HandleFunc("/parallelPx", mandelbrotHand(genMandelbrotParallelPx))
	http.HandleFunc("/parallelCol", mandelbrotHand(genMandelbrotParallelCol))
	http.HandleFunc("/workers", mandelbrotHand(genMandelbrotWorkers))
	http.HandleFunc("/workersNoChan", mandelbrotHand(genMandelbrotWorkersNoChan))
	http.ListenAndServe(":8080", nil)
}

func mandelbrotHand(mandelbrotGen func(w int, rMin, rMax, iMin, iMax float64) *image.RGBA) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		w.WriteHeader(http.StatusOK)

		start := time.Now()
		img := mandelbrotGen(1500, -2, 0.5, -1, 1)
		elapsed := time.Since(start)

		err := png.Encode(w, img)
		if err != nil {
			log.Println("Failed to encode image to png", err)
			return
		}
		log.Printf("Generated Image in %s", elapsed)
	}
}
