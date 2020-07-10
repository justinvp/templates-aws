package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/elasticsearch"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := elasticsearch.LookupDomain(ctx, &elasticsearch.LookupDomainArgs{
			DomainName: "my-domain-name",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

