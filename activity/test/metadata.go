package test

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	Method string `md:method,required`
}

type Input struct {
	Num1 string `md:"num1,required"`
	Num2 string `md:"num2,required"`
}

type Output struct {
	Output string `md:"Output"`
}

func (i *Input) FromMap(values map[string]interface{}) error {
	//var err error
	i.Num1, _ = coerce.ToString(values["num1"])
	i.Num1, _ = coerce.ToString(values["num2"])
	return nil
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Num1": i.Num1,
		"Num2": i.Num2,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {
	o.Output, _ = coerce.ToString(values["Output"])

	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Output": o.Output,
	}
}
