//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, tweeted chan Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(tweeted)
			return
		}

		tweeted <- *tweet
	}
}

func consumer(tweeted chan Tweet) {
	for {
		select {
		case tweet, ok := <- tweeted:
			if !ok {
				return
			}
			if tweet.IsTalkingAboutGo() {
				fmt.Println(tweet.Username, "\ttweets about golang")
			} else {
				fmt.Println(tweet.Username, "\tdoes not tweet about golang")
			}
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	tweeted := make(chan Tweet)
	// Producer
	go producer(stream, tweeted)

	// Consumer
	consumer(tweeted)

	fmt.Printf("Process took %s\n", time.Since(start))
}
