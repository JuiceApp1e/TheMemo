package service

import (
	tododao "wxcloudrun-golang/db/dao/todo"
	"wxcloudrun-golang/db/model"
)

func SaveTodo(todo *model.TodoModel) error {
	return tododao.TodoImp.SaveTodo(todo)
}

func DeleteTodo(id int32) error {
	return tododao.TodoImp.DeleteTodo(id)
}

func UpdateTodo(todo *model.TodoModel) error {
	return tododao.TodoImp.UpdateTodo(todo)
}

func GetTodoList() (todoList []model.TodoModel, err error) {
	return tododao.TodoImp.GetTodoList()
}
