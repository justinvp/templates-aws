using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var test = Output.Create(Aws.Efs.GetAccessPoint.InvokeAsync(new Aws.Efs.GetAccessPointArgs
        {
            AccessPointId = "fsap-12345678",
        }));
    }

}

