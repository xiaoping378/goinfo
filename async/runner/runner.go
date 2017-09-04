/*
*made by xxp 2017/09/01
 */

package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner 结构
type Runner struct {
	tasks     []func(int)
	life      <-chan time.Time
	finished  chan error
	interrupt chan os.Signal
}

// ErrTimeOut 超时错误
var ErrTimeOut = errors.New("执行者执行超时")

// ErrInterrupt 任务中断
var ErrInterrupt = errors.New("执行者被中断")

// New runner的工厂函数
func New(tm time.Duration) *Runner {
	return &Runner{
		life:      time.After(tm),
		finished:  make(chan error),
		interrupt: make(chan os.Signal, 1),
	}
}

//Add 添加要执行的任务
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

//执行任务，执行的过程中接收到中断信号时，返回中断错误
//如果任务全部执行完，还没有接收到中断信号，则返回nil
func (r Runner) run() error {
	for id, task := range r.tasks {
		if r.isInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

func (r Runner) isInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}

//Start 开始执行所有任务，并且监视通道事件
func (r *Runner) Start() error {
	//希望接收哪些系统信号
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.finished <- r.run()
	}()

	select {
	case err := <-r.finished:
		return err
	case <-r.life:
		return ErrTimeOut
	}
}
