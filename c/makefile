CC = gcc
CFLAGS = -g -std=gnu99

all:
	make clean
	make sysfork

sysfork.o: sysfork.c
	@$(CC) -c sysfork.c $(CFLAGS)

sysfork: sysfork.o
	@$(CC) -o sysfork sysfork.o

clean:
	-rm -f *.o
	-rm -f *~
