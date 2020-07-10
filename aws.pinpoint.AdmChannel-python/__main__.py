import pulumi
import pulumi_aws as aws

app = aws.pinpoint.App("app")
channel = aws.pinpoint.AdmChannel("channel",
    application_id=app.application_id,
    client_id="",
    client_secret="",
    enabled=True)

