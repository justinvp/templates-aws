import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const subnetId = config.require("subnetId");

const defaultNatGateway = aws_subnet_public.id.apply(id => aws.ec2.getNatGateway({
    subnetId: id,
}, { async: true }));

