package main

import (
	"bufio"
	"fmt"
	"os"
)

type Reindeer struct {
	Name                              string
	Speed, Uptime, Downtime, Distance int
	FlyingTime, RestingTime           int
	Score                             int
}

const (
	time int = 2503
)

var (
	reindeers       = []*Reindeer{}
	leadDistance    int
	leadingReindeer *Reindeer
)

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var (
			name                    string
			speed, uptime, downtime int
		)
		info := scanner.Text()
		if n, _ := fmt.Sscanf(info, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &name, &speed, &uptime, &downtime); n == 4 {
			reindeer := Reindeer{Name: name, Speed: speed, Uptime: uptime, Downtime: downtime}
			reindeers = append(reindeers, &reindeer)
		} else {
			panic(info)
		}
	}

	highestDistance := 0
	for _, reindeer := range reindeers {
		reindeer.calculateDistance(time)
		distance := reindeer.Distance
		if distance > highestDistance {
			highestDistance = distance
			leadingReindeer = reindeer
		}
	}
	fmt.Printf("%s has won with %d km at the end of race\n", leadingReindeer.Name, leadingReindeer.Distance)

	ResetCompition()

	highestDistance = 0
	for i := 0; i < time; i++ {
		for _, reindeer := range reindeers {
			reindeer.Tick()
			if highestDistance < reindeer.Distance {
				highestDistance = reindeer.Distance
			}
		}
		for _, reindeer := range reindeers {
			if highestDistance == reindeer.Distance {
				reindeer.Score = reindeer.Score + 1
			}
		}
	}

	highestScore := 0
	for _, reindeer := range reindeers {
		if reindeer.Score > highestScore {
			highestScore = reindeer.Score
			leadingReindeer = reindeer
		}
	}

	fmt.Printf("%s has won with %d points at the end\n", leadingReindeer.Name, leadingReindeer.Score)
}

func ResetCompition() {
	leadingReindeer = nil
	for _, reindeer := range reindeers {
		reindeer.Distance = 0
	}
}

func (r *Reindeer) Tick() {
	if r.RestingTime > 0 {
		r.RestingTime = r.RestingTime - 1
	} else {
		r.FlyingTime = r.FlyingTime + 1
		r.Distance = r.Distance + r.Speed
	}

	if r.FlyingTime == r.Uptime {
		r.FlyingTime = 0
		r.RestingTime = r.Downtime
	}
}

func (r *Reindeer) calculateDistance(time int) {

	for time > 0 {
		if time < r.Uptime {
			r.Distance = r.Distance + r.Speed*time
		} else {
			r.Distance = r.Distance + r.Speed*r.Uptime
		}
		time = time - (r.Uptime + r.Downtime)

	}
}
