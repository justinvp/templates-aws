import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myApiGatewayVpcLink = pulumi.output(aws.apigateway.getVpcLink({
    name: "my-vpc-link",
}, { async: true }));

