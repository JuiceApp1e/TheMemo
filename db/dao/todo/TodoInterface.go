package todo

type TodoInterface interface {
	saveTodos(openid string, todos interface{}) error
	getTodos(openid string) (todos interface{}, err error)
}

// TodoInterfaceImp 备忘录数据模型实现
type TodoInterfaceImp struct{}

func (imp TodoInterfaceImp) saveTodos(openid string, todos interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (imp TodoInterfaceImp) getTodos(openid string) (todos interface{}, err error) {
	//TODO implement me
	panic("implement me")
}

// TodoImp 实现实例
var TodoImp TodoInterface = &TodoInterfaceImp{}
