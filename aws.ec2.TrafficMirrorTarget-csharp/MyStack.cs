using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var nlb = new Aws.Ec2.TrafficMirrorTarget("nlb", new Aws.Ec2.TrafficMirrorTargetArgs
        {
            Description = "NLB target",
            NetworkLoadBalancerArn = aws_lb.Lb.Arn,
        });
        var eni = new Aws.Ec2.TrafficMirrorTarget("eni", new Aws.Ec2.TrafficMirrorTargetArgs
        {
            Description = "ENI target",
            NetworkInterfaceId = aws_instance.Test.Primary_network_interface_id,
        });
    }

}

