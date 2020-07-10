package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleZone, err := route53.NewZone(ctx, "exampleZone", nil)
		if err != nil {
			return err
		}
		_, err = route53.NewRecord(ctx, "exampleRecord", &route53.RecordArgs{
			AllowOverwrite: pulumi.Bool(true),
			Name:           pulumi.String("test.example.com"),
			Records: pulumi.StringArray{
				exampleZone.NameServers.ApplyT(func(nameServers []string) (string, error) {
					return nameServers[0], nil
				}).(pulumi.StringOutput),
				exampleZone.NameServers.ApplyT(func(nameServers []string) (string, error) {
					return nameServers[1], nil
				}).(pulumi.StringOutput),
				exampleZone.NameServers.ApplyT(func(nameServers []string) (string, error) {
					return nameServers[2], nil
				}).(pulumi.StringOutput),
				exampleZone.NameServers.ApplyT(func(nameServers []string) (string, error) {
					return nameServers[3], nil
				}).(pulumi.StringOutput),
			},
			Ttl:    pulumi.Int(30),
			Type:   pulumi.String("NS"),
			ZoneId: exampleZone.ZoneId,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

