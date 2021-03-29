# How to Setup - Go Application

There's two ways to execute the application, using make and docker or only docker
#### Make and Docker
```sh
    $ make
```
#### Docker
```sh
    $ docker build -t bomberman .
    $ docker run -it bomberman
```

### What use as input?
As specified in info.md, we need 4 kind of parameters

ROWS COLUMNS SECONDS
GRID
Example:
```
5 5 1
.O..X
.X..O
....X
XX.O.
XOXO.
```