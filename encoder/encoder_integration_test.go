package encoder

import (
	"github.com/onsi/gomega/gexec"
	"os/exec"
	"testing"
	"io/ioutil"
	"bytes"
)

const test_image_path = "./qr_111111.png"
const expected_image_path = "/tmp/integration_test.png"
const correction_lvl = "L"
const mask_pattern = "0"
const input_text = "111111"

var pathToBinary string

func Test_GeneratedQR(t *testing.T) {
	//Compile binary
	var err error
	pathToBinary, err = gexec.Build("github.com/anydef/qr")
	if err != nil {
		t.Fatalf("Not able to compile, reason: %s", err)
	}

	//Run compiled binary with parameters
	command := exec.Command(pathToBinary, input_text, expected_image_path, correction_lvl, mask_pattern)
	err = command.Run()
	if err != nil {
		t.Fatalf("Command exited with error")
	}

	test_image, err := ioutil.ReadFile(test_image_path)
	if err != nil {
		t.Fatalf("Could not load test image: %s. Err: %s", test_image_path, err)
	}

	expected_image, err := ioutil.ReadFile(expected_image_path)
	if err != nil {
		t.Fatalf("Could not load expected image: %s. Err: %s", expected_image_path, err)
	}

	if bytes.Compare(test_image, expected_image) != 0 {
		t.Fatalf("Image produced by programm doesnt match expected. "+
			"\nTest\t\t%b"+
			"\nExpected\t%b",
			test_image, expected_image)
	}

	//Clean up build
	gexec.CleanupBuildArtifacts()
}
