# mailjet-cli

Simple program for sending email via the Mailjet API. Built with [cobra](https://github.com/spf13/cobra),
[viper](https://github.com/spf13/viper) and [mailjet-apiv3-go](https://github.com/mailjet/mailjet-apiv3-go).

I am in no way affiliated with Mailjet.

# Installation
Just build the package to get the executable:

`$ go build`

That gives you a `mailjet-cli` binary that you can run.

# Usage
`$ MAILJET_APIKEY=mykey MAILJET_SECRET=mysecret ./mailjet-cli send --from me@mydomain.com --to someone@else.com
--subject "mailjet-cli!" --text "testing the mailjet-cli" `