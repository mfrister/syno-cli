# syno-cli

Provides a command-line interface for a Synology DiskStation allowing to list shares and lock/unlock encrypted volumes.

It allows to unlock several volumes at once by providing JSON with a list of shares and their corresponding passwords. If you have a lot of shares and install the frequent security updates (which you should), this might make your life easier.

## Example

    $ export SYNO_BASE_URL='https://myds.example.net:5001'
    $ export SYNO_USER='api'
    $ export SYNO_PASSWORD='password'
    $ syno-cli list
    [ debug messages... ]
    backup          Encrypted, locked
    homes           Not encrypted        user home
    public          Not encrypted
    sync            Encrypted, locked
    $ cat shares.json
    [
        {"name": "backup", "password": "secret"},
        {"name": "sync", "password": "secret2"}
    ]
    $ cat shares.json | syno-cli unlock --batch
    [ debug messages ...]

## Installation

You'll need [Go](https://golang.org/). Once you have Go and set a GOPATH, you can install syno-cli:

    go get frister.net/go/syno-cli

## Usage

syno-cli reads configuration from the following environment variables.

* `SYNO_BASE_URL` - your DiskStation's URL with protocol and port, e.g. `https://myds.example.net:5001`
* `SYNO_USER` - the user to log in as
* `SYNO_PASSWORD` - the user's password

```
usage: syno-cli <command> [<flags>] [<args> ...]

Flags:
  --help  Show help.

Commands:
  help [<command>]
    Show help for a command.

  list
    List shares

  lock <share name>
    Lock an encrypted volume

  unlock [<flags>] [<share name>]
    Unlock an encrypted volume
```

### JSON format for batch unlock

It's just an array of objects with a name and a password attribute:

```json
[
    {"name": "backup", "password": "secret"},
    {"name": "sync", "password": "secret2"}
]
```

## Limitations

* syno-cli logs the user in, but not out, so the session id is still valid after it quits.
* The API sometimes returns error 119. No idea why, a retry helps.
* There's still a lot of debug output, including all HTTP requests including credentials.

## Synology API documentation

I couldn't find any documentation for the part of the API this tool uses, but there's documentation for another part of the DiskStation API ([FileStation](http://ukdl.synology.com/download/Document/DeveloperGuide/Synology_File_Station_API_Guide.pdf)).

Fortunately, the web interface uses the same API, so opening the network tab of a browser's developer tools and doing the actions in the UI gives you the requests.

## Contributions

Feel free to create a pull request, being able to do more with this CLI would be nice :)

## License

    syno-cli is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    syno-cli is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with syno-cli. If not, see <http://www.gnu.org/licenses/>.
