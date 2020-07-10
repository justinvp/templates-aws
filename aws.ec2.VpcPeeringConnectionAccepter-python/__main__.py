import pulumi
import pulumi_aws as aws
import pulumi_pulumi as pulumi

peer = pulumi.providers.Aws("peer", region="us-west-2")
main = aws.ec2.Vpc("main", cidr_block="10.0.0.0/16")
peer_vpc = aws.ec2.Vpc("peerVpc", cidr_block="10.1.0.0/16",
opts=ResourceOptions(provider="aws.peer"))
peer_caller_identity = aws.get_caller_identity()
# Requester's side of the connection.
peer_vpc_peering_connection = aws.ec2.VpcPeeringConnection("peerVpcPeeringConnection",
    auto_accept=False,
    peer_owner_id=peer_caller_identity.account_id,
    peer_region="us-west-2",
    peer_vpc_id=peer_vpc.id,
    tags={
        "Side": "Requester",
    },
    vpc_id=main.id)
# Accepter's side of the connection.
peer_vpc_peering_connection_accepter = aws.ec2.VpcPeeringConnectionAccepter("peerVpcPeeringConnectionAccepter",
    auto_accept=True,
    tags={
        "Side": "Accepter",
    },
    vpc_peering_connection_id=peer_vpc_peering_connection.id,
    opts=ResourceOptions(provider="aws.peer"))

