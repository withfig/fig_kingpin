package figkingpin

import (
	"encoding/json"
	"os"

	"github.com/alecthomas/kingpin/v2"
)

func figFlagJson(model *kingpin.FlagModel) any {
	var flag = map[string]interface{}{}

	if model.Short != 0 {
		flag["name"] = []string{model.Name, string(model.Short)}
	} else {
		flag["name"] = model.Name
	}

	if model.Help != "" {
		flag["description"] = model.Help
	}

	if model.Required {
		flag["isRequired"] = model.Required
	}

	if model.Hidden {
		flag["hidden"] = model.Hidden
	}

	if !model.IsBoolFlag() {
		flag["args"] = map[string]interface{}{
			"name": "ARG",
		}
	}

	return flag

}

func figFlagsJson(models []*kingpin.FlagModel) []any {
	var flags []any
	for _, model := range models {
		flags = append(flags, figFlagJson(model))
	}
	return flags
}

func figArgJson(model *kingpin.ArgModel) any {
	var arg = map[string]interface{}{
		"name": model.Name,
	}

	if model.Help != "" {
		arg["description"] = model.Help
	}

	if model.Required {
		arg["isRequired"] = model.Required
	}

	return arg
}

func figArgsJson(models []*kingpin.ArgModel) any {
	if len(models) == 1 {
		return figArgJson(models[0])
	} else {
		var args []any
		for _, model := range models {
			args = append(args, figArgJson(model))
		}
		return args
	}
}

func figSpecJson(model *kingpin.ApplicationModel) any {
	return map[string]interface{}{
		"name":        model.Name,
		"description": model.Help,
		"subcommands": figSubcommandsJson(model.Commands),
		"options":     figFlagsJson(model.Flags),
	}
}

func figSubcommandJson(model *kingpin.CmdModel) any {
	var subcommand = map[string]interface{}{
		"name": model.Name,
	}

	if model.Help != "" {
		subcommand["description"] = model.Help
	}

	if model.Hidden {
		subcommand["hidden"] = model.Hidden
	}

	if len(model.Commands) > 0 {
		subcommand["subcommands"] = figSubcommandsJson(model.Commands)
	}

	if len(model.Flags) > 0 {
		subcommand["options"] = figFlagsJson(model.Flags)
	}

	if len(model.Args) > 0 {
		subcommand["args"] = figArgsJson(model.Args)
	}

	return subcommand
}

func figSubcommandsJson(models []*kingpin.CmdModel) []any {
	var subcommands []any
	for _, model := range models {
		subcommands = append(subcommands, figSubcommandJson(model))
	}
	return subcommands
}

func GenerateFigCompletionScript(a *kingpin.Application) func(c *kingpin.ParseContext) error {
	return func(c *kingpin.ParseContext) error {
		var marsheledJson, err = json.MarshalIndent(figSpecJson(a.Model()), "", "  ")
		if err != nil {
			return err
		}
		print("const completionSpec: Fig.Spec = ")
		print(string(marsheledJson))
		print(";\n\nexport default completionSpec;\n")

		os.Exit(0)
		return nil
	}
}
