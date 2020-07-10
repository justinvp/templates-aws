import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ec2.getLocalGatewayVirtualInterfaceGroup({
    localGatewayId: data.aws_ec2_local_gateway.example.id,
});

