### `cowlourful`

A command-line tool that pipes text and displays it in a colorful, ASCII cow-themed speech bubble.

![Screenshot of cowlourful](https://github.com/sthejas87/cowlourful/blob/cc78f73f3f052b77ae30f51f438cefa2d0db85eb/cowlourful.png)


---

### Features

* **Piped Input:** Designed to work with standard Unix pipes, making it easy to integrate into your command-line workflow.
* **Dynamic Text Wrapping:** Automatically wraps the piped message to fit the width of your terminal.
* **Dynamic Colors:** Generates a horizontal, sine-wave-based color gradient for the text inside the speech bubble.
* **Customizable:** The color spread and intensity can be easily modified by changing the `freq` and `phase` variables in the `getColour` function.

---

### Installation

To install `cowlourful`, you will need to have Go installed on your system.

The easiest way to install and manage Go binaries is with the `go install` command.

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/sthejas87/cowlourful.git
    cd cowlourful
    ```
2.  **Install the executable:**
    ```bash
    go install
    ```
    This command will compile the program and place the executable in your `$GOPATH/bin` directory (typically `~/go/bin`).
3.  **Ensure `$GOPATH/bin` is in your `$PATH`:**
    If it's not already set up, add the following line to your shell's configuration file to ensure you can run the command from any directory.
    ```bash
    export PATH="$PATH:$HOME/go/bin"
    ```
    After saving, apply the changes by running `source ~/.bashrc` (or your relevant file).

---

### Usage

The program is designed to be used with a pipe. Simply pipe any text output to the `cowlourful` executable.

**Example:**
```bash
echo "Hello World" | cowlourful
```

