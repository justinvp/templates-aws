using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var @default = new Aws.Ec2.DefaultVpc("default", new Aws.Ec2.DefaultVpcArgs
        {
            Tags = 
            {
                { "Name", "Default VPC" },
            },
        });
    }

}

