package datastore

import (
    "fmt"

    "../model"
)

type UserDataStore interface {
    AddUser(expense *model.User) error
    GetUser(expenseId string) (*model.User, error)
    DeleteUser(expenseId string) error
}


type InMemoryUserDataStore struct {
    users map[string]*model.User
}

func NewInMemoryUserDataStore() *InMemoryUserDataStore {
    return &InMemoryUserDataStore{
        users: map[string]*model.User{},
    }
}

func (c *InMemoryUserDataStore) AddUser(expense *model.User) error  {
    if _, exists := c.users[expense.Id()]; exists {
        return fmt.Errorf("expense already exists")
    }
    c.users[expense.Id()] = expense
    return nil
}

func (c *InMemoryUserDataStore) GetUser(expenseId string) (*model.User, error)  {
    if _, exists := c.users[expenseId]; !exists {
        return nil, fmt.Errorf("expense doesn't already exists")
    }
    return c.users[expenseId], nil
}

func (c *InMemoryUserDataStore) DeleteUser(expenseId string) error  {
    if _, exists := c.users[expenseId]; !exists {
        return fmt.Errorf("expense doesn't already exists")
    }
    delete(c.users, expenseId)
    return nil
}



