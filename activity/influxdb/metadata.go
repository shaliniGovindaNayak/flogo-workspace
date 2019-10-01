package influxdb

import (
	"fmt"

	"github.com/project-flogo/core/data/coerce"
	//"github.com/spf13/cast"
)

type Input struct {
	Host   string                 `md:"Host.required"`
	Schema string                 `md:"Schema.required"`
	Table  string                 `md:"Table"`
	Values map[string]interface{} `md:"Values"`
}

func (r *Input) FromMap(values map[string]interface{}) error {

	Val1, _ := coerce.ToString(values["Host"])
	r.Host = Val1
	fmt.Println(values["Host"])
	Val2, _ := coerce.ToString(values["Schema"])
	r.Schema = Val2

	Val3, _ := coerce.ToString(values["Table"])
	r.Table = Val3

	Val4, _ := coerce.ToObject("Values")
	//Val4, _ := coerce.ToParams(values["values"])
	r.Values = Val4
	fmt.Println(values["Values"])

	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Host":   r.Host,
		"Schema": r.Schema,
		"Table":  r.Table,
		"Values": r.Values,
	}
}

type Output struct {
	Output string `md:"Output"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["Output"])
	o.Output = strVal
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Output": o.Output,
	}
}
