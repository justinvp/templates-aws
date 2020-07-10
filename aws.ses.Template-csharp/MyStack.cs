using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var myTemplate = new Aws.Ses.Template("myTemplate", new Aws.Ses.TemplateArgs
        {
            Html = "<h1>Hello {{name}},</h1><p>Your favorite animal is {{favoriteanimal}}.</p>",
            Subject = "Greetings, {{name}}!",
            Text = @"Hello {{name}},
Your favorite animal is {{favoriteanimal}}.
",
        });
    }

}

