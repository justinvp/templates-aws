using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var foo = new Aws.Ec2.VpcPeeringConnection("foo", new Aws.Ec2.VpcPeeringConnectionArgs
        {
            PeerOwnerId = @var.Peer_owner_id,
            PeerVpcId = aws_vpc.Bar.Id,
            VpcId = aws_vpc.Foo.Id,
        });
    }

}

