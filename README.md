# INFINI Golang Framework Vendors

INFINI Golang framework manages all dependencies in the `vendor` directory. This approach ensures portability and facilitates offline development.

## Why Vendoring Dependencies?

Vendoring dependencies offers several benefits:

1. **Consistency and Stability**: Vendoring locks in specific versions of external libraries, ensuring that your project remains functional even if upstream dependencies change or are removed.
2. **Offline Builds**: With vendored dependencies, you can build and deploy your code without an internet connection, which is especially useful in CI/CD environments or restricted networks.
3. **Security and Compliance**: Maintaining a local copy of dependencies allows for easier auditing, reviewing, and control, which is important for security and compliance purposes.
4. **Reduced External Dependency Risks**: Vendoring protects your project from external risks like deleted or unavailable upstream packages by storing a local copy of all dependencies.

## How to Add a New Vendor

Run the following command to fetch the vendor:

```bash
go get -u github.com/vendor/repo
```

The vendor files will be located in the `~/go/src` directory.

Move the necessary vendor files to the `vendor` directory into repo.

Clean up the `.git` repositories from all submodules, run this command (exclude `.git` in the root folder):
```
find . -path './.git' -prune -o -type d -name ".git" -exec rm -rf {} +
```

After cleanup the new vendor files, commit them to your repository.

Once the vendor files are copied, you can clean up the files located in `~/go/src`, except for the `infini.sh` folder.
