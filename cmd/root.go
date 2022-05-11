package cmd

import (
	"fmt"
	"github.com/mailjet/mailjet-apiv3-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"regexp"
)

var (
	from          string
	to            string
	subject       string
	text          string
	mailjetClient *mailjet.Client
)

const (
	emailRegexp = "^[\\w-.]+@([\\w-]+\\.)+[\\w-]{2,4}$"
)

var rootCmd = &cobra.Command{
	Use: "mailjet-cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("running!")
	},
}

var sendCmd = &cobra.Command{
	Use: "send",
	Run: func(cmd *cobra.Command, args []string) {
		err := validateEmails(to, from)
		if err != nil {
			log.Fatal(err)
		}

		msg := generateMessage(from, to, subject, text)
		res, err := mailjetClient.SendMailV31(msg)
		if err != nil {
			log.Fatal(err)

		}

		if res.ResultsV31[0].Status != "success" {
			log.Fatal("something went wrong when sending email")
		}
	},
}

func setup() {
	viper.SetEnvPrefix("mailjet")
	viper.AutomaticEnv()

	if !viper.IsSet("apikey") || !viper.IsSet("secret") {
		log.Fatal("missing credentials")
	}

	mailjetClient = mailjet.NewMailjetClient(viper.GetString("apikey"), viper.GetString("secret"))
}

func init() {
	cobra.OnInitialize(setup)

	sendCmd.Flags().StringVar(&from, "from", "", "e-mail from this address")
	sendCmd.Flags().StringVar(&to, "to", "", "e-mail to this address")
	sendCmd.Flags().StringVar(&text, "text", "", "contents of e-mail")
	sendCmd.Flags().StringVar(&subject, "subject", "", "subject of e-mail")
	sendCmd.MarkFlagRequired("from")
	sendCmd.MarkFlagRequired("to")

	rootCmd.AddCommand(sendCmd)
}

func validateEmails(emails ...string) error {
	r, err := regexp.Compile(emailRegexp)
	if err != nil {
		return err
	}

	for _, e := range emails {
		t := r.MatchString(e)
		if !t {
			return fmt.Errorf("invalid email address %s\n", e)
		}
	}

	return nil
}

func generateMessage(from string, to string, subject string, text string) *mailjet.MessagesV31 {
	return &mailjet.MessagesV31{Info: []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: from,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: to,
				},
			},
			Subject:  subject,
			TextPart: text,
		},
	}}
}

func Execute() error {
	return rootCmd.Execute()
}
