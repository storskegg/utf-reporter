# utf-reporter
CLI app reporting instances of "non-standard" characters by line and column.

For the purposes of this work, a "standard" character is defined as having a base-10 value `n` where...

- `n == 10`,
- or `n == 13`, 
- or `n > 31 && n < 128`.

All other characters are identified as non-standard, and tagged as either ASCII (0 - 255) or UTF, and listed with the line and column values so they can be found in the original document.
