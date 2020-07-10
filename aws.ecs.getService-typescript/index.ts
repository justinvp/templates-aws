import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws_ecs_cluster_example.arn.apply(arn => aws.ecs.getService({
    clusterArn: arn,
    serviceName: "example",
}, { async: true }));

