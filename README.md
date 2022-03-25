# tcsf
Needs a better name : )

This CLI exists to help you with some Twilio routine task's

### Setup
`tcsf config init`
this command will create a config file on your `$HOME`
the config file have two sets of config `source` and `target`
``` json
{
	"source_api_key": "API KEY",
	"source_api_secret": "API SECRET",
	"source_workspace": "WORKSPACE SID",
	"target_api_key": "API KEY",
	"target_api_secret": "API SECRET",
	"target_workspace": "WORKSPACE SID"
}
```
If you run the command again you can overwrite the config, but only the keys you set will be overwrite.

### Commands
At any command you can pass `--target` to run the command using the **target** credentials, some command will invert the credentials like the `studio flow copy`.

Use the `--help` to get some help on a specific command

`tcsf [flags]`

`tcsf --help`

`tcsf --version || tcsf -v`

`tcsf config init` - Create or Update the config file

`tcsf taskrouter workspaces list [flags] || tcsf t w ls [flags]` - List all workspace

`tcsf taskrouter workflow list [flags] || tcsf t wf ls [flags]` - List all workflow, you can pass `--name` to search a workflow by name

`tcsf studio flow list [flags] || tcsf s f ls [flags]` - List all studio flow, you can pass `--name` to search a studio flow by name

`tcsf studio flow copy [flags] || tcsf s f ls [flags]` - Copy a studio flow from the source account to the target account, you can pass `--sid` of the flow to be copied, on this command if you pass `--target`, the credentials will be inverted, so **source become target and target become source**