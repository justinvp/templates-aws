using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var s3 = Output.Create(Aws.Ec2.GetVpcEndpointService.InvokeAsync(new Aws.Ec2.GetVpcEndpointServiceArgs
        {
            Service = "s3",
        }));
        // Create a VPC
        var foo = new Aws.Ec2.Vpc("foo", new Aws.Ec2.VpcArgs
        {
            CidrBlock = "10.0.0.0/16",
        });
        // Create a VPC endpoint
        var ep = new Aws.Ec2.VpcEndpoint("ep", new Aws.Ec2.VpcEndpointArgs
        {
            ServiceName = s3.Apply(s3 => s3.ServiceName),
            VpcId = foo.Id,
        });
    }

}

