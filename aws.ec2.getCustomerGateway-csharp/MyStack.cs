using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var foo = Output.Create(Aws.Ec2.GetCustomerGateway.InvokeAsync(new Aws.Ec2.GetCustomerGatewayArgs
        {
            Filters = 
            {
                new Aws.Ec2.Inputs.GetCustomerGatewayFilterArgs
                {
                    Name = "tag:Name",
                    Values = 
                    {
                        "foo-prod",
                    },
                },
            },
        }));
        var main = new Aws.Ec2.VpnGateway("main", new Aws.Ec2.VpnGatewayArgs
        {
            AmazonSideAsn = "7224",
            VpcId = aws_vpc.Main.Id,
        });
        var transit = new Aws.Ec2.VpnConnection("transit", new Aws.Ec2.VpnConnectionArgs
        {
            CustomerGatewayId = foo.Apply(foo => foo.Id),
            StaticRoutesOnly = false,
            Type = foo.Apply(foo => foo.Type),
            VpnGatewayId = main.Id,
        });
    }

}

