package devicetwinoperations

import (
    "fmt"
)

func main(){
	json := map[string]interface{}{
		"deviceId": "f4:d1:08:47:8b:01",
		"etag": "AAAAAAAAABs=",
		"deviceEtag": "ODgzOTc2NzE2",
		"status": "enabled",
		"statusUpdateTime": "0001-01-01T00:00:00Z",
		"connectionState": "Disconnected",
		"lastActivityTime": "0001-01-01T00:00:00Z",
		"cloudToDeviceMessageCount": 0,
		"authenticationType": "sas",
		"x509Thumbprint": {
		  "primaryThumbprint": null,
		  "secondaryThumbprint": null
		},
		"version": 28,
		"properties": {
		  "desired": {
			"high": 65,
			"cpu_family": 6,
			"Cpu_usage": "98%",
			"CPU_index_number": "0",
			"low": 23,
			"Cpu_Speed": "2900.00",
			"Cpu_VendorID": "GenuineIntel",
			"Free_disk_space": "182391029760",
			"Free_memory": "341680128",
			"Host_ID": "74619e31-ba1c-45c9-9473-c4cc05c0b558",
			"Ip_address": "192.168.43.129/24",
			"Number_of_processes_running": "316",
			"Operating_system": "linux",
			"Percentage_disk_space_usage": "23.41",
			"Percentage_used_memory": "62.43",
			"System_Uptime": "228237",
			"Time_Stamp": "2019-12-09 11:14:05",
			"Total_disk_space": "250966470656",
			"Total_memory": "250966470656",
			"Used_disk_space": "55755710464",
			"Username": "Shalu",
			"current_cpu_utilization": "[0.00,33.33,33.33,33.33]",
			"mac_address": "28:b2:bd:01:d0:24",
			"$metadata": {
			  "$lastUpdated": "2019-12-09T06:07:29.4346749Z",
			  "$lastUpdatedVersion": 27,
			  "high": {
				"$lastUpdated": "2019-12-09T06:06:44.4186182Z",
				"$lastUpdatedVersion": 25
			  },
			  "cpu_family": {
				"$lastUpdated": "2019-12-09T06:06:44.4186182Z",
				"$lastUpdatedVersion": 25
			  },
			  "Cpu_usage": {
				"$lastUpdated": "2019-12-09T06:06:44.4186182Z",
				"$lastUpdatedVersion": 25
			  },
			  "CPU_index_number": {
				"$lastUpdated": "2019-12-09T06:06:44.4186182Z",
				"$lastUpdatedVersion": 25
			  },
			  "low": {
				"$lastUpdated": "2019-12-09T06:06:44.4186182Z",
				"$lastUpdatedVersion": 25
			  },
			  "Cpu_Speed": {
				"$lastUpdated": "2019-12-09T06:06:44.4186182Z",
				"$lastUpdatedVersion": 25
			  },
			  "Cpu_VendorID": {
				"$lastUpdated": "2019-12-09T06:06:44.4186182Z",
				"$lastUpdatedVersion": 25
			  },
			  "Free_disk_space": {
				"$lastUpdated": "2019-12-09T06:06:44.4186182Z",
				"$lastUpdatedVersion": 25
			  },
			  "Free_memory": {
				"$lastUpdated": "2019-12-09T06:06:44.4186182Z",
				"$lastUpdatedVersion": 25
			  },
			  "Host_ID": {
				"$lastUpdated": "2019-12-09T06:06:44.4186182Z",
				"$lastUpdatedVersion": 25
			  },
			  "Ip_address": {
				"$lastUpdated": "2019-12-09T06:06:44.4186182Z",
				"$lastUpdatedVersion": 25
			  },
			  "Number_of_processes_running": {
				"$lastUpdated": "2019-12-09T06:06:44.4186182Z",
				"$lastUpdatedVersion": 25
			  },
			  "Operating_system": {
				"$lastUpdated": "2019-12-09T06:06:44.4186182Z",
				"$lastUpdatedVersion": 25
			  },
			  "Percentage_disk_space_usage": {
				"$lastUpdated": "2019-12-09T06:06:44.4186182Z",
				"$lastUpdatedVersion": 25
			  },
			  "Percentage_used_memory": {
				"$lastUpdated": "2019-12-09T06:06:44.4186182Z",
				"$lastUpdatedVersion": 25
			  },
			  "System_Uptime": {
				"$lastUpdated": "2019-12-09T06:06:44.4186182Z",
				"$lastUpdatedVersion": 25
			  },
			  "Time_Stamp": {
				"$lastUpdated": "2019-12-09T06:06:44.4186182Z",
				"$lastUpdatedVersion": 25
			  },
			  "Total_disk_space": {
				"$lastUpdated": "2019-12-09T06:06:44.4186182Z",
				"$lastUpdatedVersion": 25
			  },
			  "Total_memory": {
				"$lastUpdated": "2019-12-09T06:06:44.4186182Z",
				"$lastUpdatedVersion": 25
			  },
			  "Used_disk_space": {
				"$lastUpdated": "2019-12-09T06:06:44.4186182Z",
				"$lastUpdatedVersion": 25
			  },
			  "Username": {
				"$lastUpdated": "2019-12-09T06:07:29.4346749Z",
				"$lastUpdatedVersion": 27
			  },
			  "current_cpu_utilization": {
				"$lastUpdated": "2019-12-09T06:07:29.4346749Z",
				"$lastUpdatedVersion": 27
			  },
			  "mac_address": {
				"$lastUpdated": "2019-12-09T06:07:29.4346749Z",
				"$lastUpdatedVersion": 27
			  }
			},
			"$version": 27
		  },
		  "reported": {
			"$metadata": {
			  "$lastUpdated": "2019-12-06T15:03:00.2604035Z"
			},
			"$version": 1
		  }
		},
		"capabilities": {
		  "iotEdge": false
		}
	  }

	  str := fmt.Sprintf("%v",json)
	  fmt.Println(str)
}