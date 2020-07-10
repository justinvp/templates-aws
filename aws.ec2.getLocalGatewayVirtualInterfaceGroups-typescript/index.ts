import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const all = pulumi.output(aws.ec2.getLocalGatewayVirtualInterfaceGroups({ async: true }));

