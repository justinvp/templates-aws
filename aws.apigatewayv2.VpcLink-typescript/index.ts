import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.apigatewayv2.VpcLink("example", {
    securityGroupIds: [data.aws_security_group.example.id],
    subnetIds: data.aws_subnet_ids.example.ids,
    tags: {
        Usage: "example",
    },
});

