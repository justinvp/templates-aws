import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const frontEndLoadBalancer = new aws.lb.LoadBalancer("front_end", {});
const frontEndTargetGroup = new aws.lb.TargetGroup("front_end", {});
const pool = new aws.cognito.UserPool("pool", {});
const client = new aws.cognito.UserPoolClient("client", {});
const domain = new aws.cognito.UserPoolDomain("domain", {});
const frontEndListener = new aws.lb.Listener("front_end", {
    defaultActions: [
        {
            authenticateCognito: {
                userPoolArn: pool.arn,
                userPoolClientId: client.id,
                userPoolDomain: domain.domain,
            },
            type: "authenticate-cognito",
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

