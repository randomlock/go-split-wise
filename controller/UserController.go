package controller

import (
    "../datastore"
    "../model"
)

type UserController struct {
    userDataStore datastore.UserDataStore
}

func NewUserController(userDataStore datastore.UserDataStore) *UserController {
    return &UserController{userDataStore: userDataStore}
}

func (c UserController) AddUser(name string, phNo string) (user *model.User, err error) {

    user = model.NewUser(name, phNo)
    err =  c.userDataStore.AddUser(user)
    return
}

func (c UserController) GetUser(userId string) (*model.User,error) {
    return c.userDataStore.GetUser(userId)
}

