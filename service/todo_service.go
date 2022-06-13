package service

import (
	tododao "wxcloudrun-golang/db/dao/todo"
)

func SaveTodos(openid string, todos interface{}) error {
	return tododao.TodoInterfaceImp{}.SaveTodos(openid, todos)
}

func GetTodos(openid string) (todos interface{}, err error) {
	return tododao.TodoInterfaceImp{}.GetTodos(openid)
}
