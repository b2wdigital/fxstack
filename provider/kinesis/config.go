package kinesis

import (
	"log"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	RandomPartitionKey = "fxstack.provider.kinesis.randompartitionkey"
)

func init() {

	log.Println("getting configurations for fxstack provider kinesis")

	giconfig.Add(RandomPartitionKey, false, "ramdomize partition key")
}

func RandomPartitionKeyValue() bool {
	return giconfig.Bool(RandomPartitionKey)
}
