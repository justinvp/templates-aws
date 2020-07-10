import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myApiKey = pulumi.output(aws.apigateway.getKey({
    id: "ru3mpjgse6",
}, { async: true }));

