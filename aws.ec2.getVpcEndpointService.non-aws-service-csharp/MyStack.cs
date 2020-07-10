using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var custome = Output.Create(Aws.Ec2.GetVpcEndpointService.InvokeAsync(new Aws.Ec2.GetVpcEndpointServiceArgs
        {
            ServiceName = "com.amazonaws.vpce.us-west-2.vpce-svc-0e87519c997c63cd8",
        }));
    }

}

