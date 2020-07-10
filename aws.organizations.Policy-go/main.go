package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/organizations"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := organizations.NewPolicy(ctx, "example", &organizations.PolicyArgs{
			Content: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": {\n", "    \"Effect\": \"Allow\",\n", "    \"Action\": \"*\",\n", "    \"Resource\": \"*\"\n", "  }\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

