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
        var test = new Aws.LB.TargetGroup("test", new Aws.LB.TargetGroupArgs
        {
            Port = 80,
            Protocol = "HTTP",
            VpcId = main.Id,
        });
    }

}

