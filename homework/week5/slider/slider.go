package slider

import (
	"container/ring"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/eapache/queue"
)

type Counter struct {
	queue *queue.Queue
	opt   *CounterOption
}

type CounterOption struct {
	Interval time.Duration
}

func NewCounter(opt *CounterOption) *Counter {
	return &Counter{
		queue: queue.New(),
		opt:   opt,
	}
}

func Run(c *Counter) error {

	ch := make(chan interface{}, 1)

	//run timer
	go func() {
		for {
			select {
			case <-time.After(c.opt.Interval):
				c.queue.Add(0)
			case <-ch:
			}
		}
	}()

	return nil
}

func (c *Counter) Add() {
	c.queue.Get(1)
}

func (c *Counter) GetCurrentQps() {

}

func Test() {
	sigs := []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGABRT}

	signal.Notify()
}


func RingTest(){
	ring := ring.Ring{}
}