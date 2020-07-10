import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bar = new aws.ecs.Service("bar", {
    cluster: aws_ecs_cluster_foo.id,
    schedulingStrategy: "DAEMON",
    taskDefinition: aws_ecs_task_definition_bar.arn,
});

