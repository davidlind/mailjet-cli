# mailjet-cli

Simple program for sending email via the Mailjet API. Built with [cobra](https://github.com/spf13/cobra),
[viper](https://github.com/spf13/viper) and [mailjet-apiv3-go](https://github.com/mailjet/mailjet-apiv3-go).

I am in no way affiliated with Mailjet.

# Installation
You need to have Go installed to be able to build.

Just build the package to get the executable:

`$ go build`

That gives you a `mailjet-cli` binary that you can run.

# Usage

mailjet-cli only sends pure text, I've not implemented sending of html since I don't need it myself at this moment.

`$ MAILJET_APIKEY=mykey MAILJET_SECRET=mysecret ./mailjet-cli send --from me@mydomain.com --to someone@else.com
--subject "mailjet-cli!" --text "testing the mailjet-cli" `

# Bugs
Yes. This project is poorly tested(there are no tests, yet).