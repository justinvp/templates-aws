package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/autoscaling"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleLaunchTemplate, err := ec2.NewLaunchTemplate(ctx, "exampleLaunchTemplate", &ec2.LaunchTemplateArgs{
			ImageId:      pulumi.Any(data.Aws_ami.Example.Id),
			InstanceType: pulumi.String("c5.large"),
			NamePrefix:   pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		_, err = autoscaling.NewGroup(ctx, "exampleGroup", &autoscaling.GroupArgs{
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("us-east-1a"),
			},
			DesiredCapacity: pulumi.Int(1),
			MaxSize:         pulumi.Int(1),
			MinSize:         pulumi.Int(1),
			MixedInstancesPolicy: &autoscaling.GroupMixedInstancesPolicyArgs{
				LaunchTemplate: &autoscaling.GroupMixedInstancesPolicyLaunchTemplateArgs{
					LaunchTemplateSpecification: &autoscaling.GroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecificationArgs{
						LaunchTemplateId: exampleLaunchTemplate.ID(),
					},
					Override: pulumi.StringMapArray{
						pulumi.StringMap{
							"instanceType":     pulumi.String("c4.large"),
							"weightedCapacity": pulumi.String("3"),
						},
						pulumi.StringMap{
							"instanceType":     pulumi.String("c3.large"),
							"weightedCapacity": pulumi.String("2"),
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

