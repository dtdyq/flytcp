package event

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestPubSub(t *testing.T) {
	Sub("a", func(i interface{}) {
		fmt.Println(i)
	})
	Pub("a", "string")
	time.Sleep(time.Second)
}

func TestMath(t *testing.T) {
	val := math.Pow(1000, 0.25)
	fmt.Println(val)
}
