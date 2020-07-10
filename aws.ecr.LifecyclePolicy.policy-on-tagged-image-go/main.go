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
			Policy:     pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "    \"rules\": [\n", "        {\n", "            \"rulePriority\": 1,\n", "            \"description\": \"Keep last 30 images\",\n", "            \"selection\": {\n", "                \"tagStatus\": \"tagged\",\n", "                \"tagPrefixList\": [\"v\"],\n", "                \"countType\": \"imageCountMoreThan\",\n", "                \"countNumber\": 30\n", "            },\n", "            \"action\": {\n", "                \"type\": \"expire\"\n", "            }\n", "        }\n", "    ]\n", "}\n", "\n")),
			Repository: foo.Name,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

