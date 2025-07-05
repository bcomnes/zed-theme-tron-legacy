# Tron Legacy Theme Example - Python
# This file demonstrates the variety of syntax highlighting colors

# Single-line comment in dark gray (#3c4b5d)
"""
Multi-line docstring
Also serves as a comment
"""

# Import statements
import os
import sys
from typing import List, Dict, Optional, Union
from collections import defaultdict, namedtuple
from functools import wraps, lru_cache
import numpy as np

# String variations
normal_string = "Normal string in red (#FF410D)"
string_with_escape = "String with \n escape \t sequences \x1b[0m (#ff6113)"
raw_string = r"Raw string with \n no escapes"
f_string = f"Formatted string with {variable} interpolation"
multiline_string = """
    Multi-line string
    spanning several lines
"""

# Regex patterns (cyan #6ee2ff)
import re
pattern = re.compile(r'^[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z]+$')
regex_match = re.search(r'\d{3}-\d{3}-\d{4}', "Call 555-123-4567")

# Numbers in bright green (#C7F026)
integer = 42
float_num = 3.14159
scientific = 1.23e-4
hex_num = 0xFF410D
binary_num = 0b1010101
octal_num = 0o755
complex_num = 3 + 4j

# Constants and special values (orange with italic)
CONSTANT_VALUE = "GRID_CONSTANT"
MAX_POWER_LEVEL = 9000
DEBUG_MODE = True
is_active = False
nothing = None
not_implemented = NotImplemented
ellipsis_literal = ...

# Variables in light blue-gray (#d0dfe6)
user_name = "Flynn"
grid_coordinates = [10, 20, 30]
program_data = {"name": "Tron", "type": "Security"}

# Function definitions (orange #FFB20D for names)
def regular_function(param1: str, param2: int = 10) -> str:
    """Function with parameters in bright green italic (#95CC5E)"""
    return f"{param1}: {param2}"

# Built-in functions (blue italic #267fb5)
print("Built-in function example")
len([1, 2, 3, 4])
range(10)
isinstance(grid_coordinates, list)
hasattr(program_data, "name")

# Lambda functions
square = lambda x: x ** 2
add = lambda x, y: x + y

# Keywords in gray (#748aa6)
if condition:
    for i in range(10):
        while True:
            break
    else:
        continue
elif other_condition:
    pass
else:
    try:
        raise ValueError("Error in the Grid")
    except ValueError as e:
        print(f"Caught: {e}")
    finally:
        cleanup()

# Classes and inheritance (dark orange #F79D1E)
class GridEntity:
    """Base class for Grid entities"""

    def __init__(self, name: str, power: int):
        self.name = name  # Instance variables in light blue-gray
        self.power = power
        self._private = "hidden"
        self.__very_private = "more hidden"

    @property
    def power_level(self) -> int:
        """Property decorator"""
        return self.power

    @power_level.setter
    def power_level(self, value: int):
        self.power = max(0, min(value, MAX_POWER_LEVEL))

    @staticmethod
    def static_method():
        """Static method example"""
        return "No instance needed"

    @classmethod
    def class_method(cls):
        """Class method example"""
        return cls.__name__

    def __str__(self):
        """Special method (dunder)"""
        return f"{self.name} (Power: {self.power})"

# Inheritance
class Program(GridEntity):
    """Derived class"""

    def __init__(self, name: str, power: int, function: str):
        super().__init__(name, power)  # super() is special
        self.function = function

    async def execute(self):
        """Async method"""
        await self._process()
        return "Executed"

# Decorators
@lru_cache(maxsize=128)
@wraps(regular_function)
def memoized_function(x: int) -> int:
    """Decorated function"""
    return x ** 2

# Custom decorator
def grid_access(level: int = 1):
    def decorator(func):
        @wraps(func)
        def wrapper(*args, **kwargs):
            print(f"Access level {level} granted")
            return func(*args, **kwargs)
        return wrapper
    return decorator

@grid_access(level=5)
def secure_function():
    return "Secure data"

# Type hints and annotations
Vector = List[float]
Matrix = List[List[float]]
UserDict = Dict[str, Union[str, int, bool]]

def typed_function(
    users: List[str],
    settings: Optional[Dict[str, Any]] = None,
    *args: int,
    **kwargs: str
) -> Tuple[bool, str]:
    """Function with complex type hints"""
    return True, "Success"

# Generator and comprehensions
def program_generator():
    """Generator function"""
    programs = ["Tron", "CLU", "Quorra", "Flynn"]
    for program in programs:
        yield program.upper()

# List comprehension with conditionals
active_programs = [p for p in programs if p.startswith("T")]
program_powers = {p: len(p) * 100 for p in programs}
unique_letters = {char for word in programs for char in word}

# Context managers
with open("grid_data.txt", "r") as file:
    data = file.read()

# Custom context manager
class GridContext:
    def __enter__(self):
        print("Entering the Grid")
        return self

    def __exit__(self, exc_type, exc_val, exc_tb):
        print("Exiting the Grid")
        return False

# Async/await syntax
async def async_grid_operation():
    async with aiohttp.ClientSession() as session:
        async for data in fetch_stream(session):
            process(data)

# Special self keyword (purple italic #967EFB)
class IdentityDisc:
    def __init__(self):
        self.data = []

    def add_memory(self, memory):
        self.data.append(memory)
        return self  # Method chaining

# Operator overloading
class Vector3D:
    def __init__(self, x, y, z):
        self.x, self.y, self.z = x, y, z

    def __add__(self, other):
        return Vector3D(
            self.x + other.x,
            self.y + other.y,
            self.z + other.z
        )

    def __mul__(self, scalar):
        return Vector3D(
            self.x * scalar,
            self.y * scalar,
            self.z * scalar
        )

# Walrus operator (Python 3.8+)
if (n := len(programs)) > 3:
    print(f"Found {n} programs")

# Match statement (Python 3.10+)
match command:
    case "activate":
        activate_grid()
    case ["move", x, y]:
        move_to(x, y)
    case {"action": "scan", "target": target}:
        scan_target(target)
    case _:
        print("Unknown command")
