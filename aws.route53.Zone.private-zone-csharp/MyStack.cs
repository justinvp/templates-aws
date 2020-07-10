using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var @private = new Aws.Route53.Zone("private", new Aws.Route53.ZoneArgs
        {
            Vpcs = 
            {
                new Aws.Route53.Inputs.ZoneVpcArgs
                {
                    VpcId = aws_vpc.Example.Id,
                },
            },
        });
    }

}

