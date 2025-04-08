# StegoTool

StegoTool is a CLI tool for hiding and reading messages in images using steganography. It supports symmetric and asymmetric encryption and can be extended to support various image formats in the future.

## Features

-  Hide messages in PNG images using symmetric and asymmetric encryption.
-  Read hidden messages from PNG images.
-  Support for various image formats (PNG, JPEG, etc.) in future updates.

## Installation

To install StegoTool, you need to have Go installed on your system. You can download Go from [golang.org](https://golang.org/).

1. Clone the repository:

   ```sh
   git clone https://github.com/boxy-pug/stegotool.git
   cd stegotool
   ```

2. Build the project:

   ```sh
   go build -o stegotool ./cli
   ```

## Usage

### Hide a Message

To hide a message in an image, use the `hide` command:

```sh
./stegotool hide <image_path> <message> <output_path>
```

Example:

```sh
./stegotool hide input.png "Secret message" output.png
```

### Read a Hidden Message

To read a hidden message from an image, use the `read` command:

```sh
./stegotool read <image_path>
```

Example:

```sh
./stegotool read output.png
```

## Examples

### Hide a Message

```sh
./stegotool hide input.png "This is a secret message" output.png
```

### Read a Hidden Message

```sh
./stegotool read output.png
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License.

## Contact

For any questions or suggestions, please open an issue.
