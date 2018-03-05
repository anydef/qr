package intergation_test

import (
	"github.com/onsi/gomega/gexec"
	"os/exec"
	"testing"
	"io/ioutil"
	"bytes"
)

const test_image_path = "./qr_111111.png"
const expected_image_path = "/tmp/integration_test.png"
const expected_image_path_option = "-output=" + expected_image_path
const correction_lvl = "-correction-level=L"
const mask_pattern = "-mask-pattern=0"
const input_text = "-input=111111"

var pathToBinary string

func Test_GeneratedQR_Numeric(t *testing.T) {
	//Compile binary
	var err error
	pathToBinary, err = gexec.Build("github.com/anydef/qr")
	defer gexec.CleanupBuildArtifacts()
	if err != nil {
		t.Fatalf("Not able to compile, reason: %s", err)
	}

	//Run compiled binary with parameters
	command := exec.Command(pathToBinary, input_text, expected_image_path_option, correction_lvl, mask_pattern)
	err = command.Run()
	if err != nil {
		t.Fatalf("Command exited with error, %s", err)
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
}
