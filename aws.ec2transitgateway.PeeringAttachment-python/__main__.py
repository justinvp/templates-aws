import pulumi
import pulumi_aws as aws
import pulumi_pulumi as pulumi

local = pulumi.providers.Aws("local", region="us-east-1")
peer = pulumi.providers.Aws("peer", region="us-west-2")
peer_region = aws.get_region()
local_transit_gateway = aws.ec2transitgateway.TransitGateway("localTransitGateway", tags={
    "Name": "Local TGW",
},
opts=ResourceOptions(provider=aws["local"]))
peer_transit_gateway = aws.ec2transitgateway.TransitGateway("peerTransitGateway", tags={
    "Name": "Peer TGW",
},
opts=ResourceOptions(provider=aws["peer"]))
example = aws.ec2transitgateway.PeeringAttachment("example",
    peer_account_id=peer_transit_gateway.owner_id,
    peer_region=peer_region.name,
    peer_transit_gateway_id=peer_transit_gateway.id,
    transit_gateway_id=local_transit_gateway.id,
    tags={
        "Name": "TGW Peering Requestor",
    })

