#include <stdio.h>
#include <stdlib.h> // For exit() function
#include <time.h>

#define DAY1_NUM 2020
#define DAY2_NUM 30000000

unsigned int arr[DAY2_NUM] = {0}; // Faster than calloc 

int main() {
	clock_t t = clock();

	unsigned short input[7] = {6,4,12,1,20,0,16};
	unsigned char input_len = sizeof(input)/sizeof(short);
	
	//unsigned int* arr = (unsigned int*) calloc(GOAL_NUM, sizeof(int));
	//if (arr == NULL) { exit(1); };

	unsigned int i = 0; 
	for (; i < (input_len); i++){
		arr[input[i]] = i;
	}

	unsigned int next_val = 0;

	for (; i < DAY1_NUM - 1; i++){ // Dag 1, upp till 2020.
		unsigned int tmp_val = arr[next_val];
		if ( tmp_val ){
			arr[next_val] = i;
			next_val = i - tmp_val;
		} 
		else {
			if (next_val == input[0]){
				arr[next_val] = i;
				next_val = i - tmp_val;
				continue;
			}
			arr[next_val] = i;
			next_val = 0;
		}
	}
	
	printf("The %dth number spoken is: %d\n",DAY1_NUM, next_val);

	for (; i < DAY2_NUM - 1; i++){ //Dag 2, forts. upp till 30 000 000.
		unsigned int tmp_val = arr[next_val];
		if ( tmp_val ){
			arr[next_val] = i;
			next_val = i - tmp_val;
		} 
		else {
			if (next_val == input[0]){
				arr[next_val] = i;
				next_val = i - tmp_val;
				continue;
			}
			arr[next_val] = i;
			next_val = 0;
		}
	}
	//free (arr);

	printf("The %dth number spoken is: %d\n",DAY2_NUM, next_val);

 	float times = ((float) (clock() - t) / CLOCKS_PER_SEC);
	printf("Execution time = %.3f seconds.\n", times);

	return 0;
}
