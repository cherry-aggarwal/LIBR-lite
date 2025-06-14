package moderators

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/cherry-aggarwal/libr/models"
	"github.com/cherry-aggarwal/libr/modqueue"
)

var Response = [2]int{0, 1}
var Responses [3]int
var Out = 0

func AskingModsResponse() {
	rand.Seed(time.Now().UnixNano())

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("mod1")
		mut.Lock()
		Responses[0] = Response[rand.Intn(len(Response))]
		modqueue.ModChannel <- Responses[0]
		mut.Unlock()
		wg.Done()
		fmt.Println("response: ", Responses[0])

	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("mod2")
		mut.Lock()
		Responses[1] = Response[rand.Intn(len(Response))]
		modqueue.ModChannel <- Responses[1]
		mut.Unlock()
		wg.Done()
		fmt.Println("response: ", Responses[1])
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("mod3")
		mut.Lock()
		Responses[2] = Response[rand.Intn(len(Response))]
		modqueue.ModChannel <- Responses[2]
		mut.Unlock()
		wg.Done()
		fmt.Println("response: ", Responses[2])

	}(wg, mut)
	go func() {
		wg.Wait()
		close(modqueue.ModChannel)
	}()

	for val := range modqueue.ModChannel {
		Out = Out + val
	}
	fmt.Println(Out)
	// time.Sleep(1 * time.Second)

}

func SettingMsgStatus(message *models.Msg) {
	if Out > 1 {
		message.Status = "accepted"
	} else {
		message.Status = "rejected"
	}
}
