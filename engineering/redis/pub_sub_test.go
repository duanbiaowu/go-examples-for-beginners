package redis

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

func Test_PubSub(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})

	start := make(chan struct{})
	done := make(chan struct{})
	ctx := context.Background()

	//ctx, cancel := context.WithCancel(context.Background())
	//cancel() // context canceled

	go func() {
		sub := rdb.Subscribe(ctx, "dev")
		ticker := time.NewTicker(time.Second)
		ch := sub.Channel()

		defer func() {
			err := sub.Close()
			if err != nil {
				log.Println(err)
			}
			ticker.Stop()
			done <- struct{}{}
		}()

		start <- struct{}{}

		for {
			select {
			case msg := <-ch:
				fmt.Printf("channel = %s, msg = %s\n", msg.Channel, msg.Payload)
			case <-ticker.C:
				return
			}
		}
	}()

	<-start

	for i := 0; i < 10; i++ {
		err := rdb.Publish(ctx, "dev", i).Err()
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("message [%d] published\n", i)
		}
	}

	<-done
}
