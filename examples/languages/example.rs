// Tron Legacy Theme Example - Rust
// This file demonstrates the variety of syntax highlighting colors

// Single-line comment in dark gray (#3c4b5d)
/* Multi-line comment
   Also in dark gray */

/// Documentation comment
/// Shows how different comment types are highlighted
/// Can include code examples:
/// ```
/// let example = "code in docs";
/// ```

//! Module-level documentation

use std::collections::{HashMap, HashSet};
use std::io::{self, Read, Write};
use std::sync::{Arc, Mutex};
use tokio::sync::mpsc;

// Constants (orange with italic #FFB20D)
const GRID_VERSION: &str = "2.0";
const MAX_POWER_LEVEL: u32 = 9000;
const DEBUG_MODE: bool = true;

// Static variables
static PROGRAM_COUNT: AtomicUsize = AtomicUsize::new(0);
static mut GLOBAL_STATE: Option<String> = None;

// Type aliases
type ProgramId = u64;
type PowerLevel = f32;
type Result<T> = std::result::Result<T, GridError>;

// Enums (dark orange #F79D1E)
#[derive(Debug, Clone, Copy)]
enum ProgramType {
    Security,
    Utility,
    Game,
    System { privileged: bool },
}

// Structs with attributes
#[derive(Debug, Clone, PartialEq)]
pub struct GridProgram {
    name: String,           // Fields in light blue-gray (#d0dfe6)
    id: ProgramId,
    power: PowerLevel,
    program_type: ProgramType,
    active: bool,
}

// Tuple struct
struct Coordinates(f64, f64, f64);

// Unit struct
struct GridMarker;

// Implementation blocks
impl GridProgram {
    // Associated function (orange #FFB20D)
    pub fn new(name: impl Into<String>) -> Self {
        Self {
            name: name.into(),
            id: 0,
            power: 100.0,
            program_type: ProgramType::Utility,
            active: true,
        }
    }

    // Method with self (purple italic #967EFB)
    pub fn activate(&mut self) {
        self.active = true;
        self.power = self.power.max(50.0);
    }

    // Method with parameters (bright green italic #95CC5E)
    pub fn transfer_power(&mut self, amount: PowerLevel, target: &mut GridProgram) {
        let transfer = amount.min(self.power);
        self.power -= transfer;
        target.power += transfer;
    }

    // Getter with lifetime
    pub fn name(&self) -> &str {
        &self.name
    }

    // Builder pattern
    pub fn with_power(mut self, power: PowerLevel) -> Self {
        self.power = power;
        self
    }
}

// Traits (dark orange #F79D1E)
trait GridEntity {
    fn enter_grid(&self);
    fn exit_grid(&self);

    // Default implementation
    fn reboot(&mut self) {
        self.exit_grid();
        self.enter_grid();
    }
}

// Generic trait
trait Resizable<T> {
    fn resize(&mut self, new_size: T);
}

// Trait implementation
impl GridEntity for GridProgram {
    fn enter_grid(&self) {
        println!("Program {} entering the grid", self.name);
    }

    fn exit_grid(&self) {
        println!("Program {} exiting the grid", self.name);
    }
}

// Generic struct
struct Container<T> {
    items: Vec<T>,
    capacity: usize,
}

// Generic implementation with constraints
impl<T: Clone + std::fmt::Debug> Container<T> {
    fn new(capacity: usize) -> Self {
        Container {
            items: Vec::with_capacity(capacity),
            capacity,
        }
    }

    fn add(&mut self, item: T) -> Result<()> {
        if self.items.len() < self.capacity {
            self.items.push(item);
            Ok(())
        } else {
            Err(GridError::CapacityExceeded)
        }
    }
}

// Error types
#[derive(Debug, thiserror::Error)]
enum GridError {
    #[error("Access denied")]
    AccessDenied,

    #[error("Capacity exceeded")]
    CapacityExceeded,

    #[error("Program not found: {0}")]
    ProgramNotFound(String),

    #[error("IO error: {0}")]
    IoError(#[from] io::Error),
}

// Functions with different patterns
fn process_programs(programs: Vec<GridProgram>) -> Vec<String> {
    programs
        .into_iter()
        .filter(|p| p.active)
        .map(|p| p.name)
        .collect()
}

// Function with lifetime parameters
fn longest_name<'a>(x: &'a str, y: &'a str) -> &'a str {
    if x.len() > y.len() { x } else { y }
}

// Async function
async fn load_grid_data(url: &str) -> Result<Vec<u8>> {
    let response = reqwest::get(url).await?;
    let data = response.bytes().await?;
    Ok(data.to_vec())
}

// Closure examples
fn demonstrate_closures() {
    // Simple closure
    let add = |x, y| x + y;

    // Closure with type annotations
    let multiply = |x: i32, y: i32| -> i32 { x * y };

    // Closure capturing environment
    let factor = 10;
    let scale = |x| x * factor;

    // Move closure
    let data = vec![1, 2, 3];
    let process = move || {
        data.iter().sum::<i32>()
    };
}

// Pattern matching
fn analyze_program(program: &GridProgram) -> &'static str {
    match program.program_type {
        ProgramType::Security => "High priority",
        ProgramType::Utility => "Standard priority",
        ProgramType::Game => "Low priority",
        ProgramType::System { privileged: true } => "Critical",
        ProgramType::System { privileged: false } => "System",
    }
}

// Advanced pattern matching
fn process_result(result: Result<GridProgram>) {
    match result {
        Ok(program) if program.active => println!("Active: {}", program.name),
        Ok(program) => println!("Inactive: {}", program.name),
        Err(GridError::AccessDenied) => println!("Access denied!"),
        Err(e) => eprintln!("Error: {}", e),
    }
}

// Destructuring
fn destructure_examples() {
    let Coordinates(x, y, z) = Coordinates(1.0, 2.0, 3.0);

    let program = GridProgram::new("Tron");
    let GridProgram { name, power, .. } = program;

    if let Some(value) = Some(42) {
        println!("Value: {}", value);
    }

    while let Some(item) = vec![1, 2, 3].pop() {
        println!("Item: {}", item);
    }
}

// Macros
macro_rules! grid_log {
    ($($arg:tt)*) => {
        println!("[GRID] {}", format!($($arg)*));
    };
}

// Derive macros
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
struct ApiResponse {
    status: String,
    #[serde(skip_serializing_if = "Option::is_none")]
    data: Option<serde_json::Value>,
}

// Unsafe code
unsafe fn dangerous_operation(ptr: *mut u8) {
    // Unsafe operations must be in unsafe blocks
    *ptr = 42;
}

// Module system
mod grid_core {
    pub mod security {
        pub fn authenticate(user: &str) -> bool {
            // Implementation
            true
        }
    }

    pub(crate) mod internal {
        pub fn system_call() {
            // Internal only
        }
    }
}

// Tests
#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_program_creation() {
        let program = GridProgram::new("Test");
        assert_eq!(program.name(), "Test");
        assert!(program.active);
    }

    #[tokio::test]
    async fn test_async_operation() {
        let result = async_function().await;
        assert!(result.is_ok());
    }

    #[test]
    #[should_panic(expected = "Access denied")]
    fn test_panic() {
        panic!("Access denied");
    }
}

// Main function
fn main() -> Result<()> {
    // String types and formatting
    let s1 = String::from("String type");
    let s2 = "String slice";
    let s3 = r#"Raw string with "quotes""#;
    let s4 = format!("Formatted: {}", 42);

    // Number literals in bright green (#C7F026)
    let decimal = 98_222;
    let hex = 0xfff;
    let octal = 0o77;
    let binary = 0b1111_0000;
    let float = 2.71828;
    let scientific = 1e-10;

    // Character and byte literals
    let ch = 'T';
    let byte = b'A';
    let bytes = b"byte string";

    // Boolean values
    let active = true;
    let inactive = false;

    // Arrays and slices
    let array: [i32; 5] = [1, 2, 3, 4, 5];
    let slice: &[i32] = &array[1..4];

    // Vectors and iterators
    let mut vec = vec![1, 2, 3];
    vec.push(4);

    let sum: i32 = vec.iter()
        .filter(|&&x| x % 2 == 0)
        .map(|&x| x * x)
        .sum();

    // HashMap operations
    let mut map = HashMap::new();
    map.insert("Tron", 100);
    map.insert("CLU", 95);

    // Error handling with ?
    let file_content = std::fs::read_to_string("grid.conf")?;

    // Loop constructs
    loop {
        break;
    }

    for i in 0..10 {
        continue;
    }

    while condition {
        // Loop body
    }

    // Lifetime annotations
    let r: &'static str = "Static lifetime";

    Ok(())
}
