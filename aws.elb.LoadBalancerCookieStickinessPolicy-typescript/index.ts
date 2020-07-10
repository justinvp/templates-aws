import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const lb = new aws.elb.LoadBalancer("lb", {
    availabilityZones: ["us-east-1a"],
    listeners: [{
        instancePort: 8000,
        instanceProtocol: "http",
        lbPort: 80,
        lbProtocol: "http",
    }],
});
const foo = new aws.elb.LoadBalancerCookieStickinessPolicy("foo", {
    cookieExpirationPeriod: 600,
    lbPort: 80,
    loadBalancer: lb.id,
});

