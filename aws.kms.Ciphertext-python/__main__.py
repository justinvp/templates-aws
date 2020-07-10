import pulumi
import pulumi_aws as aws

oauth_config = aws.kms.Key("oauthConfig",
    description="oauth config",
    is_enabled=True)
oauth = aws.kms.Ciphertext("oauth",
    key_id=oauth_config.key_id,
    plaintext="""{
  "client_id": "e587dbae22222f55da22",
  "client_secret": "8289575d00000ace55e1815ec13673955721b8a5"
}

""")

