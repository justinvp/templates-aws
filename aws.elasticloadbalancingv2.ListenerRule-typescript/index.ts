import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const frontEndLoadBalancer = new aws.lb.LoadBalancer("front_end", {});
const frontEndListener = new aws.lb.Listener("front_end", {});
const static = new aws.lb.ListenerRule("static", {
    actions: [{
        targetGroupArn: aws_lb_target_group_static.arn,
        type: "forward",
    }],
    conditions: [
        {
            pathPattern: {
                values: ["/static/*"],
            },
        },
        {
            hostHeader: {
                values: ["example.com"],
            },
        },
    ],
    listenerArn: frontEndListener.arn,
    priority: 100,
});
const hostBasedRouting = new aws.lb.ListenerRule("host_based_routing", {
    actions: [{
        forward: {
            stickiness: {
                duration: 600,
                enabled: true,
            },
            targetGroups: [
                {
                    arn: aws_lb_target_group_main.arn,
                    weight: 80,
                },
                {
                    arn: aws_lb_target_group_canary.arn,
                    weight: 20,
                },
            ],
        },
        type: "forward",
    }],
    conditions: [{
        hostHeader: {
            values: ["my-service.*.mycompany.io"],
        },
    }],
    listenerArn: frontEndListener.arn,
    priority: 99,
});
const hostBasedWeightedRouting = new aws.lb.ListenerRule("host_based_weighted_routing", {
    actions: [{
        targetGroupArn: aws_lb_target_group_static.arn,
        type: "forward",
    }],
    conditions: [{
        hostHeader: {
            values: ["my-service.*.mydomain.io"],
        },
    }],
    listenerArn: frontEndListener.arn,
    priority: 99,
});
const redirectHttpToHttps = new aws.lb.ListenerRule("redirect_http_to_https", {
    actions: [{
        redirect: {
            port: "443",
            protocol: "HTTPS",
            statusCode: "HTTP_301",
        },
        type: "redirect",
    }],
    conditions: [{
        httpHeader: {
            httpHeaderName: "X-Forwarded-For",
            values: ["192.168.1.*"],
        },
    }],
    listenerArn: frontEndListener.arn,
});
const healthCheck = new aws.lb.ListenerRule("health_check", {
    actions: [{
        fixedResponse: {
            contentType: "text/plain",
            messageBody: "HEALTHY",
            statusCode: "200",
        },
        type: "fixed-response",
    }],
    conditions: [{
        queryStrings: [
            {
                key: "health",
                value: "check",
            },
            {
                value: "bar",
            },
        ],
    }],
    listenerArn: frontEndListener.arn,
});
const pool = new aws.cognito.UserPool("pool", {});
const client = new aws.cognito.UserPoolClient("client", {});
const domain = new aws.cognito.UserPoolDomain("domain", {});
const admin = new aws.lb.ListenerRule("admin", {
    actions: [
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
            targetGroupArn: aws_lb_target_group_static.arn,
            type: "forward",
        },
    ],
    listenerArn: frontEndListener.arn,
});

