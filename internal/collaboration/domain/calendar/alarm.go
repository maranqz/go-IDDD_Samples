package calendar

type AlarmUnitsType int

// +1 to move default value and for var RepeatType be incorrect value.
const (
	Days AlarmUnitsType = iota + 1
	Hours
	Minutes
)

type Alarm struct {
	alarmUnits     int
	alarmUnitsType AlarmUnitsType
}

func NewAlarm(unitsType AlarmUnitsType, units int) (Alarm, error) {
	return Alarm{
		alarmUnits:     units,
		alarmUnitsType: unitsType,
	}, nil
}

func (a Alarm) AlarmUnits() int {
	return a.alarmUnits
}

func (a Alarm) AlarmUnitsType() AlarmUnitsType {
	return a.alarmUnitsType
}
