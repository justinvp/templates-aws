package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := route53.LookupDelegationSet(ctx, &route53.LookupDelegationSetArgs{
			Id: "MQWGHCBFAKEID",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

