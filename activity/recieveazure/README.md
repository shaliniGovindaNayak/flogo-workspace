
# 	recieveazure - Activity

This activity provides your Flogo app the ability to recieve message to the Azure Iot Hub from a device

## Installation

```bash
flogo install github.com/shaliniGovindaNayak/flogo-workspace/activity/recieveazure
```
Link for flogo web:
```
https://github.com/shaliniGovindaNayak/flogo-workspace/activity/recieveazure
```

## Schema
Inputs and Outputs:

```json
"inputs":[
    {
      "name": "connectionString",
      "type": "string",
      "required": true
    },
    {
      "name": "Device ID",
      "type": "string"
    },
    {
      "name": "Action",
      "type": "string",
      "required": true,
      "allowed": ["recieve"]
    }
  ]
```
## Inputs
| Input                          | Description    |
|:-------------------------------|:---------------|
| Connection String              | Your Azure IOT Connection String.            |
| Device ID                      | Name of the Device  |
| Action                         | Recieve             |

## Ouputs
| Output       | Description                                            |
|:-------------|:-------------------------------------------------------|
| result       | A message from the device registered in Azure iot hub           |
| status       | The status of the request made                         |