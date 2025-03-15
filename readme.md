# 📦 Barrel Generator

The **Barrel Generator** is a CLI tool written in Go that automates the creation of export files (`barrel files`) for Dart/Flutter projects. It scans a directory and generates a `<folder_name>_exports.dart` file containing all necessary exports, organized automatically.

## 🚀 Installation

### 1️⃣ Download and Compile the Code
If you already have Go installed, you can compile the binary directly:
```sh
# Clone the repository
git clone <REPOSITORY_URL>
cd <PROJECT_FOLDER>

# Compile the binary
go build -o barrel_gen

# Move it to a globally accessible location
mv barrel_gen ~/go/bin/   # Or sudo mv barrel_gen /usr/local/bin/
```

Now, you can run the `barrel_gen` command from anywhere in the terminal.

---
## 🎯 How to Use

The **Barrel Generator** allows you to create export files for any folder in your Dart/Flutter project.

### 🔹 Automatically Generate an Export File
```sh
barrel_gen -dir=lib/features/user/models
```
If the `models` folder contains `.dart` files, it will generate:
```
lib/features/user/models/models_exports.dart
```

### 🔹 Specify a Custom Output Name
```sh
barrel_gen -dir=lib/features/user/models -o=lib/features/user/models.dart
```
This will create a file with a different name:
```
lib/features/user/models.dart
```

---
## 🛠 Example Usage
If we run:
```sh
barrel_gen -dir=lib/models
```
And the `lib/models` folder contains the following files:
```
lib/models/
│── user.dart
│── product.dart
│── order.dart
│── user.freezed.dart  (ignored)
│── user.g.dart        (ignored)
```
The generated `models_exports.dart` file will be:
```dart
// 📦 Automatically generated file. DO NOT EDIT MANUALLY.

export 'user.dart';
export 'product.dart';
export 'order.dart';
```

---
## 🎯 Features
✔ **Automatically generates a `<folder_name>_exports.dart` file**  
✔ **Allows setting a custom output name with `-o`**  
✔ **Ignores generated files (`.freezed.dart`, `.g.dart`)**  
✔ **Automatically sorts exports**  
✔ **Color-coded logs for better feedback**  

Organizing your files in Flutter has never been easier! 🚀

