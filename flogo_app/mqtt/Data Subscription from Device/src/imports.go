package main

import (
	_ "github.com/shaliniGovindaNayak/flogo-workspace/activity/stringtojson"
	_ "github.com/project-flogo/legacybridge"
	_ "github.com/retgits/flogo-components/activity/dynamodbinsert"
	_ "github.com/pradyuz3rocool/flogo-workspace/activity/sendazureiot"
	_ "github.com/TIBCOSoftware/flogo-contrib/activity/rest"
	_ "github.com/TIBCOSoftware/flogo-contrib/activity/log"
	_ "github.com/TIBCOSoftware/flogo-contrib/action/flow"
	_ "github.com/TIBCOSoftware/flogo-contrib/trigger/mqtt"
)
