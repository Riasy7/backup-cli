# Backup CLI Tool using Go


## Description

This is a simple CLI tool is a command-line app designed to backup files and directories to a specified destination using Go.
The tool allows users to create backups of files and directories. It's a straightfoward way to ensure important data is safely stored.

## 📦 Technologies

- Go
- Go Modules
- os and filepath packages

## ✅ Features

The main features of this application consist of the following:

- Backup files and directories to a specified destination
- Create a backup of a single file
- Create a backup of a directory
- Create a backup of multiple files and directories
- Logging of backup operations for the purpose of tracking

## 🚀 Installation

To install the Backup CLI tool, clone the repository and build the tool:

```
git clone https://github.com/Riasy7/backup-cli.git
cd backup-cli
go build -o backup-cli ./cmd/main.go
```

## 🤔🔑 Future Improvements

Many things can be improved to scale this application into a more robust and feature-rich tool.
Imagine if this was a real CLI tool you could use what would you want to be able to do with it?
Here are some improvements I'd like to make:

- First I would improve the user experience by adding things such as progress bars, more detailed logging (log levels, rotation, format), etc.
    - Using libraries such as promptui or blessed for a nice looking CLI interface.
- GitHub integration for backups: Either backup to GitHub or backup GitHub repositories.
- Different file types: images, documents, databases, or even entire drives.
- File Encryption: Encrypt files before backing them up.
- Another optional feature I would add which would be interesting is Cloud Storage integration.
    - Backup to Google Drive, Dropbox, AWS, etc,
- Different OS compatibility would also be great.
- The last thing would be Version Control and also better Algorithms for the backup process.
    - For example, only backup files that have changed since the last backup. (incremental, differential, etc)

## 📝 Comments

This project allowed me to deepen my understanding of Go and how to create a CLI tool using it.
Interestingly enough, I found that using Go for a CLI tool is quite simple and straightforward.
Of course my application is simple, but I designed it in a way where it has room for scalability and future features.
Go's concurrency model is also a great feature that I would like to explore more in the future.
Go's standard library is also very powerful and I was able to use it to create a simple yet effective CLI tool.
Error handling and test cases were also very easy to implement in Go.
Overall, making the improvements I mentioned would be a great way to further my understanding of Go and CLI tools in general.

## Example usage:

```
go run cmd/main.go --source="/mnt/c/Users/riasy/Desktop/sample" --dest="/mnt/c/Users/riasy/Desktop/backups"
```

