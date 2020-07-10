using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleRecorder = new Aws.Cfg.Recorder("exampleRecorder", new Aws.Cfg.RecorderArgs
        {
        });
        var exampleFunction = new Aws.Lambda.Function("exampleFunction", new Aws.Lambda.FunctionArgs
        {
        });
        var examplePermission = new Aws.Lambda.Permission("examplePermission", new Aws.Lambda.PermissionArgs
        {
            Action = "lambda:InvokeFunction",
            Function = exampleFunction.Arn,
            Principal = "config.amazonaws.com",
        });
        var exampleRule = new Aws.Cfg.Rule("exampleRule", new Aws.Cfg.RuleArgs
        {
            Source = new Aws.Cfg.Inputs.RuleSourceArgs
            {
                Owner = "CUSTOM_LAMBDA",
                SourceIdentifier = exampleFunction.Arn,
            },
        }, new CustomResourceOptions
        {
            DependsOn = 
            {
                "aws_config_configuration_recorder.example",
                "aws_lambda_permission.example",
            },
        });
    }

}

