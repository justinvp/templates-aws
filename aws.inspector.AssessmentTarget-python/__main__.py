import pulumi
import pulumi_aws as aws

bar = aws.inspector.ResourceGroup("bar", tags={
    "Env": "bar",
    "Name": "foo",
})
foo = aws.inspector.AssessmentTarget("foo", resource_group_arn=bar.arn)

