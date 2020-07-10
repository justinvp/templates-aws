package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lambda.LookupLayerVersion(ctx, &lambda.LookupLayerVersionArgs{
			LayerName: layerName,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

