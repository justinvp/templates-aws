import pulumi
import pulumi_aws as aws

my_template = aws.ses.Template("myTemplate",
    html="<h1>Hello {{name}},</h1><p>Your favorite animal is {{favoriteanimal}}.</p>",
    subject="Greetings, {{name}}!",
    text="""Hello {{name}},
Your favorite animal is {{favoriteanimal}}.
""")

