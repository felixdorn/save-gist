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
