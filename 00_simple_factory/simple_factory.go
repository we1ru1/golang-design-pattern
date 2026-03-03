package simpleFactory

import "fmt"

// API is interface
type API interface {
	Say(name string) string
}

// 为什么这里可以把 API 作为 NewAPI() 的返回值？
//     1. 本质是“隐式赋值”。实际是告诉编译器：“这个函数会返回一个值，
// 我不告诉你具体是什么类型（可能是 hiAPI，也可能是 helloAPI），但我保
// 证它一定实现了 API 接口里的所有方法。”
//     2. 把接口作为返回值，是面向对象编程中非常重要的一种思想，叫做
// “面向接口编程”。能够隐藏接口的具体实现细节，并且提供统一的调用方式。

// NewAPI return Api instance by type
func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

// hiAPI is one of API implement
type hiAPI struct{}

// Say hi to name
func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

// helloAPI is another API implement
type helloAPI struct{}

// Say hello to name
func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
