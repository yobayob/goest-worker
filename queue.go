package goest_worker

import "context"

var ChannelQueue PoolQueue = func (ctx context.Context, capacity int) Queue {
	queue := &channelQueue{
		ctx: ctx,
		ch: make(chan jobCall, capacity),
	}
	return queue
}

type channelQueue struct {
	ctx context.Context
	ch chan jobCall
}

func (c *channelQueue) Pop () (jobCall) {
	select {
	case j := <- c.ch:
		return j
	case <- c.ctx.Done():
		return nil
	}
}

func (c *channelQueue) Insert (job jobCall) () {

	select {
	case c.ch <- job:
		return
	case <- c.ctx.Done():
		return
	 }
}

func (c *channelQueue) Len () (uint64) {
	return uint64(len(c.ch))
}