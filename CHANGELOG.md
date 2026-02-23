# Changelog v1.0.1

## Changes between v1.0.0 and v1.0.1

- Changed permissions of `buildcommands` file meant for building targets ( linux-x64 and win-x64 )

- Printing of server commands with log configuration

- Ensuring user changes or is aware of security level of his header key value

## Changes between v1.0.1 and v1.0.2

- Templates schema changed to include reasons mapping

- Moved templating struct to util package

- Changed buildcommands to bash script (I'm developing on debian 13 server via ssh as my 'workspace'). Alternative could be Makefile.