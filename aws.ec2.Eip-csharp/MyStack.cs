using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var lb = new Aws.Ec2.Eip("lb", new Aws.Ec2.EipArgs
        {
            Instance = aws_instance.Web.Id,
            Vpc = true,
        });
    }

}

