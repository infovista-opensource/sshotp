# SSHOTP: Automatic entry of non-interactive passwords

Autopass is essentially a go implementation of [sshpass](https://linux.die.net/man/1/sshpass), though unlike sshpass it doesn't restrict itself to SSH logins. It can supply a password to any process with an identifiable password prompt.

**Do not use this unless you understand the risks involved - ssh prompts for a password interactively for a reason!**

The original use case for this was needing to automate the acquisition and use of an SSH OTP (via vault) in a nice script.

## Requirements

- Mac/Linux
- Go 1.11+ (to build)

## Install

```bash
go install github.com/timestee/sshotp@latest
```

## Example

```bash
sshotp --password mypassword123 "ssh me@myserver.mine -p 2222"
```

## Usage

```
USAGE: 
      sshpass <flags> <args>

DESCRIPTION:
      sshotp is essentially a go implementation of sshpass (https://linux.die.net/man/1/sshpass).
      Though unlike sshpass it doesn't restrict itself to SSH logins.
      It can supply a password to any process with an identifiable password prompt.

OPTIONS GLOBAL:
      -----------------------------------------------------------------------------------------------------------------------------------------------------------
      FLAG                         TYPE      USAGE
      -----------------------------------------------------------------------------------------------------------------------------------------------------------
      --xconf_flag_files           string    |M| xconf files provided by flag, file slice, split by ,
      -----------------------------------------------------------------------------------------------------------------------------------------------------------

OPTIONS LOCAL:
      -----------------------------------------------------------------------------------------------------------------------------------------------------------
      FLAG                         TYPE      USAGE
      -----------------------------------------------------------------------------------------------------------------------------------------------------------
      --disable-ssh-host-confirm   bool      |Y| sshpass will automatically confirm the authenticity of SSH hosts unless this option is specified (default false)
      --env_name                   string    |Y| use value environment variable as password
      --expected_failure           string    |Y| the string to treat as an indication of failure (default "denied")
      --expected_prompt            string    |Y| the string to treat as the password prompt (default "password:")
      --password                   string    |Y| plaintext password (not recommended)
      --shell                      string    |Y| Shell is a path to the shell to use e.g. /bin/bash - leave blank to use user shell
      --timeout                    Duration  |Y| timeout length to wait for prompt/confirmation (default 10s)
      -----------------------------------------------------------------------------------------------------------------------------------------------------------

```
