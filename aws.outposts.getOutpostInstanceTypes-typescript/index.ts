import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.outposts.getOutpostInstanceTypes({
    arn: data.aws_outposts_outpost.example.arn,
});

