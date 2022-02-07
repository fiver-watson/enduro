package batch

import (
	"context"
	"io/ioutil"
	"time"

	"github.com/artefactual-labs/enduro/internal/collection"
	wferrors "github.com/artefactual-labs/enduro/internal/workflow/errors"

	"go.uber.org/cadence/workflow"
)

const (
	BatchWorkflowName = "batch-workflow"
	BatchWorkflowID   = "batch-workflow"
	BatchActivityName = "batch-activity"
)

type BatchProgress struct {
	CurrentID uint
}

type BatchWorkflowInput struct {
	Path             string
	PipelineName     string
	ProcessingConfig string
	CompletedDir     string
	RetentionPeriod  *time.Duration
}

func BatchWorkflow(ctx workflow.Context, params BatchWorkflowInput) error {
	opts := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Hour * 24 * 365,
		StartToCloseTimeout:    time.Hour * 24 * 365,
		WaitForCancellation:    true,
	})
	return workflow.ExecuteActivity(opts, BatchActivityName, params).Get(opts, nil)
}

type BatchActivity struct {
	batchsvc Service
}

func NewBatchActivity(batchsvc Service) *BatchActivity {
	return &BatchActivity{
		batchsvc: batchsvc,
	}
}

func (a *BatchActivity) Execute(ctx context.Context, params BatchWorkflowInput) error {
	files, err := ioutil.ReadDir(params.Path)
	if err != nil {
		return wferrors.NonRetryableError(err)
	}
	for _, file := range files {
		req := collection.ProcessingWorkflowRequest{
			BatchDir:         params.Path,
			Key:              file.Name(),
			IsDir:            file.IsDir(),
			PipelineName:     params.PipelineName,
			ProcessingConfig: params.ProcessingConfig,
			CompletedDir:     params.CompletedDir,
			RetentionPeriod:  params.RetentionPeriod,
		}
		_ = a.batchsvc.InitProcessingWorkflow(ctx, &req)
	}
	return nil
}
