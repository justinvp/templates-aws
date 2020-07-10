import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myTemplate = new aws.ses.Template("MyTemplate", {
    html: "<h1>Hello {{name}},</h1><p>Your favorite animal is {{favoriteanimal}}.</p>",
    subject: "Greetings, {{name}}!",
    text: `Hello {{name}},
Your favorite animal is {{favoriteanimal}}.`,
});

