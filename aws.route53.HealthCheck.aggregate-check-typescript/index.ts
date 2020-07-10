import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const parent = new aws.route53.HealthCheck("parent", {
    childHealthThreshold: 1,
    childHealthchecks: [aws_route53_health_check_child.id],
    tags: {
        Name: "tf-test-calculated-health-check",
    },
    type: "CALCULATED",
});

