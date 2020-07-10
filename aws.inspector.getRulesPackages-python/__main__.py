import pulumi
import pulumi_aws as aws

rules = aws.inspector.get_rules_packages()
# e.g. Use in aws_inspector_assessment_template
group = aws.inspector.ResourceGroup("group", tags={
    "test": "test",
})
assessment_assessment_target = aws.inspector.AssessmentTarget("assessmentAssessmentTarget", resource_group_arn=group.arn)
assessment_assessment_template = aws.inspector.AssessmentTemplate("assessmentAssessmentTemplate",
    duration="60",
    rules_package_arns=rules.arns,
    target_arn=assessment_assessment_target.arn)

