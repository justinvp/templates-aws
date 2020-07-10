using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var peer = new Aws.Provider("peer", new Aws.ProviderArgs
        {
            Region = "us-west-2",
        });
        var main = new Aws.Ec2.Vpc("main", new Aws.Ec2.VpcArgs
        {
            CidrBlock = "10.0.0.0/16",
        });
        var peerVpc = new Aws.Ec2.Vpc("peerVpc", new Aws.Ec2.VpcArgs
        {
            CidrBlock = "10.1.0.0/16",
        }, new CustomResourceOptions
        {
            Provider = "aws.peer",
        });
        var peerCallerIdentity = Output.Create(Aws.GetCallerIdentity.InvokeAsync());
        // Requester's side of the connection.
        var peerVpcPeeringConnection = new Aws.Ec2.VpcPeeringConnection("peerVpcPeeringConnection", new Aws.Ec2.VpcPeeringConnectionArgs
        {
            AutoAccept = false,
            PeerOwnerId = peerCallerIdentity.Apply(peerCallerIdentity => peerCallerIdentity.AccountId),
            PeerRegion = "us-west-2",
            PeerVpcId = peerVpc.Id,
            Tags = 
            {
                { "Side", "Requester" },
            },
            VpcId = main.Id,
        });
        // Accepter's side of the connection.
        var peerVpcPeeringConnectionAccepter = new Aws.Ec2.VpcPeeringConnectionAccepter("peerVpcPeeringConnectionAccepter", new Aws.Ec2.VpcPeeringConnectionAccepterArgs
        {
            AutoAccept = true,
            Tags = 
            {
                { "Side", "Accepter" },
            },
            VpcPeeringConnectionId = peerVpcPeeringConnection.Id,
        }, new CustomResourceOptions
        {
            Provider = "aws.peer",
        });
    }

}

