import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const rotation_example = new aws.secretsmanager.Secret("rotation-example", {
    rotationLambdaArn: aws_lambda_function_example.arn,
    rotationRules: {
        automaticallyAfterDays: 7,
    },
});

