import * as ts from "typescript";

// Function to visit TypeScript nodes and extract information
const visit = (node: ts.Node) => {
    // Check if it's an object type
    if (ts.isTypeLiteralNode(node)) {
        const properties: { [key: string]: string[] } = {};
        // Iterate through each member of the type
        node.members.forEach(member => {
            // Check if it's a property signature
            if (ts.isPropertySignature(member)) {
                // Get the property name
                const propertyName = member.name.getText();
                // Get the property type
                const propertyType = member.type?.getText();
                // Add property name and type to the properties object
                if (propertyName && propertyType) {
                    properties[propertyName] = getTypeValues(propertyType);
                }
            }
        });
        return properties;
    }
    return {};
};

// Helper function to get values of a union type
const getTypeValues = (type: string): string[] => {
    if (type.includes("|")) {
        return type.split("|").map(value => value.trim());
    }
    return [type.trim()];
};

// Function to convert a TypeScript type to an object literal
const convertToObject = (type: string) => {
    // Create a TypeScript program from the provided code
    const sourceFile = ts.createSourceFile("temp.ts", type, ts.ScriptTarget.Latest);
    const typeNode = sourceFile.statements[0] as ts.TypeAliasDeclaration;

    // Visit the type node and get the object literal
    const result = visit(typeNode.type!);
    return result;
};

// Example usage
const result = convertToObject(`type Button = {
    variant: "solid" | "text";
};`);
console.log(result);