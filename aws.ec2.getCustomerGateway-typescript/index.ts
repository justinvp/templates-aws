import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = pulumi.output(aws.ec2.getCustomerGateway({
    filters: [{
        name: "tag:Name",
        values: ["foo-prod"],
    }],
}, { async: true }));
const main = new aws.ec2.VpnGateway("main", {
    amazonSideAsn: "7224",
    vpcId: aws_vpc_main.id,
});
const transit = new aws.ec2.VpnConnection("transit", {
    customerGatewayId: foo.id!,
    staticRoutesOnly: false,
    type: foo.type,
    vpnGatewayId: main.id,
});

