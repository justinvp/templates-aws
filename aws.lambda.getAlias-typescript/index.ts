import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const production = pulumi.output(aws.lambda.getAlias({
    functionName: "my-lambda-func",
    name: "production",
}, { async: true }));

