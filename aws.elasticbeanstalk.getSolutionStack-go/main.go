package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/elasticbeanstalk"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := true
		_, err := elasticbeanstalk.GetSolutionStack(ctx, &elasticbeanstalk.GetSolutionStackArgs{
			MostRecent: &opt0,
			NameRegex:  fmt.Sprintf("%v%v", "^64bit Amazon Linux (.*) Multi-container Docker (.*)", "$"),
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

