package influxdb

import (
	"github.com/project-flogo/core/data/coerce"
)

type Input struct {
	Host string `md:"host.required`
	Schema string `md:"schema.required"`
	Table string `md:"table.required"`
	Values map[string]interface{} `md:"values"`
}

func (r *Input) FromMap(values map[string]interface{}) error {

	Val1, _ := coerce.ToString(values["host"])
	r.Host = Val1

	Val2, _ := coerce.ToString(values["schema"])
	r.Schema = Val2

	Val3, _ := coerce.ToString(values["table"])
	r.Table = Val3

	Val4, _ := coerce.ToParams(values["values"])
	r.Values = Val4
	
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"host":r.Host,
		"schema":r.Schema,
		"table":r.Table,
		"values":r.Values,
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
