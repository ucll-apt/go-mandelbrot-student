package main

type Job func()

type Planner interface {
	jobCount() int
	getJob(index int) Job
}

type PixelPlanner struct {
	mandelbrots []Mandelbrot
}

func (p PixelPlanner) jobCount() int {
	nfractals := len(p.mandelbrots)
	width := p.mandelbrots[0].Width()
	height := p.mandelbrots[0].Height()
	return nfractals * width * height
}

func (p PixelPlanner) getJob(index int) Job {
	width := p.mandelbrots[0].Width()
	height := p.mandelbrots[0].Height()
	pixelCount := width * height

	frameIndex := index / pixelCount
	pixelIndex := index % pixelCount
	x := pixelIndex % width
	y := pixelIndex / width

	return func() {
		p.mandelbrots[frameIndex].ComputeSingle(x, y)
	}
}

type RowPlanner struct {
	mandelbrots []Mandelbrot
}

func (p RowPlanner) jobCount() int {
	// TODO
	return -1
}

func (p RowPlanner) getJob(index int) Job {
	/*
		One job corresponds to one row of one frame
	*/
	return nil
}

type FramePlanner struct {
	mandelbrots []Mandelbrot
}

func (p FramePlanner) jobCount() int {
	// TODO
	return -1
}

func (p FramePlanner) getJob(index int) Job {
	/*
		One full frame per job.
	*/
	return nil
}

type MonolithPlanner struct {
	mandelbrots []Mandelbrot
}

func (p MonolithPlanner) jobCount() int {
	// TODO
	return -1
}

func (p MonolithPlanner) getJob(index int) Job {
	/*
		Everything in one job
	*/
	return nil
}
