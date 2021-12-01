package main

type Scheduler interface {
	Schedule(planner Planner)
}

type SerialScheduler struct{}

func (s SerialScheduler) Schedule(planner Planner) {
	/*
		Runs all jobs on the current thread
	*/
}

type ParallelScheduler struct{}

func (s ParallelScheduler) Schedule(planner Planner) {
	/*
		Runs each job in a separate goroutine.
		Find a way to find out when each goroutine has finished.
	*/
}
