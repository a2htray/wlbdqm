package wlbdqm

import (
	"log"
	"testing"
)

func TestRunDiskQuota(t *testing.T) {
	output, err := RunDiskQuota()
	if err != nil {
		t.Fatal(err)
	}

	log.Println(output)
}

func TestParseDiskQuotaOutput(t *testing.T) {
	var mockOutput = `Filesystem	type	blocks	quota	limit	in_doubt	grace	|	files	quota	limit	in_doubt	grace	Remarks
public	USR	15.39T	45T	50T	11.71G	expired	|	645614	7500000	8000000	14238	none`
	dq, err := ParseDiskQuotaOutput(mockOutput)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(dq.HTMLTable())

}
