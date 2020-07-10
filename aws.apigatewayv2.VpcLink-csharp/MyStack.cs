using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.ApiGatewayV2.VpcLink("example", new Aws.ApiGatewayV2.VpcLinkArgs
        {
            SecurityGroupIds = 
            {
                data.Aws_security_group.Example.Id,
            },
            SubnetIds = data.Aws_subnet_ids.Example.Ids,
            Tags = 
            {
                { "Usage", "example" },
            },
        });
    }

}

