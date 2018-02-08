package main

import (
	"image"
	"image/color"
	"math/cmplx"
	"runtime"
	"sync"
)

const MAX_ITERATIONS = 100

func mandelbrot(a complex128) float64 {
	z := a
	for i := 0; i < MAX_ITERATIONS; i++ {
		z = z*z + a
		if cmplx.Abs(z) > 2 {
			return float64(MAX_ITERATIONS-i) / MAX_ITERATIONS
		}
	}
	return 0
}

func toColor(v float64) color.Color {
	const contrast = 255
	return color.Gray{255 - uint8(v*contrast)}
}

func calcPx(img *image.RGBA, x, y int, scale, rMin, iMin float64) {
	val := mandelbrot(complex(float64(x)/scale+rMin, float64(y)/scale+iMin))
	img.Set(x, y, toColor(val))
}

func genMandelbrotSync(width int, rMin, rMax, iMin, iMax float64) *image.RGBA {
	scale := float64(width) / (rMax - rMin)
	height := int(scale * (iMax - iMin))

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			calcPx(img, x, y, scale, rMin, iMin)
		}
	}

	return img
}

func genMandelbrotParallelPx(width int, rMin, rMax, iMin, iMax float64) *image.RGBA {
	scale := float64(width) / (rMax - rMin)
	height := int(scale * (iMax - iMin))

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	wg := sync.WaitGroup{}
	wg.Add(width * height)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			go func(x, y int) {
				calcPx(img, x, y, scale, rMin, iMin)
				wg.Done()
			}(x, y)
		}
	}
	wg.Wait()

	return img
}

func genMandelbrotParallelCol(width int, rMin, rMax, iMin, iMax float64) *image.RGBA {
	scale := float64(width) / (rMax - rMin)
	height := int(scale * (iMax - iMin))

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	wg := sync.WaitGroup{}
	wg.Add(width)
	for x := 0; x < width; x++ {
		go func(x int) {
			for y := 0; y < height; y++ {
				calcPx(img, x, y, scale, rMin, iMin)
			}
			wg.Done()
		}(x)
	}
	wg.Wait()

	return img
}

func genMandelbrotWorkers(width int, rMin, rMax, iMin, iMax float64) *image.RGBA {
	scale := float64(width) / (rMax - rMin)
	height := int(scale * (iMax - iMin))

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	cols := make(chan int, width)
	for x := 0; x < width; x++ {
		cols <- x
	}
	close(cols)

	numCPU := runtime.NumCPU()
	wg := sync.WaitGroup{}
	wg.Add(numCPU)
	for i := 0; i < numCPU; i++ {
		go func() {
			for x := range cols {
				for y := 0; y < height; y++ {
					calcPx(img, x, y, scale, rMin, iMin)
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()

	return img
}

func genMandelbrotWorkersNoChan(width int, rMin, rMax, iMin, iMax float64) *image.RGBA {
	scale := float64(width) / (rMax - rMin)
	height := int(scale * (iMax - iMin))

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	numCPU := runtime.NumCPU()
	wg := sync.WaitGroup{}
	wg.Add(numCPU)
	for i := 0; i < numCPU; i++ {
		go func(xBegin int) {
			for x := xBegin; x < width; x += numCPU {
				for y := 0; y < height; y++ {
					calcPx(img, x, y, scale, rMin, iMin)
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	return img
}
