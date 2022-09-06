package types

import (
	"time"
)

type Manager struct {
	ID        int       `json:"id"`
	EmpID     int       `json:"emp_id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Position struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	SalaryPerHour float64   `json:"salary_per_hr"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Employee struct {
	ID         int       `json:"id"`
	Firstname  string    `json:"firstname"`
	Lastname   string    `json:"lastname"`
	PositionID int       `json:"position_id"`
	ManagerID  int       `json:"manager_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Shift struct {
	ID          string    `json:"id"`
	EmpID       int       `json:"emp_id"`
	HoursWorked int       `json:"hours_worked"`
	ShiftDate   time.Time `json:"shift_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
