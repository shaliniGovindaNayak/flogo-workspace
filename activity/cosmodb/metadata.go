package cosmodb

import "github.com/project-flogo/core/data/coerce"

type Input struct {
	database     string `md:"database,required"`
	username    string `md:"username,required"`
	password    string `md:"password,required"`
	url 		string `md:"url,required"`
	data		string `md:"data`
}

type Output struct {
	Output string `md:"Output"`
}

func (i *Input) FromMap(values map[string]interface{}) error {
	//var err error
	i.database, _ = coerce.ToString(values["database"])
	i.username, _ = coerce.ToString(values["username"])
	i.password, _ = coerce.ToString(values["password"])
	i.url, _ = coerce.ToString(values["url"])
	i.data, _ = coerce.ToString(values["data"])

	return nil
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"database":     i.database,
		"username":    i.username,
		"password":    i.password,
		"url": i.url,
		"data": i.data,
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
