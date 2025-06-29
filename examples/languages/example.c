/* Tron Legacy Theme Example - C
 * This file demonstrates the variety of syntax highlighting colors
 */

// Single-line comment in dark gray (#3c4b5d)

/* Multi-line comment
 * spanning several lines
 * Also in dark gray
 */

// Preprocessor directives (cyan #6ee2ff)
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>
#include <stdint.h>
#include <pthread.h>

#define GRID_VERSION "2.0"
#define MAX_PROGRAMS 100
#define DEBUG 1

// Macro with parameters
#define MIN(a, b) ((a) < (b) ? (a) : (b))
#define SQUARE(x) ((x) * (x))

// Conditional compilation
#ifdef DEBUG
    #define LOG(msg) printf("[DEBUG] %s\n", msg)
#else
    #define LOG(msg) /* no-op */
#endif

// Include guards
#ifndef GRID_CORE_H
#define GRID_CORE_H
// Header content would go here
#endif

// Type definitions
typedef unsigned int program_id_t;
typedef float power_level_t;

// Enums (dark orange #F79D1E)
typedef enum {
    PROGRAM_SECURITY,
    PROGRAM_UTILITY,
    PROGRAM_GAME,
    PROGRAM_SYSTEM
} program_type_t;

// Enum with explicit values
enum grid_status {
    GRID_OFFLINE = 0,
    GRID_INITIALIZING = 1,
    GRID_ONLINE = 2,
    GRID_ERROR = -1
};

// Struct definitions (dark orange #F79D1E)
typedef struct {
    char name[64];              // Fields in light blue-gray
    program_id_t id;
    power_level_t power;
    program_type_t type;
    bool active;
} GridProgram;

// Nested struct
struct grid_node {
    struct {
        float x, y, z;
    } position;
    struct grid_node *next;
    struct grid_node *prev;
    void *data;
};

// Union
typedef union {
    uint32_t raw;
    struct {
        uint8_t r;
        uint8_t g;
        uint8_t b;
        uint8_t a;
    } channels;
} color_t;

// Function declarations (orange #FFB20D)
void initialize_grid(void);
GridProgram* create_program(const char *name, program_type_t type);
void activate_program(GridProgram *program);
int transfer_power(GridProgram *source, GridProgram *target, power_level_t amount);

// Function with variable arguments
void grid_log(const char *format, ...);

// Function pointers
typedef void (*program_callback_t)(GridProgram *);
typedef int (*compare_func_t)(const void *, const void *);

// Static function
static void cleanup_resources(void) {
    LOG("Cleaning up resources");
}

// Inline function
inline int max(int a, int b) {
    return (a > b) ? a : b;
}

// Constants (orange with italic #FFB20D)
const double PI = 3.14159265358979323846;
const char *GRID_NAME = "ENCOM OS-12";
static const int BUFFER_SIZE = 1024;

// Global variables (light blue-gray #d0dfe6)
GridProgram *active_programs[MAX_PROGRAMS];
int program_count = 0;
volatile bool grid_running = false;
extern int external_value;

// Function implementation
GridProgram* create_program(const char *name, program_type_t type) {
    // Allocate memory
    GridProgram *program = malloc(sizeof(GridProgram));
    if (!program) {
        perror("Failed to allocate program");
        return NULL;
    }

    // Initialize fields
    strncpy(program->name, name, sizeof(program->name) - 1);
    program->name[sizeof(program->name) - 1] = '\0';
    program->id = program_count++;
    program->power = 100.0f;
    program->type = type;
    program->active = false;

    return program;
}

// Function with multiple parameters (bright green italic #95CC5E)
int transfer_power(GridProgram *source, GridProgram *target, power_level_t amount) {
    if (!source || !target) {
        return -1;
    }

    if (source->power < amount) {
        amount = source->power;
    }

    source->power -= amount;
    target->power += amount;

    return 0;
}

// Main function
int main(int argc, char *argv[]) {
    // String literals (red #FF410D)
    printf("Tron Legacy Grid System v%s\n", GRID_VERSION);

    // String with escape sequences (light red #ff6113)
    printf("Initializing...\n\tPhase 1\n\tPhase 2\n");
    printf("Special chars: \x1B[1mBold\x1B[0m\n");

    // Character literals
    char grid_marker = 'G';
    char newline = '\n';
    char tab = '\t';
    char hex_char = '\x41';

    // Number literals (bright green #C7F026)
    int decimal = 42;
    int hex = 0xFF;
    int octal = 0755;
    int binary = 0b1010;  // C99
    long big_num = 1234567890L;
    unsigned int flags = 0xDEADBEEFU;

    // Floating point numbers
    float power = 95.5f;
    double precision = 3.14159265358979;
    double scientific = 1.23e-4;

    // Boolean values (C99)
    bool is_active = true;
    bool is_corrupt = false;

    // Null pointer
    void *null_ptr = NULL;

    // Arrays
    int coordinates[3] = {10, 20, 30};
    char buffer[BUFFER_SIZE];
    GridProgram *programs[10] = {NULL};

    // Pointers and addresses
    int value = 42;
    int *ptr = &value;
    int **ptr_to_ptr = &ptr;

    // Array indexing and pointer arithmetic
    programs[0] = create_program("Tron", PROGRAM_SECURITY);
    *(programs + 1) = create_program("CLU", PROGRAM_SYSTEM);

    // Control structures (keywords in gray #748aa6)
    if (programs[0] != NULL) {
        activate_program(programs[0]);
    } else {
        fprintf(stderr, "Failed to create program\n");
    }

    // Loops
    for (int i = 0; i < 10; i++) {
        printf("Iteration %d\n", i);
    }

    while (grid_running) {
        // Process events
        if (program_count > MAX_PROGRAMS) {
            break;
        }
        continue;
    }

    do {
        // At least once
    } while (0);

    // Switch statement
    switch (programs[0]->type) {
        case PROGRAM_SECURITY:
            printf("Security program\n");
            break;
        case PROGRAM_UTILITY:
        case PROGRAM_GAME:
            printf("User program\n");
            break;
        default:
            printf("Unknown type\n");
    }

    // Goto (rarely used)
    goto cleanup;

    // Operators
    int a = 10, b = 20;
    int sum = a + b;
    int diff = a - b;
    int product = a * b;
    int quotient = b / a;
    int remainder = b % a;

    // Bitwise operations
    unsigned int mask = 0xFF00;
    unsigned int result = flags & mask;
    result |= 0x0F;
    result ^= 0xAA;
    result = ~result;
    result = result << 4;
    result = result >> 2;

    // Logical operations
    if (a > 0 && b < 100) {
        // Both conditions true
    }

    if (ptr == NULL || program_count == 0) {
        // Either condition true
    }

    // Ternary operator
    int max_value = (a > b) ? a : b;

    // sizeof operator
    size_t program_size = sizeof(GridProgram);
    size_t array_size = sizeof(coordinates) / sizeof(coordinates[0]);

    // Type casting
    double precise_power = (double)programs[0]->power;
    void *generic_ptr = (void *)programs[0];
    GridProgram *cast_back = (GridProgram *)generic_ptr;

    // Compound literals (C99)
    GridProgram temp = (GridProgram){
        .name = "Temp",
        .id = 999,
        .power = 50.0f,
        .type = PROGRAM_UTILITY,
        .active = true
    };

    // Variable length arrays (C99)
    int n = 5;
    int vla[n];
    for (int i = 0; i < n; i++) {
        vla[i] = i * i;
    }

    // Static assertions (C11)
    _Static_assert(sizeof(int) == 4, "int must be 4 bytes");

    // Generic selection (C11)
    #define TYPE_NAME(x) _Generic((x), \
        int: "int", \
        float: "float", \
        double: "double", \
        default: "other")

    printf("Type of power: %s\n", TYPE_NAME(power));

cleanup:
    // Free allocated memory
    for (int i = 0; i < 10; i++) {
        if (programs[i]) {
            free(programs[i]);
        }
    }

    return EXIT_SUCCESS;
}

// Thread function example
void* grid_worker(void *arg) {
    pthread_t tid = pthread_self();
    printf("Worker thread %lu started\n", (unsigned long)tid);

    // Thread-local storage
    static __thread int thread_local_counter = 0;
    thread_local_counter++;

    return NULL;
}

// Variadic function
void grid_log(const char *format, ...) {
    va_list args;
    va_start(args, format);

    printf("[GRID] ");
    vprintf(format, args);
    printf("\n");

    va_end(args);
}

// Restrict keyword (C99)
void copy_data(char * restrict dest, const char * restrict src, size_t n) {
    for (size_t i = 0; i < n; i++) {
        dest[i] = src[i];
    }
}

// Complex numbers (C99)
#include <complex.h>
void complex_example(void) {
    double complex z1 = 1.0 + 2.0 * I;
    double complex z2 = 3.0 - 4.0 * I;
    double complex sum = z1 + z2;

    printf("Real: %f, Imaginary: %f\n", creal(sum), cimag(sum));
}
