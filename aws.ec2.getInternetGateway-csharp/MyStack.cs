using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var config = new Config();
        var vpcId = config.RequireObject<dynamic>("vpcId");
        var @default = Output.Create(Aws.Ec2.GetInternetGateway.InvokeAsync(new Aws.Ec2.GetInternetGatewayArgs
        {
            Filters = 
            {
                new Aws.Ec2.Inputs.GetInternetGatewayFilterArgs
                {
                    Name = "attachment.vpc-id",
                    Values = 
                    {
                        vpcId,
                    },
                },
            },
        }));
    }

}

