import pulumi
import pulumi_aws as aws

example_layer_version = aws.lambda_.LayerVersion("exampleLayerVersion")
example_function = aws.lambda_.Function("exampleFunction", layers=[example_layer_version.arn])

