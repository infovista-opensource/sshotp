package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/timestee/sshotp/app"

	"github.com/sandwich-go/xconf/xcmd"
)

func Execute() {
	cc := app.NewOptions()
	var rootCmd = xcmd.NewCommand("sshpass",
		xcmd.WithShort("Enter passwords to commands non-interactively"),
		xcmd.WithDescription(`
		sshpass is essentially a go implementation of sshpass (https://linux.die.net/man/1/sshpass).
		Though unlike sshpass it doesn't restrict itself to SSH logins.
		It can supply a password to any process with an identifiable password prompt.`),
	).Use(func(ctx context.Context, cmd *xcmd.Command, next xcmd.Executer) error {
		if len(cmd.FlagSet.Args()) == 0 {
			fmt.Println("You must specify a command.")
			os.Exit(1)
		}
		return next(ctx, cmd)
	}).SetExecuter(func(ctx context.Context, cmd *xcmd.Command) error {
		args := strings.Join(cmd.FlagSet.Args(), " ")
		password := ""
		if cc.EnvName != "" {
			password = os.Getenv(cc.EnvName)
		}
		if err := app.Run(args, []string{password}, cc); err != nil {
			fmt.Println("Error: " + err.Error())
			os.Exit(1)
		}
		return nil
	}).BindSet(cc)

	if err := rootCmd.Execute(context.Background(), os.Args[1:]...); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
