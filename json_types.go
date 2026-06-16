package datev_api

import (
	"encoding/json"
	"log"
	"strconv"
	"time"
)

type Date struct {
	time.Time
}

func (d Date) MarshalSchema() string {
	return d.Time.Format("2006-01-02")
}

func (d Date) IsEmpty() bool {
	return d.Time.IsZero()
}

func (d *Date) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(d.Time.Format("2006-01-02"))
}

func (d *Date) UnmarshalJSON(text []byte) (err error) {
	var value string
	err = json.Unmarshal(text, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	d.Time, err = time.Parse(time.RFC3339, value)
	if err == nil {
		return nil
	}

	d.Time, err = time.Parse("2006-01-02", value)
	if err == nil {
		return nil
	}

	d.Time, err = time.Parse("2006-01-02T15:04:05", value)
	log.Println(d.Time)
	return err
}

type IntString string

func (f *IntString) UnmarshalJSON(text []byte) (err error) {
	var str string
	err = json.Unmarshal(text, &str)
	if err == nil {
		*f = IntString(str)
		return err
	}

	// error, so try int
	var integer int
	err = json.Unmarshal(text, &integer)
	if err != nil {
		return err
	}

	str = strconv.Itoa(integer)
	*f = IntString(str)
	return nil
}

type StringInt int

func (f *StringInt) UnmarshalJSON(text []byte) (err error) {
	var integer int
	err = json.Unmarshal(text, &integer)
	if err == nil {
		*f = StringInt(integer)
		return err
	}

	// error, so try string
	var str string
	err = json.Unmarshal(text, &str)
	if err != nil {
		return err
	}

	integer, err = strconv.Atoi(str)
	if err != nil {
		return err
	}

	*f = StringInt(integer)
	return nil
}
