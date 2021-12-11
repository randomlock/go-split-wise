package model

import (
    "github.com/google/uuid")

type BaseModel struct {
    id string
}

func (b BaseModel) Id() string {
    return b.id
}

func NewBaseModel() *BaseModel {
    return &BaseModel{ id: uuid.New().String()}
}
