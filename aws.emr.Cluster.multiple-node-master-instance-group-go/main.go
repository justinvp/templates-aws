package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/emr"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleSubnet, err := ec2.NewSubnet(ctx, "exampleSubnet", &ec2.SubnetArgs{
			MapPublicIpOnLaunch: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		_, err = emr.NewCluster(ctx, "exampleCluster", &emr.ClusterArgs{
			CoreInstanceGroup: nil,
			Ec2Attributes: &emr.ClusterEc2AttributesArgs{
				SubnetId: exampleSubnet.ID(),
			},
			MasterInstanceGroup: &emr.ClusterMasterInstanceGroupArgs{
				InstanceCount: pulumi.Int(3),
			},
			ReleaseLabel:          pulumi.String("emr-5.24.1"),
			TerminationProtection: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

