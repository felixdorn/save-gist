# Save gist
This tool was built in 30 minutes (check out the livestream, https://www.twitch.tv/videos/1583737595, if you like to see a panicked developer looking at a timer).

## Usage
```
SG_GITHUB_TOKEN=... sg [FILES...]
```
> The access token needs the `gist` scope, it's not recommended to run the command this way. Instead, set the environment variable and then call `sg`.  

## Installation
```bash
curl -L https://github.com/felixdorn/save-gist/releases/latest/download/sg -o /usr/bin/sg
chmod +x ./usr/bin/sg
```

To update `sg` to the latest version, run `sg --update`. Note, you can not downgrade or update to a specific version with this command. You'll have to use cURL.

## Configuration

* `SG_DIRECTORY_SEPARATOR=-`,  defaults to a dash `-`, I like to use a `:` instead. Note, slashes are not allowed by GitHub.