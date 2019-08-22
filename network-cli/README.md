CLI written in Go to search 
- Name Servers
- IP
- Mail Server 
- CNAME

package used: github.com/urfave/cli

help output:
```
NAME:
   Website Lookup CLI - Let's you query IPs, CNAMEs, MX records and Name Servers

USAGE:
   cli [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
   ns       Looks up the Name Servers for a Particular Host
   ip       Looks up the IP addresses for a partucular host
   cname    Looks up the CNAME for a partucular host
   mx       Looks up the MX records for a partucular host
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

URL: https://tutorialedge.net/golang/building-a-cli-in-go/