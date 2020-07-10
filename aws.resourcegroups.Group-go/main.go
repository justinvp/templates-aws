package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/resourcegroups"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := resourcegroups.NewGroup(ctx, "test", &resourcegroups.GroupArgs{
			ResourceQuery: &resourcegroups.GroupResourceQueryArgs{
				Query: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"ResourceTypeFilters\": [\n", "    \"AWS::EC2::Instance\"\n", "  ],\n", "  \"TagFilters\": [\n", "    {\n", "      \"Key\": \"Stage\",\n", "      \"Values\": [\"Test\"]\n", "    }\n", "  ]\n", "}\n", "\n")),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

