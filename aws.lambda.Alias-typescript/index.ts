import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testLambdaAlias = new aws.lambda.Alias("test_lambda_alias", {
    description: "a sample description",
    functionName: aws_lambda_function_lambda_function_test.arn,
    functionVersion: "1",
    routingConfig: {
        additionalVersionWeights: {
            "2": 0.5,
        },
    },
});

