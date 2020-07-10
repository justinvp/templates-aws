import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.lambda.FunctionEventInvokeConfig("example", {
    functionName: aws_lambda_alias.example.function_name,
    qualifier: aws_lambda_alias.example.name,
});
// ... other configuration ...

