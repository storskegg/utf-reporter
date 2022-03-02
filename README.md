# utf-reporter

CLI app reporting instances of "non-standard" characters by line and column.

![screenshot](media/screenshot.gif)

For the purposes of this work, a "standard" character is defined as having a base-10 value `n` where...

- `n == 10`,
- or `n == 13`, 
- or `n > 31 && n < 128`.

All other characters are identified as non-standard, and tagged as either ASCII (0 - 255) or UTF, and listed with the line and column values so they can be found in the original document.

## Why

A recurring issue I've experienced throughout my engineering career is the presence of "non-standard" characters in places where there are expected to be none. This often happens when...

- Text is copy/pasted from a popular word processor or team chat application (often coming from a stakeholder). These applications commonly replace double quotes with fancy quotes, commas with fancy commas, etc.
- Bad input allowed from somewhere upstream
- You're the victim of an engineering practical joke when you left your computer unlocked (maybe a semi-colon was replaced by a Greek questionmark, or zero-width space inserted into a variable declaration)

Regardless, you're now dealing with some unexpected behavior in your system.

Obviously, this could be solved by "simply supporting extended charsets," but in the real world, on real systems, that's not always practical, let along feasible. So, for those cases, I originally wrote a small, janky CodePen app that would detect and highlight these characters.

I needed something to parse through massive CSV's, so I ported the pen to this little CLI app.

## Running

`utf-reporter` supports supports both piped stdin and a file flag.

### Piped Stdin

```bash
cat test.txt | utf-reporter
```

### File Flag

```bash
utf-reporter -f path/to/my/test.txt
```
