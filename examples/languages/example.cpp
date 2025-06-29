// Tron Legacy Theme Example - C++
// This file demonstrates the variety of syntax highlighting colors

// Single-line comment in dark gray (#3c4b5d)
/* Multi-line comment
   Also in dark gray */

/// Documentation comment (Doxygen style)
/// @brief Example class for the Grid system
/// @author Flynn

//! Another documentation style

// Headers and includes
#include <iostream>
#include <string>
#include <vector>
#include <memory>
#include <algorithm>
#include <functional>
#include <thread>
#include <chrono>
#include <type_traits>
#include <optional>
#include <variant>

// Using declarations
using namespace std;
using std::string;
using std::vector;

// Namespace definition
namespace Grid {
    namespace Core {
        constexpr int VERSION = 2;
    }

    // Inline namespace (C++11)
    inline namespace V2 {
        void initialize();
    }
}

// Preprocessor directives (cyan #6ee2ff)
#define GRID_DEBUG 1
#define MAX_POWER_LEVEL 9000
#define STRINGIFY(x) #x
#define CONCAT(a, b) a##b

#ifdef GRID_DEBUG
    #define LOG(msg) std::cout << "[DEBUG] " << msg << std::endl
#else
    #define LOG(msg)
#endif

// Forward declarations
class GridProgram;
struct ProgramData;
template<typename T> class Container;

// Type aliases (C++11)
using ProgramID = std::uint64_t;
using PowerLevel = float;
using ProgramPtr = std::shared_ptr<GridProgram>;
using Callback = std::function<void(const GridProgram&)>;

// Template alias
template<typename T>
using Vec = std::vector<T>;

// Constants (orange with italic #FFB20D)
const double PI = 3.14159265358979323846;
constexpr int MAX_PROGRAMS = 100;
inline constexpr auto GRID_NAME = "ENCOM OS-12";

// Global variables (light blue-gray #d0dfe6)
std::atomic<int> program_count{0};
thread_local int thread_id = 0;

// Enums (dark orange #F79D1E)
enum class ProgramType {
    Security,
    Utility,
    Game,
    System
};

// Enum with underlying type
enum class Status : uint8_t {
    Offline = 0,
    Initializing = 1,
    Online = 2,
    Error = 255
};

// Struct with attributes
struct [[nodiscard]] ProgramData {
    std::string name;           // Fields in light blue-gray
    ProgramID id;
    PowerLevel power{100.0f};
    bool active = false;

    // Default constructor
    ProgramData() = default;

    // Constructor with initializer list
    ProgramData(std::string n, ProgramID i)
        : name(std::move(n)), id(i) {}
};

// Class definition (dark orange #F79D1E)
class GridProgram {
private:
    ProgramData data_;
    mutable std::mutex mutex_;
    static inline int instance_count_ = 0;

public:
    // Constructors and destructor
    GridProgram() : GridProgram("Unknown", 0) {}

    explicit GridProgram(std::string name, ProgramID id = 0)
        : data_{std::move(name), id} {
        ++instance_count_;
    }

    // Copy constructor
    GridProgram(const GridProgram& other)
        : data_(other.data_) {
        ++instance_count_;
    }

    // Move constructor
    GridProgram(GridProgram&& other) noexcept
        : data_(std::move(other.data_)) {
        ++instance_count_;
    }

    // Destructor
    virtual ~GridProgram() {
        --instance_count_;
    }

    // Copy and move assignment
    GridProgram& operator=(const GridProgram& other) {
        if (this != &other) {
            data_ = other.data_;
        }
        return *this;
    }

    GridProgram& operator=(GridProgram&& other) noexcept {
        if (this != &other) {
            data_ = std::move(other.data_);
        }
        return *this;
    }

    // Member functions (orange #FFB20D)
    void activate() {
        std::lock_guard<std::mutex> lock(mutex_);
        data_.active = true;
    }

    [[nodiscard]] bool isActive() const {
        std::lock_guard<std::mutex> lock(mutex_);
        return data_.active;
    }

    // Function with parameters (bright green italic #95CC5E)
    void transferPower(GridProgram& target, PowerLevel amount) {
        std::scoped_lock lock(mutex_, target.mutex_);

        amount = std::min(amount, data_.power);
        data_.power -= amount;
        target.data_.power += amount;
    }

    // Const member function
    [[nodiscard]] const std::string& getName() const noexcept {
        return data_.name;
    }

    // Static member function
    [[nodiscard]] static int getInstanceCount() noexcept {
        return instance_count_;
    }

    // Virtual functions
    virtual void execute() = 0;
    virtual void shutdown() { data_.active = false; }

    // Operator overloading
    bool operator==(const GridProgram& other) const {
        return data_.id == other.data_.id;
    }

    bool operator<(const GridProgram& other) const {
        return data_.id < other.data_.id;
    }

    // Friend function
    friend std::ostream& operator<<(std::ostream& os, const GridProgram& prog);
};

// Derived class
class SecurityProgram : public GridProgram {
private:
    int security_level_;

public:
    using GridProgram::GridProgram;  // Inherit constructors

    explicit SecurityProgram(std::string name, int level = 1)
        : GridProgram(std::move(name)), security_level_(level) {}

    // Override virtual function
    void execute() override {
        LOG("Executing security scan...");
    }

    // Final function
    void shutdown() final {
        LOG("Security program shutting down");
        GridProgram::shutdown();
    }
};

// Template class (dark orange #F79D1E)
template<typename T, size_t N = 10>
class Container {
private:
    std::array<T, N> items_;
    size_t count_ = 0;

public:
    // Template member function
    template<typename... Args>
    void emplace(Args&&... args) {
        if (count_ < N) {
            items_[count_++] = T(std::forward<Args>(args)...);
        }
    }

    // Iterator support
    auto begin() { return items_.begin(); }
    auto end() { return items_.begin() + count_; }

    // Const iterators
    auto begin() const { return items_.cbegin(); }
    auto end() const { return items_.cbegin() + count_; }
};

// Template specialization
template<>
class Container<bool, 8> {
    uint8_t bits_ = 0;
public:
    void set(size_t index, bool value) {
        if (value) {
            bits_ |= (1 << index);
        } else {
            bits_ &= ~(1 << index);
        }
    }
};

// Function templates
template<typename T>
[[nodiscard]] constexpr T square(T value) noexcept {
    return value * value;
}

// Variadic template
template<typename... Args>
void gridLog(Args&&... args) {
    std::cout << "[GRID] ";
    ((std::cout << args << " "), ...);  // Fold expression (C++17)
    std::cout << std::endl;
}

// Concepts (C++20)
template<typename T>
concept Numeric = std::is_arithmetic_v<T>;

template<Numeric T>
T add(T a, T b) {
    return a + b;
}

// Lambda expressions
auto createProgram = [](const std::string& name) -> ProgramPtr {
    return std::make_shared<SecurityProgram>(name);
};

// Generic lambda (C++14)
auto print = [](const auto& value) {
    std::cout << value << std::endl;
};

// Mutable lambda
auto counter = [count = 0]() mutable {
    return ++count;
};

// Main function
int main() {
    // String literals (red #FF410D)
    std::string grid_name = "ENCOM Digital Grid";
    const char* version = "2.0";

    // String with escape sequences (light red #ff6113)
    std::cout << "Initializing...\n\tPhase 1\n\tPhase 2\n";
    std::cout << "Special: \x1B[1mBold\x1B[0m\n";

    // Raw string literal (C++11)
    auto query = R"(
        SELECT * FROM programs
        WHERE status = 'active'
        AND power > 50
    )";

    // Character literals
    char grid_marker = 'G';
    wchar_t wide_char = L'Ω';
    char16_t utf16 = u'€';
    char32_t utf32 = U'🎮';

    // Number literals (bright green #C7F026)
    int decimal = 42;
    int hex = 0xFF;
    int octal = 0755;
    int binary = 0b1010'1010;  // Digit separator (C++14)

    // Integer suffixes
    auto small = 42u;
    auto large = 42ull;
    auto bignum = 18'446'744'073'709'551'615ull;

    // Floating point
    float power = 95.5f;
    double precision = 3.141'592'653'589'793;
    long double extended = 1.23e-4L;

    // User-defined literals
    using namespace std::chrono_literals;
    auto duration = 100ms;
    auto wait_time = 2.5s;

    // Boolean values
    bool is_active = true;
    bool is_corrupt = false;

    // Null pointers
    void* null_ptr = nullptr;
    GridProgram* prog_ptr = nullptr;

    // Auto type deduction
    auto programs = std::vector<ProgramPtr>{};
    auto [x, y, z] = std::tuple{1, 2, 3};  // Structured binding (C++17)

    // Range-based for loop
    for (const auto& program : programs) {
        if (program) {
            program->execute();
        }
    }

    // Algorithms and lambdas
    std::sort(programs.begin(), programs.end(),
        [](const auto& a, const auto& b) {
            return a->getName() < b->getName();
        });

    // STL usage
    std::vector<int> values = {1, 2, 3, 4, 5};
    auto sum = std::accumulate(values.begin(), values.end(), 0);

    // Smart pointers
    auto unique_prog = std::make_unique<SecurityProgram>("Tron");
    auto shared_prog = std::make_shared<SecurityProgram>("CLU");
    std::weak_ptr<GridProgram> weak_ref = shared_prog;

    // Exception handling (keywords in gray #748aa6)
    try {
        if (!unique_prog) {
            throw std::runtime_error("Program creation failed");
        }
        unique_prog->execute();
    } catch (const std::exception& e) {
        std::cerr << "Error: " << e.what() << std::endl;
    } catch (...) {
        std::cerr << "Unknown error occurred" << std::endl;
    }

    // RAII and scope guards
    {
        std::lock_guard<std::mutex> lock(global_mutex);
        // Critical section
    }

    // Optional and variant (C++17)
    std::optional<int> maybe_value = 42;
    if (maybe_value.has_value()) {
        std::cout << *maybe_value << std::endl;
    }

    std::variant<int, float, std::string> data = "Grid";
    std::visit([](const auto& v) { print(v); }, data);

    // Attributes
    [[maybe_unused]] int unused_var = 10;
    [[deprecated("Use new API")]] void old_function();

    switch (Status::Online) {
        [[likely]] case Status::Online:
            gridLog("System online");
            break;
        [[unlikely]] case Status::Error:
            gridLog("System error");
            [[fallthrough]];
        default:
            gridLog("Unknown status");
    }

    // Coroutines (C++20)
    // co_await some_async_operation();
    // co_return result;

    return EXIT_SUCCESS;
}

// Friend function implementation
std::ostream& operator<<(std::ostream& os, const GridProgram& prog) {
    os << "Program[" << prog.data_.name << ", ID:" << prog.data_.id << "]";
    return os;
}

// Function using self/this (purple italic #967EFB)
class GridEntity {
    std::string name_;
public:
    GridEntity& setName(std::string name) {
        this->name_ = std::move(name);  // 'this' in purple
        return *this;  // Method chaining
    }

    void process() const {
        auto lambda = [this]() {  // Capture 'this'
            std::cout << this->name_ << std::endl;
        };
        lambda();
    }
};

// Compile-time computation
template<int N>
struct Factorial {
    static constexpr int value = N * Factorial<N-1>::value;
};

template<>
struct Factorial<0> {
    static constexpr int value = 1;
};

// Using constexpr
constexpr int fibonacci(int n) {
    return (n <= 1) ? n : fibonacci(n-1) + fibonacci(n-2);
}

// Type traits and SFINAE
template<typename T>
typename std::enable_if<std::is_integral<T>::value, T>::type
safe_divide(T a, T b) {
    return b != 0 ? a / b : 0;
}

// Trailing return type
auto multiply(int a, int b) -> decltype(a * b) {
    return a * b;
}

// Initializer list
void process_values(std::initializer_list<int> values) {
    for (auto v : values) {
        std::cout << v << " ";
    }
}

// Usage: process_values({1, 2, 3, 4, 5});
