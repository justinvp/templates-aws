package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/elasticsearch"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := elasticsearch.NewDomain(ctx, "example", &elasticsearch.DomainArgs{
			ElasticsearchVersion: pulumi.String("2.3"),
		})
		if err != nil {
			return err
		}
		_, err = elasticsearch.NewDomainPolicy(ctx, "main", &elasticsearch.DomainPolicyArgs{
			AccessPolicies: example.Arn.ApplyT(func(arn string) (string, error) {
				return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "    \"Version\": \"2012-10-17\",\n", "    \"Statement\": [\n", "        {\n", "            \"Action\": \"es:*\",\n", "            \"Principal\": \"*\",\n", "            \"Effect\": \"Allow\",\n", "            \"Condition\": {\n", "                \"IpAddress\": {\"aws:SourceIp\": \"127.0.0.1/32\"}\n", "            },\n", "            \"Resource\": \"", arn, "/*\"\n", "        }\n", "    ]\n", "}\n", "\n"), nil
			}).(pulumi.StringOutput),
			DomainName: example.DomainName,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

