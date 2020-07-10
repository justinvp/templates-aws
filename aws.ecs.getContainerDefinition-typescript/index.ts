import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ecs_mongo = aws_ecs_task_definition_mongo.id.apply(id => aws.ecs.getContainerDefinition({
    containerName: "mongodb",
    taskDefinition: id,
}, { async: true }));

