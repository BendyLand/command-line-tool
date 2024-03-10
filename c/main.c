#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_CHARS 100
#define MAX_TASKS 100

typedef struct {
    int id;
    char* body;
} Task;

Task tasks[MAX_TASKS];
int numTasks = 0;

void display() {
    for (int i = 0; i < MAX_TASKS; i++) {
        if (tasks[i].body != NULL) {
            printf("Task ID: %d\nTask body: %s\n", tasks[i].id, tasks[i].body);
        }
    }
}

void createTask() {

    // check if the task list is full
    if (numTasks < MAX_TASKS) {

        // increments the currentTask, gets the pointer to the `tasks` list, and assigns to newTask
        Task* newTask = &tasks[numTasks++];

        // sets newTask's id to the current task number
        newTask->id = numTasks; 

        // prompts the user for text to enter into the task body
        printf("Enter the task body for task %d: ", newTask->id);

        // create a temporary storage location for the text
        char buffer[100]; 

        // get user input of sizeof(buffer), 
        fgets(buffer, sizeof(buffer), stdin); 

        // finds the index of '\n' with strcspn, 
        // uses it as the index to buffer to find the end of the input, 
        // and replace the newline with a NULL character
        buffer[strcspn(buffer, "\n")] = '\0'; 

        // sets the body of newTask to the input, copied from buffer.
        // strdup allocates memory as if malloc was called, so it needs to be free()d
        newTask->body = strdup(buffer);
    } else {
        printf("Task list if full!\n");
    }
}

void cleanupTasks() {
    for (int i = 0; i < MAX_TASKS; i++) {
        free(tasks[i].body);
    }
}

int main(void) {

    Task test;
    test.id = 1;
    
    createTask(); 
    display();
    cleanupTasks();

    return 0;
}