using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var fooVpc = new Aws.Ec2.Vpc("fooVpc", new Aws.Ec2.VpcArgs
        {
            CidrBlock = "10.0.0.0/16",
            Tags = 
            {
                { "Name", "tf-test" },
            },
        });
        var fooSubnet = new Aws.Ec2.Subnet("fooSubnet", new Aws.Ec2.SubnetArgs
        {
            AvailabilityZone = "us-west-2a",
            CidrBlock = "10.0.0.0/24",
            Tags = 
            {
                { "Name", "tf-test" },
            },
            VpcId = fooVpc.Id,
        });
        var bar = new Aws.ElastiCache.SubnetGroup("bar", new Aws.ElastiCache.SubnetGroupArgs
        {
            SubnetIds = 
            {
                fooSubnet.Id,
            },
        });
    }

}

