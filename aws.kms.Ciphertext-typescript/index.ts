import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const oauthConfig = new aws.kms.Key("oauth_config", {
    description: "oauth config",
    isEnabled: true,
});
const oauth = new aws.kms.Ciphertext("oauth", {
    keyId: oauthConfig.keyId,
    plaintext: `{
  "client_id": "e587dbae22222f55da22",
  "client_secret": "8289575d00000ace55e1815ec13673955721b8a5"
}
`,
});

