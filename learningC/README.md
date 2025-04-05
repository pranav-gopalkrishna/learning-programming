# Learning C programming
My notes, resources and programs from my learning of C programming.

---

**Setting up the IDE to run C programs...**

---

# Downloading and installing Visual Code (a solid choice for an IDE)

Visual studio code can be installed from [here](https://code.visualstudio.com/Download)

After instaling this, [C/C++ extention](https://marketplace.visualstudio.com/items?itemName=ms-vscode.cpptools) must be installed.

Next, XCode command line tools need to be installed. To do this, go to terminal and run the command below.

```sh
 xcode-select --install
 ```
 
## Common errors
 
 Without these tools, you will get the following error while compiling programs using the command-line code.
 
 ```sh
 xcrun: error: invalid active developer path (/Library/Developer/CommandLineTools), missing xcrun at:
 /Library/Developer/CommandLineTools/usr/bin/xcrun
```

[<< Go to table of contents](/doc/cPrograms/readme.md)

[<< Go to previous chapter (1.1 Downloading and installing Visual Code)](/doc/1-1-downloading-and-installing-visual-code.md)

# 2. Writing and compiling `helloWorld.c`

Here is the C program outputting the string "Hello world".

```c
#include <stdio.h>
int main()
{
    printf("Hello world \n");
    return 0;
}
```

The Unix command for compiling is gcc. GCC is a compilier from the [GNU project](https://www.gnu.org/), and it stands for GNU Compilier Collections. To compile the C program in command-line, enter the following command. 

```sh
gcc helloWorld.c
```

The full syntax is as follows.

```sh
gcc helloWorld.c -o helloWorld
```

The difference is that in the first command, an executable file *a.out* is automatically created, while in the second command, you create the executable file *helloWorld*. To see the creation of these files with your own eyes, enter the command ```ls```.

Execute the program by typing ```helloWorld``` in the terminal.

If trying to execute the program leads to an error, it is because the program's location is not included in the PATH variable. PATH is an [environment variable](https://en.wikipedia.org/wiki/Environment_variable) that tells the shell, i.e. the terminal interface, which directories to search for executables in response to the commands issued by the user. For example, ```ls, mkdir, cd, rm, cp, mv...``` commands, which are stored in the *usr/bin* folder, are executable throughout the file system because the *usr/bin* folder is incuded in the PATH. 

To view the contents of the PATH variable, enter the following command.

```sh
echo $PATH
```

The input and output would be something like the following.

```sh
Pranavs-MacBook-Air:cPrograms pranav$ echo $PATH
/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:
/Applications/VMware Fusion.app/Contents/Public:
/usr/local/share/dotnet:
~/.dotnet/tools:
/Library/Frameworks/Mono.framework/Versions/Current/Commands:
/Users/pranav/cPrograms/
```

If you want to add a directory to PATH, enter the following command.

```sh
export PATH=$PATH:/<yourWorkingDirectory>
```

## Additional information
To see all the environment variables in the file system, enter the command ```printenv```.

