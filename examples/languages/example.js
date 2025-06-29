// Tron Legacy Theme Example - JavaScript
// This file demonstrates the variety of syntax highlighting colors

// Single-line comment in dark gray (#3c4b5d)
/* Multi-line comment
   Also in dark gray */

// String variations
const normalString = "Normal string in red (#FF410D)";
const stringWithEscape = "String with \n escape \t sequences (#ff6113)";
const regexString = /regex.*pattern[0-9]+/gi; // Regex in cyan (#6ee2ff)
const templateString = `Template with ${variable} interpolation`;

// Numbers in bright green (#C7F026)
const integer = 42;
const float = 3.14159;
const hex = 0xFF410D;
const binary = 0b1010;

// Constants and booleans (orange with italic)
const MY_CONSTANT = "CONSTANT_VALUE";
const isTrue = true;
const isFalse = false;
const nothing = null;
const notDefined = undefined;

// Variables in light blue-gray (#d0dfe6)
let mutableVariable = "can change";
var oldStyleVar = "old school";

// Function variations
function regularFunction(param1, param2) { // Function name in orange (#FFB20D)
    // Parameters in bright green italic (#95CC5E)
    return param1 + param2;
}

// Built-in functions in blue italic (#267fb5)
console.log("Built-in function example");
Array.isArray([1, 2, 3]);
Object.keys({ a: 1, b: 2 });

// Arrow function with parameters
const arrowFunc = (x, y, z) => x * y * z;

// Keywords in gray (#748aa6)
if (condition) {
    for (let i = 0; i < 10; i++) {
        while (true) {
            break;
        }
    }
} else {
    switch (value) {
        case 1:
            return true;
        default:
            throw new Error("Error message");
    }
}

// Classes and constructors (dark orange #F79D1E)
class TronEntity extends BaseEntity {
    constructor(name, power) {
        super(name);
        this.power = power; // Properties in light blue-gray
    }

    static classMethod() {
        return "static method";
    }

    get powerLevel() {
        return this.power;
    }
}

// Object with properties
const tronObject = {
    name: "Flynn",
    program: "CLU",
    grid: true,
    coordinates: [10, 20, 30],
    // Method definition
    enter() {
        return "Entering the Grid";
    },
    // Computed property
    ["dynamic" + "Key"]: "value"
};

// Array operations with punctuation
const programs = ["Tron", "CLU", "Quorra"]; // Brackets in bright blue (#2f9ee2)
const [first, second, ...rest] = programs;
const newPrograms = [...programs, "Flynn"];

// Special operators and punctuation
const result = (x > y) ? x : y;
const combined = str1 + str2;
const logical = a && b || c;
const bitwise = flags & MASK;

// Async/await
async function loadGrid() {
    try {
        const data = await fetch("/api/grid");
        return await data.json();
    } catch (error) {
        console.error("Grid loading failed:", error);
    }
}

// Special this keyword (purple italic #967EFB)
const program = {
    name: "Tron",
    announce() {
        return `I am ${this.name}`;
    }
};

// Import/export statements
import { Component } from 'react';
export default TronEntity;
export { regularFunction, arrowFunc };

// JSX example (if supported)
const GridComponent = () => {
    return (
        <div className="grid-container">
            <h1>Welcome to the Grid</h1>
            <p>Programs: {programs.length}</p>
        </div>
    );
};

// Decorators (if supported)
@deprecated
@memoize({ maxAge: 5000 })
class LegacyProgram {
    @readonly
    version = "1.0";
}

// Symbol usage
const gridSymbol = Symbol("grid");
const programSymbol = Symbol.for("program");

// Generator function
function* generatePrograms() {
    yield "Tron";
    yield "CLU";
    yield "Quorra";
}

// Map and Set
const programMap = new Map([
    ["Tron", { type: "Security" }],
    ["CLU", { type: "Admin" }]
]);

const uniquePrograms = new Set(["Tron", "CLU", "Tron"]);

// Promise chain
Promise.resolve("Start")
    .then(data => processData(data))
    .catch(error => handleError(error))
    .finally(() => cleanup());
