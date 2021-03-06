package engine

import "sync"

type commandQueue struct {
	sync.Mutex

	cmds     []Command
	noCmd    bool
	pushDone chan struct{}
}

func (cmdQ *commandQueue) push(cmd Command) {
	cmdQ.Lock()

	cmdQ.cmds = append(cmdQ.cmds, cmd)

	cmdQ.Unlock()

	if cmdQ.noCmd {
		cmdQ.noCmd = false
		cmdQ.pushDone <- struct{}{}
	}
}

func (cmdQ *commandQueue) pull() Command {
	cmdQ.Lock()

	if len(cmdQ.cmds) == 0 {
		cmdQ.noCmd = true
		cmdQ.Unlock()
		<-cmdQ.pushDone
		cmdQ.Lock()
	}

	cmd := cmdQ.cmds[0]
	cmdQ.cmds[0] = nil
	cmdQ.cmds = cmdQ.cmds[1:]

	cmdQ.Unlock()

	return cmd
}

func (cmdQ *commandQueue) length() int {
	cmdQ.Lock()
	defer cmdQ.Unlock()

	return len(cmdQ.cmds)
}

type cmdExecutor struct {
	executor func()
}

func (cmdEx *cmdExecutor) Execute(h Handler) {
	cmdEx.executor()
}

type EventLoop struct {
	q          *commandQueue
	finishWait bool
	finishDone chan struct{}
}

func (eventl *EventLoop) Start() {
	eventl.q = &commandQueue{
		pushDone: make(chan struct{}),
	}
	eventl.finishDone = make(chan struct{})

	go func() {
		for !(eventl.finishWait && eventl.q.length() == 0) {
			cmd := eventl.q.pull()
			cmd.Execute(eventl)
		}
		eventl.finishDone <- struct{}{}
	}()
}

func (eventl *EventLoop) Post(cmd Command) {
	eventl.q.push(cmd)
}

func (eventl *EventLoop) AwaitFinish() {
	finish := &cmdExecutor{func() { eventl.finishWait = true }}
	eventl.Post(finish)

	<-eventl.finishDone
}
