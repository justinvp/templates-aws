import pulumi
import pulumi_aws as aws

example = aws.directconnect.GatewayAssociationProposal("example",
    associated_gateway_id=aws_vpn_gateway["example"]["id"],
    dx_gateway_id=aws_dx_gateway["example"]["id"],
    dx_gateway_owner_account_id=aws_dx_gateway["example"]["owner_account_id"])

