// 2667. Create Hello WOrld Functions
// https://leetcode.com/problems/create-hello-world-function
// Time Complexity: O(1)

function createHelloWorld() {
    
    return function(...args): string {
        return "Hello World"
    };
};

/**
 * const f = createHelloWorld();
 * f(); // "Hello World"
 */