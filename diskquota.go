package wlbdqm

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// Filesystem	type	blocks	quota	limit	in_doubt	grace	|	files	quota	limit	in_doubt	grace	Remarks

const (
	dqFieldFilesystemIdx = iota
	dqFieldTypeIdx
	dqFieldDBlocksIdx
	dqFieldDQuotaIdx
	dqFieldDLimitIdx
	dqFieldDInDoubtIdx
	dqFieldDGraceIdx
	dqFieldUnused
	dqFieldFFilesIdx
	dqFieldFQuotaIdx
	dqFieldFLimitIdx
	dqFieldFInDoubtIdx
	dqFieldFGraceIdx
	dqFieldRemarksIdx
)

var dqHeader = "Filesystem\ttype\tblocks\tquota\tlimit\tin_doubt\tgrace\t|\tfiles\tquota\tlimit\tin_doubt\tgrace\tRemarks"

func RunDiskQuota() (string, error) {
	cmd := exec.Command("diskquota")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

// DQOutput
// Filesystem	type	blocks	quota	limit	in_doubt	grace
// files	quota	limit	in_doubt	grace	Remarks
type DQOutput struct {
	Filesystem string
	Type       string
	DBlocks    string
	DQuota     string
	DLimit     string
	DInDoubt   string
	DGrace     string
	FFiles     string
	FQuota     string
	FLimit     string
	FInDoubt   string
	FGrace     string
	Remarks    string
}

func (dq DQOutput) ByteOfBlocks() float64 {
	f, err := ParseSizeToByte(dq.DBlocks)
	if err != nil {
		panic(err)
	}
	return f
}

func (dq DQOutput) ByteOfDQuota() float64 {
	f, err := ParseSizeToByte(dq.DQuota)
	if err != nil {
		panic(err)
	}
	return f
}

func (dq DQOutput) NumOfFFiles() int64 {
	DebugPrintln("dq.FFiles", dq.FFiles)
	i, err := strconv.ParseInt(dq.FFiles, 10, 64)
	if err != nil {
		panic(err)
	}

	DebugPrintln("dq.FFiles returns", i)
	return i
}

func (dq DQOutput) NumOfFQuota() int64 {
	DebugPrintln("dq.FQuota", dq.FQuota)
	i, err := strconv.ParseInt(dq.FQuota, 10, 64)
	if err != nil {
		panic(err)
	}

	DebugPrintln("dq.FQuota returns", i)
	return i
}

func (dq DQOutput) String() string {
	s1 := dqHeader
	s2 := fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s", dq.Filesystem, dq.Type, dq.DBlocks,
		dq.DQuota, dq.DLimit, dq.DInDoubt, dq.DGrace, "|", dq.FFiles, dq.FQuota, dq.FLimit, dq.FInDoubt, dq.FGrace, dq.Remarks)

	return s1 + "\n" + s2
}

// Filesystem	type	blocks	quota	limit	in_doubt	grace
// files	quota	limit	in_doubt	grace	Remarks
func (dq DQOutput) HTMLTable() string {
	return `<table>
	<tr>
		<th>Filesystem</th>
		<th>type</th>
		<th>blocks</th>
		<th>quota</th>
		<th>limit</th>
		<th>in_doubt</th>
		<th>grace</th>
		<th>|</th>
		<th>files</th>
		<th>quota</th>
		<th>limit</th>
		<th>in_doubt</th>
		<th>grace</th>
		<th>Remarks</th>
	</tr>
	<tr>
		<td>` + dq.Filesystem + `</td>
		<td>` + dq.Type + `</td>
		<td>` + dq.DBlocks + `</td>
		<td>` + dq.DQuota + `</td>
		<td>` + dq.DLimit + `</td>
		<td>` + dq.DInDoubt + `</td>
		<td>` + dq.DGrace + `</td>
		<td>|</td>
		<td>` + dq.FFiles + `</td>
		<td>` + dq.FQuota + `</td>
		<td>` + dq.FLimit + `</td>
		<td>` + dq.FInDoubt + `</td>
		<td>` + dq.FGrace + `</td>
		<td>` + dq.Remarks + `</td>
	</tr>
</table>`
}

func ParseDiskQuotaOutput(output string) (*DQOutput, error) {
	lines := strings.Split(output, "\n")
	DebugPrintln("output line number", len(lines))
	if len(lines) < 2 {
		return nil, errors.New("output line number error")
	}

	fieldValues := strings.Split(lines[1], "\t")
	DebugPrintln(fieldValues)

	remarks := ""
	if len(fieldValues) > 13 {
		remarks = fieldValues[dqFieldRemarksIdx]
	}

	dqOutput := &DQOutput{
		Filesystem: fieldValues[dqFieldFilesystemIdx],
		Type:       fieldValues[dqFieldTypeIdx],
		DBlocks:    fieldValues[dqFieldDBlocksIdx],
		DQuota:     fieldValues[dqFieldDQuotaIdx],
		DLimit:     fieldValues[dqFieldDLimitIdx],
		DInDoubt:   fieldValues[dqFieldDInDoubtIdx],
		DGrace:     fieldValues[dqFieldDGraceIdx],
		FFiles:     fieldValues[dqFieldFFilesIdx],
		FQuota:     fieldValues[dqFieldFQuotaIdx],
		FLimit:     fieldValues[dqFieldFLimitIdx],
		FInDoubt:   fieldValues[dqFieldFInDoubtIdx],
		FGrace:     fieldValues[dqFieldFGraceIdx],
		Remarks:    remarks,
	}

	return dqOutput, nil
}
