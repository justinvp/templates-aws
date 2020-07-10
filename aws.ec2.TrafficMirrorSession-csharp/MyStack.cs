using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var filter = new Aws.Ec2.TrafficMirrorFilter("filter", new Aws.Ec2.TrafficMirrorFilterArgs
        {
            Description = "traffic mirror filter - example",
            NetworkServices = 
            {
                "amazon-dns",
            },
        });
        var target = new Aws.Ec2.TrafficMirrorTarget("target", new Aws.Ec2.TrafficMirrorTargetArgs
        {
            NetworkLoadBalancerArn = aws_lb.Lb.Arn,
        });
        var session = new Aws.Ec2.TrafficMirrorSession("session", new Aws.Ec2.TrafficMirrorSessionArgs
        {
            Description = "traffic mirror session - example",
            NetworkInterfaceId = aws_instance.Test.Primary_network_interface_id,
            TrafficMirrorFilterId = filter.Id,
            TrafficMirrorTargetId = target.Id,
        });
    }

}

