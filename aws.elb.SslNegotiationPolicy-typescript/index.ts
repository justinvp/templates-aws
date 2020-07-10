import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const lb = new aws.elb.LoadBalancer("lb", {
    availabilityZones: ["us-east-1a"],
    listeners: [{
        instancePort: 8000,
        instanceProtocol: "https",
        lbPort: 443,
        lbProtocol: "https",
        sslCertificateId: "arn:aws:iam::123456789012:server-certificate/certName",
    }],
});
const foo = new aws.elb.SslNegotiationPolicy("foo", {
    attributes: [
        {
            name: "Protocol-TLSv1",
            value: "false",
        },
        {
            name: "Protocol-TLSv1.1",
            value: "false",
        },
        {
            name: "Protocol-TLSv1.2",
            value: "true",
        },
        {
            name: "Server-Defined-Cipher-Order",
            value: "true",
        },
        {
            name: "ECDHE-RSA-AES128-GCM-SHA256",
            value: "true",
        },
        {
            name: "AES128-GCM-SHA256",
            value: "true",
        },
        {
            name: "EDH-RSA-DES-CBC3-SHA",
            value: "false",
        },
    ],
    lbPort: 443,
    loadBalancer: lb.id,
});

