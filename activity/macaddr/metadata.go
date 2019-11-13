package macaddr

import "github.com/project-flogo/core/data/coerce"


type Output struct {
	Output string `md:"Output"`
	MacAddr string `md:mac_address`
	Battery1 string `md:battery1`
	Battery2 string `md:battery2`

}

func (o *Output) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["Output"])
	o.Output = strVal
	strVal1, _ := coerce.ToString(values["MacAddr"])
	o.MacAddr = strVal1
	strVal2, _ := coerce.ToString(values["Battery1"])
	o.Battery1 = strVal2
	strVal3, _ := coerce.ToString(values["Battery2"])
	o.Battery2 = strVal3
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Output": o.Output,
		"mac_address":o.MacAddr,
		"battery1":o.Battery1,
		"battery2":o.Battery2,
	}
}
