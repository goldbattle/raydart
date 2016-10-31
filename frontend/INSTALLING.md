* Full instructions for installing are located [here](https://golang.org/doc/install)
* Download the golang toolset from [golang.org](https://golang.org/dl/)
  * Go is normally installed at `C:/Go/`
* On windows create a "go workspace"
  * This is where all working files, and downloads are stored
  * The recommended folder to make is `C:/Users/<username>/Go Workspace/`
  * Inside the workspace create a folder called `src/`
* Next open the windows [environmental variables](https://www.java.com/en/download/help/path.xml)
  * Create the following entries and point them to your respective locations
  * `GOROOT` - `C:/Go/`
  * `GOPATH` - `C:/Users/<username>/Go Workspace/`
  * Append the following to the `Path` variable, separate with semi-colons
  * `Path` - `C:\Go\bin`
* Next download any version of a Jetbrain
  * I recommend [webstorm](https://www.jetbrains.com/webstorm/)
  * Install this normally with the given installer
  * If you are a student get a [free](https://www.jetbrains.com/student/) account,
  * If you are not a student use the [community edition](https://www.jetbrains.com/idea/) of idea
* After loading into the editor install the golang plugin
  * Github is located here: https://github.com/go-lang-plugin-org/go-lang-idea-plugin
  * Plugin repo is here: https://plugins.jetbrains.com/plugin/5047
  * To install a plugin, search for "plugins", or on the splash screen click on "configure" in the bottom right
  * Next click "Browse Repositories"
  * Next click "Manage Repositories"
  * Click the plus button and paste in the plugin repo url above
  * Then click ok, and search for the plugin called "Go"
  * Install it
* Next clone this repository into your go workspace
  * In go, the package structure is based on the domain name it is downloaded from
  * In our case the folder "web" should be placed in the following
  * `C:/Users/<username>/Go Workspace/src/git.pgeneva.com/pistachio/web`
* Finally open the folder `web` in the editor
* Create a new building/running application
  * Select "go application"
  * Set the file to `web.go`
  * Set the output folder to be a folder called `out` inside the `web` folder
* Finally download all the missing packages needed
  * Open the terminal (bottom left of the editor)
  * Run `go get` if this command is missing, you have not set your `Path` variable correctly
  * Try restarting the editor refresh your command prompt