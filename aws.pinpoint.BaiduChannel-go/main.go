package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/pinpoint"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		app, err := pinpoint.NewApp(ctx, "app", nil)
		if err != nil {
			return err
		}
		_, err = pinpoint.NewBaiduChannel(ctx, "channel", &pinpoint.BaiduChannelArgs{
			ApiKey:        pulumi.String(""),
			ApplicationId: app.ApplicationId,
			SecretKey:     pulumi.String(""),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

