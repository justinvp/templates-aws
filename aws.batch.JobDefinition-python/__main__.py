import pulumi
import pulumi_aws as aws

test = aws.batch.JobDefinition("test",
    container_properties="""{
	"command": ["ls", "-la"],
	"image": "busybox",
	"memory": 1024,
	"vcpus": 1,
	"volumes": [
      {
        "host": {
          "sourcePath": "/tmp"
        },
        "name": "tmp"
      }
    ],
	"environment": [
		{"name": "VARNAME", "value": "VARVAL"}
	],
	"mountPoints": [
		{
          "sourceVolume": "tmp",
          "containerPath": "/tmp",
          "readOnly": false
        }
	],
    "ulimits": [
      {
        "hardLimit": 1024,
        "name": "nofile",
        "softLimit": 1024
      }
    ]
}

""",
    type="container")

