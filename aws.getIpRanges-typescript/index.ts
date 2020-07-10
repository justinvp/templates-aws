import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const europeanEc2 = aws.getIpRanges({
    regions: [
        "eu-west-1",
        "eu-central-1",
    ],
    services: ["ec2"],
});
const fromEurope = new aws.ec2.SecurityGroup("fromEurope", {
    ingress: [{
        fromPort: "443",
        toPort: "443",
        protocol: "tcp",
        cidrBlocks: europeanEc2.then(europeanEc2 => europeanEc2.cidrBlocks),
        ipv6CidrBlocks: europeanEc2.then(europeanEc2 => europeanEc2.ipv6CidrBlocks),
    }],
    tags: {
        CreateDate: europeanEc2.then(europeanEc2 => europeanEc2.createDate),
        SyncToken: europeanEc2.then(europeanEc2 => europeanEc2.syncToken),
    },
});

