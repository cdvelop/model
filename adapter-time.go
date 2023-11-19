package model

type TimeAdapter interface {
	TimeNow
	TimeWeek
	UnixTimeHandler
}

type TimeNow interface {
	// layout ej: 2006-01-02
	ToDay(layout string) string
}

type TimeWeek interface {
	//date ej: 2006-01-02,  0 para Domingo, 1 para Lunes, etc.
	WeekDayNumber(date_in string) (int, error)
}

// ej: time.Now()
type UnixTimeHandler interface {
	UnixNano() int64
}
