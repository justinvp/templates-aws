import pulumi
import pulumi_aws as aws

config = pulumi.Config()
layer_name = config.require_object("layerName")
existing = aws.lambda.get_layer_version(layer_name=layer_name)

