package v2

import (
	"os"

	"code.cloudfoundry.org/cli/cf/cmd"
	"code.cloudfoundry.org/cli/commands"
	"code.cloudfoundry.org/cli/commands/flags"
)

type UnbindSecurityGroupCommand struct {
	RequiredArgs flags.BindSecurityGroupArgs `positional-args:"yes"`
	usage        interface{}                 `usage:"CF_NAME unbind-security-group SECURITY_GROUP ORG SPACE\n\nTIP: Changes will not apply to existing running applications until they are restarted."`
}

func (_ UnbindSecurityGroupCommand) Setup(config commands.Config) error {
	return nil
}

func (_ UnbindSecurityGroupCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
