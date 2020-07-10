import pulumi
import pulumi_aws as aws

example_recorder = aws.cfg.Recorder("exampleRecorder")
example_function = aws.lambda_.Function("exampleFunction")
example_permission = aws.lambda_.Permission("examplePermission",
    action="lambda:InvokeFunction",
    function=example_function.arn,
    principal="config.amazonaws.com")
example_rule = aws.cfg.Rule("exampleRule", source={
    "owner": "CUSTOM_LAMBDA",
    "sourceIdentifier": example_function.arn,
},
opts=ResourceOptions(depends_on=[
        "aws_config_configuration_recorder.example",
        "aws_lambda_permission.example",
    ]))

