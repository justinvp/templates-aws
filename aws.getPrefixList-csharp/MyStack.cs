using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var privateS3VpcEndpoint = new Aws.Ec2.VpcEndpoint("privateS3VpcEndpoint", new Aws.Ec2.VpcEndpointArgs
        {
            ServiceName = "com.amazonaws.us-west-2.s3",
            VpcId = aws_vpc.Foo.Id,
        });
        var privateS3PrefixList = privateS3VpcEndpoint.PrefixListId.Apply(prefixListId => Aws.GetPrefixList.InvokeAsync(new Aws.GetPrefixListArgs
        {
            PrefixListId = prefixListId,
        }));
        var bar = new Aws.Ec2.NetworkAcl("bar", new Aws.Ec2.NetworkAclArgs
        {
            VpcId = aws_vpc.Foo.Id,
        });
        var privateS3NetworkAclRule = new Aws.Ec2.NetworkAclRule("privateS3NetworkAclRule", new Aws.Ec2.NetworkAclRuleArgs
        {
            CidrBlock = privateS3PrefixList.Apply(privateS3PrefixList => privateS3PrefixList.CidrBlocks[0]),
            Egress = false,
            FromPort = 443,
            NetworkAclId = bar.Id,
            Protocol = "tcp",
            RuleAction = "allow",
            RuleNumber = 200,
            ToPort = 443,
        });
    }

}

