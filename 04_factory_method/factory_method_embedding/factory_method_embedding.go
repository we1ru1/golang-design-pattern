package factoryMethodEmbedding

import "fmt"

// 定义 Interviewer 接口（产品）
type Interviewer interface {
	AskQuestion()
}

// 定义具体产品：技术面试官
type Developer struct{}

func (d *Developer) AskQuestion() {
	fmt.Println("Ask about deisign pattern!")
}

// 定义具体产品：运营面试官
type Operation struct{}

func (o *Operation) AskQuestion() {
	fmt.Println("Ask about users and content!")
}

// ManagerCreator 定义工厂方法接口。
type ManagerCreator interface {
	MakeInterviewer() Interviewer
}

// HiringManager 提供公共的面试流程。
// 通过匿名嵌入，这个方法会被“提升”到具体 Manager 上。
type HiringManager struct {
	creator ManagerCreator
}

func (hm *HiringManager) TakeInterview() {
	interviewer := hm.creator.MakeInterviewer()
	interviewer.AskQuestion()
}

// DeveloperManager 通过嵌入复用 TakeInterview。
type DeveloperManager struct {
	HiringManager
}

func NewDeveloperManager() *DeveloperManager {
	dm := &DeveloperManager{}
	dm.HiringManager.creator = dm
	return dm
}

func (dm *DeveloperManager) MakeInterviewer() Interviewer {
	return &Developer{}
}

// OperationManager 通过嵌入复用 TakeInterview。
type OperationManager struct {
	HiringManager
}

func NewOperationManager() *OperationManager {
	om := &OperationManager{}
	om.HiringManager.creator = om
	return om
}

func (om *OperationManager) MakeInterviewer() Interviewer {
	return &Operation{}
}
