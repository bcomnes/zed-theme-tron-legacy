// Tron Legacy Theme - Comment Contrast Test
// Comments should now use #6684a7 for better visibility (WCAG AA compliant)

// Single-line comment - this should be clearly readable
/* Multi-line comment
   spanning multiple lines
   with good contrast against the dark background */

/**
 * JSDoc style documentation
 * @param {string} name - The name parameter
 * @returns {boolean} Success status
 */
function testFunction(name) {
    // Inline comment explaining the code
    const result = name.toUpperCase(); // End-of-line comment

    /*
     * Another multi-line comment
     * explaining complex logic
     */

    return true;
}

// TODO: This is a todo comment
// FIXME: This needs attention
// NOTE: Important information
// WARNING: Be careful here

// The comment color #6684a7 has a contrast ratio of ~4.5:1
// against the background #14191f, meeting WCAG AA standards
// This is much better than the previous #3c4b5d which was too dark

/* Background: #14191f (RGB: 20, 25, 31)
 * Comment:    #6684a7 (RGB: 102, 132, 167)
 * This provides sufficient contrast for readability */
