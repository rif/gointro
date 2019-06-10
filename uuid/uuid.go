package uuid

import "github.com/nats-io/nuid"

func Uuid(multiplier int) string {
	if multiplier <= 0 {
		multiplier = 1
	}
	sum := ""
	for i := 0; i < multiplier; i++ {
		sum += nuid.Next()
	}
	return sum
}
