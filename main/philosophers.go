package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

const philos_number int = 5

// Each philosopher should eat only 3 times (
const max_eatings int = 3

// The host allows no more than 2 philosophers to eat concurrently.
const max_eating_philos int = 2

type Host struct {
	// philo name : eating count
	eatings       map[string]int
	full_philos   int
	eating_philos int
	mt            sync.Mutex
}

func (host *Host) feed(philos []*Philo, done *func()) {
	for host.full_philos < philos_number {
		for i := 0; i < philos_number; i++ {
			if !host.get_permission(philos[i].name) {
				continue
			}
			var done = func(name string) {
				host.done(name)
			}
			go philos[i].eat(&done)
		}
	}
	(*done)()
}

func (host *Host) get_permission(name string) bool {
	if host.eating_philos == max_eating_philos {
		return false
	}
	if host.eatings[name] >= max_eatings {
		return false
	}
	host.mt.Lock()
	host.eating_philos = host.eating_philos + 1
	host.mt.Unlock()
	return true
}

func (host *Host) done( /*wg *sync.WaitGroup, */ name string) {
	host.mt.Lock()
	host.eating_philos = host.eating_philos - 1
	host.eatings[name] = host.eatings[name] + 1
	if host.eatings[name] >= max_eatings {
		host.full_philos = host.full_philos + 1
	}
	// if host.full_philos == philos_number {
	// 	wg.Done()
	// 	fmt.Println("Done!")
	// }
	host.mt.Unlock()
}

type ChopS struct {
	sync.Mutex
}

func makeChopSticks(size int) []*ChopS {
	chop_sticks := make([]*ChopS, size)
	for i := 0; i < size; i++ {
		chop_sticks[i] = new(ChopS)
	}
	return chop_sticks
}

type Philo struct {
	name            string
	leftCS, rightCS *ChopS
	eat_attempt     int
}

func (p *Philo) eat(done *func(name string)) {
	if p.eat_attempt >= max_eatings {
		return
	}

	p.leftCS.Lock()
	p.rightCS.Lock()

	fmt.Println("starting to eat " + p.name)
	time.Sleep(time.Millisecond)
	fmt.Println("finishing eating " + p.name)

	p.eat_attempt = p.eat_attempt + 1

	p.rightCS.Unlock()
	p.leftCS.Unlock()
	(*done)(p.name)
}

func makePhilos(size int, chop_sticks []*ChopS) []*Philo {
	philos := make([]*Philo, philos_number)
	for i := 0; i < philos_number; i++ {
		philos[i] = &Philo{
			name:    strconv.Itoa(i + 1),
			leftCS:  chop_sticks[i],
			rightCS: chop_sticks[(i+1)%philos_number],
		}
	}
	return philos
}

func main() {
	var wg sync.WaitGroup

	philos := makePhilos(philos_number, makeChopSticks(philos_number))
	host := Host{eatings: make(map[string]int)}
	var done = func() {
		wg.Done()
	}
	wg.Add(1)
	go host.feed(philos, &done)
	wg.Wait()
}
