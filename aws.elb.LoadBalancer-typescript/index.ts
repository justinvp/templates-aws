import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Create a new load balancer
const bar = new aws.elb.LoadBalancer("bar", {
    accessLogs: {
        bucket: "foo",
        bucketPrefix: "bar",
        interval: 60,
    },
    availabilityZones: [
        "us-west-2a",
        "us-west-2b",
        "us-west-2c",
    ],
    connectionDraining: true,
    connectionDrainingTimeout: 400,
    crossZoneLoadBalancing: true,
    healthCheck: {
        healthyThreshold: 2,
        interval: 30,
        target: "HTTP:8000/",
        timeout: 3,
        unhealthyThreshold: 2,
    },
    idleTimeout: 400,
    instances: [aws_instance_foo.id],
    listeners: [
        {
            instancePort: 8000,
            instanceProtocol: "http",
            lbPort: 80,
            lbProtocol: "http",
        },
        {
            instancePort: 8000,
            instanceProtocol: "http",
            lbPort: 443,
            lbProtocol: "https",
            sslCertificateId: "arn:aws:iam::123456789012:server-certificate/certName",
        },
    ],
    tags: {
        Name: "foobar-elb",
    },
});

