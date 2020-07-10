package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/datasync"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datasync.NewNfsLocation(ctx, "example", &datasync.NfsLocationArgs{
			OnPremConfig: &datasync.NfsLocationOnPremConfigArgs{
				AgentArns: pulumi.StringArray{
					pulumi.Any(aws_datasync_agent.Example.Arn),
				},
			},
			ServerHostname: pulumi.String("nfs.example.com"),
			Subdirectory:   pulumi.String("/exported/path"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

