package pool

import (
	"sync"
)

// init
func init() {
}

// WaitGroup 一个异步结构体

type WaitGroup struct {
	workChan chan int
	wg       sync.WaitGroup
}

// NewPool 生成一个工作池, coreNum 限制

func NewPool(coreNum int) *WaitGroup {
	ch := make(chan int, coreNum)
	return &WaitGroup{
		workChan: ch,
		wg:       sync.WaitGroup{},
	}
}

// Add 添加

func (p *WaitGroup) Add(num int) {
	for i := 0; i < num; i++ {
		p.workChan <- i
		p.wg.Add(1)
	}
}

// Done

func (p *WaitGroup) Done() {
LOOP:
	for {
		select {
		case <-p.workChan:
			break LOOP
		}
	}
	p.wg.Done()
}

// Wait 等待

func (p *WaitGroup) Wait() {
	p.wg.Wait()
}
