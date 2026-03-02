# frep

Fuzzy grep. 
Name is combination fo fuzzy and grep which does not make sense because grep is global, regex, print and adding f makes is fuzzy regex print, but this does not use regex but levenshtein distance. 
I wanted to make is as close to `grep` as possible for muscle memory, but there is a lot of `fre` clis so `fre-TAB>` does not work.

Since it is recursive by nature and kind clunky to use it can be: **F**uzzy **R**ecursive **E**valuation of **P**atterns

# Installation

1. Clone the repository
   ```bash
   git clone https://github.com/TomislavRigo/frep.git
   ```

2. Build

   2.1. Make
   ```bash
   make build
   ```

   2.2 Manual
   ```bash
   go build ./frep/cli/main.go
   ```

4. Add path or move to `/bin/`
