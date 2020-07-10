using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var centos = Output.Create(Aws.Ssm.GetPatchBaseline.InvokeAsync(new Aws.Ssm.GetPatchBaselineArgs
        {
            NamePrefix = "AWS-",
            OperatingSystem = "CENTOS",
            Owner = "AWS",
        }));
    }

}

