#define _GNU_SOURCE
#include <unistd.h>
#include <errno.h>
#include <sched.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <fcntl.h>

void ns_exec(void) {
    fprintf(stdout, "ns_exec............");

	char *pid;
	pid = getenv("pid");
	if (pid) {
		fprintf(stdout, "got pid=%s\n", pid);
	} else {
		fprintf(stdout, "missing pid env skip nsenter");
		return;
	}
	char *cmd;
	cmd = getenv("cmd");
	if (cmd) {
		fprintf(stdout, "got cmd=%s\n", cmd);
	} else {
		fprintf(stdout, "missing cmd env skip nsenter");
		return;
	}
	int i;
	char nspath[1024];
	char *namespaces[] = { "ipc", "uts", "net", "pid", "mnt" };

	for (i=0; i<5; i++) {
		sprintf(nspath, "/proc/%s/ns/%s", pid, namespaces[i]);
		int fd = open(nspath, O_RDONLY);

		if (setns(fd, 0) == -1) {
			fprintf(stderr, "setns on %s namespace failed: %s\n", namespaces[i], strerror(errno));
		} else {
			fprintf(stdout, "setns on %s namespace succeeded\n", namespaces[i]);
		}
		close(fd);
	}
	int res = system(cmd);
	fprintf(stdout, "xx  %d\n", res);
	exit(0);
	return;
}

