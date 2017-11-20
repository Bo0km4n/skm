package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestUsage(t *testing.T) {
	tmp := prepareTest(t)
	defer os.RemoveAll(tmp)
	t.Log(tmp)
	cmd := exec.Command(fmt.Sprintf("%s/bin/skm", tmp), "-h")
	_, err := cmd.CombinedOutput()
	if err != nil {
		t.Log(err)
		t.Fatal("Expected exit code 1 bot 0")
	}
}

func TestInvalidArgs(t *testing.T) {
	tmp := prepareTest(t)
	expectString := "No help topic for 'hogehoge'\n"
	defer os.RemoveAll(tmp)
	cmd := exec.Command(fmt.Sprintf("%s/bin/skm", tmp), "hogehoge")
	b, _ := cmd.CombinedOutput()

	if expectString != string(b) {
		t.Log(string(b))
		t.Fatalf("Expected string is : %s", expectString)
	}
}

func prepareTest(t *testing.T) (tmpPath string) {
	tmp := os.TempDir()
	tmp = filepath.Join(tmp, uuid.NewV4().String())
	runCmd(t, "go", "build", "-o", filepath.Join(tmp, "bin", "skm"), "github.com/TimothyYe/skm")
	os.MkdirAll(filepath.Join(tmp, "bin"), 0755)
	return tmp
}

func runCmd(t *testing.T, cmd string, args ...string) []byte {
	b, err := exec.Command(cmd, args...).CombinedOutput()
	if err != nil {
		t.Fatalf("Expected %v, but %v: %v", nil, err, string(b))
	}
	return b
}
