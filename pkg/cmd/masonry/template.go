/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package masonry

var usageTemplate = `{{if gt (len .Aliases) 0}}

Aliases:
{{.NameAndAliases}}{{end}}Usage:
  {{.CommandPath}} [global-options]{{if .HasAvailableSubCommands}} COMMAND{{end}}{{if .HasAvailableFlags}} [command-options]{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Options:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableLocalFlags}}

Options:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}
{{if .HasHelpSubCommands}}
Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}
See "{{.CommandPath}} help <TOPIC>" for more information on a specific topic.
See "{{.CommandPath}} <COMMAND> --help" for more information about a command.
{{end}}
`
