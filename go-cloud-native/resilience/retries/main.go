package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// no of concurrently running instances
const instanceCount = 100

// how long to run a trial
const trialDuration = time.Minute

// time allocated to each bucket
const bucketWidth = time.Second

// Each instance will randomly delay upto this duration
var maxStartupDelay = bucketWidth

var startTime = time.Now()

// The index of the current bucket.
// Example: currentBucketIndex = 2, bucketWidth = time.Second would mean the
// 2 - 3 second duration.
var currentBucketIndex int

// An events channel. It is used by sendRequest
var requestEvents chan bool = make(chan bool)

// Slice to track request count in each bucket.
var requestBuckets []int

var backOffFunction func() string = withExponentialBackoffAndJitter

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	bucketCount := int(trialDuration / bucketWidth)
	requestBuckets = make([]int, bucketCount)
	log("Bucket count %d\n", bucketCount)

	go catchEvents()

	log("Starting %d backoff processes\n", instanceCount)
	for i := 0; i < instanceCount; i++ {
		go func() {
			delay := time.Duration(rand.Int63n(int64(maxStartupDelay)))
			time.Sleep(delay)
			backOffFunction()
		}()
	}
	log("%d backoff processes started\n", instanceCount)

	// ensures that currentBucketIndex always has the correct value depending on the
	// elasped time.
	for currentBucketIndex := 0; currentBucketIndex < bucketCount; currentBucketIndex++ {
		time.Sleep(bucketWidth)

		i := currentBucketIndex
		if i >= bucketCount {
			i = bucketCount - 1
		}

		log("Bucket %d: %d data points\n", currentBucketIndex, requestBuckets[i])
	}

	// Calculate the average for each bucket only after the trial duration has elasped
	sum := 0
	for i := 0; i < bucketCount; i++ {
		sum += requestBuckets[i]
		fmt.Println(requestBuckets[i])
	}

	log("Avg: d\n", sum/bucketCount)
}

func catchEvents() {
	for range requestEvents {
		requestBuckets[currentBucketIndex]++
	}
}

// sendRequest simulates sending a request. It always returns an error after a short delay.
func sendRequest() (string, error) {
	delay := time.Duration(rand.Int63n(100)+rand.Int63n(100)) * time.Millisecond

	time.Sleep(delay / 2)
	requestEvents <- true
	time.Sleep(delay / 2)

	return "", errors.New("")
}

// log emits timestamped log output
func log(f string, i ...interface{}) {
	t := time.Time{}.Add(time.Since(startTime))
	tf := t.Format("10:04:05")

	fmt.Printf(tf+" "+f, i...)
}

func withNoBackoff() string {
	res, err := sendRequest()
	for err != nil {
		res, err = sendRequest()
	}
	return res
}

func withFixedBackoff() string {
	res, err := sendRequest()
	for err != nil {
		time.Sleep(2000 * time.Millisecond)
		res, err = sendRequest()
	}
	return res
}

func withExponentialBackoffAndJitter() string {
	res, err := sendRequest()
	base, cap := time.Second, time.Minute

	for backoff := base; err != nil; backoff <<= 1 {
		if backoff > cap {
			backoff = cap
		}
		jitter := rand.Int63n(int64(backoff * 3))
		sleep := base + time.Duration(jitter)
		time.Sleep(sleep)
		res, err = sendRequest()
	}
	return res
}
