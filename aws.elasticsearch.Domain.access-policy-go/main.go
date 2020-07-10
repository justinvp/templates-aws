package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/elasticsearch"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		currentRegion, err := aws.GetRegion(ctx, nil, nil)
		if err != nil {
			return err
		}
		currentCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = elasticsearch.NewDomain(ctx, "example", &elasticsearch.DomainArgs{
			AccessPolicies: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": \"es:*\",\n", "      \"Principal\": \"*\",\n", "      \"Effect\": \"Allow\",\n", "      \"Resource\": \"arn:aws:es:", currentRegion.Name, ":", currentCallerIdentity.AccountId, ":domain/", domain, "/*\",\n", "      \"Condition\": {\n", "        \"IpAddress\": {\"aws:SourceIp\": [\"66.193.100.22/32\"]}\n", "      }\n", "    }\n", "  ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

