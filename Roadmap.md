# Frontend Engineer Interview Prep Plan (4 YOE | FAANG + Top Startups)

> **Profile:** 4 years frontend experience | React | Targeting FAANG + Tier-1 startups
> **Pace:** 1-2 hours/day | ~16 weeks (flexible)
> **Key insight:** At 4 YOE, you'll be interviewing for **L4/E4 (mid-level)** at most companies. Some may push for L5/E5 (senior). Frontend interviews at top companies have 5-6 distinct round types -- you need to be ready for ALL of them.

---

## Table of Contents

1. [Interview Round Types at Top Companies](#1-interview-round-types-at-top-companies)
2. [Phase 1: JavaScript & TypeScript Foundations (Weeks 1-3)](#phase-1-javascript--typescript-foundations-weeks-1-3)
3. [Phase 2: HTML, CSS & Browser Internals (Weeks 3-4)](#phase-2-html-css--browser-internals-weeks-3-4)
4. [Phase 3: React Deep Dive (Weeks 4-6)](#phase-3-react-deep-dive-weeks-4-6)
5. [Phase 4: DSA for Frontend (Weeks 5-12)](#phase-4-dsa-for-frontend-weeks-5-12)
6. [Phase 5: Frontend System Design (Weeks 8-13)](#phase-5-frontend-system-design-weeks-8-13)
7. [Phase 6: Web Performance & Optimization (Weeks 10-11)](#phase-6-web-performance--optimization-weeks-10-11)
8. [Phase 7: Testing Strategies (Week 12)](#phase-7-testing-strategies-week-12)
9. [Phase 8: Behavioral & Leadership (Weeks 13-16)](#phase-8-behavioral--leadership-weeks-13-16)
10. [Phase 9: Machine Coding / UI Rounds (Weeks 6-16)](#phase-9-machine-coding--ui-rounds-weeks-6-16)
11. [Company-Specific Breakdowns](#company-specific-breakdowns)
12. [Weekly Schedule Template](#weekly-schedule-template)
13. [Resources Master List](#resources-master-list)
14. [Progress Tracker](#progress-tracker)

---

## 1. Interview Round Types at Top Companies

Every FAANG/top startup frontend loop has some combination of these:

| Round | What They Test | Weight |
|-------|---------------|--------|
| **DSA / Coding** | Problem-solving, algorithms, data structures, time/space analysis | High |
| **JavaScript / Language** | Deep JS knowledge, closures, promises, prototypes, event loop | High |
| **UI / Machine Coding** | Build a component/feature from scratch in 45-60 min | Very High |
| **Frontend System Design** | Design large-scale frontend apps (like Google Docs, Twitter Feed) | High (especially L5) |
| **Behavioral / Culture Fit** | Past experience, leadership, conflict resolution | Medium-High |
| **CSS / HTML Quiz** | Accessibility, semantic HTML, CSS layout, animations | Medium |

### What Each Company Emphasizes

| Company | DSA | JS Deep Dive | UI/Machine Coding | System Design | Behavioral |
|---------|-----|-------------|-------------------|---------------|------------|
| **Google** | Very Heavy | Medium | Medium | Heavy | Medium |
| **Meta** | Heavy | Heavy | Very Heavy | Heavy (E5+) | Medium |
| **Amazon** | Heavy | Medium | Medium | Medium | Very Heavy (LPs) |
| **Apple** | Heavy | Heavy | Heavy | Medium | Medium |
| **Netflix** | Medium | Very Heavy | Heavy | Heavy | Heavy |
| **Stripe** | Heavy | Heavy | Very Heavy | Heavy | Medium |
| **Airbnb** | Medium | Heavy | Very Heavy | Heavy | Heavy |
| **Uber** | Heavy | Medium | Heavy | Heavy | Medium |

---

## Phase 1: JavaScript & TypeScript Foundations (Weeks 1-3)

### Why This Matters
JS trivia/deep-dive rounds are **the most common rejection reason** for frontend candidates. Companies like Meta and Netflix will specifically test if you truly understand the language, not just the framework.

### Week 1: Core JavaScript Concepts

#### Day 1-2: Execution Context & Scope
- [ ] How JavaScript engine executes code (creation phase vs execution phase)
- [ ] Global execution context vs function execution context
- [ ] Variable hoisting (`var` vs `let` vs `const`) -- be able to predict output
- [ ] Temporal Dead Zone (TDZ)
- [ ] Scope chain and lexical scoping

**Practice questions:**
```javascript
// Q1: What's the output?
console.log(a);
var a = 1;
console.log(a);
let b = 2;

// Q2: What's the output?
for (var i = 0; i < 3; i++) {
  setTimeout(() => console.log(i), 1000);
}

// Q3: Fix Q2 without changing var to let (3 different ways)

// Q4: What's the output?
function foo() {
  console.log(this);
}
const obj = { foo };
const bar = obj.foo;
obj.foo();  // ?
bar();      // ?
```

#### Day 3-4: Closures & Functions
- [ ] What closures are and how they work (explain with memory model)
- [ ] Practical uses: data privacy, function factories, partial application
- [ ] IIFE pattern and why it exists
- [ ] Implement: `once()`, `memoize()`, `curry()`, `debounce()`, `throttle()`
- [ ] Arrow functions vs regular functions (5 differences)

**Must-implement from scratch:**
```javascript
// 1. debounce
function debounce(fn, delay) { /* implement */ }

// 2. throttle
function throttle(fn, limit) { /* implement */ }

// 3. curry - should support both curry(1)(2)(3) and curry(1,2)(3)
function curry(fn) { /* implement */ }

// 4. memoize - with support for multiple args
function memoize(fn) { /* implement */ }

// 5. once - function that only executes once
function once(fn) { /* implement */ }

// 6. pipe & compose
function pipe(...fns) { /* implement */ }
function compose(...fns) { /* implement */ }
```

#### Day 5-6: `this`, Prototypes & Inheritance
- [ ] 4 rules of `this` binding (default, implicit, explicit, new)
- [ ] `call`, `apply`, `bind` -- implement all three from scratch
- [ ] Prototype chain: `__proto__` vs `prototype` vs `Object.getPrototypeOf()`
- [ ] How `new` keyword works (implement `myNew()`)
- [ ] ES6 classes are syntactic sugar -- understand what they compile to
- [ ] `instanceof` -- implement it from scratch
- [ ] Inheritance patterns: prototypal vs classical vs functional

**Implement from scratch:**
```javascript
// 1. Function.prototype.myBind
Function.prototype.myBind = function(context, ...args) { /* implement */ }

// 2. Function.prototype.myCall
Function.prototype.myCall = function(context, ...args) { /* implement */ }

// 3. Function.prototype.myApply
Function.prototype.myApply = function(context, args) { /* implement */ }

// 4. instanceof
function myInstanceOf(obj, Constructor) { /* implement */ }

// 5. Object.create
function myObjectCreate(proto) { /* implement */ }

// 6. new keyword
function myNew(Constructor, ...args) { /* implement */ }
```

#### Day 7: Event Loop & Async Model
- [ ] Call stack, Web APIs, callback queue, microtask queue
- [ ] `setTimeout` vs `setInterval` vs `requestAnimationFrame`
- [ ] Microtasks (Promise callbacks, queueMicrotask, MutationObserver) vs Macrotasks (setTimeout, setInterval, I/O)
- [ ] Event loop order: call stack -> microtask queue -> render -> macrotask queue
- [ ] Starvation: what happens if microtask queue never empties?

**Output prediction questions (practice 20+ of these):**
```javascript
// Q1:
console.log('1');
setTimeout(() => console.log('2'), 0);
Promise.resolve().then(() => console.log('3'));
console.log('4');
// Output: ?

// Q2:
setTimeout(() => console.log('A'), 0);
Promise.resolve()
  .then(() => {
    console.log('B');
    setTimeout(() => console.log('C'), 0);
  })
  .then(() => console.log('D'));
Promise.resolve().then(() => console.log('E'));
console.log('F');
// Output: ?

// Q3:
async function foo() {
  console.log('1');
  await Promise.resolve();
  console.log('2');
}
console.log('3');
foo();
console.log('4');
// Output: ?
```

### Week 2: Advanced JavaScript

#### Day 8-9: Promises & Async/Await Deep Dive
- [ ] Promise states: pending, fulfilled, rejected
- [ ] Promise chaining and error propagation
- [ ] `Promise.all` vs `Promise.allSettled` vs `Promise.race` vs `Promise.any`
- [ ] Implement `Promise` from scratch (this is asked at Google, Meta)
- [ ] Implement `Promise.all`, `Promise.race`, `Promise.allSettled` from scratch
- [ ] async/await -- how it desugars to generators + promises
- [ ] Error handling in async/await (try/catch vs .catch())
- [ ] Concurrent vs sequential async operations
- [ ] Implement: retry with exponential backoff, parallel limit executor

**Implement from scratch:**
```javascript
// 1. Full Promise implementation (A+ spec)
class MyPromise {
  constructor(executor) { /* implement */ }
  then(onFulfilled, onRejected) { /* implement */ }
  catch(onRejected) { /* implement */ }
  finally(callback) { /* implement */ }
  static resolve(value) { /* implement */ }
  static reject(reason) { /* implement */ }
  static all(promises) { /* implement */ }
  static race(promises) { /* implement */ }
  static allSettled(promises) { /* implement */ }
  static any(promises) { /* implement */ }
}

// 2. promisify - convert callback-based function to promise-based
function promisify(fn) { /* implement */ }

// 3. Retry with exponential backoff
function retry(fn, retries = 3, delay = 1000) { /* implement */ }

// 4. Execute promises with concurrency limit
function parallelLimit(tasks, limit) { /* implement */ }

// 5. Auto-retry fetch with timeout
function fetchWithRetry(url, options, retries = 3) { /* implement */ }
```

#### Day 10-11: Iterators, Generators & Advanced Patterns
- [ ] Symbol.iterator protocol
- [ ] Generators: `function*`, `yield`, `yield*`
- [ ] Async iterators and `for await...of`
- [ ] Proxy and Reflect -- how Vue 3 reactivity works
- [ ] WeakMap, WeakSet, WeakRef -- use cases
- [ ] Symbols and well-known symbols
- [ ] Tagged template literals
- [ ] Structured clone vs JSON.parse(JSON.stringify) vs spread

**Implement from scratch:**
```javascript
// 1. Deep clone (handle Date, RegExp, Map, Set, circular refs)
function deepClone(obj) { /* implement */ }

// 2. Deep equal comparison
function deepEqual(a, b) { /* implement */ }

// 3. EventEmitter (on, off, emit, once)
class EventEmitter { /* implement */ }

// 4. Observable pattern
class Observable { /* implement */ }

// 5. Make any object iterable with a range
function range(start, end) { /* implement -- should work with for...of */ }
```

#### Day 12-13: DOM & Browser APIs
- [ ] DOM traversal: parentNode, children, nextSibling, querySelector
- [ ] Event propagation: capturing, target, bubbling phases
- [ ] `addEventListener` options (capture, once, passive)
- [ ] Event delegation and why it matters at scale
- [ ] `stopPropagation` vs `stopImmediatePropagation` vs `preventDefault`
- [ ] IntersectionObserver, MutationObserver, ResizeObserver
- [ ] `requestAnimationFrame` vs `requestIdleCallback`
- [ ] Web Storage API: localStorage, sessionStorage, cookies, IndexedDB
- [ ] Fetch API: headers, CORS, credentials, AbortController

**Implement from scratch:**
```javascript
// 1. Event delegation system
function delegate(parent, selector, event, handler) { /* implement */ }

// 2. DOM manipulation: createElement with nested children
function createElement(tag, props, ...children) { /* implement */ }

// 3. Virtual DOM diffing (simplified)
function diff(oldVTree, newVTree) { /* implement */ }

// 4. Simple data binding (two-way)
function createReactiveObject(obj, onChange) { /* implement -- use Proxy */ }

// 5. Infinite scroll using IntersectionObserver
function createInfiniteScroll(container, loadMore) { /* implement */ }

// 6. Custom fetch wrapper with timeout, retry, interceptors
class HttpClient { /* implement */ }
```

#### Day 14: Module Systems & Tooling
- [ ] CommonJS vs ES Modules (differences, import/export syntax)
- [ ] Tree shaking -- how it works, what prevents it
- [ ] Module resolution in Node.js vs browsers
- [ ] Bundlers: Webpack, Vite, esbuild -- high-level understanding
- [ ] Source maps, code splitting, lazy loading
- [ ] Babel -- what it does and why (transpilation vs polyfilling)

### Week 3: TypeScript

#### Day 15-17: TypeScript Essentials
- [ ] Basic types: string, number, boolean, null, undefined, void, never, unknown, any
- [ ] `unknown` vs `any` (why `unknown` is preferred)
- [ ] `never` type -- exhaustive checks, unreachable code
- [ ] Type narrowing: typeof, instanceof, discriminated unions, type predicates
- [ ] Union types, intersection types, literal types
- [ ] Type aliases vs interfaces (when to use which)
- [ ] Generics: functions, classes, constraints, defaults
- [ ] Generic constraints with `extends`
- [ ] Conditional types: `T extends U ? X : Y`
- [ ] Mapped types: `Partial`, `Required`, `Readonly`, `Pick`, `Omit`, `Record`
- [ ] Template literal types
- [ ] `infer` keyword in conditional types
- [ ] Declaration merging
- [ ] Module augmentation

**Practice: Implement these utility types from scratch:**
```typescript
// 1. MyPartial<T>
type MyPartial<T> = { [K in keyof T]?: T[K] };

// 2. MyRequired<T>
type MyRequired<T> = /* implement */

// 3. MyReadonly<T>
type MyReadonly<T> = /* implement */

// 4. MyPick<T, K>
type MyPick<T, K extends keyof T> = /* implement */

// 5. MyOmit<T, K>
type MyOmit<T, K extends keyof T> = /* implement */

// 6. MyReturnType<T>
type MyReturnType<T extends (...args: any[]) => any> = /* implement */

// 7. MyParameters<T>
type MyParameters<T extends (...args: any[]) => any> = /* implement */

// 8. DeepPartial<T>
type DeepPartial<T> = /* implement */

// 9. DeepReadonly<T>
type DeepReadonly<T> = /* implement */

// 10. MyAwaited<T> (unwrap Promise type)
type MyAwaited<T> = /* implement */
```

#### Day 18-21: TypeScript with React Patterns
- [ ] Typing component props (FC vs function declaration)
- [ ] Typing state with `useState<T>`
- [ ] Typing refs with `useRef<T>`
- [ ] Typing events: `React.ChangeEvent<HTMLInputElement>`, `React.MouseEvent`, etc.
- [ ] Typing context with `createContext<T>`
- [ ] Generic components
- [ ] Discriminated unions for component props
- [ ] Typing HOCs and render props
- [ ] Typing custom hooks with generics
- [ ] Typing reducers with discriminated union actions
- [ ] `as const` assertions for action types
- [ ] React.ComponentProps, React.PropsWithChildren, React.HTMLAttributes

---

## Phase 2: HTML, CSS & Browser Internals (Weeks 3-4)

### Why This Matters
CSS/HTML rounds eliminate candidates who "just use a component library." You MUST know layout from first principles.

### Week 3 (continued) & Week 4

#### CSS Layout Mastery
- [ ] Box model: content-box vs border-box, margin collapsing rules
- [ ] Display: block, inline, inline-block, none, contents
- [ ] **Flexbox** (know every property cold):
  - `flex-direction`, `justify-content`, `align-items`, `align-self`
  - `flex-grow`, `flex-shrink`, `flex-basis` and the `flex` shorthand
  - How to center anything (horizontal, vertical, both)
  - Common layouts: navbar, card grid, holy grail, sticky footer
- [ ] **Grid** (know every property cold):
  - `grid-template-columns/rows`, `grid-template-areas`
  - `fr` unit, `minmax()`, `repeat()`, `auto-fill` vs `auto-fit`
  - `grid-column`, `grid-row`, span syntax
  - Named grid lines and areas
  - Implicit vs explicit grid
- [ ] **Positioning**: static, relative, absolute, fixed, sticky
  - Stacking context: what creates one, how `z-index` actually works
  - Containing block rules
- [ ] BFC (Block Formatting Context) -- what creates one, why it matters

#### CSS Deep Topics
- [ ] Specificity calculation (0,0,0,0 model)
- [ ] Cascade: importance > specificity > source order
- [ ] CSS custom properties (variables) -- scoping, inheritance, fallbacks
- [ ] CSS-in-JS tradeoffs (styled-components, emotion, CSS modules, Tailwind)
- [ ] Responsive design: media queries, container queries, clamp(), fluid typography
- [ ] CSS animations vs transitions vs Web Animations API
- [ ] `transform`, `opacity`, `will-change` -- what triggers GPU compositing
- [ ] Pseudo-elements (::before, ::after) and pseudo-classes (:nth-child, :has, :is)
- [ ] CSS logical properties (inline/block vs left/right)

#### HTML & Accessibility
- [ ] Semantic HTML: when to use `section`, `article`, `aside`, `nav`, `main`, `header`, `footer`
- [ ] ARIA roles, `aria-label`, `aria-describedby`, `aria-live`
- [ ] Focus management: tabindex, focus trapping in modals
- [ ] Keyboard navigation requirements (WCAG 2.1)
- [ ] Form accessibility: labels, fieldsets, error announcements
- [ ] `<picture>`, `srcset`, responsive images
- [ ] `<dialog>` element
- [ ] Open Graph & meta tags for SEO

#### Browser Internals (High-Level)
- [ ] Critical rendering path: HTML parse -> DOM -> CSSOM -> Render Tree -> Layout -> Paint -> Composite
- [ ] Reflow vs repaint -- what triggers each
- [ ] `<script>` loading: regular vs `async` vs `defer`
- [ ] Preload, prefetch, preconnect, dns-prefetch
- [ ] How CSS and JS block rendering
- [ ] CORS: simple vs preflight requests, when OPTIONS is sent

**CSS coding challenges (practice these):**
```
1. Center a div (5 different ways)
2. Create a responsive card grid (no media queries, using grid)
3. Build a tooltip with pure CSS (using ::after and :hover)
4. Create a modal overlay with backdrop blur
5. Build a responsive navbar with hamburger menu (CSS only)
6. Implement a CSS-only accordion
7. Create a loading spinner animation
8. Build a sticky header that shrinks on scroll (CSS + minimal JS)
9. Implement a masonry layout
10. Create a theme switcher using CSS custom properties
```

---

## Phase 3: React Deep Dive (Weeks 4-6)

### Why This Matters
As a React developer with 4 years experience, interviewers expect **expert-level** understanding, not just usage knowledge.

### Core React Concepts (Must Know Cold)

#### Rendering & Reconciliation
- [ ] Virtual DOM and fiber architecture (high-level)
- [ ] Reconciliation algorithm: diffing, keys, and why they matter
- [ ] When does React re-render? (state change, parent re-render, context change)
- [ ] Batching: automatic batching in React 18 vs React 17
- [ ] `React.memo` -- when to use, when NOT to use (overhead of shallow comparison)
- [ ] `useMemo` vs `useCallback` -- actual use cases, premature optimization traps
- [ ] `useRef` -- not just for DOM refs (storing mutable values without re-renders)
- [ ] Strict Mode: why it double-renders, what it catches

#### Hooks Deep Dive
- [ ] `useState`: functional updates, lazy initialization
- [ ] `useEffect`: dependency array rules, cleanup timing, closure gotchas
- [ ] `useLayoutEffect` vs `useEffect` -- when to use which (DOM measurements)
- [ ] `useReducer`: complex state, state machines
- [ ] `useContext`: performance pitfalls, splitting contexts
- [ ] `useId`, `useDeferredValue`, `useTransition` (React 18)
- [ ] `useSyncExternalStore` -- for external state integration
- [ ] Custom hooks: composition patterns, rules

**Implement these custom hooks from scratch:**
```javascript
// 1. useDebounce(value, delay)
// 2. useThrottle(value, delay)
// 3. useLocalStorage(key, initialValue)
// 4. usePrevious(value)
// 5. useIntersectionObserver(ref, options)
// 6. useFetch(url, options) -- with loading, error, caching
// 7. useEventListener(event, handler, element)
// 8. useMediaQuery(query)
// 9. useClickOutside(ref, handler)
// 10. useAsync(asyncFn) -- with execute, loading, error, value
// 11. useWindowSize()
// 12. useKeyPress(targetKey)
```

#### State Management Patterns
- [ ] Local state vs lifted state vs global state (when to use each)
- [ ] Context API: performance problems at scale, how to mitigate
- [ ] Redux: core concepts (actions, reducers, store, middleware)
- [ ] Redux Toolkit vs classic Redux
- [ ] Zustand / Jotai / Recoil -- understand mental models
- [ ] Server state: React Query / TanStack Query vs SWR
  - [ ] Query invalidation, caching strategies, optimistic updates
  - [ ] Stale-while-revalidate pattern

#### Advanced React Patterns
- [ ] Compound components
- [ ] Render props pattern
- [ ] Higher-Order Components (HOCs) -- and why hooks replaced most uses
- [ ] Controlled vs uncontrolled components
- [ ] State reducer pattern
- [ ] Prop getter pattern
- [ ] Slots pattern (children composition)

#### React 18+ Features
- [ ] Concurrent features: `startTransition`, `useDeferredValue`
- [ ] Suspense for data fetching
- [ ] Streaming SSR with `renderToPipeableStream`
- [ ] React Server Components (RSC) -- conceptual understanding
- [ ] `useOptimistic` (React 19)

#### React Performance Optimization
- [ ] Identifying unnecessary re-renders (React DevTools profiler)
- [ ] Code splitting with `React.lazy` and `Suspense`
- [ ] Virtualization for long lists (react-window, tanstack-virtual)
- [ ] Image optimization (lazy loading, srcset, next/image)
- [ ] Avoiding prop drilling without over-using context
- [ ] Selective context subscription patterns

---

## Phase 4: DSA for Frontend (Weeks 5-12)

### Why This Matters
DSA is still the **highest-weight round** at Google, Meta, and Amazon. You said you're weak here -- this is your biggest risk area. Plan for **8 weeks** of consistent practice.

### Strategy for Someone Starting Weak

**Week 5-6: Foundations**
Focus on understanding, not speed. Solve EASY problems only.

**Week 7-9: Build Pattern Recognition**
Transition to MEDIUM. Focus on the 15 most common patterns.

**Week 10-12: Simulate Interviews**
Timed practice. Mix of easy + medium. Occasional hard.

### Week 5-6: Data Structure Foundations

#### Arrays & Strings (MOST COMMON in frontend interviews)
- [ ] Two pointers technique
- [ ] Sliding window
- [ ] Prefix sum
- [ ] Hash map for O(1) lookups
- [ ] String manipulation (reversal, palindromes, anagrams)

**Must-solve problems:**
```
Easy:
1. Two Sum (#1)
2. Valid Palindrome (#125)
3. Valid Anagram (#242)
4. Best Time to Buy and Sell Stock (#121)
5. Merge Sorted Arrays (#88)
6. Contains Duplicate (#217)
7. Maximum Subarray (#53)
8. Valid Parentheses (#20)

Medium:
9. 3Sum (#15)
10. Longest Substring Without Repeating (#3)
11. Group Anagrams (#49)
12. Product of Array Except Self (#238)
```

#### Hash Maps & Sets
- [ ] When to use Map vs Object in JS
- [ ] Frequency counting pattern
- [ ] Two-pass vs one-pass solutions

**Must-solve:**
```
13. Ransom Note (#383)
14. First Unique Character (#387)
15. Intersection of Two Arrays II (#350)
16. Top K Frequent Elements (#347)
17. LRU Cache (#146) -- VERY commonly asked
```

#### Stacks & Queues
- [ ] Stack: LIFO, monotonic stack pattern
- [ ] Queue: FIFO, BFS foundation
- [ ] Deque for sliding window problems

**Must-solve:**
```
18. Min Stack (#155)
19. Implement Queue using Stacks (#232)
20. Next Greater Element (#496)
21. Daily Temperatures (#739)
22. Evaluate Reverse Polish Notation (#150)
```

### Week 7-8: Trees & Graphs

#### Trees (Binary Trees & BST)
- [ ] Tree traversals: inorder, preorder, postorder (iterative + recursive)
- [ ] Level-order traversal (BFS with queue)
- [ ] BST properties and search
- [ ] DFS vs BFS for trees
- [ ] Tree recursion pattern: process left, process right, combine

**Must-solve:**
```
23. Maximum Depth of Binary Tree (#104)
24. Invert Binary Tree (#226)
25. Same Tree (#100)
26. Symmetric Tree (#101)
27. Binary Tree Level Order Traversal (#102)
28. Validate BST (#98)
29. Lowest Common Ancestor (#236)
30. Binary Tree Right Side View (#199)
31. Serialize and Deserialize Binary Tree (#297)
32. Path Sum (#112)
```

#### Graphs
- [ ] Adjacency list representation
- [ ] BFS and DFS on graphs
- [ ] Cycle detection
- [ ] Topological sort
- [ ] Connected components

**Must-solve:**
```
33. Number of Islands (#200)
34. Clone Graph (#133)
35. Course Schedule (#207)
36. Pacific Atlantic Water Flow (#417)
37. Word Ladder (#127)
```

#### Tries (relevant for autocomplete features)
- [ ] Trie insert, search, startsWith

**Must-solve:**
```
38. Implement Trie (#208)
39. Design Add and Search Words (#211)
```

### Week 9-10: Intermediate Patterns

#### Linked Lists
- [ ] Pointer manipulation, fast/slow pointers
- [ ] Reversal patterns

**Must-solve:**
```
40. Reverse Linked List (#206)
41. Merge Two Sorted Lists (#21)
42. Linked List Cycle (#141)
43. Remove Nth Node From End (#19)
44. Reorder List (#143)
```

#### Binary Search
- [ ] Classic binary search
- [ ] Search in rotated array
- [ ] Finding boundaries

**Must-solve:**
```
45. Binary Search (#704)
46. Search in Rotated Sorted Array (#33)
47. Find Minimum in Rotated Array (#153)
48. Search a 2D Matrix (#74)
```

#### Recursion & Backtracking
- [ ] Subsets, permutations, combinations
- [ ] Decision tree visualization

**Must-solve:**
```
49. Subsets (#78)
50. Permutations (#46)
51. Combination Sum (#39)
52. Letter Combinations of Phone Number (#17)
53. Generate Parentheses (#22)
```

### Week 11-12: Dynamic Programming & Advanced

#### Dynamic Programming (know at least these patterns)
- [ ] 1D DP: climbing stairs, house robber
- [ ] 2D DP: unique paths, longest common subsequence
- [ ] Knapsack pattern
- [ ] Recognize when a problem is DP (overlapping subproblems + optimal substructure)

**Must-solve:**
```
54. Climbing Stairs (#70)
55. House Robber (#198)
56. Coin Change (#322)
57. Longest Increasing Subsequence (#300)
58. Unique Paths (#62)
59. Word Break (#139)
60. Decode Ways (#91)
```

#### Heap / Priority Queue
**Must-solve:**
```
61. Kth Largest Element (#215)
62. Merge K Sorted Lists (#23)
63. Find Median from Data Stream (#295)
```

### DSA Tips for JavaScript/Frontend Interviews
- Always clarify: input types, edge cases, expected time complexity
- Start by talking through brute force, then optimize
- Use JavaScript idioms (Map, Set, array methods) -- interviewers notice
- For frontend-specific DSA, you might get DOM-tree-related problems:
  - Flatten a nested DOM tree
  - Find all text nodes
  - Implement `document.querySelectorAll` (simplified)
  - DOM tree diffing

### LeetCode Practice Plan

| Week | Focus | Problems/Day | Difficulty |
|------|-------|-------------|------------|
| 5 | Arrays, Strings, Hash Maps | 2-3 easy | Easy only |
| 6 | Stacks, Queues, more Arrays | 2-3 easy | Easy only |
| 7 | Trees (BFS/DFS) | 1-2 | Easy + Medium |
| 8 | Graphs, Tries | 1-2 | Easy + Medium |
| 9 | Linked Lists, Binary Search | 1-2 | Medium |
| 10 | Backtracking, Recursion | 1-2 | Medium |
| 11 | Dynamic Programming | 1 | Medium |
| 12 | Mixed practice + review | 2 | Medium + Hard |

---

## Phase 5: Frontend System Design (Weeks 8-13)

### Why This Matters
This round is what separates **senior** candidates. Even at L4 level, showing system design thinking puts you in the top bucket. For Airbnb, Uber, and Stripe this round is heavily weighted.

### Framework for Frontend System Design Answers

**Use this structure for every answer (practice until it's second nature):**

```
1. REQUIREMENTS CLARIFICATION (3-5 min)
   - Functional requirements (what features)
   - Non-functional requirements (performance, scale, offline, accessibility)
   - Ask about target devices, browser support, user base

2. HIGH-LEVEL ARCHITECTURE (5 min)
   - Component hierarchy diagram
   - Data flow (client-server interaction)
   - API design (REST/GraphQL endpoints you'd need)

3. DATA MODEL (5 min)
   - Client-side state shape
   - What goes in local state vs global state vs server cache
   - Normalization strategy

4. COMPONENT DESIGN (10 min)
   - Key component breakdown
   - Reusable vs feature-specific
   - State management approach

5. API DESIGN (5 min)
   - Endpoints needed
   - Pagination strategy
   - Real-time data (WebSocket? SSE? Polling?)

6. OPTIMIZATION & EDGE CASES (5-10 min)
   - Performance: virtualization, lazy loading, caching
   - Offline support
   - Error handling & fallbacks
   - Accessibility
   - SEO (if applicable)

7. TRADE-OFFS (ongoing)
   - Always mention alternatives and why you chose what you chose
```

### Must-Practice System Design Questions

#### Tier 1: Most Commonly Asked
- [ ] **Design a News Feed / Twitter Feed**
  - Infinite scroll, virtualization, optimistic updates
  - Real-time updates (polling vs WebSocket vs SSE)
  - Feed ranking on client side, caching strategies
  - Image lazy loading, skeleton screens

- [ ] **Design an Autocomplete / Typeahead Search**
  - Debouncing input, caching previous results
  - Trie vs API-based suggestions
  - Keyboard navigation, accessibility
  - Handling network race conditions (stale results)

- [ ] **Design a Chat Application**
  - WebSocket connection management
  - Message ordering, delivery receipts
  - Offline message queue
  - Scroll position management, infinite scroll up

- [ ] **Design Google Docs (Collaborative Editor)**
  - Real-time collaboration: OT vs CRDT
  - Cursor position sharing
  - Conflict resolution
  - Undo/redo stack

#### Tier 2: Frequently Asked
- [ ] **Design an Image Carousel / Gallery**
  - Touch gestures, swipe
  - Lazy loading, preloading adjacent images
  - Responsive images, different resolutions
  - Virtualization for 1000s of images

- [ ] **Design a Kanban Board (Trello)**
  - Drag and drop implementation
  - Optimistic reordering
  - Real-time sync across users
  - State management for complex nested data

- [ ] **Design a Spreadsheet (Google Sheets)**
  - Cell rendering and virtualization
  - Formula parsing and dependency graph
  - Undo/redo, copy/paste
  - Performance with 1M+ cells

- [ ] **Design an E-commerce Product Page**
  - Image zoom, color/size variants
  - Inventory real-time checks
  - Reviews with pagination
  - SEO optimization, structured data

#### Tier 3: Advanced
- [ ] **Design a Video Player (YouTube)**
  - Adaptive bitrate streaming
  - Custom controls, keyboard shortcuts
  - Subtitles, picture-in-picture
  - Buffer management

- [ ] **Design a Map Application (Google Maps)**
  - Tile-based rendering
  - Pin clustering for performance
  - Panning, zooming, geolocation
  - Offline map caching

- [ ] **Design a Design Tool (Figma)**
  - Canvas rendering vs DOM
  - Layer system, z-ordering
  - Multi-user cursor & collaboration
  - Infinite canvas, zoom levels

### Key Topics to Weave Into Every Design

| Topic | What to Say |
|-------|------------|
| **Performance** | Lazy loading, code splitting, virtualization, debounce/throttle, Web Workers |
| **Caching** | Service Worker, Cache API, HTTP cache, in-memory LRU cache, stale-while-revalidate |
| **Real-time** | WebSocket vs SSE vs Long Polling -- know tradeoffs |
| **Offline** | Service Workers, IndexedDB, background sync, optimistic UI |
| **Accessibility** | ARIA, keyboard nav, screen readers, color contrast, focus management |
| **Error Handling** | Error boundaries, retry logic, graceful degradation, fallback UI |
| **Security** | XSS, CSRF, CSP, sanitization, auth token storage |
| **Monitoring** | Error tracking (Sentry), performance monitoring (Web Vitals), analytics |
| **SEO** | SSR/SSG, meta tags, structured data, sitemap |
| **Testing** | Component tests, integration tests, E2E, visual regression |

---

## Phase 6: Web Performance & Optimization (Weeks 10-11)

### Why This Matters
Performance questions come up in almost every loop -- either as a dedicated round or woven into system design. Netflix and Airbnb particularly care about this.

### Core Web Vitals (Must Know)
- [ ] **LCP** (Largest Contentful Paint) < 2.5s
  - What affects it: slow server, render-blocking CSS/JS, slow resource load, client-side rendering
  - How to fix: optimize critical path, preload LCP image, use SSR/SSG, CDN
- [ ] **INP** (Interaction to Next Paint) < 200ms
  - What affects it: long tasks, heavy JS execution, layout thrashing
  - How to fix: break up long tasks, use `requestIdleCallback`, debounce, Web Workers
- [ ] **CLS** (Cumulative Layout Shift) < 0.1
  - What affects it: images without dimensions, dynamic content insertion, web fonts
  - How to fix: explicit dimensions, font-display: swap, content placeholders

### Performance Optimization Techniques

#### Loading Performance
- [ ] Code splitting: route-based, component-based
- [ ] Tree shaking: dead code elimination
- [ ] Bundle analysis (webpack-bundle-analyzer)
- [ ] Dynamic `import()` and `React.lazy`
- [ ] Resource hints: `preload`, `prefetch`, `preconnect`, `dns-prefetch`
- [ ] Image optimization: WebP/AVIF, srcset, lazy loading, blur-up
- [ ] Font optimization: subsetting, font-display, preload
- [ ] HTTP/2 multiplexing, HTTP/3

#### Runtime Performance
- [ ] Debounce & throttle for scroll/resize/input handlers
- [ ] Virtualization for long lists (render only visible items)
- [ ] `requestAnimationFrame` for smooth animations
- [ ] Web Workers for CPU-intensive tasks
- [ ] `will-change` and composite layers for GPU acceleration
- [ ] Avoid layout thrashing (batch DOM reads and writes)
- [ ] React-specific: `useMemo`, `useCallback`, `React.memo`, state colocation

#### Caching Strategies
- [ ] Browser cache: `Cache-Control`, `ETag`, `Last-Modified`
- [ ] Service Worker cache: cache-first, network-first, stale-while-revalidate
- [ ] Application cache: React Query, SWR, Apollo Client cache
- [ ] CDN caching and cache invalidation

#### Measuring Performance
- [ ] Lighthouse
- [ ] Chrome DevTools Performance tab
- [ ] Web Vitals JS library
- [ ] Performance Observer API
- [ ] Real User Monitoring (RUM) vs Synthetic monitoring

### Questions You Might Get
```
1. "Our page takes 8 seconds to load. Walk me through how you'd diagnose and fix this."
2. "Users report jank when scrolling through a feed of 10,000 items. What do you do?"
3. "How would you reduce our bundle size from 2MB to under 500KB?"
4. "Explain the critical rendering path and how to optimize it."
5. "What's the difference between SSR, SSG, and CSR? When would you use each?"
```

---

## Phase 7: Testing Strategies (Week 12)

### Why This Matters
Testing questions are increasingly common, especially at Stripe, Airbnb, and Google. You don't need to be an expert, but you need to speak intelligently about it.

### Testing Pyramid for Frontend
```
        /  E2E Tests  \          <-- Few (Cypress/Playwright)
       / Integration    \        <-- Some (Testing Library)
      /  Unit Tests      \       <-- Many (Jest/Vitest)
     /  Static Analysis   \      <-- Always (TypeScript, ESLint)
```

### What to Know

#### Unit Testing
- [ ] Jest / Vitest: describe, it, expect, mocks, spies
- [ ] Mocking: `jest.fn()`, `jest.mock()`, module mocking
- [ ] Testing pure functions, utility functions
- [ ] Snapshot testing (when it's useful and when it's not)

#### Component Testing
- [ ] React Testing Library philosophy: test behavior, not implementation
- [ ] `render`, `screen`, `fireEvent`, `userEvent`, `waitFor`
- [ ] Testing user interactions (click, type, submit)
- [ ] Testing async behavior (API calls, loading states)
- [ ] Testing custom hooks with `renderHook`
- [ ] Avoid testing implementation details (don't test state directly)

#### Integration Testing
- [ ] Testing component composition
- [ ] MSW (Mock Service Worker) for API mocking
- [ ] Testing routing, form flows
- [ ] Testing error boundaries

#### E2E Testing
- [ ] Playwright / Cypress basics
- [ ] When to write E2E vs integration tests
- [ ] Visual regression testing (Percy, Chromatic)

#### What Interviewers Ask
```
1. "How would you test this component?" (given a React component)
2. "What's your testing strategy for a new feature?"
3. "Should you mock API calls in tests? Why or why not?"
4. "What's the difference between unit, integration, and E2E tests?"
5. "How do you decide what to test?"
```

---

## Phase 8: Behavioral & Leadership (Weeks 13-16)

### Why This Matters
Amazon weights behavioral rounds **equal to technical rounds** (Leadership Principles). Meta, Google, and others use behavioral to break ties between technically equal candidates. At 4 YOE, they expect leadership signals.

### STAR Method (Use This for Every Answer)

```
S - Situation: Set the context (1-2 sentences)
T - Task: What was your responsibility (1 sentence)
A - Action: What YOU specifically did (bulk of the answer, 60%)
R - Result: Quantifiable outcome (numbers, metrics, impact)
```

### Prepare Stories for These 15 Themes

**Write 2-3 STAR stories for each theme.** Mix stories from different projects/roles.

#### Technical Leadership
- [ ] 1. A time you made a critical technical decision for the team
- [ ] 2. A time you introduced a new technology/tool/process
- [ ] 3. A time you improved developer experience or team productivity
- [ ] 4. A time you handled technical debt

#### Problem Solving
- [ ] 5. A time you debugged a critical production issue
- [ ] 6. A time you solved a complex technical problem
- [ ] 7. A time you had to make a decision with incomplete information

#### Collaboration & Conflict
- [ ] 8. A time you disagreed with a teammate/manager (and how it resolved)
- [ ] 9. A time you mentored someone or helped them grow
- [ ] 10. A time you worked with a difficult stakeholder

#### Delivery & Impact
- [ ] 11. Your most impactful project (explain end-to-end)
- [ ] 12. A time you had to deliver under tight deadlines
- [ ] 13. A time a project failed or had setbacks (what you learned)

#### Growth & Self-Awareness
- [ ] 14. Your biggest professional weakness and how you're addressing it
- [ ] 15. A time you received tough feedback and what you did

### Amazon Leadership Principles (if targeting Amazon)

You MUST have stories for all 16 LPs. The most tested ones for SDE-2:
- [ ] **Customer Obsession**: decisions driven by customer impact
- [ ] **Ownership**: went above and beyond your scope
- [ ] **Dive Deep**: dug into metrics/data to find root cause
- [ ] **Bias for Action**: acted without perfect information
- [ ] **Deliver Results**: shipped under constraints
- [ ] **Earn Trust**: admitted mistakes, gave credit
- [ ] **Have Backbone; Disagree and Commit**: pushed back respectfully, then committed

### Questions to Prepare For

```
1. Tell me about yourself (2-min elevator pitch: journey, key skills, what you're looking for)
2. Why do you want to work at [Company]?
3. Tell me about the most complex frontend feature you've built
4. How do you approach a new codebase?
5. How do you handle disagreements about technical approaches?
6. Tell me about a time you had to learn something quickly
7. How do you prioritize tasks when everything seems urgent?
8. What's your approach to code reviews?
9. How do you ensure code quality on your team?
10. Where do you see yourself in 3-5 years?
```

---

## Phase 9: Machine Coding / UI Rounds (Weeks 6-16)

### Why This Matters
This is the **MOST frontend-specific round** and often the **deciding factor**. You'll be asked to build a functional UI component or small application in 45-60 minutes. Meta, Airbnb, and Stripe heavily use this round.

### What They Evaluate
- Clean, modular code architecture
- React/JS proficiency under time pressure
- CSS skills (usually no component library allowed)
- Edge case handling
- Accessibility
- Code organization (separation of concerns)

### Must-Practice Machine Coding Challenges

#### Tier 1: Very High Frequency (Practice until you can do each in < 30 min)
- [ ] **Autocomplete / Typeahead**
  - Debounced input, API call, keyboard navigation (up/down/enter)
  - Highlighting matching text, loading/error states
  - Accessibility: aria-combobox, aria-activedescendant

- [ ] **Star Rating Component**
  - Click to set rating, hover preview
  - Half stars, read-only mode
  - Accessibility: radio group pattern

- [ ] **Modal / Dialog**
  - Portal rendering, backdrop click to close
  - Focus trapping, Escape to close
  - Animation on enter/exit
  - Body scroll lock

- [ ] **Todo App (Advanced)**
  - CRUD operations, filters (all/active/completed)
  - Local storage persistence
  - Drag and drop reordering
  - Undo functionality

- [ ] **Infinite Scroll / Pagination**
  - IntersectionObserver for scroll detection
  - Loading indicators, error retry
  - Scroll position restoration

- [ ] **Accordion / Collapsible Sections**
  - Single/multiple expand modes
  - Animated expand/collapse
  - Keyboard accessible

#### Tier 2: High Frequency
- [ ] **Data Table with Sorting, Filtering, Pagination**
  - Column sort (asc/desc/none)
  - Text search filter
  - Client-side pagination
  - Multi-select rows

- [ ] **Image Carousel / Slider**
  - Auto-play, manual navigation
  - Touch/swipe support
  - Dot indicators, thumbnail navigation
  - Lazy load off-screen images

- [ ] **Multi-step Form / Wizard**
  - Step navigation, validation per step
  - Form state persistence across steps
  - Progress indicator
  - Review step before submit

- [ ] **Dropdown / Select Component**
  - Search/filter options
  - Multi-select with chips
  - Keyboard navigation
  - Virtual scroll for large lists

- [ ] **File Uploader**
  - Drag and drop area
  - Progress bar for each file
  - File type/size validation
  - Preview for images

- [ ] **Tabs Component**
  - Lazy/eager tab content loading
  - Keyboard navigation (arrow keys)
  - ARIA: tablist, tab, tabpanel roles

#### Tier 3: Challenging
- [ ] **Kanban Board** (drag and drop across columns)
- [ ] **Calendar / Date Picker**
- [ ] **Rich Text Editor** (basic: bold, italic, lists)
- [ ] **Spreadsheet** (basic grid with editable cells)
- [ ] **Drawing App** (canvas-based)
- [ ] **Nested Comments** (Reddit-style, with collapse)
- [ ] **Poll / Survey Builder**
- [ ] **Tic-Tac-Toe / Memory Game**

### Machine Coding Round Template

Use this approach in every UI round:

```
Minutes 0-3:   Clarify requirements, ask about edge cases
Minutes 3-8:   Plan component structure, sketch on paper/whiteboard
Minutes 8-40:  Build the core functionality FIRST (get it working)
Minutes 40-50: Add polish (error handling, edge cases, accessibility)
Minutes 50-55: Refactor if needed, add comments on tradeoffs
Minutes 55-60: Explain what you'd improve given more time
```

### What to Practice Building With
- **Vanilla React** (no component libraries -- they're usually banned)
- **CSS** (plain CSS, CSS modules, or styled-components -- ask what's allowed)
- **No UI libraries** (no Material UI, Ant Design, Chakra, etc.)

---

## Company-Specific Breakdowns

### Google (Frontend SWE L4)
```
Interview Loop:
1. Phone Screen: 1 coding round (DSA, 45 min)
2. Onsite (4-5 rounds):
   - 2x Coding/DSA (LeetCode Medium-Hard)
   - 1x Frontend-specific (DOM, JS, UI implementation)
   - 1x System Design (may be frontend-focused)
   - 1x Behavioral (Googleyness & Leadership)

Focus: DSA is king. Must solve mediums comfortably, some hards.
Unique: They use Google Docs for coding (no autocomplete/running code).
Tip: Think aloud, explain your approach before coding.
```

### Meta (Frontend E4/E5)
```
Interview Loop:
1. Phone Screen: 1 coding round (DSA, 45 min)
2. Onsite (4-5 rounds):
   - 2x Coding/DSA (LeetCode Medium)
   - 1x Product Sense + System Design (E5 heavy, E4 lighter)
   - 1x Behavioral
   - 1x UI/Machine Coding (build a feature in React)

Focus: Balanced loop. UI round is critical for frontend role.
Unique: They have a dedicated "Product Engineering" interview.
Tip: For E4, showing E5 signals in system design will get you up-leveled.
```

### Amazon (Frontend SDE-2)
```
Interview Loop:
1. Phone Screen: 1 DSA + LP behavioral
2. Onsite (4-5 rounds):
   - 2x Coding/DSA (LeetCode Medium)
   - 1x System Design (scale-focused)
   - 1x Frontend Deep Dive (JS, React, architecture)
   - ALL rounds have LP behavioral questions (15 min each)

Focus: Leadership Principles dominate. Prepare 2 stories per LP.
Unique: Every single round ends with LP questions.
Tip: Use the STAR format. Quantify results ("reduced load time by 40%").
```

### Stripe (Frontend)
```
Interview Loop:
1. Phone Screen: JS/debugging round
2. Onsite (4 rounds):
   - 1x Bug Squash (debug a broken frontend app)
   - 1x UI Implementation (build a component)
   - 1x Integration (build feature with API calls)
   - 1x System Design / Architecture

Focus: Practical coding over algorithm puzzles.
Unique: Their "Bug Squash" round is unique -- practice debugging.
Tip: Write clean, production-quality code. They care about code quality.
```

### Airbnb (Frontend)
```
Interview Loop:
1. Phone Screen: Coding challenge
2. Onsite (4-5 rounds):
   - 1x Coding/DSA
   - 1x UI Implementation (very detailed, CSS matters)
   - 1x System Design
   - 1x Cross-functional collaboration
   - 1x Core Values (behavioral)

Focus: UI implementation round is extremely detailed.
Unique: They care deeply about design taste and pixel-perfection.
Tip: Practice CSS layouts without any library. Pixel-perfect matters.
```

---

## Weekly Schedule Template

Since you have 1-2 hours/day, here's how to split each week:

### Weekdays (1 hour/day)

| Day | Focus | Activity |
|-----|-------|----------|
| **Monday** | DSA | Solve 1-2 problems on the week's topic |
| **Tuesday** | DSA | Solve 1-2 problems, review yesterday's |
| **Wednesday** | JS/React/CSS | Study concepts, implement from scratch |
| **Thursday** | JS/React/CSS | Continue concepts, practice coding |
| **Friday** | Machine Coding | Build one UI component timed (45 min) |

### Weekends (2 hours/day if possible)

| Day | Focus | Activity |
|-----|-------|----------|
| **Saturday** | System Design | Study one design topic (2 hours) |
| **Sunday** | Review + Behavioral | Review week's learning (1 hr) + Write STAR stories (1 hr) |

---

## Resources Master List

### JavaScript
- [javascript.info](https://javascript.info) -- best free JS resource
- [You Don't Know JS](https://github.com/getify/You-Dont-Know-JS) (book series, free on GitHub)
- [33 JS Concepts](https://github.com/leonardomso/33-js-concepts)
- Lydia Hallie's JavaScript Visualized series

### TypeScript
- [TypeScript Handbook](https://www.typescriptlang.org/docs/handbook/)
- [type-challenges](https://github.com/type-challenges/type-challenges) (GitHub)
- Matt Pocock's Total TypeScript (YouTube + course)

### React
- [New React Docs](https://react.dev) (official, excellent)
- [React TypeScript Cheatsheet](https://react-typescript-cheatsheet.netlify.app/)
- Kent C. Dodds' blog (advanced patterns)

### DSA
- [NeetCode.io](https://neetcode.io) -- best structured roadmap (NeetCode 150)
- [Grind 75](https://www.techinterviewhandbook.org/grind75) -- customizable problem set
- LeetCode (filter by company tags if you have premium)
- [Structy](https://www.structy.net) -- great for visual learners

### System Design
- [GreatFrontEnd System Design](https://www.greatfrontend.com/system-design)
- [Frontend System Design by Evgeniy](https://www.youtube.com/c/user) (YouTube)
- [Designing Large-Scale Frontend Apps](https://frontendmastery.com/)
- [System Design Primer](https://github.com/donnemartin/system-design-primer)

### Machine Coding
- [GreatFrontEnd](https://www.greatfrontend.com) -- frontend-specific problems
- [FrontendMentor](https://www.frontendmentor.io) -- UI challenges
- [BigFrontEnd.dev](https://bigfrontend.dev) -- JS/React interview problems

### Behavioral
- [Tech Interview Handbook - Behavioral](https://www.techinterviewhandbook.org/behavioral-interview/)
- Amazon LP examples on YouTube (search "Amazon LP interview examples")
- Keep a running document of STAR stories

### Mock Interviews
- [Pramp](https://www.pramp.com) -- free peer mock interviews
- [Interviewing.io](https://interviewing.io) -- anonymous mock interviews with FAANG engineers
- Practice with friends / colleagues

---

## Progress Tracker

### Phase 1: JavaScript & TypeScript (Weeks 1-3)
- [ ] Execution context, scope, hoisting
- [ ] Closures -- can explain with memory diagram
- [ ] Implemented: debounce, throttle, curry, memoize, once, pipe, compose
- [ ] `this` -- can predict output for any scenario
- [ ] Implemented: bind, call, apply, instanceof, Object.create, new
- [ ] Event loop -- can predict output for 10+ scenarios
- [ ] Implemented: Promise from scratch (with all static methods)
- [ ] Implemented: parallelLimit, retry, promisify
- [ ] Implemented: deepClone, deepEqual, EventEmitter
- [ ] DOM APIs: event delegation, observers, fetch patterns
- [ ] TypeScript: can implement utility types from scratch
- [ ] TypeScript + React: can type any component/hook pattern

### Phase 2: HTML, CSS & Browser (Weeks 3-4)
- [ ] Flexbox: can solve any layout problem
- [ ] Grid: can create complex responsive layouts
- [ ] Positioning & stacking context
- [ ] Specificity and cascade rules
- [ ] CSS animations and transitions
- [ ] Semantic HTML and ARIA
- [ ] Critical rendering path
- [ ] CORS fully understood

### Phase 3: React (Weeks 4-6)
- [ ] Can explain reconciliation and fiber
- [ ] Know all hooks and their edge cases
- [ ] Implemented 12 custom hooks from scratch
- [ ] State management patterns clear
- [ ] React 18+ features understood
- [ ] Advanced patterns: compound components, render props, etc.

### Phase 4: DSA (Weeks 5-12)
- [ ] Solved 30+ Easy problems
- [ ] Solved 25+ Medium problems
- [ ] Solved 5+ Hard problems
- [ ] Arrays/Strings: comfortable with two pointers, sliding window
- [ ] Trees: can do BFS/DFS without thinking
- [ ] Graphs: BFS, DFS, topological sort
- [ ] DP: can recognize and solve basic DP problems
- [ ] Can solve a medium in 25-30 minutes consistently

### Phase 5: System Design (Weeks 8-13)
- [ ] Can use the framework for any question
- [ ] Practiced: News Feed, Autocomplete, Chat App
- [ ] Practiced: Collaborative Editor, E-commerce Page
- [ ] Can discuss tradeoffs for caching, real-time, offline

### Phase 6: Performance (Weeks 10-11)
- [ ] Core Web Vitals: can explain and improve each
- [ ] Loading performance optimization techniques
- [ ] Runtime performance optimization techniques
- [ ] Can answer "our page is slow, diagnose it"

### Phase 7: Testing (Week 12)
- [ ] Can write React Testing Library tests
- [ ] Know testing strategy for any component
- [ ] Understand MSW for API mocking

### Phase 8: Behavioral (Weeks 13-16)
- [ ] 15 STAR stories written and practiced
- [ ] Amazon LPs covered (if targeting Amazon)
- [ ] 2-minute "tell me about yourself" pitch polished
- [ ] Did 3+ mock behavioral interviews

### Phase 9: Machine Coding (Weeks 6-16)
- [ ] Can build Autocomplete in 30 min
- [ ] Can build Modal with focus trap in 25 min
- [ ] Can build Data Table with sort/filter in 40 min
- [ ] Can build Infinite Scroll in 25 min
- [ ] Built 10+ components from scratch under time pressure
- [ ] Did 3+ mock UI coding interviews

---

## Final Tips

1. **Don't skip DSA.** Yes, it feels disconnected from frontend work. But it's still how FAANG gates candidates. Dedicate consistent time.

2. **Machine coding is your secret weapon.** This is where you can truly shine as a frontend developer. Practice building components without any UI library until it's muscle memory.

3. **System design wins senior offers.** Even at 4 YOE (mid-level), showing system design maturity can get you up-leveled to senior.

4. **Behavioral prep is underrated.** Most candidates wing it. Preparing STAR stories puts you in the top 10%.

5. **Do mock interviews.** The gap between "knowing" and "performing under pressure" is massive. Do at least 5-10 mocks before real interviews.

6. **Track your progress.** Use the tracker above. Check items off. Seeing progress builds confidence.

7. **Start applying by week 10-12.** You don't need to be 100% ready. Apply to lower-priority companies first as practice. Save your top choices for when you're peak ready.

8. **Rest is part of prep.** Burnout kills interview performance. Take at least one day completely off per week.

---

*Last updated: March 14, 2026*
*Total estimated prep time: 150-200 hours over 16 weeks*
