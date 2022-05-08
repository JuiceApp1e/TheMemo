package todo

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

const tableName = "todo_model"

func (imp TodoInterfaceImp) SaveTodo(todo *model.TodoModel) error {
	cli := db.Get()
	return cli.Table(tableName).Create(todo).Error
}

func (imp TodoInterfaceImp) UpdateTodo(todo *model.TodoModel) error {
	cli := db.Get()
	return cli.Table(tableName).Save(todo).Error
}

func (imp TodoInterfaceImp) DeleteTodo(id int32) error {
	cli := db.Get()
	return cli.Table(tableName).Delete(&model.TodoModel{Id: id}).Error
}

func (imp TodoInterfaceImp) GetTodoList() (todoList []model.TodoModel, err error) {
	cli := db.Get()
	err = cli.Table(tableName).Find(&todoList).Error
	return
}
