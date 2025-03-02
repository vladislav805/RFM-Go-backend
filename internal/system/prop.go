package system

import (
	"fmt"
	"os/exec"
)

func GetProp(name string) (string, error) {
	cmd := exec.Command("/system/bin/getprop", name)

	out, err := cmd.Output()

	if err != nil {
		fmt.Printf("GetProp(%s)err = %v\n", name, err)
		return "", err
	}

	// Значение возвращается вместе с \n на конце
	value := string(out[:len(out)-1])

	fmt.Printf("GetProp(%s)out = %v\n", name, value)

	return value, nil
}

func SetProp(name, value string) error {
	cmd := exec.Command("/system/bin/setprop", name, value)

	_, err := cmd.Output()

	if err != nil {
		return err
	}

	exitCode := cmd.ProcessState.ExitCode()

	fmt.Printf("SetProp(%s, %s) = %d, %v\n", name, value, exitCode, err)

	if exitCode > 0 {
		return fmt.Errorf("SetProp: exit code %d", exitCode)
	}

	return nil
}
