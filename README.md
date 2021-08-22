# Strongbox

A simple password manager written in Go. Encrypt and store all your data into a binary file that can be exported and imported.

Accounts can have their passwords generated using a specific argument see examples below. 

On account fetching passwords will be copied to clipboard.
( currently only works on macOs and Windows )

## How to Install :

Linux bash script:


<br>

Windows 10 powershell script:
(Open powershell in administrator mode)
```
$ScriptFromGithHub = Invoke-WebRequest https://raw.githubusercontent.com/JSONhilder/strongbox/main/install_strongbox.ps1

Invoke-Expression $($ScriptFromGithHub.Content)
```
Windows 10 direct download:
* Download latest release from the github repository
* Extract it to desired directory
* Add alias to powershell with:
```
notepad $PROFILE

then add: 

Set-Alias -Name strongbox -Value path/to/your/extracted/exe
```


## Usage :
```

strongbox [command]

Use "strongbox [command] --help" for more information about a command.

Available Commands:
  delete      Delete account by name from strongbox
  edit        Edit specified account details from strongbox
  export-db   Exported encrypted db file.
  get         Get account by name from strongbox
  help        Help about any command
  import-db   Import encrypted db file.
  list        Lists all accounts in strongbox
  save        Save new entry to strongbox accounts
  version     Current Strongbox version
  help        Stronbox cli documentation and help

```
<br>

## Examples :

* Add new account:
```
strongbox save Twitter myusername mypassword
```
* Add new account with optional url: 
```
strongbox save Twitter myusername mypassword www.twitter.com
```
* Auto generate a password of length 30 characters:
```
strongbox save Twitter myusername gen=30
```
* Export the strongbox binary file:
```
strongbox export-db full/path/to/directory

or for windows:

strongbox export-db full\path\to\directory
```
* Import the strongbox binary file:
```
strongbox import-db full/path/to/strongbox_file

or for windows:

strongbox import-db full\path\to\strongbox_file
```