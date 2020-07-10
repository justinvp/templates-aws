import pulumi
import pulumi_aws as aws

gcm_application = aws.sns.PlatformApplication("gcmApplication",
    platform="GCM",
    platform_credential="<GCM API KEY>")

