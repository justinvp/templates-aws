using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var s3 = Output.Create(Aws.Ec2.GetVpcEndpoint.InvokeAsync(new Aws.Ec2.GetVpcEndpointArgs
        {
            ServiceName = "com.amazonaws.us-west-2.s3",
            VpcId = aws_vpc.Foo.Id,
        }));
        var privateS3 = new Aws.Ec2.VpcEndpointRouteTableAssociation("privateS3", new Aws.Ec2.VpcEndpointRouteTableAssociationArgs
        {
            RouteTableId = aws_route_table.Private.Id,
            VpcEndpointId = s3.Apply(s3 => s3.Id),
        });
    }

}

