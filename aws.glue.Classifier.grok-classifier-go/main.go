package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewClassifier(ctx, "example", &glue.ClassifierArgs{
			GrokClassifier: &glue.ClassifierGrokClassifierArgs{
				Classification: pulumi.String("example"),
				GrokPattern:    pulumi.String("example"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

