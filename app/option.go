package app

import "time"

//go:generate optiongen --xconf=true --usage_tag_name=usage
func OptionsOptionDeclareWithDefault() interface{} {
	return map[string]interface{}{
		// annotation@EnvName(comment="use value environment variable as password")
		"EnvName": "",
		// annotation@Password(comment="plaintext password (not recommended)")
		"Password": "",
		// annotation@Timeout(comment="timeout length to wait for prompt/confirmation")
		"Timeout": time.Duration(time.Second * 10),
		// annotation@DisableConfirmHostAuthenticity(xconf="disable-ssh-host-confirm",comment="sshpass will automatically confirm the authenticity of SSH hosts unless this option is specified")
		"DisableConfirmHostAuthenticity": false,
		// annotation@Shell(comment="Shell is a path to the shell to use e.g. /bin/bash - leave blank to use user shell")
		"Shell": "",
		// annotation@ExpectedPrompt(comment="the string to treat as the password prompt")
		"ExpectedPrompt": "password:",
		// annotation@ExpectedFailure(comment="the string to treat as an indication of failure")
		"ExpectedFailure": "denied",
	}
}
