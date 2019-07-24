package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	criteria := Criteria{Age: Age(35)}
	membership := criteria.Membership(NewDayType(now), NewTimeType(now))
	fmt.Printf("+v, +v", membership, membership.TicketPrice())
}

type Membership interface {
	TicketPrice() Yen
}

type cinemaCitizen struct {
	dayType  DayType
	timeType TimeType
}

func (c cinemaCitizen) TicketPrice() Yen {
	switch {
	case c.dayType == Weekday || (c.dayType == Holiday && c.timeType == LateTime):
		return Yen(1000)
	case c.dayType == Movieday:
		return Yen(1100)
	default:
		return Yen(1300)
	}
}

type cinemaCitizenSenior struct {
}

func (cinemaCitizenSenior) TicketPrice() Yen {
	return Yen(1000)
}

type normal struct {
	dayType  DayType
	timeType TimeType
}

func (n normal) TicketPrice() Yen {
	switch {
	case n.dayType == Movieday:
		return Yen(1100)
	case n.timeType == LateTime:
		return Yen(1300)
	default:
		return Yen(1800)
	}
}

type senior struct {
}

func (senior) TicketPrice() Yen {
	return Yen(1000)
}

type seniorStudent struct {
	dayType  DayType
	timeType TimeType
}

func (s seniorStudent) TicketPrice() Yen {
	switch {
	case s.dayType == Movieday:
		return Yen(1100)
	case s.timeType == LateTime:
		return Yen(1300)
	default:
		return Yen(1500)
	}
}

type highStudent struct {
}

func (highStudent) TicketPrice() Yen {
	return Yen(1000)
}

type junior struct {
}

func (junior) TicketPrice() Yen {
	return Yen(1000)
}

type seniorHandicapped struct {
}

func (seniorHandicapped) TicketPrice() Yen {
	return Yen(1000)
}

type juniorHandicapped struct {
}

func (juniorHandicapped) TicketPrice() Yen {
	return Yen(900)
}

type Criteria struct {
	Age
	StudentType
	HasMemberCard            bool
	HasCertificate           bool
	HasStudentCertificate    bool
	HasDisabilityCertificate bool
	WithSeniorHandicapped    bool
	WithJuniorHandicapped    bool
}

func (c Criteria) Membership(dayType DayType, timeType TimeType) Membership {
	switch {
	case c.HasMemberCard && c.Age >= Age(60):
		return cinemaCitizenSenior{}
	case c.HasMemberCard && c.Age < Age(60):
		return cinemaCitizen{dayType: dayType, timeType: timeType}
	case (c.StudentType <= highStudentType && c.HasDisabilityCertificate) || c.WithJuniorHandicapped:
		return juniorHandicapped{}
	case (c.StudentType > highStudentType && c.HasDisabilityCertificate) || c.WithSeniorHandicapped:
		return seniorHandicapped{}
	case c.StudentType == lowStudentType || c.StudentType == juiorStudentType:
		return junior{}
	case (c.StudentType == juiorHighStudentType || c.StudentType == highStudentType) && c.HasStudentCertificate:
		return highStudent{}
	case c.StudentType == seniorStudentType && c.HasStudentCertificate:
		return seniorStudent{}
	case c.Age >= Age(70) && c.HasStudentCertificate:
		return senior{}
	default:
		return normal{dayType: dayType, timeType: timeType}
	}
}

type Age int

type StudentType int

const (
	lowStudentType       = StudentType(1)
	juiorStudentType     = StudentType(2)
	juiorHighStudentType = StudentType(3)
	highStudentType      = StudentType(4)
	seniorStudentType    = StudentType(5)
)

type description string
type DayType struct {
	description
}

var (
	Weekday  = DayType{"平日"}
	Holiday  = DayType{"土日祝"}
	Movieday = DayType{"映画の日"}
)

type TimeType struct {
	description
}

var (
	NormalTime = TimeType{"通常時間"}
	LateTime   = TimeType{"レイトショー"}
)

func NewDayType(t time.Time) DayType {
	// 一旦祝日考慮なし
	switch {
	case t.Day() == 1:
		return Movieday
	case t.Weekday() == time.Saturday:
		return Holiday
	case t.Weekday() == time.Sunday:
		return Holiday
	default:
		return Weekday
	}
}

func NewTimeType(t time.Time) TimeType {
	switch {
	case t.Hour() >= 20:
		return LateTime
	default:
		return NormalTime
	}
}

type Yen int
