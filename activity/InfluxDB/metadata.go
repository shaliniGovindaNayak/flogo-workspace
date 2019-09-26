package influxdb

import "github.com/project-flogo/core/data/coerce"

type Input struct {
	Host     string `md:"host,required"`
	Schema   string `md:"schema"`
	Table    string `md:"table"`
	Username string `md:"username"`
	Password string `md:"password"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	hval, _ := coerce.ToString(values["host"])
	r.Host = hval

	sval, _ := coerce.ToString(values("schema"))
	r.Schema = sval

	tval, _ := coerce.ToString(values("table"))
	r.Table = tval

	uval, _ := coerce.ToString(values("username"))
	r.Username = uval

	psval := coerce.ToString(values["password"])
	r.Password = psval
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"host":     r.Host,
		"schema":   r.Schema,
		"username": r.Username,
		"password": r.Password,
		"table":    r.Table,
	}
}

type Output struct {
	Output string `md:"Output"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["Output"])
	o.AnOutput = strVal
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Output": o.Output,
	}
}
