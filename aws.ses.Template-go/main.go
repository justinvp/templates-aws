package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ses"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ses.NewTemplate(ctx, "myTemplate", &ses.TemplateArgs{
			Html:    pulumi.String("<h1>Hello {{name}},</h1><p>Your favorite animal is {{favoriteanimal}}.</p>"),
			Subject: pulumi.String("Greetings, {{name}}!"),
			Text: pulumi.String(fmt.Sprintf("%v%v", "Hello {{name}},\n", "Your favorite animal is {{favoriteanimal}}.\n")),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

