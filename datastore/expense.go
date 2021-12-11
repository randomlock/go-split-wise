package datastore

import (
    "fmt"

    "../model"
)

type ExpenseDataStore interface {
    AddExpense(expense *model.Expense) error
    GetExpense(expenseId string) (*model.Expense, error)
    DeleteExpense(expenseId string) error
}


type InMemoryExpenseDataStore struct {
    expenses map[string]*model.Expense
}

func NewInMemoryExpenseDataStore() *InMemoryExpenseDataStore {
    return &InMemoryExpenseDataStore{
        expenses: map[string]*model.Expense{},
    }
}

func (c *InMemoryExpenseDataStore) AddExpense(expense *model.Expense) error  {
    if _, exists := c.expenses[expense.Id()]; exists {
        return fmt.Errorf("expense already exists")
    }
    c.expenses[expense.Id()] = expense
    return nil
}

func (c *InMemoryExpenseDataStore) GetExpense(expenseId string) (*model.Expense, error)  {
    if _, exists := c.expenses[expenseId]; !exists {
        return nil, fmt.Errorf("expense doesn't already exists")
    }
    return c.expenses[expenseId], nil
}

func (c *InMemoryExpenseDataStore) DeleteExpense(expenseId string) error  {
    if _, exists := c.expenses[expenseId]; !exists {
        return fmt.Errorf("expense doesn't already exists")
    }
    delete(c.expenses, expenseId)
    return nil
}

