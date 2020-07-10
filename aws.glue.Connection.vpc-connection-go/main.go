package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewConnection(ctx, "example", &glue.ConnectionArgs{
			ConnectionProperties: pulumi.StringMap{
				"JDBC_CONNECTION_URL": pulumi.String(fmt.Sprintf("%v%v%v", "jdbc:mysql://", aws_rds_cluster.Example.Endpoint, "/exampledatabase")),
				"PASSWORD":            pulumi.String("examplepassword"),
				"USERNAME":            pulumi.String("exampleusername"),
			},
			PhysicalConnectionRequirements: &glue.ConnectionPhysicalConnectionRequirementsArgs{
				AvailabilityZone: pulumi.Any(aws_subnet.Example.Availability_zone),
				SecurityGroupIdList: pulumi.AnyArray{
					pulumi.Any(aws_security_group.Example.Id),
				},
				SubnetId: pulumi.Any(aws_subnet.Example.Id),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

