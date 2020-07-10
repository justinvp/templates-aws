import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const service = pulumi.output(aws.ecr.getRepository({
    name: "ecr-repository",
}, { async: true }));

