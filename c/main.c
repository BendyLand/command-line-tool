#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_CHARS 100
#define MAX_TASKS 10

typedef struct {
    int id;
    char* body;
} Task;

Task tasks[MAX_TASKS];
int numTasks = 0;

void run();
void display();
void cleanup();
void createTask();
void deleteTask();

int main(void) {
    printf("Welcome to the C Task Manager!\n");
    run();

    return 0;
}

void run() {
    printf("\nWhat would you like to do?\n1.) Create task\n2.) View tasks\n3.) Delete task\n4.) Exit\n\n");
    char buffer[MAX_TASKS];
    fgets(buffer, sizeof(buffer), stdin);
    buffer[strcspn(buffer, "\n")] = '\0';
    int id = atoi(&buffer[0]);
    switch (id) {
        case 1:
            createTask();
            run();
            break;
        case 2:
            display();
            run();
            break;
        case 3:
            deleteTask();
            run();
            break;
        default:
            printf("Shutting down todo list...\n");
            cleanup();
            break;
    }
}

void deleteTask() {
    printf("Which Task (ID) would you like to delete?\n\n");
    display();
    // temporary storage for user input
    char buffer[MAX_TASKS];
    // get user input from stdin and store it in buffer
    fgets(buffer, sizeof(buffer), stdin);
    // change the newline to a null byte
    buffer[strcspn(buffer, "\n")] = '\0';
    // convert user input to an int
    int id = atoi(&buffer[0]);
    // loop through the list of tasks
    for (int i = 0; i < MAX_TASKS; i++) {
        // if the task id matches the user input...
        if (tasks[i].id == id) {
            // free the memory associated with current pointer
            free(tasks[i].body);
            // set the value of the memory address to NULL to prevent double free in cleanup()
            tasks[i].body = NULL;
            // print success message and return
            printf("Task deleted successfully!\n");
            return;
        }
    }
    // if you reach this point, an id matching the user input was not found
    printf("There was a problem deleting the task.\n");
}

void cleanup() {
    printf("Cleaning up remaining tasks...\n");
    for (int i = 0; i < MAX_TASKS; i++) {
        if (tasks[i].body != NULL) {
            printf("Removing Task ID %d\n", i+1);
            free(tasks[i].body);
        }
    }
    printf("Tasks cleared successfully!\n\n");
}

void createTask() {
    // check if the task list is full
    if (numTasks < MAX_TASKS) {
        // increments the currentTask, gets the pointer to the `tasks` list at that index, and assigns to newTask
        Task* newTask = &tasks[numTasks++];
        // sets newTask's id to the current task number
        newTask->id = numTasks;
        // prompts the user for text to enter into the task body
        printf("Enter the task body for task %d: ", newTask->id);
        // create a temporary storage location for the text
        char buffer[MAX_CHARS];
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
        printf("Task list is full!\n");
    }
}

void display() {
    printf("Here are your tasks:\n");
    for (int i = 0; i < MAX_TASKS; i++) {
        if (tasks[i].body != NULL) {
            printf("Task ID %d: %s\n", tasks[i].id, tasks[i].body);
        }
    }
    printf("\n");
}