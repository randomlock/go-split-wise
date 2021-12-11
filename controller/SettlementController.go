package controller

import (
    "../datastore"
    "../model"
)

type SettlementController struct {
    settlementDataStore datastore.SettlementDataStore
}

func NewSettlementController(settlementDataStore datastore.SettlementDataStore) *SettlementController {
    return &SettlementController{settlementDataStore: settlementDataStore}
}

func (c *SettlementController) CreateSettlement(amount float64, userId string) (settlement *model.Settlement, err error) {

    // expense, err := c.expenseDataStore.GetExpense(expenseId)
    // if err != nil {
    //     return err
    // }
    //
    // if !expense.CanAddSettlement() {
    //     return fmt.Errorf("cannnot add settlement to settled expense")
    // }

    settlement = model.NewSettlement(amount, userId)
    err = c.settlementDataStore.AddSettlement(settlement)
    return
}

func (c SettlementController) GetSettlement(settlementId string) (*model.Settlement,error) {
    return c.settlementDataStore.GetSettlement(settlementId)
}


