PREFIX = /opt
MANPREFIX = $(PREFIX)/man
X11INC = /usr/X11R6/include
X11LIB = /usr/X11R6/lib
CC = tcc
CPPFLAGS = -I$(X11INC) -D_DEFAULT_SOURCE
CFLAGS   = -std=c99 -pedantic -Wall -Wextra -Os
LDFLAGS  = -L$(X11LIB) -s
LDLIBS   = -lX11
VERSION = 0
