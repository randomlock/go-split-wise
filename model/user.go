package model

type User struct {
    *BaseModel
    name string
    phNo string
}

func (u User) Name() string {
    return u.name
}

func (u User) PhNo() string {
    return u.phNo
}

func NewUser(name string, phNo string) *User {
    return &User{
        BaseModel: NewBaseModel(),
        name: name,
        phNo: phNo,
    }
}
