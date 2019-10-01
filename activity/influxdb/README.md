
# 	InfluxDB - Activity
This activity provides your Flogo app the ability to insert data on InfluxDB

## Installation

```bash
flogo install github.com/shaliniGovindaNayak/flogo-workspace/activity/influxdb
```
Link for flogo web:
```
github.com/shaliniGovindaNayak/flogo-workspace/activity/influxdb
```

## Schema
Inputs and Outputs:

```json
"inputs":[
    {
      "name": "Host",
      "type": "string",
      "required": true
    },
    {
      "name":"Schema",
      "type":"string",
      "required":true
    },
    {
      "name":"Table",
      "type":"string"
    },
    {
       "name":"Value",
       "type":"any"
    }
  ]
```
## Inputs
| Input                          | Description    |
|:-------------------------------|:---------------|
| Host                           | Host name along with port where influxDB is running.           |
| Schema                         | Schema(database name) that You want to use   |
| Table                          | Table that needs to be inserted        |
| Value                          | Data that needs to be inserted |


## Ouputs
| Output       | Description                                            |
|:-------------|:-------------------------------------------------------|
| Output       | The sucess message indicating the insertion |
