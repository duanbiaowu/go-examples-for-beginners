package redis

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	config = redis.Options{
		Addr: ":6379",
	}
)

func Test_Nil(t *testing.T) {
	rdb := redis.NewClient(&config)
	ctx := context.Background()

	val, err := rdb.Get(ctx, "not_exist_key").Result()
	switch {
	case err == redis.Nil:
		fmt.Println("key does not exist")
	case err != nil:
		fmt.Println("Get failed", err)
	case val == "":
		fmt.Println("value is empty")
	}
}

func Test_Info(t *testing.T) {
	rdb := redis.NewClient(&config)
	ctx := context.Background()

	info, err := rdb.Info(ctx, "Memory").Result()
	switch {
	case err == redis.Nil:
		fmt.Println("key does not exist")
	case err != nil:
		fmt.Println("Get failed", err)
	case info == "":
		fmt.Println("value is empty")
	default:
		fmt.Println(info)
	}
}

func Test_Universal_Client(t *testing.T) {
	t.SkipNow()

	// rdb is *redis.Client.
	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{":6379"},
	})
	fmt.Printf("%T\n", client)

	// rdb is *redis.ClusterClient.
	clusterClient := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{":6379", ":6380"},
	})
	fmt.Printf("%T\n", clusterClient)

	// rdb is *redis.FailoverClient.
	failOverClient := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:      []string{":6379"},
		MasterName: "master",
	})
	fmt.Printf("%T\n", failOverClient)
}

func Test_Pipeline(t *testing.T) {
	rdb := redis.NewClient(&config)
	ctx := context.Background()

	pipe := rdb.Pipeline()
	key := "pipeline_counter"
	incr := pipe.Incr(ctx, key)
	pipe.Expire(ctx, key, time.Hour)
	pipe.Del(ctx, key)

	cmds, err := pipe.Exec(ctx)
	if err != nil {
		panic(err)
	}
	// The value is available only after Exec is called.
	fmt.Println(incr.Val())

	for i := range cmds {
		fmt.Println(cmds[i])
	}
}

// Alternatively, you can use Pipelined which calls Exec when the function exits:
func Test_Pipelined(t *testing.T) {
	rdb := redis.NewClient(&config)
	ctx := context.Background()

	var incr *redis.IntCmd
	key := "pipelined_counter"

	cmds, err := rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		incr = pipe.Incr(ctx, key)
		pipe.Expire(ctx, key, time.Hour)
		pipe.Del(ctx, key)
		return nil
	})

	if err != nil {
		panic(err)
	}
	// The value is available only after the pipeline is executed.
	fmt.Println(incr.Val())

	for i := range cmds {
		fmt.Println(cmds[i].String())
	}
}

// Pipelines also return the executed commands so can iterate over them to retrieve results:
func Test_Pipelined2(t *testing.T) {
	rdb := redis.NewClient(&config)
	ctx := context.Background()

	cmds, err := rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for i := 0; i < 10; i++ {
			pipe.Set(ctx, fmt.Sprintf("key%d", i), i, 5*time.Second)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	for i := range cmds {
		fmt.Println(cmds[i].(*redis.StatusCmd).Val())
	}
}

func Test_PoolSize(t *testing.T) {
	rdb := redis.NewClient(&config)
	fmt.Println(rdb.Options().PoolSize)
}

func Test_HashFieldIntoAStruct(t *testing.T) {
	type Model struct {
		Str1    string   `redis:"str1"`
		Str2    string   `redis:"str2"`
		Int     int      `redis:"int"`
		Bool    bool     `redis:"bool"`
		Ignored struct{} `redis:"-"`
	}

	rdb := redis.NewClient(&config)
	ctx := context.Background()

	key := "hash_field"
	if _, err := rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		rdb.HSet(ctx, key, "str1", "hello")
		rdb.HSet(ctx, key, "str2", "world")
		rdb.HSet(ctx, key, "int", 123)
		rdb.HSet(ctx, key, "bool", 1)
		rdb.Expire(ctx, key, 5*time.Second)
		return nil
	}); err != nil {
		panic(err)
	}

	// After that we are ready to scan the data using HGetAll:
	var model1 Model
	// Scan all fields into the model.
	if err := rdb.HGetAll(ctx, key).Scan(&model1); err != nil {
		panic(err)
	}
	fmt.Println(model1)

	var model2 Model
	// Scan a subset of the fields.
	if err := rdb.HMGet(ctx, key, "str2", "int").Scan(&model2); err != nil {
		panic(err)
	}
	fmt.Println(model2)
}

func Test_HyperLog(t *testing.T) {
	rdb := redis.NewClient(&config)
	ctx := context.Background()

	key := "hyper_log_set"

	for i := 0; i < 10; i++ {
		if err := rdb.PFAdd(ctx, key, fmt.Sprint(i)).Err(); err != nil {
			panic(err)
		}
	}
	if _, err := rdb.Expire(ctx, key, 5*time.Second).Result(); err != nil {
		panic(err)
	}

	card, err := rdb.PFCount(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("set cardinality", card)
}
