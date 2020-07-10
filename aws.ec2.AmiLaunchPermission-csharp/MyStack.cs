using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Ec2.AmiLaunchPermission("example", new Aws.Ec2.AmiLaunchPermissionArgs
        {
            AccountId = "123456789012",
            ImageId = "ami-12345678",
        });
    }

}

