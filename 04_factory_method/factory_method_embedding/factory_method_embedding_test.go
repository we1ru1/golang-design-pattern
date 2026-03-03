package factoryMethodEmbedding

import "testing"

func TestInterview(t *testing.T) {
	developerManager := NewDeveloperManager()
	developerManager.TakeInterview()

	operationManager := NewOperationManager()
	operationManager.TakeInterview()
}
