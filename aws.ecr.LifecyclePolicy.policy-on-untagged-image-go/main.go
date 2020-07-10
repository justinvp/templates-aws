package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ecr"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		foo, err := ecr.NewRepository(ctx, "foo", nil)
		if err != nil {
			return err
		}
		_, err = ecr.NewLifecyclePolicy(ctx, "foopolicy", &ecr.LifecyclePolicyArgs{
			Policy:     pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "    \"rules\": [\n", "        {\n", "            \"rulePriority\": 1,\n", "            \"description\": \"Expire images older than 14 days\",\n", "            \"selection\": {\n", "                \"tagStatus\": \"untagged\",\n", "                \"countType\": \"sinceImagePushed\",\n", "                \"countUnit\": \"days\",\n", "                \"countNumber\": 14\n", "            },\n", "            \"action\": {\n", "                \"type\": \"expire\"\n", "            }\n", "        }\n", "    ]\n", "}\n", "\n")),
			Repository: foo.Name,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

