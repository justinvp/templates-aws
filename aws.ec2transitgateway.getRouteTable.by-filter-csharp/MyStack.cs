using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Ec2TransitGateway.GetRouteTable.InvokeAsync(new Aws.Ec2TransitGateway.GetRouteTableArgs
        {
            Filters = 
            {
                new Aws.Ec2TransitGateway.Inputs.GetRouteTableFilterArgs
                {
                    Name = "default-association-route-table",
                    Values = 
                    {
                        "true",
                    },
                },
                new Aws.Ec2TransitGateway.Inputs.GetRouteTableFilterArgs
                {
                    Name = "transit-gateway-id",
                    Values = 
                    {
                        "tgw-12345678",
                    },
                },
            },
        }));
    }

}

