package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/kinesis"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kinesis.NewVideoStream(ctx, "_default", &kinesis.VideoStreamArgs{
			DataRetentionInHours: pulumi.Int(1),
			DeviceName:           pulumi.String("kinesis-video-device-name"),
			MediaType:            pulumi.String("video/h264"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("kinesis-video-stream"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

