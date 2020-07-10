import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const lb = new aws.elb.LoadBalancer("lb", {
    availabilityZones: ["us-east-1a"],
    listeners: [
        {
            instancePort: 25,
            instanceProtocol: "tcp",
            lbPort: 25,
            lbProtocol: "tcp",
        },
        {
            instancePort: 587,
            instanceProtocol: "tcp",
            lbPort: 587,
            lbProtocol: "tcp",
        },
    ],
});
const smtp = new aws.ec2.ProxyProtocolPolicy("smtp", {
    instancePorts: [
        "25",
        "587",
    ],
    loadBalancer: lb.name,
});

