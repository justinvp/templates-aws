import pulumi
import pulumi_aws as aws

foo_with_lifecyle_policy = aws.efs.FileSystem("fooWithLifecylePolicy", lifecycle_policy={
    "transitionToIa": "AFTER_30_DAYS",
})

