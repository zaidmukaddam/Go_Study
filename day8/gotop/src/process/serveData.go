package process

import (
	"context"

	proc "github.com/shirou/gopsutil/process"
	"github.com/zaidmukaddam/Golang/day8/gotop/src/utils"
)

func Serve(process *Process, dataChannel chan *Process, ctx context.Context, refreshRate int64) error {
	return utils.TickUntilDone(ctx, refreshRate, func() error {
		process.UpdateProcInfo()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case dataChannel <- process:
		}

		return nil
	})
}

func ServeProcs(dataChannel chan []*proc.Process, ctx context.Context, refreshRate int64) error {
	return utils.TickUntilDone(ctx, refreshRate, func() error {
		procs, err := proc.Processes()
		if err != nil {
			return err
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case dataChannel <- procs:
		}

		return err
	})
}
