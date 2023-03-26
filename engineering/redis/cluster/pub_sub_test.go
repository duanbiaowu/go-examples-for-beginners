package cluster

import (
	"context"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func TestPubSub(t *testing.T) {
	t.SkipNow()

	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{":6379", ":6380", ":6381", ":6382", ":6383", ":6384"},
		Password: "123456",
	})
	ctx := context.TODO()

	err := rdb.Ping(ctx).Err()
	assert.Nil(t, err)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		sub := rdb.Subscribe(ctx, "test-channel")
		ch := sub.Channel()

		defer func() {
			err = sub.Close()
			assert.Nil(t, err)
			wg.Done()
		}()

		for {
			select {
			case msg := <-ch:
				t.Logf("received: %s\n", msg.Payload)
			case <-time.After(time.Second * 3):
				return
			}
		}
	}()

	time.Sleep(time.Second * 1)

	go func() {
		defer func() {
			wg.Done()
		}()

		for i := 0; i < 10; i++ {
			rdb.Publish(ctx, "test-channel", strconv.Itoa(i*10000))
		}
	}()

	wg.Wait()
	assert.Nil(t, rdb.Close())
}
