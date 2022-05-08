package todo

import "wxcloudrun-golang/db/model"

type TodoInterface interface {
	SaveTodo(todo *model.TodoModel) error
	UpdateTodo(todo *model.TodoModel) error
	DeleteTodo(id int32) error
	GetTodoList() (todoList []model.TodoModel, err error)
}

// TodoInterfaceImp 备忘录数据模型实现
type TodoInterfaceImp struct{}

// TodoImp 实现实例
var TodoImp TodoInterface = &TodoInterfaceImp{}
