# TypeScript Type to Object Literal Converter

This TypeScript script inputs a string containing a TypeScript type and outputs an object literal representing the type.

## Requirements

- Node.js and npm installed on your machine.

## Installation

1. Clone this repository to your local machine:
git clone https://github.com/jesus87/frameplay.git


2. Navigate to the project directory:

cd frameplay
3. Install dependencies:
npm install


## Usage

1. Open the `convert.ts` file in a text editor.

2. Modify the `type` variable at the bottom of the file to include your TypeScript type definition.

3. Save the file.

4. Run the script using Node.js:

node convert.js


5. The output object literal will be printed to the console.

## Example

For example, if you have the following TypeScript type definition:

```typescript
type Button = {
    variant: "solid" | "text";
};
Running the script will output:

{
    "variant": ["solid", "text"]
}
