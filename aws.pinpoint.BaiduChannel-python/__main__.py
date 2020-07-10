import pulumi
import pulumi_aws as aws

app = aws.pinpoint.App("app")
channel = aws.pinpoint.BaiduChannel("channel",
    api_key="",
    application_id=app.application_id,
    secret_key="")

