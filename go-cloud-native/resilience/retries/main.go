package main

import (
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

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	bucketCount := int(trialDuration / bucketWidth)
	requestBuckets := make([]int, bucketCount)
	log("Bucket count %d\n", bucketCount)

	for i := 0; i < instanceCount; i++ {
		go func() {

		}()
	}

	for currentBucketIndex := 0; currentBucketIndex < bucketCount; currentBucketIndex++ {
		time.Sleep(bucketWidth)
	}
}

// log emits timestamped log output
func log(f string, i ...interface{}) {
	since := time.Now().Sub(startTime)
	t := time.Time{}.Add(since)
	tf := t.Format("10:04:05")

	fmt.Printf(tf+" "+f, i...)
}
