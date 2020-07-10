import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.apigatewayv2.ApiMapping("example", {
    apiId: aws_apigatewayv2_api_example.id,
    domainName: aws_apigatewayv2_domain_name_example.id,
    stage: aws_apigatewayv2_stage_example.id,
});

