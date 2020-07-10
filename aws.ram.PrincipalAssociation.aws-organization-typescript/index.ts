import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ram.PrincipalAssociation("example", {
    principal: aws_organizations_organization_example.arn,
    resourceShareArn: aws_ram_resource_share_example.arn,
});

