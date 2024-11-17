package main

import "sync"

type ChannelCounter struct {
	count int
	mutex sync.Mutex
}

func (counter *ChannelCounter) increase() {
	counter.mutex.Lock()
	counter.count++
	counter.mutex.Unlock()
}

func (counter *ChannelCounter) isZero() bool {
	var res bool

	counter.mutex.Lock()
	res = counter.count == 0
	counter.mutex.Unlock()

	return res
}

func (counter *ChannelCounter) decrease() {
	counter.mutex.Lock()
	counter.count--
	counter.mutex.Unlock()
}

var matchThreadCounter, bbbThreadCounter ChannelCounter

var TeamLock sync.Mutex
var PlayerLock sync.Mutex
var TourLock sync.Mutex
var TournamentLock sync.Mutex
var SeriesLock sync.Mutex
var SeriesSquadLock sync.Mutex
var SeasonLock sync.Mutex
var HostNationLock sync.Mutex
var GroundLock sync.Mutex
var CityLock sync.Mutex
