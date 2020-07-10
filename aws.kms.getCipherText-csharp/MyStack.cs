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
        var oauth = oauthConfig.KeyId.Apply(keyId => Aws.Kms.GetCipherText.InvokeAsync(new Aws.Kms.GetCipherTextArgs
        {
            KeyId = keyId,
            Plaintext = @"{
  ""client_id"": ""e587dbae22222f55da22"",
  ""client_secret"": ""8289575d00000ace55e1815ec13673955721b8a5""
}

",
        }));
    }

}

