import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.secretsmanager.SecretRotation("example", {
    rotationLambdaArn: aws_lambda_function_example.arn,
    rotationRules: {
        automaticallyAfterDays: 30,
    },
    secretId: aws_secretsmanager_secret_example.id,
});

