# ECS Task Definition Generator

This is a command-line tool to generate ECS task definition files with environment-specific values embedded in the template. Using this tool, users can easily create ECS task definitions based on a template and environment file (.env), replacing placeholder values in the template with real environment variables.

## Features

Generates ECS task definition files based on a JSON template with placeholder values.
Replaces placeholders in the template file (e.g., ${VAR_NAME}) with actual values from an environment file (.env).
Allows users to specify template and environment file paths, as well as the output directory.

## Installation

1. Clone this repository:

```bash
git clone <https://github.com/yourusername/ECSbuilder.git>
```

1. Navigate to the project directory:

```bash
cd ECSbuilder
```

1. Install dependencies:

```bash
go mod tidy
```

## Usage

```bash
go install github.com/rendaman0215/ECSbuilder
ECSbuilder --template PATH-TO-TEMPLATE --env PATH-TO-ENV --output PATH-TO-OUTPUT_DIR
```

### Command-line Options

The following options are available when running the program:

- `--input`, {`-i`} {default: `.`}: Path to the definition and env file input directory.
- `--template`, {`-t`} (default: `sample/taskdef.template.json`): Path to the task definition template file.
- `--env`, `-e` (default: `.env`): Path to the environment file containing variable values.
- `--output`, `-o` (default: `taskdef.json`): Path to the output file for the generated task definition.

If no arguments are provided, the program will prompt you for the template file, environment file, and output path.

### Example

To generate an ECS task definition from a template and environment file in the sample/ directory:

```bash
go run main.go --template sample/taskdef.template.json --env .env --output output/taskdef.json
```

### .env File Structure

The environment file (`.env`) should contain key-value pairs representing environment-specific variables. For example:

```dotenv
ENVIRONMENT=production
CONTAINER_IMAGE=myapp-image:latest
```

### Template File Structure

The template file should use `${VAR_NAME}` format for placeholders. For example:

```json
{
  "family": "myapp-${ENVIRONMENT}",
  "containerDefinitions": [
    {
      "name": "app-container",
      "image": "${CONTAINER_IMAGE}"
    }
  ]
}
```

## Project Structure

The project is organized based on DDD and Clean Architecture principles:

- `cmd`: Command-line interface handling arguments and user prompts.
- `internal/app`: Main application logic for generating the task definition file.
- `internal/config`: Handles loading environment variables from the .env file.
- `pkg/template_parser`: Contains the template parser logic for replacing placeholders with environment values.

## Testing

Tests are written using `stretchr/testify`. To run tests:

```bash
go test -v -cover ./...
```

## License

This project is licensed under the MIT License. See the LICENSE file for details.
