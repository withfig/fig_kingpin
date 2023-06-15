package figkingpin

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

func figFlagJson(model *kingpin.FlagModel, specName string, fullCmd string) any {
	var flag = map[string]interface{}{}

	if model.Short != 0 {
		flag["name"] = []string{fmt.Sprintf("--%s", model.Name), fmt.Sprintf("-%c", model.Short)}
	} else {
		flag["name"] = fmt.Sprintf("--%s", model.Name)
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
			"name": model.FormatPlaceHolder(),
			"generator": map[string]interface{}{
				"script": fmt.Sprintf("%s --completion-bash %s --%s", specName, fullCmd, model.Name),
			},
		}
	}

	return flag

}

func figFlagsJson(models []*kingpin.FlagModel, specName string, fullCmd string) []any {
	var flags []any
	for _, model := range models {
		flags = append(flags, figFlagJson(model, specName, fullCmd))
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

func figSpecJson(model *kingpin.ApplicationModel, name string) any {
	return map[string]interface{}{
		"name":        name,
		"description": model.Help,
		"subcommands": figSubcommandsJson(model.Commands, name),
		"options":     figFlagsJson(model.Flags, name, ""),
	}
}

func figSubcommandJson(model *kingpin.CmdModel, specName string) any {
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
		subcommand["subcommands"] = figSubcommandsJson(model.Commands, specName)
	}

	if len(model.Flags) > 0 {
		subcommand["options"] = figFlagsJson(model.Flags, specName, model.FullCommand)
	}

	if len(model.Args) > 0 {
		subcommand["args"] = figArgsJson(model.Args)
	}

	return subcommand
}

func figSubcommandsJson(models []*kingpin.CmdModel, specName string) []any {
	var subcommands []any
	for _, model := range models {
		subcommands = append(subcommands, figSubcommandJson(model, specName))
	}
	return subcommands
}

func GenerateFigCompletionSpec(a *kingpin.Application) func(c *kingpin.ParseContext) error {
	return func(c *kingpin.ParseContext) error {
		var name = c.Elements[len(c.Elements)-1].Value

		var marsheledJson, err = json.MarshalIndent(figSpecJson(a.Model(), *name), "", "  ")
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
