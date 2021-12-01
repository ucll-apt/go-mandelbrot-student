package main

import (
	"fmt"
	"time"
)

const outputPath = "g:/temp/mandel.wif"
const initialWidth float64 = 3
const minimumWidth float64 = 0.0001
const zoomFactor float64 = 0.98
const x float64 = -0.761574
const y float64 = -0.0847596
const horizontalResolution int = 1920
const verticalResolution int = 1080
const maxIterations int = 255
const maxMagnitude float64 = 5.0

func main() {
	mandelbrots := createFrames()
	fmt.Printf("Rendering %d frames\n", len(mandelbrots))

	// Create planner
	planner := PixelPlanner{mandelbrots: mandelbrots}

	// Create scheduler
	scheduler := SerialScheduler{}

	before := time.Now()
	scheduler.Schedule(planner)
	elapsed := time.Since(before)

	fmt.Printf("Used %s seconds\n", elapsed)

	// Use for WPF based viewer
	ExportBinary(outputPath, mandelbrots)

	// User for Python based viewer
	// ExportText(outputPath, mandelbrots)
}

func createFrames() []Mandelbrot {
	result := []Mandelbrot{}
	ratio := float64(horizontalResolution) / float64(verticalResolution)
	width := initialWidth

	for width > minimumWidth {
		rectangle := NewRectangle(&Point{X: x, Y: y}, ratio, float64(width))
		mandelbrot := NewMandelbrot(horizontalResolution, verticalResolution, rectangle, maxIterations, maxMagnitude)
		result = append(result, *mandelbrot)
		width *= zoomFactor
	}

	return result
}
