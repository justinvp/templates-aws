package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/datasync"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datasync.NewLocationSmb(ctx, "example", &datasync.LocationSmbArgs{
			AgentArns: pulumi.StringArray{
				pulumi.Any(aws_datasync_agent.Example.Arn),
			},
			Password:       pulumi.String("ANotGreatPassword"),
			ServerHostname: pulumi.String("smb.example.com"),
			Subdirectory:   pulumi.String("/exported/path"),
			User:           pulumi.String("Guest"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

