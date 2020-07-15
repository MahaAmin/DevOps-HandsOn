# Make

## Table of Content:

1. [What is make?](#what-is-make?)
2. [Makefile](#makefile)


<hr>

## What Is Make?

- Make is an automation utility. It determines which pieces of a large program need to be recompiled, and issues commands to recompile them.

- GNU make was developed by Richard Stallman and Roland McGrath. [GNU-Make-Manual](https://www.gnu.org/software/make/manual/make.pdf)


<hr>

## Makefile:

- Basic Syntax:

    ```
    target: prerequisites
    <TAB> recipe
    ```

- Example:
    ```
    cmd1:
	    @echo "This is cmd1"

    cmd2:
        @echo "This is cmd2"

    all: cmd1 cmd2
    ```

- Run (This will only executes the first target):

    ```
    make
    ```

- Make default target. Add this to the beginning of the makefile:

    ```
    .DEFAULT_GOAL := cmd2
        ...
    ```

- Run all targets:
    ```
    make all
    ```

- Run specific target:
    ```
    make cmd2
    ```
<hr>




