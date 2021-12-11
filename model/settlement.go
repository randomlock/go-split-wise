package model

type SettlementStatus int


const (
    SettlementSettled SettlementStatus = iota
    SettlementUnSettled
)

type Settlement struct {
    *BaseModel
    status SettlementStatus
    amount float64
    userId string
    expenseId string
}

func (s *Settlement) UserId() string {
    return s.userId
}

func (s *Settlement) ExpenseId() string {
    return s.expenseId
}

func (s *Settlement) SetStatus(status SettlementStatus) {
    s.status = status
}

func (s Settlement) Status() SettlementStatus {
    return s.status
}

func (s Settlement) Amount() float64 {
    return s.amount
}

func NewSettlement(amount float64, userId string) *Settlement {
    return &Settlement{
        amount: amount,
        BaseModel: NewBaseModel(),
        status: SettlementUnSettled,
        userId: userId,
    }
}

