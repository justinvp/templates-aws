using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var foo = new Aws.Ec2.Vpc("foo", new Aws.Ec2.VpcArgs
        {
            CidrBlock = "10.0.0.0/16",
        });
        var alphaSubnet = new Aws.Ec2.Subnet("alphaSubnet", new Aws.Ec2.SubnetArgs
        {
            AvailabilityZone = "us-west-2a",
            CidrBlock = "10.0.1.0/24",
            VpcId = foo.Id,
        });
        var alphaMountTarget = new Aws.Efs.MountTarget("alphaMountTarget", new Aws.Efs.MountTargetArgs
        {
            FileSystemId = aws_efs_file_system.Foo.Id,
            SubnetId = alphaSubnet.Id,
        });
    }

}

