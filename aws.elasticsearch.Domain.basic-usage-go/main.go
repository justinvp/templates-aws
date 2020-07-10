package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/elasticsearch"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := elasticsearch.NewDomain(ctx, "example", &elasticsearch.DomainArgs{
			ClusterConfig: &elasticsearch.DomainClusterConfigArgs{
				ClusterConfig: pulumi.String("r4.large.elasticsearch"),
			},
			ElasticsearchVersion: pulumi.String("1.5"),
			SnapshotOptions: &elasticsearch.DomainSnapshotOptionsArgs{
				SnapshotOptions: pulumi.Float64(23),
			},
			Tags: pulumi.StringMap{
				"Domain": pulumi.String("TestDomain"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

