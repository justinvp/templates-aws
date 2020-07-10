package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudformation"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudformation.NewStack(ctx, "network", &cloudformation.StackArgs{
			Parameters: pulumi.StringMap{
				"VPCCidr": pulumi.String("10.0.0.0/16"),
			},
			TemplateBody: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Parameters\" : {\n", "    \"VPCCidr\" : {\n", "      \"Type\" : \"String\",\n", "      \"Default\" : \"10.0.0.0/16\",\n", "      \"Description\" : \"Enter the CIDR block for the VPC. Default is 10.0.0.0/16.\"\n", "    }\n", "  },\n", "  \"Resources\" : {\n", "    \"myVpc\": {\n", "      \"Type\" : \"AWS::EC2::VPC\",\n", "      \"Properties\" : {\n", "        \"CidrBlock\" : { \"Ref\" : \"VPCCidr\" },\n", "        \"Tags\" : [\n", "          {\"Key\": \"Name\", \"Value\": \"Primary_CF_VPC\"}\n", "        ]\n", "      }\n", "    }\n", "  }\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

