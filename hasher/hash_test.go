package hasher

import (
	"log"
	"os/exec"
	"testing"
)

func TestHash(t *testing.T) {
	var strInput1 string = "MySwissBankAccountPassword"
	expected := bashWrapper(strInput1)
	actual := MD5Hash(strInput1)
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	var strInput2 string = "MyPanamaPapersAccountPassword"
	expected = bashWrapper(strInput2)
	actual = MD5Hash(strInput2)
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

// bashWrapper : wrapper for BashPipedCommand
func bashWrapper(strInput string) string {

	echo := exec.Command("echo", strInput)
	md5 := exec.Command("md5sum")
	awk := exec.Command("awk", `{printf $1}`) // avoid newline pitfall from awk print output

	output, _, err := BashPipedCommand(echo, md5, awk)
	if err != nil {
		log.Println(err.Error())
	}

	if len(output) > 0 {
		return string(output)
	}
	return ""
}
