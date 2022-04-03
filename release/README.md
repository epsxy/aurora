# Installation

The installation process will copy the executable to `/usr/local/bin/aurora`. It will also create the alias `git gz`, using the file `/usr/local/bin/git-gz` and make it executable. You'll maybe need to source your bash after install: `source ~/.bashrc`, `source ~/.zshrc`, etc.

```bash
# Install
sudo make install

# Uninstall
sudo make uninstall
```
## Todo

The installation process for other platforms than OSX must be tweaked from the process above. The linux process should roughly be the same. Windows and freebsd are probably different because paths and commands may be different.