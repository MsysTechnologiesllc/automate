// +build integration

package integration_test

import (
	"context"
	"sync"
	"time"

	"github.com/chef/automate/lib/cereal"
)

func (suite *CerealTestSuite) TestDelayedTask() {
	taskName := randName("delayed")
	workflowName := randName("delayed")
	instanceName := randName("instance")

	wgTask := sync.WaitGroup{}
	wgTask.Add(2)

	expectedTime := time.Now().Add(10 * time.Second)
	m := suite.newManager(
		WithTaskExecutorF(
			taskName,
			func(context.Context, cereal.Task) (interface{}, error) {
				// Tasks are scheduled every 2 seconds. A little extra time is provided
				suite.Assert().WithinDuration(expectedTime, time.Now(), 3*time.Second)
				wgTask.Done()
				return nil, nil
			}),
		WithWorkflowExecutor(
			workflowName,
			&workflowExecutorWrapper{
				onStart: func(w cereal.WorkflowInstance, ev cereal.StartEvent) cereal.Decision {
					err := w.EnqueueTask(taskName, nil, cereal.StartAfter(expectedTime))
					suite.Require().NoError(err, "failed to enqueue task")
					return w.Continue(nil)
				},
				onTaskComplete: func(w cereal.WorkflowInstance, ev cereal.TaskCompleteEvent) cereal.Decision {
					suite.Assert().Equal(1, w.TotalCompletedTasks())
					suite.Assert().Equal(1, w.TotalEnqueuedTasks())
					wgTask.Done()
					return w.Complete()
				},
			},
		),
	)
	defer m.Stop()
	err := m.EnqueueWorkflow(context.Background(), workflowName, instanceName, nil)
	suite.Require().NoError(err, "Failed to enqueue workflow")
	wgTask.Wait()
	time.Sleep(20 * time.Millisecond)
	err = m.Stop()
	suite.NoError(err)
}