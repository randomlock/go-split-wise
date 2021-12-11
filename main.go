package main

import (
    "fmt"

    "./controller"
    "./datastore"
    "./model"
)

func main() {

    userDataStore := datastore.NewInMemoryUserDataStore()
    expenseDataStore := datastore.NewInMemoryExpenseDataStore()
    settlementDataStore := datastore.NewInMemorySettlementDataStore()

    userController :=  controller.NewUserController(userDataStore)
    settlementController := controller.NewSettlementController(settlementDataStore)
    expenseController := controller.NewExpenseController(expenseDataStore, userController, settlementController)

    user1, _ := userController.AddUser("foo", "222222")
    user2, _ := userController.AddUser("bar", "111111")
    user3, _ := userController.AddUser("var", "11111")
    expense1, err := expenseController.CreateExpense(100, user1.Id(), model.SettlementTypeExact, *model.NewActivity(model.ActivityTypeFood, "burger"))

    if err != nil {
        println(err.Error())
        return
    }
    expenseController.AddSettlementToExpense(expense1.Id(), user2.Id(), 50)
    expenseController.AddSettlementToExpense(expense1.Id(), user3.Id(), 50)
    err = expenseController.AddSettlementToExpense(expense1.Id(), user3.Id(), 24)
    if err != nil {
        fmt.Println(err.Error())
    }
    expenseController.CheckPendingBalance(expense1.Id())

    expense2, err := expenseController.CreateExpense(200, user1.Id(), model.SettlementTypePercentage, *model.NewActivity(model.ActivityTypeFood, "pizza"))

    if err != nil {
        println(err.Error())
        return
    }
    expenseController.AddSettlementToExpense(expense2.Id(), user2.Id(), 20)
    expenseController.AddSettlementToExpense(expense2.Id(), user3.Id(), 20)
    err = expenseController.AddSettlementToExpense(expense2.Id(), user3.Id(), 60)
    if err != nil {
        fmt.Println(err.Error())
    }


    expenseController.CheckPendingBalance(expense2.Id())
}