
# 	ServiceNow - Activity
This activity provides your Flogo app the ability to raise an incident in service now 

## Installation

```bash
flogo install github.com/shaliniGovindaNayak/flogo-workspace/activity/servicenow
```
Link for flogo web:
```
github.com/shaliniGovindaNayak/flogo-workspace/activity/servicenow
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
      "name":"insident value",
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
| insidence value                | json object containing the insident details |


## Ouputs
| Output       | Description                                            |
|:-------------|:-------------------------------------------------------|
| Output       | The sucess message indicating the insident creation |
