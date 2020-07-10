import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.inspector.AssessmentTemplate("example", {
    duration: 3600,
    rulesPackageArns: [
        "arn:aws:inspector:us-west-2:758058086616:rulespackage/0-9hgA516p",
        "arn:aws:inspector:us-west-2:758058086616:rulespackage/0-H5hpSawc",
        "arn:aws:inspector:us-west-2:758058086616:rulespackage/0-JJOtZiqQ",
        "arn:aws:inspector:us-west-2:758058086616:rulespackage/0-vg5GGHSD",
    ],
    targetArn: aws_inspector_assessment_target_example.arn,
});

