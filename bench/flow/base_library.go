package flow

import (
	"context"
	"time"

	"github.com/isucon/isucandar"
	"github.com/isucon/isucandar/worker"
	"github.com/logica0419/gasshuku-isucon/bench/utils"
)

const (
	libraryFlowCycle = 1000 * time.Millisecond
	booksPostNum     = 100
)

// 図書館職員フロー
func (c *Controller) baseLibraryFlow(step *isucandar.BenchmarkStep) worker.WorkerFunc {
	return func(ctx context.Context, _ int) {
		select {
		case <-ctx.Done():
			return
		default:
		}

		timer := time.After(libraryFlowCycle)

		runner := utils.WeightedSelect(
			[]utils.Choice[flow]{
				{Val: c.getMembersFlow("", step)},
				{Val: c.getMembersFlow(utils.RandString(26), step), Weight: 2},
				{Val: c.getMembersFlow(c.mr.GetRandomMember().ID, step), Weight: 2},
				{Val: c.postBooksFlow(booksPostNum, step), Weight: 8},
			},
		)
		runner(ctx)

		select {
		case <-ctx.Done():
			return
		case <-timer:
			return
		default:
		}

		c.addLibInCycleCount()

		select {
		case <-ctx.Done():
			return
		case <-timer:
			return
		}
	}
}
