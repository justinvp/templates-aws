import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.route53.ResolverEndpoint("foo", {
    direction: "INBOUND",
    ipAddresses: [
        {
            subnetId: aws_subnet_sn1.id,
        },
        {
            ip: "10.0.64.4",
            subnetId: aws_subnet_sn2.id,
        },
    ],
    securityGroupIds: [
        aws_security_group_sg1.id,
        aws_security_group_sg2.id,
    ],
    tags: {
        Environment: "Prod",
    },
});

