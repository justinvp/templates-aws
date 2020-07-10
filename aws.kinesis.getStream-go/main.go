package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/kinesis"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kinesis.LookupStream(ctx, &kinesis.LookupStreamArgs{
			Name: "stream-name",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

