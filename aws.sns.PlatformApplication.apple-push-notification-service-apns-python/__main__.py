import pulumi
import pulumi_aws as aws

apns_application = aws.sns.PlatformApplication("apnsApplication",
    platform="APNS",
    platform_credential="<APNS PRIVATE KEY>",
    platform_principal="<APNS CERTIFICATE>")

