package utils

import (
	"database/sql/driver"
	"errors"
	"github.com/json-iterator/go"
)

type GormString string

func (static GormString) Value() (driver.Value, error) {
	return string(static), nil
}

func (this *GormString) Scan(v interface{}) error {
	val, flag := v.(string)
	if !flag {
		return errors.New("not string")
	}

	*this = GormString(val)
	return nil
}

type GormInt int

func (static GormInt) Value() (driver.Value, error) {
	return int(static), nil
}

func (this *GormInt) Scan(v interface{}) error {
	val, flag := v.(int)
	if !flag {
		return errors.New("not string")
	}

	*this = GormInt(val)
	return nil
}

type GormInt64 int64

func (static GormInt64) Value() (driver.Value, error) {
	return int64(static), nil
}

func (this *GormInt64) Scan(v interface{}) error {
	val, flag := v.(int64)
	if !flag {
		return errors.New("not string")
	}

	*this = GormInt64(val)
	return nil
}

type GormFloat64 float64

func (static GormFloat64) Value() (driver.Value, error) {
	return float64(static), nil
}

func (this *GormFloat64) Scan(v interface{}) error {
	val, flag := v.(float64)
	if !flag {
		return errors.New("not string")
	}

	*this = GormFloat64(val)
	return nil
}

type GormStrings []string

func (static GormStrings) Value() (driver.Value, error) {
	return jsoniter.Marshal(static)
}

func (this *GormStrings) Scan(v interface{}) error {
	var strs []string
	if err := jsoniter.Unmarshal(v.([]byte), &strs); err != nil {
		return err
	}

	*this = strs
	return nil
}

type GormInts []int

func (static GormInts) Value() (driver.Value, error) {
	return jsoniter.Marshal(static)
}

func (this *GormInts) Scan(v interface{}) error {
	var strs []int
	if err := jsoniter.Unmarshal(v.([]byte), &strs); err != nil {
		return err
	}

	*this = strs
	return nil
}

type GormInt64s []int64

func (static GormInt64s) Value() (driver.Value, error) {
	return jsoniter.Marshal(static)
}

func (this *GormInt64s) Scan(v interface{}) error {
	var strs []int64
	if err := jsoniter.Unmarshal(v.([]byte), &strs); err != nil {
		return err
	}

	*this = strs
	return nil
}

type GormFloat64s []float64

func (static GormFloat64s) Value() (driver.Value, error) {
	return jsoniter.Marshal(static)
}

func (this *GormFloat64s) Scan(v interface{}) error {
	var strs []float64
	if err := jsoniter.Unmarshal(v.([]byte), &strs); err != nil {
		return err
	}

	*this = strs
	return nil
}

type GormMap map[string]interface{}

func (static GormMap) Value() (driver.Value, error) {
	return jsoniter.Marshal(static)
}

func (this *GormMap) Scan(v interface{}) error {
	var strs map[string]interface{}
	if err := jsoniter.Unmarshal(v.([]byte), &strs); err != nil {
		return err
	}

	*this = strs
	return nil
}
