package designpatterns

// 策略模式设计
type Strategy interface {
	add(n int) int
}

// 策略1
type addOneStep struct{}

func (add *addOneStep) add(a int) int {
	return a + 1
}

// 策略2
type addTwoStep struct{}

func (add *addTwoStep) add(a int) int {
	return a + 2
}

// 策略3
type addThreeStep struct{}

func (add *addThreeStep) add(a int) int {
	return a + 3
}

type RunStep struct {
	n     int
	ahead Strategy
}

func newRunStep(n int, ahead Strategy) *RunStep {
	return &RunStep{
		n:     n,
		ahead: ahead,
	}
}

func (run *RunStep) setStrategy(ahead Strategy) {
	run.ahead = ahead
}

func (run *RunStep) Run() {
	run.n = run.ahead.add(run.n)
}
