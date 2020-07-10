using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var main = new Aws.Ec2.Vpc("main", new Aws.Ec2.VpcArgs
        {
            CidrBlock = "10.0.0.0/16",
        });
        var foo = new Aws.Ec2.Subnet("foo", new Aws.Ec2.SubnetArgs
        {
            AvailabilityZone = "us-west-2a",
            CidrBlock = "10.0.1.0/24",
            VpcId = main.Id,
        });
        var barSubnet = new Aws.Ec2.Subnet("barSubnet", new Aws.Ec2.SubnetArgs
        {
            AvailabilityZone = "us-west-2b",
            CidrBlock = "10.0.2.0/24",
            VpcId = main.Id,
        });
        var barDirectory = new Aws.DirectoryService.Directory("barDirectory", new Aws.DirectoryService.DirectoryArgs
        {
            Password = "SuperSecretPassw0rd",
            Size = "Small",
            Tags = 
            {
                { "Project", "foo" },
            },
            VpcSettings = new Aws.DirectoryService.Inputs.DirectoryVpcSettingsArgs
            {
                SubnetIds = 
                {
                    foo.Id,
                    barSubnet.Id,
                },
                VpcId = main.Id,
            },
        });
    }

}

