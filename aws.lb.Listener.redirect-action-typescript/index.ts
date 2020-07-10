import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const frontEndLoadBalancer = new aws.lb.LoadBalancer("front_end", {});
const frontEndListener = new aws.lb.Listener("front_end", {
    defaultActions: [{
        redirect: {
            port: "443",
            protocol: "HTTPS",
            statusCode: "HTTP_301",
        },
        type: "redirect",
    }],
    loadBalancerArn: frontEndLoadBalancer.arn,
    port: 80,
    protocol: "HTTP",
});

