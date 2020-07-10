package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/msk"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := msk.NewConfiguration(ctx, "example", &msk.ConfigurationArgs{
			KafkaVersions: pulumi.StringArray{
				pulumi.String("2.1.0"),
			},
			ServerProperties: pulumi.String(fmt.Sprintf("%v%v%v", "auto.create.topics.enable = true\n", "delete.topic.enable = true\n", "\n")),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

