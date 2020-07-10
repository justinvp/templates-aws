import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const frontEndLoadBalancer = new aws.lb.LoadBalancer("front_end", {});
const frontEndTargetGroup = new aws.lb.TargetGroup("front_end", {});
const frontEndListener = new aws.lb.Listener("front_end", {
    certificateArn: "arn:aws:iam::187416307283:server-certificate/test_cert_rab3wuqwgja25ct3n4jdj2tzu4",
    defaultActions: [{
        targetGroupArn: frontEndTargetGroup.arn,
        type: "forward",
    }],
    loadBalancerArn: frontEndLoadBalancer.arn,
    port: 443,
    protocol: "HTTPS",
    sslPolicy: "ELBSecurityPolicy-2016-08",
});

