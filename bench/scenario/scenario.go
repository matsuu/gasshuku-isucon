package scenario

import (
	"github.com/isucon/isucandar"
	"github.com/isucon/isucandar/worker"
	"github.com/logica0419/gasshuku-isucon/bench/flow"
)

type Scenario struct {
	wc <-chan worker.WorkerFunc
	fc *flow.Controller
}

var (
	_ isucandar.PrepareScenario = &Scenario{}
	_ isucandar.LoadScenario    = &Scenario{}
)

func NewScenario(wc chan worker.WorkerFunc, fc *flow.Controller) *Scenario {
	return &Scenario{
		wc: wc,
		fc: fc,
	}
}
