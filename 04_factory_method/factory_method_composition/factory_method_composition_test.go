package factoryMethodComposition

import "testing"

func TestInterview(t *testing.T) {

	hiringManager := &HiringManager{}

	developerManager := &DeveloperManager{}
	hiringManager.TakeInterview(developerManager)

	operationManager := &OperationManager{}
	hiringManager.TakeInterview(operationManager)
}
