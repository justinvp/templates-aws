import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleLoadBalancer = new aws.lb.LoadBalancer("example", {
    internal: true,
    loadBalancerType: "network",
    subnetMappings: [{
        subnetId: "12345",
    }],
});
const exampleVpcLink = new aws.apigateway.VpcLink("example", {
    description: "example description",
    targetArn: exampleLoadBalancer.arn,
});

