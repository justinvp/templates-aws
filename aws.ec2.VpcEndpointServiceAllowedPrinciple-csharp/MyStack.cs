using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var current = Output.Create(Aws.GetCallerIdentity.InvokeAsync());
        var allowMeToFoo = new Aws.Ec2.VpcEndpointServiceAllowedPrinciple("allowMeToFoo", new Aws.Ec2.VpcEndpointServiceAllowedPrincipleArgs
        {
            PrincipalArn = current.Apply(current => current.Arn),
            VpcEndpointServiceId = aws_vpc_endpoint_service.Foo.Id,
        });
    }

}

