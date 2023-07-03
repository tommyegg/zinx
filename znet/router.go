package znet

import "github.com/tommyegg/zinx/ziface"

type BaseRouter struct{}

// 在處理conn之前的hook
func (br *BaseRouter) PreHandler(request ziface.IRequest) {}

// 在處理conn的主方法hook
func (br *BaseRouter) Handler(request ziface.IRequest) {}

// 在處理conn之後的hook
func (br *BaseRouter) PostHandler(request ziface.IRequest) {}
