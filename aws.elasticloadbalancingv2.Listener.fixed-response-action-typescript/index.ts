import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const frontEndLoadBalancer = new aws.lb.LoadBalancer("front_end", {});
const frontEndListener = new aws.lb.Listener("front_end", {
    defaultActions: [{
        fixedResponse: {
            contentType: "text/plain",
            messageBody: "Fixed response content",
            statusCode: "200",
        },
        type: "fixed-response",
    }],
    loadBalancerArn: frontEndLoadBalancer.arn,
    port: 80,
    protocol: "HTTP",
});

