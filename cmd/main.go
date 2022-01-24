package main

import (
	"ablib"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	config, err := ablib.MakeConfig("example_config.yml")
	if err != nil {
		log.Fatalf("failed to make config: %v", err)
	}
	fmt.Printf("%v", config.String())
	n := 6_000_000
	start := time.Now()
	experiment, err := config.Experiment("foo")
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
	for _, cmp := range experiment.Comp {
		fmt.Printf("\t%s=%0.4f%%\n", cmp.Name, (float64(ctr[cmp.Name])*100.0)/float64(n))
	}
}



