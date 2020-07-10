using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.GuardDuty.GetDetector.InvokeAsync());
    }

}

