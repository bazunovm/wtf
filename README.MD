**WTF** is a convenient CLI tool for explaining Linux and Mac command errors.  
It analyzes error outputs and provides clear explanations and recommendations.


---

## 📦 Installation

### Using the `install.sh` script

Clone the repository:

```bash
git clone https://github.com/bazunovm/wtf.git
cd wtf
```

Give the script execution permissions:

```bash
chmod +x install.sh
```

Run the script with superuser privileges:

```bash
sudo ./install.sh
```

To use AI get your Gemini API key on https://aistudio.google.com/ and run ai_add.sh script with your API key as an argument:

```bash
sudo ./ai_add.sh <your_api_key>
```

## 🚀 Usage

### After installation, you can run commands using wtf:

```bash
wtf <command>
```

### Examples:

Command:

```bash
wtf make
```

Output:

```bash
Error: make is not installed
Meaning: The make utility is missing
Suggestions:
  - sudo apt install make
```

### 📝 Adding new rules

All error rules are stored in internal/explainer/rules.yaml.
To add a new rule:

```yaml
- id: example_error
  match: "example: command not found"
  title: "Example Error"
  meaning: "Description of the error"
  suggestions:
    - "Recommendation 1"
    - "Recommendation 2"
```
