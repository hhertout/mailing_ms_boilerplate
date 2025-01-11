package main

import (
	"flag"
	"fmt"
	"log"
	"mailer_ms/pkg/mailer"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/joho/godotenv"
)

func main() {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start()
	s.Prefix = "⌛ Loading... retrieving email configuration"

	from := flag.String("from", "", "sender email address")
	to := flag.String("to", "", "receiver email address, separated by comma")
	subject := flag.String("subject", "", "email subject")
	attachement := flag.String("attachement", "", "attachement file path, separated by comma")
	defaultTemplate := "defaults.html"

	flag.Parse()

	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(1)
	}

	checkFlags(map[string]*string{
		"from":    from,
		"to":      to,
		"subject": subject,
	})

	s.Prefix = fmt.Sprintf("⌛ Sending email from %s to %s...", *from, *to)
	time.Sleep(2 * time.Second)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	receivers := strings.Split(*to, ",")
	r := mailer.NewRequest(*subject, receivers)
	_, err = r.ParseHTMLTemplate(defaultTemplate, nil)
	if err != nil {
		fmt.Println("Error parsing template")
		panic(err)
	}

	m := mailer.NewMailer()

	if *attachement == "" {
		err = m.SendEmail(r)
		if err != nil {
			fmt.Println("Error during sending email")
			panic(err)
		}
	} else {
		for _, file := range strings.Split(*attachement, ",") {
			if err := r.AddAttachement(file); err != nil {
				fmt.Println("Error while adding attachement")
				panic(err)
			}
		}

		err = m.SendMailWithAttachment(r)
		if err != nil {
			fmt.Println("Error during sending email with attachement")
			panic(err)
		}
	}

	s.Stop()
	fmt.Printf("✅ Email sent successfully to %s\n", *to)
}

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
