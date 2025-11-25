package calendar

import "errors"

type Date struct {
	Year  int
	Month int
	Day   int
}

func (d *Date) RtYear() int {
	return d.Year
}
func (d *Date) RtMonth() int {
	return d.Month
}
func (d *Date) RtDay() int {
	return d.Day
}

//구조체 속 자료형 이름이 첫 글자(대소문)로 지랄이고, 함수명이랑 겹치기까지하면서 난항이었음..
//+함수명 첫 글자(또소문 씨발)

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.Year = year
	return nil
}
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.Month = month
	return nil
}
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.Day = day
	return nil
}
