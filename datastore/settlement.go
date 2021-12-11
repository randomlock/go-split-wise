package datastore

import (
    "fmt"

    "../model"
)

type SettlementDataStore interface {
    AddSettlement(expense *model.Settlement) error
    GetSettlement(expenseId string) (*model.Settlement, error)
    DeleteSettlement(expenseId string) error
}

type InMemorySettlementDataStore struct {
    settlements map[string]*model.Settlement
}

func NewInMemorySettlementDataStore() *InMemorySettlementDataStore {
    return &InMemorySettlementDataStore{
        settlements: map[string]*model.Settlement{},
    }
}

func (c *InMemorySettlementDataStore) AddSettlement(expense *model.Settlement) error  {
    if _, exists := c.settlements[expense.Id()]; exists {
        return fmt.Errorf("expense already exists")
    }
    c.settlements[expense.Id()] = expense
    return nil
}

func (c *InMemorySettlementDataStore) GetSettlement(expenseId string) (*model.Settlement, error)  {
    if _, exists := c.settlements[expenseId]; !exists {
        return nil, fmt.Errorf("expense doesn't already exists")
    }
    return c.settlements[expenseId], nil
}

func (c *InMemorySettlementDataStore) DeleteSettlement(expenseId string) error  {
    if _, exists := c.settlements[expenseId]; !exists {
        return fmt.Errorf("expense doesn't already exists")
    }
    delete(c.settlements, expenseId)
    return nil
}

