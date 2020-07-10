using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var fooVpc = new Aws.Ec2.Vpc("fooVpc", new Aws.Ec2.VpcArgs
        {
            CidrBlock = "10.1.0.0/16",
        });
        var fooSubnet = new Aws.Ec2.Subnet("fooSubnet", new Aws.Ec2.SubnetArgs
        {
            AvailabilityZone = "us-west-2a",
            CidrBlock = "10.1.1.0/24",
            Tags = 
            {
                { "Name", "tf-dbsubnet-test-1" },
            },
            VpcId = fooVpc.Id,
        });
        var bar = new Aws.Ec2.Subnet("bar", new Aws.Ec2.SubnetArgs
        {
            AvailabilityZone = "us-west-2b",
            CidrBlock = "10.1.2.0/24",
            Tags = 
            {
                { "Name", "tf-dbsubnet-test-2" },
            },
            VpcId = fooVpc.Id,
        });
        var fooSubnetGroup = new Aws.RedShift.SubnetGroup("fooSubnetGroup", new Aws.RedShift.SubnetGroupArgs
        {
            SubnetIds = 
            {
                fooSubnet.Id,
                bar.Id,
            },
            Tags = 
            {
                { "environment", "Production" },
            },
        });
    }

}

