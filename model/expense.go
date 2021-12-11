package model

type ExpenseStatus int
type SettlementType int

const (
    SettlementTypeExact SettlementType = iota
    SettlementTypePercentage
)


const (
    ExpenseStatusInitiated ExpenseStatus = iota
    ExpenseStatusInSettlement
    ExpenseStatusSettled
)

type Expense struct {
    *BaseModel
    activity Activity
    createdBy *User
    totalAmount float64
    settlements []*Settlement
    status ExpenseStatus
    settlementType SettlementType
}

func (e *Expense) SettlementType() SettlementType {
    return e.settlementType
}


func (e *Expense) AddSettlements(settlement *Settlement) {
    e.settlements = append(e.settlements, settlement)
}

func (e *Expense) UpdateStatus(status ExpenseStatus) {
    e.status = status
}

func (e Expense) Activity() Activity {
    return e.activity
}

func (e Expense) CreatedBy() *User {
    return e.createdBy
}

func (e Expense) TotalAmount() float64 {
    return e.totalAmount
}

func (e Expense) Settlements() []*Settlement {
    return e.settlements
}

func (e Expense) Status() ExpenseStatus {
    return e.status
}

func (e Expense) CanAddSettlement() bool {
    return e.status != ExpenseStatusSettled
}

func NewExpense(activity Activity, createdBy *User, totalAmount float64, settlementType SettlementType) *Expense {
    return &Expense{
        BaseModel: NewBaseModel(),
        activity: activity,
        createdBy: createdBy,
        totalAmount: totalAmount,
        status: ExpenseStatusInitiated,
        settlementType: settlementType,
    }
}