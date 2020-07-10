using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var @default = Output.Create(Aws.Ec2.GetLaunchTemplate.InvokeAsync(new Aws.Ec2.GetLaunchTemplateArgs
        {
            Name = "my-launch-template",
        }));
    }

}

