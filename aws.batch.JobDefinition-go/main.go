package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/batch"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := batch.NewJobDefinition(ctx, "test", &batch.JobDefinitionArgs{
			ContainerProperties: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "	\"command\": [\"ls\", \"-la\"],\n", "	\"image\": \"busybox\",\n", "	\"memory\": 1024,\n", "	\"vcpus\": 1,\n", "	\"volumes\": [\n", "      {\n", "        \"host\": {\n", "          \"sourcePath\": \"/tmp\"\n", "        },\n", "        \"name\": \"tmp\"\n", "      }\n", "    ],\n", "	\"environment\": [\n", "		{\"name\": \"VARNAME\", \"value\": \"VARVAL\"}\n", "	],\n", "	\"mountPoints\": [\n", "		{\n", "          \"sourceVolume\": \"tmp\",\n", "          \"containerPath\": \"/tmp\",\n", "          \"readOnly\": false\n", "        }\n", "	],\n", "    \"ulimits\": [\n", "      {\n", "        \"hardLimit\": 1024,\n", "        \"name\": \"nofile\",\n", "        \"softLimit\": 1024\n", "      }\n", "    ]\n", "}\n", "\n")),
			Type: pulumi.String("container"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

