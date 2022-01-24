package main

import (
	"ablib"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	config, err := ablib.MakeConfig("config.yml")
	if err != nil {
		log.Fatalf("failed to make config: %v", err)
	}
	n := 2_000_000
	start := time.Now()
	experiment, err := config.Experiment("hello_world")
	if err != nil {
		log.Fatalf("failed to get config: %v", err)
	}
	ctr := make(map[string]int)
	for i := 0; i < n; i++ {
		randInput := fmt.Sprintf("%v", rand.Int63())
		t, err := experiment.Treatment(randInput)
		if err != nil {
			log.Fatalf("error getting treatment: %v", err)
		}
		ctr[t]++
	}
	fmt.Printf("took %v for %d trials of %v\n", time.Since(start), n, experiment)
	for _, comp := range experiment.Comp {
		fmt.Printf("\t%s=%0.4f%%\n", comp.Name, (float64(ctr[comp.Name])*100.0)/float64(n))
	}
}



