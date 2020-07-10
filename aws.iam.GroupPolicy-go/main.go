package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		myDevelopers, err := iam.NewGroup(ctx, "myDevelopers", &iam.GroupArgs{
			Path: pulumi.String("/users/"),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewGroupPolicy(ctx, "myDeveloperPolicy", &iam.GroupPolicyArgs{
			Group:  myDevelopers.ID(),
			Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": [\n", "        \"ec2:Describe*\"\n", "      ],\n", "      \"Effect\": \"Allow\",\n", "      \"Resource\": \"*\"\n", "    }\n", "  ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

