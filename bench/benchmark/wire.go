//go:build wireinject
// +build wireinject

package benchmark

import (
	"github.com/google/wire"
	"github.com/isucon/isucandar/worker"
	"github.com/logica0419/gasshuku-isucon/bench/action"
	"github.com/logica0419/gasshuku-isucon/bench/config"
	"github.com/logica0419/gasshuku-isucon/bench/flow"
	"github.com/logica0419/gasshuku-isucon/bench/repository"
	"github.com/logica0419/gasshuku-isucon/bench/scenario"
)

func NewBenchmark(c chan worker.WorkerFunc) (*Benchmark, error) {
	wire.Build(
		config.NewConfig,

		repository.NewRepository,
		wire.Bind(new(repository.MemberRepository), new(*repository.Repository)),

		action.NewController,
		wire.Bind(new(action.MemberController), new(*action.Controller)),
		wire.Bind(new(action.InitializeController), new(*action.Controller)),

		flow.NewController,

		scenario.NewScenario,

		newBenchmark,
	)

	return nil, nil
}
