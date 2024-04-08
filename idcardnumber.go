package idcardnumber

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/bluesky335/IDCheck/IDNumber"
)

type IDCardNumber struct {
	number IDNumber.IDNumber
}

func (i IDCardNumber) Valid() bool {
	return i.number.IsValid()
}

func (i IDCardNumber) Birthday() string {
	b := i.number.GetBirthday()

	return fmt.Sprintf("%s-%s-%s", b.Year, b.Month, b.Day)
}

func (i IDCardNumber) Age() (int, error) {
	b := i.number.GetBirthday()

	year, err := strconv.Atoi(b.Year)
	if err != nil {
		return 0, errors.New("出生年份不正确")
	}

	month, err := strconv.Atoi(b.Month)
	if err != nil {
		return 0, errors.New("出生月份不正确")
	}

	day, err := strconv.Atoi(b.Day)
	if err != nil {
		return 0, errors.New("出生日期不正确")
	}

	age := CalculateAgeFromBirthday(year, month, day)

	return age, nil
}

// CalculateAgeFromBirthday 根据出生日期计算年龄
func CalculateAgeFromBirthday(year, month, day int) int {
	now := time.Now()

	age := now.Year() - year

	if now.Month() < time.Month(month) || (now.Month() == time.Month(month) && now.Day() < day) {
		age--
	}

	return age
}

// ParseAgeFromBirthday 根据出生日期计算年龄
// birthdayLayout: 出生日期格式, 如: "2006-01-02"
// s: 出生日期
func ParseAgeFromBirthday(birthdayLayout, s string) (int, error) {
	t, err := time.Parse(birthdayLayout, s)
	if err != nil {
		return 0, errors.New("出生日期格式不正确")
	}

	age := CalculateAgeFromBirthday(t.Year(), int(t.Month()), t.Day())

	return age, nil
}

func New(number string) *IDCardNumber {
	return &IDCardNumber{
		number: IDNumber.New(number),
	}
}

// Validator 验证身份证号码
// number: 身份证号码
func Validator(number string) error {
	if !New(number).Valid() {
		return errors.New("身份证号码不合法")
	}

	return nil
}
