#include <stdio.h>
#include <stdlib.h>

#include <sys/types.h>
#include <sys/ipc.h>
#include <sys/shm.h>

#include <string.h>

int main(int argc, char *argv[]) {
	int shmID = atoi(argv[1]);

	// if (shmget(shmID, 4096, 0666) < 0) {
	// 	printf("Failed to open shared memory regio %d!\n", shmID);
	//	exit(1);
	// }

	double* pBase = (double*) shmat(shmID, NULL, 0);
	if (pBase == (double*) -1) {
		printf("Failed to open shared memory region %d!\n", shmID);
		exit(1);
	}

	for (int i = 0; i < 100; i++) {
		printf("NUM: %lf\n", pBase[i]);
	}

	return 0;
}
