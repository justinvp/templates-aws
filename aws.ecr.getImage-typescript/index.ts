import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const serviceImage = pulumi.output(aws.ecr.getImage({
    imageTag: "latest",
    repositoryName: "my/service",
}, { async: true }));

