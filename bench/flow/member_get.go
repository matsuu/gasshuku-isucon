package flow

import (
	"context"
	"fmt"
	"net/http"

	"github.com/isucon/isucandar"
	"github.com/isucon/isucandar/failure"
	"github.com/isucon/isucandar/worker"
	"github.com/logica0419/gasshuku-isucon/bench/action"
	"github.com/logica0419/gasshuku-isucon/bench/model"
	"github.com/logica0419/gasshuku-isucon/bench/validator"
)

func (c *FlowController) MemberGetFlow(step *isucandar.BenchmarkStep) worker.WorkerFunc {
	return func(ctx context.Context, _ int) {
		res, err := c.ma.GetMembers(ctx, action.GetMembersQuery{})
		if model.IsErrTimeout(err) {
			step.AddError(fmt.Errorf("GET /api/members: %w", failure.NewError(model.ErrTimeout, nil)))
			return
		}
		if err != nil {
			step.AddError(fmt.Errorf("GET /api/members: %w", failure.NewError(model.ErrRequestFailed, err)))
			return
		}

		err = validator.Validate(res,
			validator.WithStatusCode(http.StatusOK),
		)
		if err != nil {
			step.AddError(fmt.Errorf("GET /api/members: %w", err))
			return
		}
	}
}
