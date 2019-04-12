package main

import (
	"encoding/json"
	"errors"
	"log"
	"sync"
	"time"
)

var inputJSON = `{
	"gophers": [{
		"name": "A",
		"sleep": 1,
		"eat": 1
	},
	{
		"name": "B",
		"sleep": 4,
		"eat": 3
	},
	{
		"name": "C",
		"sleep": 1,
		"eat": 4
	},
	{
		"name": "D",
		"sleep": 5,
		"eat": 2
	 },
	 {
		"name": "E",
		"sleep": 2,
		"eat": 3
	 }
	],
	"totalFood": 30
}`

// Gopher is our structure for Gopher - tags can be helpful for unmarshalling the json
type Gopher struct {
	Name  string        `json:"name"`
	Sleep time.Duration `json:"sleep"`
	Eat   int           `json:"eat"`
}

// Farm is our structure for Farm - tags can be helpful for unmarshalling the json
type Farm struct {
	Gophers   []Gopher `json:"gophers"`
	TotalFood int      `json:"totalFood"`
	mu        sync.Mutex
}

func (gopher *Gopher) gopherLive(farm *Farm, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		time.Sleep(time.Second * gopher.Sleep)
		err := farm.eatFood(gopher)
		if err != nil {
			return
		}
	}
}

func (f *Farm) eatFood(gopher *Gopher) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.TotalFood < gopher.Eat {
		log.Printf("gopher %s wants to eat %v food unit(s) but there's not enough food!", gopher.Name, gopher.Eat)
		return errors.New("there's not enough food")
	}
	f.TotalFood = f.TotalFood - gopher.Eat
	log.Printf("gopher %s eats %v food unit(s). %v food unit(s) left.", gopher.Name, gopher.Eat, f.TotalFood)
	return nil
}

func main() {
	gopherFarm := &Farm{}
	err := json.Unmarshal([]byte(inputJSON), gopherFarm)
	if err != nil {
		log.Fatal("cannot unmarshal: ", err)
	}

	wg := new(sync.WaitGroup)

	for i := range gopherFarm.Gophers {
		wg.Add(1)
		log.Printf("gopher %s joins the farm!", gopherFarm.Gophers[i].Name)
		go (&gopherFarm.Gophers[i]).gopherLive(gopherFarm, wg) // stuff to think about here

	}
	time.Sleep(time.Second * 20)
}
