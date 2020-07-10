package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/datasync"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datasync.NewTask(ctx, "example", &datasync.TaskArgs{
			DestinationLocationArn: pulumi.Any(aws_datasync_location_s3.Destination.Arn),
			Options: &datasync.TaskOptionsArgs{
				BytesPerSecond: pulumi.Int(-1),
			},
			SourceLocationArn: pulumi.Any(aws_datasync_location_nfs.Source.Arn),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

