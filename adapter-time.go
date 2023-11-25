package model

type TimeAdapter interface {
	TimeNow
	TimeWeek
	UnixTimeHandler

	//esperar en milliseconds ej: 500, 2000 ..
	WaitFor(milliseconds int, callback func())

	// new_date ej: 2003-11-22 o "" para volver al estado normal
	SetDate(new_date string)
}

type TimeNow interface {
	// layout ej: 2006-01-02
	ToDay(layout string) string
}

type TimeWeek interface {
	//date ej: 2006-01-02,  0 para Domingo, 1 para Lunes, etc.
	WeekDayNumber(date_in string) (n int, err string)
}

// ej: time.Now()
type UnixTimeHandler interface {
	UnixNano() int64
}
