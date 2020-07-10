import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const frontEndLoadBalancer = new aws.lb.LoadBalancer("front_end", {});
const frontEndTargetGroup = new aws.lb.TargetGroup("front_end", {});
const frontEndListener = new aws.lb.Listener("front_end", {
    defaultActions: [
        {
            authenticateOidc: {
                authorizationEndpoint: "https://example.com/authorization_endpoint",
                clientId: "client_id",
                clientSecret: "client_secret",
                issuer: "https://example.com",
                tokenEndpoint: "https://example.com/token_endpoint",
                userInfoEndpoint: "https://example.com/user_info_endpoint",
            },
            type: "authenticate-oidc",
        },
        {
            targetGroupArn: frontEndTargetGroup.arn,
            type: "forward",
        },
    ],
    loadBalancerArn: frontEndLoadBalancer.arn,
    port: 80,
    protocol: "HTTP",
});

