# godot for spx


## Getting Started (WIP, Linux only for now)

### start and run 
```
    git clone git@github.com:realdream-ai/gdspx.git
    cd gdspx 
    cd cmd/gdspx/
    go install .
    cd ../../
    gdspx editor tutorial/01_aircraft
    # gdspx run tutorial/01_aircraft  
```

### how to use
    
    - gdspx init            # Create a gdspx project in the current directory
    - gdspx run             # Run the current project
    - gdspx build           # Build the dynamic library
    - gdspx export          # Export the PC package (supports macOS, Windows, Linux)
    - gdspx buildweb        # Build for WebAssembly (WASM)
    - gdspx exportweb       # Export the web package




## Build from source
### prepare env
1. install "**make**" for build
2. install python
3. run script 
    ```
    make init
    ```

### make cmd
eg:  make pc 

    - make init          # Initialize the environment
    - make engine        # Rebuild engine
    - make editor        # Run in editor mode
    - make build         # Build DLL for PC (Windows, macOS, Linux)
    - make run           # Run on PC (Windows, macOS, Linux)
    - make wasm          # Build for WebAssembly (WASM)
    - make web           # Build and export the web package to the directory (./build/games)
    - make server        # Launch a local web server (http://127.0.0.1:8005/)
    - make fmt           # Format the code


