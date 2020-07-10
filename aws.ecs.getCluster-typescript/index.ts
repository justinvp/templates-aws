import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ecs_mongo = pulumi.output(aws.ecs.getCluster({
    clusterName: "ecs-mongo-production",
}, { async: true }));

