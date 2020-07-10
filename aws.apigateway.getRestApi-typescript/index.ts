import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myRestApi = pulumi.output(aws.apigateway.getRestApi({
    name: "my-rest-api",
}, { async: true }));

