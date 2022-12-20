# Description of the PR

... (e.g. what value does it bring? Why should it be merged?)

# Review checklist

(based on the work of Karl E. Wiegers)

**Structure**

- [ ] Does the code completely and correctly implement the design?
- [ ] Is the code well-structured, consistent in style, and consistently formatted?
- [ ] Are there any uncalled or unneeded procedures or any unreachable code?
- [ ] Are there any leftover stubs or test routines in the code?
- [ ] Can any code be replaced by calls to external reusable components or library functions?
- [ ] Are there any blocks of repeated code that could be condensed into a single procedure?
- [ ] Are symbolics used rather than “magic number” constants or string constants?
- [ ] Are any modules excessively complex and should be restructured or split into multiple routines?

**Documentation**

- [ ] Is the code clearly and adequately documented with an easy-to-maintain commenting style?
- [ ] Are all comments consistent with the code?

**Variables**

- [ ] Are all variables properly defined with meaningful, consistent, and clear names?
- [ ] Are there any redundant or unused variables?

**Arithmetic Operations**

- [ ] Does the code avoid comparing floating-point numbers for equality?
- [ ] Are divisors tested for zero or noise?

**Loops and Branches**

- [ ] Are all loops, branches, and logic constructs complete and correct?
- [ ] Are all loops, branches, and logic constructs properly nested and avoid unnecessary nesting?
- [ ] Can any statements that are enclosed within loops be placed outside the loops?

**Defensive Programming**

- [ ] Are imported data and input arguments tested for validity and completeness?
- [ ] Are files checked for existence before attempting to access them?
