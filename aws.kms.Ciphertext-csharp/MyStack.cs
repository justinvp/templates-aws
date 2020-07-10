using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var oauthConfig = new Aws.Kms.Key("oauthConfig", new Aws.Kms.KeyArgs
        {
            Description = "oauth config",
            IsEnabled = true,
        });
        var oauth = new Aws.Kms.Ciphertext("oauth", new Aws.Kms.CiphertextArgs
        {
            KeyId = oauthConfig.KeyId,
            Plaintext = @"{
  ""client_id"": ""e587dbae22222f55da22"",
  ""client_secret"": ""8289575d00000ace55e1815ec13673955721b8a5""
}

",
        });
    }

}

