import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.lambda.ProvisionedConcurrencyConfig("example", {
    functionName: aws_lambda_function.example.function_name,
    provisionedConcurrentExecutions: 1,
    qualifier: aws_lambda_function.example.version,
});

