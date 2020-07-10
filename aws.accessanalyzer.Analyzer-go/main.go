package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/accessanalyzer"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := accessanalyzer.NewAnalyzer(ctx, "example", &accessanalyzer.AnalyzerArgs{
			AnalyzerName: pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

