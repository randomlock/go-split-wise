package model

type ActivityType int

const (
    ActivityTypeTrip ActivityType = iota
    ActivityTypeShopping
    ActivityTypeFood
)

type Activity struct {
    *BaseModel
    activityType ActivityType
    name string
}

func (a Activity) ActivityType() ActivityType {
    return a.activityType
}

func (a Activity) Name() string {
    return a.name
}

func NewActivity(activityType ActivityType, name string) *Activity {
    return &Activity{
        activityType: activityType,
        name: name,
        BaseModel: NewBaseModel(),
    }
}
