import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = pulumi.output(aws.codecommit.getRepository({
    repositoryName: "MyTestRepository",
}, { async: true }));

