import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleTransitGateway = new aws.ec2transitgateway.TransitGateway("example", {});
const exampleCustomerGateway = new aws.ec2.CustomerGateway("example", {
    bgpAsn: 65000,
    ipAddress: "172.0.0.1",
    type: "ipsec.1",
});
const exampleVpnConnection = new aws.ec2.VpnConnection("example", {
    customerGatewayId: exampleCustomerGateway.id,
    transitGatewayId: exampleTransitGateway.id,
    type: exampleCustomerGateway.type,
});

