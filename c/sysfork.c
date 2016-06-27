#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <sys/stat.h>

#define BUF_SIZE 1024
#define CHILD 0

typedef struct {
    char filename[BUF_SIZE];
    off_t size;
} SizeMessage;

int main(int argc, const char ** argv) {
    /* fork process id */
    pid_t pid = -1;

    /* Start index at 1 to ignore prog name in args */
    int startIndex = 1;
    int filesCount = argc - 1;

    /* Pipe to communicate sizes from the children to the parent */
    int pipe_fd[2];
    pipe(pipe_fd);

    /* Fork for each file given as an argument */
    int childIndex = 0;
    for(int i = startIndex; i < argc; i++) {
        if((pid = fork()) == -1) {
            perror("fork");
            exit(EXIT_FAILURE);
        }
        if(pid == 0) {
            /* Child will process current index */
            childIndex = i;
            break;
        }
    }

    if(pid == CHILD) {
        /* Child process closes up input side of pipe */
        close(pipe_fd[0]);
        SizeMessage msg;
        strcpy(msg.filename, argv[childIndex]);

        /* Get file size */
        struct stat st;
        if(stat(msg.filename, &st) == -1) {
            perror("stat");
            msg.size = -1;
        } else {
            msg.size = st.st_size;
        }

        /* Send size message through the output side of pipe */
        write(pipe_fd[1], &msg, sizeof(msg));
        exit(EXIT_SUCCESS);
    } else {
        /* Parent process closes up output side of pipe */
        close(pipe_fd[1]);

        /* Read in size messages from the pipe */
        SizeMessage biggest = {0};
        SizeMessage evens[filesCount];
        int countEvens = 0;
        for(int i = startIndex; i < argc; i++) {
            SizeMessage msg;
            if(read(pipe_fd[0], &msg, sizeof(msg)) == -1) {
                perror("read");
            }
            if(msg.size > biggest.size) {
                biggest = msg;
                countEvens = 0;
            }
            if(msg.size == biggest.size) {
                evens[countEvens++] = msg;
            }
        }

        /* Display result */
        if(countEvens == filesCount) {
            printf("All files are even\n");
        } else if(countEvens > 1) {
            printf("Biggest files are:");
            for(int i = 0; i < countEvens; i++) {
                printf(" %s", evens[i].filename);
            }
            printf("\n");
        } else {
            printf("Biggest file is: %s\n", biggest.filename);
        }
    }

    return EXIT_SUCCESS;
}
