package main

import (
	"os"
	"testing"
)

func TestReadDir(t *testing.T) {
	var envFiles = make(map[string][]byte, 3)
	envFiles["TEST_ENV_NAME_1"] = []byte("TEST_ENV_VAL_1")
	envFiles["TEST_ENV_NAME_2"] = []byte("TEST_ENV_VAL_2")
	envFiles["TEST_ENV_NAME_3"] = []byte("TEST_ENV_VAL_3")

	err := os.Mkdir("test_env_dir", 0777)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll("test_env_dir")

	for name, content := range envFiles {
		if err = createFile("test_env_dir/"+name, content); err != nil {
			t.Fatal(err)
		}
	}

	env, err := ReadDir("test_env_dir/")
	if err != nil {
		t.Error(err)
	}
	for name, val := range env {
		if string(envFiles[name]) != val {
			t.Error("unexpected result: ", envFiles[name], []byte(val))
		}
	}
}

func TestRunCmd(t *testing.T) {
	stdout, err := RunCmd([]string{"env"}, map[string]string{"USER": "ne_gena"})
	if err != nil {
		t.Fatal(stdout)
	}
	if string(stdout) != "USER=ne_gena\n" {
		t.Error("unexpected result: ", stdout, []byte("USER=ne_gena"))
	}
}

func createFile(path string, content []byte) error {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		return err
	}

	_, err = file.Write(content)
	return err
}
