package discovery

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestConsulDeregister(t *testing.T) {
	s := strings.Replace("1.2.3.4:8080", ":", "_", 1)
	fmt.Println(s)
}

func Test1(t *testing.T) {
	rand.Seed(time.Now().Unix())

	for i := 0; i < 4; i++ {
		fmt.Println(rand.Int())
	}
}
