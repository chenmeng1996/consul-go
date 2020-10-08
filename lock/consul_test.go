package lock

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestLock(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	for i := 0; i < 2; i++ {
		go fuck(&wg)
		time.Sleep(time.Second * 1)
	}

	wg.Wait()
}

func fuck(wg *sync.WaitGroup) {
	fmt.Println("prepare fuck")
	sessionId := CreateSession("fuck")
	for {
		ok := Lock("fuck", sessionId)
		if !ok {
			time.Sleep(time.Millisecond * 100)
			continue
		}
		fmt.Println("start fuck")
		time.Sleep(time.Second * 2)
		fmt.Println("end fuck")
		UnLock("fuck", sessionId)
		break
	}
	DeleteSession(sessionId)
	wg.Done()
}
