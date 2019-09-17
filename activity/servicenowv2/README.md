

# 	recieveazure - Activity
This activity provides your Flogo app the ability to recieve message from Azure Iot Hub from a device

## Installation

```bash
flogo install github.com/shaliniGovindaNayak/flogo-workspace/activity/servicenowv2
```
Link for flogo web:
```
github.com/shaliniGovindaNayak/flogo-workspace/activity/servicenowv2
```

## Schema
Inputs and Outputs:

```json
"inputs":[
    {
      "name": "Instance url",
      "type": "string",
      "required": true
    },
    {
      "name":"Username",
      "type":"string",
      "required":true
    },
    {
      "name":"password",
      "type":"string",
      "required":true
    },
    {
      "name":"content",
      "type":"Object",
      "required":true
    }
  ]
```
## Inputs
| Input                          | Description    |
|:-------------------------------|:---------------|
| Instance URL                   | Your Service now instance url.            |
| Username                       | Login username   |
| Password                       | password         |
| content                        | json object containing the insident details |


## Ouputs
| Output       | Description                                            |
|:-------------|:-------------------------------------------------------|
| Output       | The sucess message indicating the insident creation |
