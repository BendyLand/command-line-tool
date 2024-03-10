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

void display(Task* task) {
    printf("Task ID: %d\nTask body: %s\n", task->id, task->body);
}

void createTask(Task* task) {
    printf("Enter the task body: ");
    char buffer[MAX_CHARS]; 
    fgets(buffer, sizeof(buffer), stdin); 
    buffer[strcspn(buffer, "\n")] = '\0';
    task->body = malloc(strlen(buffer) + 1); 
    strcpy(task->body, buffer); 
}

void deleteTask(Task* task) {
    free(task->body);
}

int main(void) {

    Task test;
    test.id = 1;
    
    createTask(&test); 
    display(&test);
    deleteTask(&test);

    return 0;
}