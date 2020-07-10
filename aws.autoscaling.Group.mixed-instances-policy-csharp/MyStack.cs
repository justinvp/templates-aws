using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleLaunchTemplate = new Aws.Ec2.LaunchTemplate("exampleLaunchTemplate", new Aws.Ec2.LaunchTemplateArgs
        {
            ImageId = data.Aws_ami.Example.Id,
            InstanceType = "c5.large",
            NamePrefix = "example",
        });
        var exampleGroup = new Aws.AutoScaling.Group("exampleGroup", new Aws.AutoScaling.GroupArgs
        {
            AvailabilityZones = 
            {
                "us-east-1a",
            },
            DesiredCapacity = 1,
            MaxSize = 1,
            MinSize = 1,
            MixedInstancesPolicy = new Aws.AutoScaling.Inputs.GroupMixedInstancesPolicyArgs
            {
                LaunchTemplate = new Aws.AutoScaling.Inputs.GroupMixedInstancesPolicyLaunchTemplateArgs
                {
                    LaunchTemplateSpecification = new Aws.AutoScaling.Inputs.GroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecificationArgs
                    {
                        LaunchTemplateId = exampleLaunchTemplate.Id,
                    },
                    Override = 
                    {
                        
                        {
                            { "instanceType", "c4.large" },
                            { "weightedCapacity", "3" },
                        },
                        
                        {
                            { "instanceType", "c3.large" },
                            { "weightedCapacity", "2" },
                        },
                    },
                },
            },
        });
    }

}

