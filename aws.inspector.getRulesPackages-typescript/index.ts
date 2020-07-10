import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Declare the data source
const rules = pulumi.output(aws.inspector.getRulesPackages({ async: true }));
// e.g. Use in aws_inspector_assessment_template
const group = new aws.inspector.ResourceGroup("group", {
    tags: {
        test: "test",
    },
});
const assessmentAssessmentTarget = new aws.inspector.AssessmentTarget("assessment", {
    resourceGroupArn: group.arn,
});
const assessmentAssessmentTemplate = new aws.inspector.AssessmentTemplate("assessment", {
    duration: 60,
    rulesPackageArns: rules.arns,
    targetArn: assessmentAssessmentTarget.arn,
});

