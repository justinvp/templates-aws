package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		main, err := route53.NewZone(ctx, "main", nil)
		if err != nil {
			return err
		}
		dev, err := route53.NewZone(ctx, "dev", &route53.ZoneArgs{
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("dev"),
			},
		})
		if err != nil {
			return err
		}
		_, err = route53.NewRecord(ctx, "dev_ns", &route53.RecordArgs{
			Name: pulumi.String("dev.example.com"),
			Records: pulumi.StringArray{
				dev.NameServers.ApplyT(func(nameServers []string) (string, error) {
					return nameServers[0], nil
				}).(pulumi.StringOutput),
				dev.NameServers.ApplyT(func(nameServers []string) (string, error) {
					return nameServers[1], nil
				}).(pulumi.StringOutput),
				dev.NameServers.ApplyT(func(nameServers []string) (string, error) {
					return nameServers[2], nil
				}).(pulumi.StringOutput),
				dev.NameServers.ApplyT(func(nameServers []string) (string, error) {
					return nameServers[3], nil
				}).(pulumi.StringOutput),
			},
			Ttl:    pulumi.Int(30),
			Type:   pulumi.String("NS"),
			ZoneId: main.ZoneId,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

