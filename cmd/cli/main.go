package main

import (
	"flag"
	"fmt"
	"log"
	"mailer_ms/pkg/mailer"
	"mailer_ms/pkg/mailer/static"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/joho/godotenv"
)

const (
	ERR_LOADING_ENV                   = "Error loading .env file"
	ERR_PARSING_TEMPLATE              = "Error parsing template"
	ERR_SENDING_EMAIL                 = "Error during sending email"
	ERR_ADDING_ATTACHMENT             = "Error while adding attachment"
	ERR_SENDING_EMAIL_WITH_ATTACHMENT = "Error during sending email with attachment"
)

func main() {
	// Initialize spinner for loading indication
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start()
	s.Prefix = "⌛ Loading... retrieving email configuration"

	// Define command-line flags
	from := flag.String("from", "", "sender email address")
	to := flag.String("to", "", "receiver email address, separated by comma")
	subject := flag.String("subject", "", "email subject")
	attachment := flag.String("attachment", "", "attachment file path, separated by comma")

	flag.Parse()

	// Check if no flags are provided
	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(1)
	}

	// Check required flags
	checkFlags(map[string]*string{
		"from":    from,
		"to":      to,
		"subject": subject,
	})

	s.Prefix = fmt.Sprintf("⌛ Sending email from %s to %s...", *from, *to)
	time.Sleep(2 * time.Second)

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal(ERR_LOADING_ENV)
	}

	receivers := strings.Split(*to, ",")
	r := mailer.NewRequest(*subject, receivers).AddTemplateFromString(static.DefaultsTemplate)

	if err != nil {
		log.Fatal(ERR_PARSING_TEMPLATE)
	}

	m := mailer.NewMailer()

	// Send email with or without attachment
	if *attachment == "" {
		err = m.SendEmail(r)
		if err != nil {
			log.Fatal(ERR_SENDING_EMAIL)
		}
	} else {
		for _, file := range strings.Split(*attachment, ",") {
			if err := r.AddAttachement(file); err != nil {
				log.Fatal(ERR_ADDING_ATTACHMENT)
			}
		}

		err = m.SendMailWithAttachment(r)
		if err != nil {
			log.Fatal(ERR_SENDING_EMAIL_WITH_ATTACHMENT)
		}
	}

	s.Stop()
	fmt.Printf("✅ Email sent successfully to %s\n", strings.Join(receivers, ", "))
}

// checkFlags ensures that all required flags are provided
func checkFlags(requiredFlags map[string]*string) {
	missingFlags := []string{}

	for name, value := range requiredFlags {
		if *value == "" {
			missingFlags = append(missingFlags, fmt.Sprintf("-%s", name))
		}
	}

	if len(missingFlags) > 0 {
		fmt.Println("Error: Missing required flag(s):", strings.Join(missingFlags, ", "))
		fmt.Println()
		flag.PrintDefaults()
		fmt.Println()
		os.Exit(1)
	}
}
