import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ec = new aws.sagemaker.EndpointConfiguration("ec", {
    productionVariants: [{
        initialInstanceCount: 1,
        instanceType: "ml.t2.medium",
        modelName: aws_sagemaker_model_m.name,
        variantName: "variant-1",
    }],
    tags: {
        Name: "foo",
    },
});

