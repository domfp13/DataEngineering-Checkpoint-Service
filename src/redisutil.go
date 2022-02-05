package src

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

// GetAllCheckpoints Pulls ALL checkpoint keys.
// Output:
// 		[]string Returns a list of strings with the keys
// 		error Returns an error if something goes wrong with the function.
func GetAllCheckpoints() ([]string, error) {
	var keys []string
	iter := rdb.Scan(ctx, 0, "*", 0).Iterator()
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}
	return keys, nil
}

// GetCheckpoint Pulls a checkpoint struct out of the redis server
// Inputs:
//     tableName: TableName key.
// Output:
// 		*CheckpointObject Returns a CheckpointObject pointer
// 		error Returns an error if something goes wrong with the function.
func GetCheckpoint(tableName string) (*CheckpointObject, error) {
	result, redisErr := rdb.Get(ctx, tableName).Result()
	if redisErr != nil && redisErr != redis.Nil {
		log.Println(redisErr)
		return nil, fmt.Errorf("find: redis error: %w", redisErr)
	}
	if result == "" {
		log.Println("find: not found")
		return nil, fmt.Errorf("find: not found")
	}
	checkpoint := &CheckpointObject{}
	if err := checkpoint.UnmarshalBinary([]byte(result)); err != nil {
		log.Println(err)
		return nil, fmt.Errorf("find: unmarshal error: %w", err)
	}
	return checkpoint, nil
}

// SetCheckpoint Takes a CheckpointObject struct, creates a binary and stores it in the redis server.
// Inputs:
//		tableName: TableName key.
//		CheckpointObject Receives a CheckpointObject struct.
// Output:
// 		error Returns an error if something goes wrong with the function.
func SetCheckpoint(tableName string, checkpoint CheckpointObject) error {
	value, err := checkpoint.MarshalBinary()
	if err != nil {
		log.Println(err)
		return fmt.Errorf("set: marshalling eror: %w", err)
	}
	redisErr := rdb.Set(ctx, tableName, value, 0).Err()
	if redisErr != nil {
		log.Println(redisErr)
		return fmt.Errorf("set: redis Error: %w", redisErr)
	} else {
		return nil
	}
}
