# mailjet-cli

Simple program for sending email via the Mailjet API. Built with [cobra](https://github.com/spf13/cobra),
[viper](https://github.com/spf13/viper) and [mailjet-apiv3-go](https://github.com/mailjet/mailjet-apiv3-go).

I am in no way affiliated with Mailjet.

# Usage
`$ MAILJET_APIKEY=mykey MAILJET_SECRET=mysecret ./mailjet-cli send --from me@mydomain.com --to someone@else.com
--subject "mailjet-cli!" --text "testing the mailjet-cli" `