package factoryMethodComposition

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

// 在 Go 中，我们通常把 takeInterview（具体方法）定义
// 在结构体上，而把 makeInterviewer（抽象方法）定义在接口里。
// --- 模拟抽象类 ---
type ManagerCreator interface {
	// 工厂方法：必须返回一个实现了 Interviewer 接口的对象
	MakeInterviewer() Interviewer
}

// 具体的面试行为的“逻辑载体”
type HiringManager struct{}

// 接收一个 ManagerCreator 接口作为参数，从而调用“抽象方法”
func (hm *HiringManager) TakeInterview(creator ManagerCreator) {
	interviewer := creator.MakeInterviewer() // 调用“子类”的实现
	interviewer.AskQuestion()
}

// Developer招聘经理
type DeveloperManager struct{}

func (dm *DeveloperManager) MakeInterviewer() Interviewer {
	return &Developer{}
}

// Operation招聘经理
type OperationManager struct{}

func (om *OperationManager) MakeInterviewer() Interviewer {
	return &Operation{}
}
