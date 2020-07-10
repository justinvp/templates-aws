using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Ec2.Fleet("example", new Aws.Ec2.FleetArgs
        {
            LaunchTemplateConfig = new Aws.Ec2.Inputs.FleetLaunchTemplateConfigArgs
            {
                LaunchTemplateSpecification = new Aws.Ec2.Inputs.FleetLaunchTemplateConfigLaunchTemplateSpecificationArgs
                {
                    LaunchTemplateId = aws_launch_template.Example.Id,
                    Version = aws_launch_template.Example.Latest_version,
                },
            },
            TargetCapacitySpecification = new Aws.Ec2.Inputs.FleetTargetCapacitySpecificationArgs
            {
                DefaultTargetCapacityType = "spot",
                TotalTargetCapacity = 5,
            },
        });
    }

}

