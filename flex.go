package golazada

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type FlexString string

func (f *FlexString) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		*f = ""
		return nil
	}
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		*f = FlexString(s)
		return nil
	}
	var n json.Number
	if err := json.Unmarshal(data, &n); err == nil {
		*f = FlexString(n.String())
		return nil
	}
	var b bool
	if err := json.Unmarshal(data, &b); err == nil {
		*f = FlexString(strconv.FormatBool(b))
		return nil
	}
	return fmt.Errorf("cannot unmarshal %s into FlexString", data)
}

type FlexFloat float64

func (f *FlexFloat) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" || string(data) == `""` {
		*f = 0
		return nil
	}
	var n json.Number
	if err := json.Unmarshal(data, &n); err == nil {
		v, err := n.Float64()
		if err != nil {
			return err
		}
		*f = FlexFloat(v)
		return nil
	}
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("cannot unmarshal %s into FlexFloat", data)
	}
	s = strings.TrimSpace(s)
	if s == "" {
		*f = 0
		return nil
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	*f = FlexFloat(v)
	return nil
}

type FlexInt int64

func (i *FlexInt) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" || string(data) == `""` {
		*i = 0
		return nil
	}
	var n json.Number
	if err := json.Unmarshal(data, &n); err == nil {
		v, err := n.Int64()
		if err != nil {
			return err
		}
		*i = FlexInt(v)
		return nil
	}
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("cannot unmarshal %s into FlexInt", data)
	}
	s = strings.TrimSpace(s)
	if s == "" {
		*i = 0
		return nil
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*i = FlexInt(v)
	return nil
}
