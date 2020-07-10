import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.lb.LoadBalancer("example", {
    loadBalancerType: "network",
    subnetMappings: [
        {
            allocationId: aws_eip_example1.id,
            subnetId: aws_subnet_example1.id,
        },
        {
            allocationId: aws_eip_example2.id,
            subnetId: aws_subnet_example2.id,
        },
    ],
});

