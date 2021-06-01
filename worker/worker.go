// The worker package runs the worker that checks for remaining checks every 5 minutes
package worker

import (
	"fmt"

	"github.com/bamzi/jobrunner"
)

func Run() {
	jobrunner.Start()
	jobrunner.Schedule("@every 5s", CheckStatus{})
}

type CheckStatus struct{}

func (c CheckStatus) Run() {
	fmt.Println("Runs every 5 seconds")
}
