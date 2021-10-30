package wlbdqm

import (
	"log"
	"os"
	"testing"
)

func TestPrepareDiskInsufficientContent(t *testing.T) {
	contentItem := ContentItem{
		MaxPercentage: 80,
		DPercentage:   81.22,
		FPercentage:   15.01,
		DPOutput: `Filesystem	type	blocks	quota	limit	in_doubt	grace	|	files	quota	limit	in_doubt	grace	Remarks
public	USR	15.39T	45T	50T	11.71G	expired	|	645614	7500000	8000000	14238	none`,
	}

	content, err := PrepareDiskInsufficientContent(contentItem)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(content)
}

func TestSendEmail(t *testing.T) {
	var mockOutput = `Filesystem	type	blocks	quota	limit	in_doubt	grace	|	files	quota	limit	in_doubt	grace	Remarks
public	USR	15.39T	45T	50T	11.71G	expired	|	645614	7500000	8000000	14238	none`
	dq, err := ParseDiskQuotaOutput(mockOutput)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(dq.HTMLTable())

	contentItem := ContentItem{
		MaxPercentage: 80,
		DPercentage:   81.22,
		FPercentage:   15.01,
		DPOutput:      dq.HTMLTable(),
	}

	content, err := PrepareDiskInsufficientContent(contentItem)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Setenv(MailKeyFrom, "a2htray@outlook.com")
	err = os.Setenv(MailKeyPassword, "MY PASSWORD")
	err = os.Setenv(MailKeyHost, "smtp-mail.outlook.com")
	err = os.Setenv(MailKeyPort, "587")

	err = SendEmail([]string{
		"a2htray@outlook.com",
	}, content)

	if err != nil {
		log.Fatal(err)
	}

}
