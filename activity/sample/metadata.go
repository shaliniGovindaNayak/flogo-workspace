package influxdb

import "github.com/project-flogo/core/data/coerce"

type Input struct {
	AnInput string `md:"anInput,required"`
}

func (s *Settings) FromMap(values map[string]interface{}) error {
	val, _ := coerce.ToString(values["aSetting"])
	s.ASetting = val
	return nil
}

func (s *Settings) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"aSettings": s.ASetting,
	}
}

func (r *Input) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["anInput"])
	r.AnInput = strVal
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"anInput": r.AnInput,
	}
}

type Output struct {
	AnOutput string `md:"anOutput"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["anOutput"])
	o.AnOutput = strVal
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"anOutput": o.AnOutput,
	}
}
