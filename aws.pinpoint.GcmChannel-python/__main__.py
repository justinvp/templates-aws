import pulumi
import pulumi_aws as aws

app = aws.pinpoint.App("app")
gcm = aws.pinpoint.GcmChannel("gcm",
    api_key="api_key",
    application_id=app.application_id)

