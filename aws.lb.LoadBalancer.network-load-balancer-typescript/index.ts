import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.lb.LoadBalancer("test", {
    enableDeletionProtection: true,
    internal: false,
    loadBalancerType: "network",
    subnets: [aws_subnet_public.map(v => v.id)],
    tags: {
        Environment: "production",
    },
});

