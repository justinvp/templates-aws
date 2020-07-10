using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.SecretsManager.SecretRotation("example", new Aws.SecretsManager.SecretRotationArgs
        {
            RotationLambdaArn = aws_lambda_function.Example.Arn,
            RotationRules = new Aws.SecretsManager.Inputs.SecretRotationRotationRulesArgs
            {
                AutomaticallyAfterDays = 30,
            },
            SecretId = aws_secretsmanager_secret.Example.Id,
        });
    }

}

