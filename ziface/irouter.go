package ziface

/*
	路由抽象接口
	路由裡的數據都是IRequest
*/

type IRouter interface {
	//在處理conn之前的hook
	PreHandler(request IRequest)
	//在處理conn的主方法hook
	Handler(request IRequest)
	//在處理conn之後的hook
	PostHandler(request IRequest)
}
