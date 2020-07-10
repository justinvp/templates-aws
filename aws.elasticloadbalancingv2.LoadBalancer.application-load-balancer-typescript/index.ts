import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.lb.LoadBalancer("test", {
    accessLogs: {
        bucket: aws_s3_bucket_lb_logs.bucket,
        enabled: true,
        prefix: "test-lb",
    },
    enableDeletionProtection: true,
    internal: false,
    loadBalancerType: "application",
    securityGroups: [aws_security_group_lb_sg.id],
    subnets: [aws_subnet_public.map(v => v.id)],
    tags: {
        Environment: "production",
    },
});

