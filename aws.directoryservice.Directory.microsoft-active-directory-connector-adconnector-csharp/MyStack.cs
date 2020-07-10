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
        var bar = new Aws.Ec2.Subnet("bar", new Aws.Ec2.SubnetArgs
        {
            AvailabilityZone = "us-west-2b",
            CidrBlock = "10.0.2.0/24",
            VpcId = main.Id,
        });
        var connector = new Aws.DirectoryService.Directory("connector", new Aws.DirectoryService.DirectoryArgs
        {
            ConnectSettings = new Aws.DirectoryService.Inputs.DirectoryConnectSettingsArgs
            {
                CustomerDnsIps = 
                {
                    "A.B.C.D",
                },
                CustomerUsername = "Admin",
                SubnetIds = 
                {
                    foo.Id,
                    bar.Id,
                },
                VpcId = main.Id,
            },
            Password = "SuperSecretPassw0rd",
            Size = "Small",
            Type = "ADConnector",
        });
    }

}

