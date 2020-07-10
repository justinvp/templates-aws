package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssm.NewDocument(ctx, "foo", &ssm.DocumentArgs{
			Content:      pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "  {\n", "    \"schemaVersion\": \"1.2\",\n", "    \"description\": \"Check ip configuration of a Linux instance.\",\n", "    \"parameters\": {\n", "\n", "    },\n", "    \"runtimeConfig\": {\n", "      \"aws:runShellScript\": {\n", "        \"properties\": [\n", "          {\n", "            \"id\": \"0.aws:runShellScript\",\n", "            \"runCommand\": [\"ifconfig\"]\n", "          }\n", "        ]\n", "      }\n", "    }\n", "  }\n", "\n")),
			DocumentType: pulumi.String("Command"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

