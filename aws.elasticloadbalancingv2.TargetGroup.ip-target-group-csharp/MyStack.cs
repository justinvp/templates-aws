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
        var ip_example = new Aws.LB.TargetGroup("ip-example", new Aws.LB.TargetGroupArgs
        {
            Port = 80,
            Protocol = "HTTP",
            TargetType = "ip",
            VpcId = main.Id,
        });
    }

}

