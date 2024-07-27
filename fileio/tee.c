
#include <unistd.h>
#include <sys/stat.h>
#include <stdio.h>
#include <fcntl.h>
#include "tlpi_hdr.h"

#ifndef BUF_SIZE /* Allow "cc -D" to override definition */
#define BUF_SIZE 1024
#endif

int main(int argc, char *argv[])
{
    int outputFd, openFlags;
    ssize_t numRead;
    char buf[BUF_SIZE];

    if (argc != 2 || strcmp(argv[1], "--help") == 0)
    {
        usageErr("%s output-file\n", argv[0]);
    }

    outputFd = open(argv[1],
                    O_CREAT | O_WRONLY | O_TRUNC,
                    S_IRUSR | S_IWUSR | S_IRGRP | S_IWGRP | S_IROTH | S_IWOTH);

    if (outputFd == -1)
    {
        errExit("opening file %s", argv[1]);
    }

    while ((numRead = read(STDIN_FILENO, buf, BUF_SIZE)) > 0)
    {
        if (write(STDOUT_FILENO, buf, numRead) != numRead)
        {
            fatal("couldn't write whole buffer");
        }

        if (write(outputFd, buf, numRead) != numRead)
        {
            fatal("couldn't write whole buffer");
        }
    }

    close(outputFd);
    if (close(outputFd) == -1)
    {
        errExit("close output");
    }

    return 0;
}
