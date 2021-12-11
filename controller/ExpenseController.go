package controller

import (
    "fmt"

    "../datastore"
    "../model"
)

type ExpenseController struct {
    expenseDataStore datastore.ExpenseDataStore
    userController *UserController
    settlementController *SettlementController
}

func NewExpenseController(expenseDataStore datastore.ExpenseDataStore, userController *UserController, settlementController *SettlementController) *ExpenseController {
    return &ExpenseController{
        expenseDataStore: expenseDataStore,
        userController: userController,
        settlementController: settlementController,
    }
}

func (c *ExpenseController) CreateExpense(amount float64, userId string, settlementType model.SettlementType, activity model.Activity) (expense *model.Expense, err error) {

    // expense, err := c.expenseDataStore.GetExpense(expenseId)
    // if err != nil {
    //     return err
    // }
    // 
    // if !expense.CanAddExpense() {
    //     return fmt.Errorf("cannnot add expense to settled expense")
    // }

    user, err := c.userController.GetUser(userId)
    if err != nil {
        return nil, err
    }
    expense = model.NewExpense(activity, user, amount, settlementType)
    err =  c.expenseDataStore.AddExpense(expense)
    return expense, err
}

func (c *ExpenseController) AddSettlementToExpense(expenseId string, userId string, value float64) error {
    expense, err := c.expenseDataStore.GetExpense(expenseId)
    if err != nil {
        return err
    }

    var existingAmount float64

    for _, settlement := range expense.Settlements() {
        existingAmount += settlement.Amount()
    }

    if expense.SettlementType() == model.SettlementTypeExact {
        if existingAmount + value > expense.TotalAmount() {
            return fmt.Errorf("cannot add more than total amount")
        }
        settlement, err := c.settlementController.CreateSettlement(value, userId)
        if err != nil {
            return err
        }
        expense.AddSettlements(settlement)
    } else {

        value := value * expense.TotalAmount() / 100
        if value > expense.TotalAmount() {
            return fmt.Errorf("cannot add more than total amount")
        }
        settlement, err := c.settlementController.CreateSettlement(value, userId)
        if err != nil {
            return err
        }
        expense.AddSettlements(settlement)
    }

    return nil
}


func (c *ExpenseController) CheckPendingBalance(expenseId string) (float64, error) {
    expense, err := c.expenseDataStore.GetExpense(expenseId)
    if err != nil {
        return float64(0), err
    }
    amount := float64(0)
    for _, settlement := range expense.Settlements() {
        if settlement.Status() == model.SettlementUnSettled {
            fmt.Printf("expense - %s user - %s pending amount - %f \n", expenseId, settlement.UserId(), settlement.Amount())
            amount += settlement.Amount()
        }
    }
    return amount, nil
}


func (c ExpenseController) GetExpense(expenseId string) (*model.Expense,error) {
    return c.expenseDataStore.GetExpense(expenseId)
}


func (c ExpenseController) SettleExpense(userId string, amount float64, expenseId string) error  {
    expense, err := c.GetExpense(expenseId)
    if err != nil {
        return err
    }
    var userSettlement *model.Settlement

    for _, settlement := range expense.Settlements() {
        if settlement.Status() == model.SettlementUnSettled && settlement.UserId() == userId {
            userSettlement = settlement
            break
        }
    }

    if userSettlement == nil {
        return fmt.Errorf("user is not part of expense or expense is already settled")
    }

    if userSettlement.Amount() != amount {
        return fmt.Errorf("amount is different. Please pay the exact amount")
    }

    userSettlement.SetStatus(model.SettlementSettled)
    return nil
}
