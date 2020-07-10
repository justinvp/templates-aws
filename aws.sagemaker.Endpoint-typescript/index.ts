import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const endpoint = new aws.sagemaker.Endpoint("e", {
    endpointConfigName: aws_sagemaker_endpoint_configuration_ec.name,
    tags: {
        Name: "foo",
    },
});

