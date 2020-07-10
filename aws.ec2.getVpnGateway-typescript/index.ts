import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const selected = pulumi.output(aws.ec2.getVpnGateway({
    filters: [{
        name: "tag:Name",
        values: ["vpn-gw"],
    }],
}, { async: true }));

export const vpnGatewayId = selected.id!;

