package influxdb

import "github.com/project-flogo/core/data/coerce"

type Input struct {
	Host   string `md:"host,required"`
	Schema string `md:"schema"`
	Table  string `md:"table"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	hval, _ := coerce.ToString(values["host"])
	r.Host = hval

	sval, _ := coerce.ToString(values["schema"])
	r.Schema = sval

	tval, _ := coerce.ToString(values["table"])
	r.Table = tval

	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"host":   r.Host,
		"schema": r.Schema,
		"table":  r.Table,
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
