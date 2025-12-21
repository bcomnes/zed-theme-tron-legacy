# Diff Syntax Highlighting Example

This file demonstrates diff/patch syntax highlighting in the Tron Legacy theme for Zed.

## Standard Diff Format

```diff
diff --git a/src/example.js b/src/example.js
index 1234567..89abcdef 100644
--- a/src/example.js
+++ b/src/example.js
@@ -1,10 +1,12 @@
 function greet(name) {
-  console.log("Hello, " + name);
+  console.log(`Hello, ${name}!`);
+  return name;
 }
 
 const users = [
-  { id: 1, name: "Alice" },
-  { id: 2, name: "Bob" }
+  { id: 1, name: "Alice", active: true },
+  { id: 2, name: "Bob", active: true },
+  { id: 3, name: "Charlie", active: false }
 ];
 
 greet("World");
```

## Color Guide

In the Tron Legacy theme, diff syntax uses these colors:

- **Green** (`#C7F026ff`) - Added lines (lines starting with `+`)
- **Red** (`#F92672ff`) - Removed lines (lines starting with `-`)
- **Yellow** (`#ffd12cff`) - Modified sections and context markers
- **Cyan** (`#4a95b3ff`) - Moved content

## Git Diff Example

```diff
diff --git a/package.json b/package.json
index abc123..def456 100644
--- a/package.json
+++ b/package.json
@@ -1,6 +1,7 @@
 {
   "name": "my-project",
-  "version": "1.0.0",
+  "version": "1.1.0",
+  "description": "An example project",
   "scripts": {
     "start": "node index.js",
-    "test": "jest"
+    "test": "jest --coverage"
   }
 }
```



## Patch File Example

```diff
--- old-file.txt	2024-01-15 10:30:00
+++ new-file.txt	2024-01-15 11:45:00
@@ -5,7 +5,8 @@
 This is line 5
 This is line 6
-This is the old line 7
+This is the new line 7
+This is an additional line 8
 This is line 8 (now 9)
 This is line 9 (now 10)
```

## Expected Appearance

When viewing this file in Zed with the Tron Legacy theme:

1. Lines starting with `-` should appear in **red** (deletions)
2. Lines starting with `+` should appear in **green** (additions)
3. Context lines (no prefix or space prefix) should use default text color
4. Diff headers (`@@`, `---`, `+++`, `diff --git`) should use contextual colors

This rational color scheme makes it easy to quickly identify:
- What was removed (red = stop/delete)
- What was added (green = go/add)
- What changed (yellow = caution/modified)