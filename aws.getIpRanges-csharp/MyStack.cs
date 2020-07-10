using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var europeanEc2 = Output.Create(Aws.GetIpRanges.InvokeAsync(new Aws.GetIpRangesArgs
        {
            Regions = 
            {
                "eu-west-1",
                "eu-central-1",
            },
            Services = 
            {
                "ec2",
            },
        }));
        var fromEurope = new Aws.Ec2.SecurityGroup("fromEurope", new Aws.Ec2.SecurityGroupArgs
        {
            Ingress = 
            {
                new Aws.Ec2.Inputs.SecurityGroupIngressArgs
                {
                    FromPort = 443,
                    ToPort = 443,
                    Protocol = "tcp",
                    CidrBlocks = europeanEc2.Apply(europeanEc2 => europeanEc2.CidrBlocks),
                    Ipv6CidrBlocks = europeanEc2.Apply(europeanEc2 => europeanEc2.Ipv6CidrBlocks),
                },
            },
            Tags = 
            {
                { "CreateDate", europeanEc2.Apply(europeanEc2 => europeanEc2.CreateDate) },
                { "SyncToken", europeanEc2.Apply(europeanEc2 => europeanEc2.SyncToken) },
            },
        });
    }

}

