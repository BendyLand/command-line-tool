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

void display() {
    for (int i = 0; i < MAX_TASKS; i++) {
        if (tasks[i].body != NULL) {
            printf("Task ID: %d\nTask body: %s\n", tasks[i].id, tasks[i].body);
        }
    }
}

void createTask() {
    if (numTasks < MAX_TASKS) {
        Task* newTask = &tasks[numTasks++];
        newTask->id = numTasks; 
        printf("Enter the task body for task %d: ", newTask->id);
        char buffer[100]; 
        fgets(buffer, sizeof(buffer), stdin); 
        buffer[strcspn(buffer, "\n")] = '\0'; 
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