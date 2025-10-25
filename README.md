## Build Yourself
### Requirements
- [go](https://go.dev/doc/install) 1.24.3 or higher
### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/syedhuzaif199/shell-calendar.git
   ```
2. Navigate to the shell-calendar directory:
   ```bash
   cd shell-calendar
   ```
3. Build
   - For system-wide installation, run:
      ```bash
      go install
      ```
   - To build an executable, run:
      ```bash
      go build
      ```


### Usage
If you used `go install`, the program will be installed system-wide. You can run it from any terminal:
   - Requires that the GOPATH/bin directory be included in the PATH environment variable
      ```bash
      shell-calendar
      ```
   - To run a built executable, run:
     ```bash
     ./shell-calendar
     ```

### What it looks like
<img width="682" height="589" alt="image" src="https://github.com/user-attachments/assets/a6cfa080-effb-4e92-8543-f2d1ca701064" />

### Todo
- Add flags for customization (e.g., start day of the week, highlight current day)
- Add releases
