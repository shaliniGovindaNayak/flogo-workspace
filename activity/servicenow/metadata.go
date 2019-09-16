package servicenow

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	Username    string `md:"Username"`
	Password    string `md:"Password"`
	Instanceurl string `md:"Instanceurl`
}

type Input struct {
	Content string `md:"content"`
}

type Output struct {
	Output string `md:"output"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	//var err error
	r.Content, _ = coerce.ToString(values["content"])

	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"content": r.Content,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {
	o.Output, _ = coerce.ToString(values["output"])

	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"output": o.Output,
	}
}
