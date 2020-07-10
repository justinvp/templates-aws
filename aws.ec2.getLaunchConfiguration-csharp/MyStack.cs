using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var ubuntu = Output.Create(Aws.Ec2.GetLaunchConfiguration.InvokeAsync(new Aws.Ec2.GetLaunchConfigurationArgs
        {
            Name = "test-launch-config",
        }));
    }

}

