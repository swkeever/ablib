package main

import (
	"ablib"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

func makeConfig(filename string) ablib.Experiments {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	c := ablib.Experiments{}
	if err = json.Unmarshal(dat, &c); err != nil {
		log.Fatalf("failed to unmarshal: %v", err)
	}
	return c
}

func main() {
	experiments := makeConfig("testdata/example_config.json")
	n := 6_000_000
	experiment, ok := experiments["testing"]
	if !ok {
		log.Fatalf("failed to get experiment")
	}
	ctr := make(map[string]int)
	start := time.Now()
	for i := 0; i < n; i++ {
		randInput := fmt.Sprintf("%v", rand.Int63())
		t, err := experiment.Treatment(randInput)
		if err != nil {
			log.Fatalf("error getting treatment: %v", err)
		}
		ctr[t]++
	}
	fmt.Printf("took %v for %d trials\n", time.Since(start), n)
	for _, comp := range experiment {
		fmt.Printf("%s=%0.4f%%\n", comp.Name, (float64(ctr[comp.Name])*100.0)/float64(n))
	}
}



