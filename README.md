# dImports
Package for importing and running Go dApps.

1. [About](#about) 
2. [StartApp](#startapp)
3. [Importing](#importing)
4. [Build](#build) 
5. [Donations](#donations) 
6. [Licensing](#licensing) 

### About
![goMod](https://img.shields.io/github/go-mod/go-version/dReam-dApps/dImports.svg)![goReport](https://goreportcard.com/badge/github.com/dReam-dApps/dImports)[![goDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/dReam-dApps/dImports)
The idea behind **dImports** is to be able to import and run Go dApps or code that is external to the executing program. Primarily designed for use with [Dero](https://dero.io) dApps, **dImports** aims to give users easy access to libraries of powerful dApps. 

Currently the `dimport` package is using [gore](https://github.com/x-motemen/gore) to import a Go package from *any given path* and run its `StartApp()`. It includes a [Fyne](https://fyne.io) widget import for GUI use which is demonstrated in `cmd/dImporter-gui`.



View latest [release](https://github.com/dReam-dApps/dImports/releases)

![windowsOS](https://raw.githubusercontent.com/SixofClubsss/dreamdappsite/main/assets/os-windows-green.svg)![macOS](https://raw.githubusercontent.com/SixofClubsss/dreamdappsite/main/assets/os-macOS-green.svg)![linuxOS](https://raw.githubusercontent.com/SixofClubsss/dreamdappsite/main/assets/os-linux-green.svg)

### StartApp
Simply put, `StartApp()` is a exportable version of a `main()`. It contains the overall logic and control flow of a application. To build a `StartApp()` that can be dimported, write a function called StartApp in a Go package and publish it. 
```
package mydapp

func StartApp() {
    // Your logic
}
```
- *mydapp must not be in top level of repo*

### Importing
To run a external `StartApp()` from your code.
```
package main

import "github.com/dReam-dApps/dImports/dimport"

func main() {
	path := "github.com/user/repo/mydapp"

	err := dimport.ImportAndStartApp(path)
	if err != nil {
		// handle error
	}
}
```

### Build
dImports contains a CLI and GUI app which can be built from source.
- Install latest [Go version](https://go.dev/doc/install)
- Install [Fyne](https://developer.fyne.io/started/) dependencies (*this step is only required for GUI build*)
- Clone repo and build using:
```
git clone https://github.com/dReam-dApps/dImports.git
cd dImports/cmd/dImporter-gui
go build .
./dImporter-gui
```

### Donations
- **Dero Address**: dero1qyr8yjnu6cl2c5yqkls0hmxe6rry77kn24nmc5fje6hm9jltyvdd5qq4hn5pn

![DeroDonations](https://raw.githubusercontent.com/SixofClubsss/dreamdappsite/main/assets/DeroDonations.jpg)

---

#### Licensing

dImports is free and open source.   
The source code is published under the [MIT](https://github.com/dReam-dApps/dImports/blob/main/LICENSE) License.   
Copyright © 2023 dReam dApps   
