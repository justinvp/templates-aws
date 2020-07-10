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
        var ruleout = new Aws.Ec2.TrafficMirrorFilterRule("ruleout", new Aws.Ec2.TrafficMirrorFilterRuleArgs
        {
            Description = "test rule",
            DestinationCidrBlock = "10.0.0.0/8",
            RuleAction = "accept",
            RuleNumber = 1,
            SourceCidrBlock = "10.0.0.0/8",
            TrafficDirection = "egress",
            TrafficMirrorFilterId = filter.Id,
        });
        var rulein = new Aws.Ec2.TrafficMirrorFilterRule("rulein", new Aws.Ec2.TrafficMirrorFilterRuleArgs
        {
            Description = "test rule",
            DestinationCidrBlock = "10.0.0.0/8",
            DestinationPortRange = new Aws.Ec2.Inputs.TrafficMirrorFilterRuleDestinationPortRangeArgs
            {
                FromPort = 22,
                ToPort = 53,
            },
            Protocol = 6,
            RuleAction = "accept",
            RuleNumber = 1,
            SourceCidrBlock = "10.0.0.0/8",
            SourcePortRange = new Aws.Ec2.Inputs.TrafficMirrorFilterRuleSourcePortRangeArgs
            {
                FromPort = 0,
                ToPort = 10,
            },
            TrafficDirection = "ingress",
            TrafficMirrorFilterId = filter.Id,
        });
    }

}

