import pulumi
import pulumi_aws as aws

default = aws.iam.OpenIdConnectProvider("default",
    client_id_lists=["266362248691-342342xasdasdasda-apps.googleusercontent.com"],
    thumbprint_lists=[],
    url="https://accounts.google.com")

