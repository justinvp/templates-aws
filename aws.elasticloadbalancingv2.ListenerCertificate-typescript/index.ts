import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleCertificate = new aws.acm.Certificate("example", {});
const frontEndLoadBalancer = new aws.lb.LoadBalancer("front_end", {});
const frontEndListener = new aws.lb.Listener("front_end", {});
const exampleListenerCertificate = new aws.lb.ListenerCertificate("example", {
    certificateArn: exampleCertificate.arn,
    listenerArn: frontEndListener.arn,
});

