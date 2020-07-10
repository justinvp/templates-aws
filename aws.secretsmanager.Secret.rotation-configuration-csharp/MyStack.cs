using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var rotation_example = new Aws.SecretsManager.Secret("rotation-example", new Aws.SecretsManager.SecretArgs
        {
            RotationLambdaArn = aws_lambda_function.Example.Arn,
            RotationRules = new Aws.SecretsManager.Inputs.SecretRotationRulesArgs
            {
                AutomaticallyAfterDays = 7,
            },
        });
    }

}

