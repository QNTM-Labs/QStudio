package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"io/ioutil"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Renderer string `toml:"renderer"`
	Esync    bool   `toml:"esync"`
	Fsync    bool   `toml:"fsync"`
	DpiFix   bool   `toml:"dpi_fix"`
}

func func1() Config {
	homeDir, _ := os.UserHomeDir()
	configPath := filepath.Join(homeDir, ".qstudio", "config.toml")
	
	var config Config
	config.Renderer = "vulkan"
	config.Esync = true
	config.Fsync = true
	config.DpiFix = true
	
	if _, err := os.Stat(configPath); err == nil {
		toml.DecodeFile(configPath, &config)
	}
	
	return config
}

func func2() {
	homeDir, _ := os.UserHomeDir()
	qstudioDir := filepath.Join(homeDir, ".qstudio")
	prefixDir := filepath.Join(qstudioDir, "wineprefix")
	overlayDir := filepath.Join(qstudioDir, "overlay")
	
	os.MkdirAll(qstudioDir, 0755)
	os.MkdirAll(prefixDir, 0755)
	os.MkdirAll(overlayDir, 0755)
	
	configPath := filepath.Join(qstudioDir, "config.toml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		defaultConfig := `renderer = "vulkan"
esync = true
fsync = true
dpi_fix = true`
		ioutil.WriteFile(configPath, []byte(defaultConfig), 0644)
	}
	
	func7()
}

func func3(config Config) {
	homeDir, _ := os.UserHomeDir()
	prefixDir := filepath.Join(homeDir, ".qstudio", "wineprefix")
	
	env := os.Environ()
	env = append(env, "WINEPREFIX="+prefixDir)
	
	if config.Esync {
		env = append(env, "WINEESYNC=1")
	}
	
	if config.Fsync {
		env = append(env, "WINEFSYNC=1")
	}
	
	switch config.Renderer {
	case "vulkan":
		env = append(env, "DXVK_HUD=1")
		env = append(env, "__GL_SHADER_DISK_CACHE=1")
		env = append(env, "__GL_THREADED_OPTIMIZATIONS=1")
	case "dx11":
		env = append(env, "WINED3D_CONFIG=renderer=gl")
	case "opengl":
		env = append(env, "WINED3D_CONFIG=renderer=gl")
		env = append(env, "__GL_SHADER_DISK_CACHE=1")
	}
	
	if config.DpiFix {
		env = append(env, "WINEDLLOVERRIDES=winemenubuilder.exe=d")
	}
	
	fmt.Println("Setting up Wine prefix...")
	cmd := exec.Command("wineboot", "--init")
	cmd.Env = env
	cmd.Dir = prefixDir
	cmd.Run()
	
	fmt.Println("Installing Wine components...")
	cmd = exec.Command("winetricks", "--unattended", "corefonts", "vcrun2019", "dotnet48")
	cmd.Env = env
	cmd.Dir = prefixDir
	cmd.Run()
	
	if config.Renderer == "vulkan" {
		fmt.Println("Installing DXVK...")
		dxvkPath := filepath.Join(homeDir, ".local", "share", "dxvk", "dxvk-2.3.1")
		if _, err := os.Stat(dxvkPath); err == nil {
			cmd = exec.Command("bash", filepath.Join(dxvkPath, "setup_dxvk.sh"), "install")
			cmd.Env = env
			cmd.Dir = prefixDir
			cmd.Run()
		}
	}
	
	overlayDir := filepath.Join(homeDir, ".qstudio", "overlay")
	if entries, err := ioutil.ReadDir(overlayDir); err == nil {
		for _, entry := range entries {
			if !entry.IsDir() {
				sourcePath := filepath.Join(overlayDir, entry.Name())
				targetPath := filepath.Join(prefixDir, "drive_c", "users", os.Getenv("USER"), "AppData", "Local", "Roblox", "Versions")
				os.MkdirAll(targetPath, 0755)
				
				sourceData, _ := ioutil.ReadFile(sourcePath)
				ioutil.WriteFile(filepath.Join(targetPath, entry.Name()), sourceData, 0644)
			}
		}
	}
	
	fmt.Println("Wine prefix setup complete!")
}

func main() {
	func6()
}

func func4() {
	homeDir, _ := os.UserHomeDir()
	prefixDir := filepath.Join(homeDir, ".qstudio", "wineprefix")
	
	env := os.Environ()
	env = append(env, "WINEPREFIX="+prefixDir)
	
	studioPath := filepath.Join(prefixDir, "drive_c", "users", os.Getenv("USER"), "AppData", "Local", "Roblox", "Versions")
	
	if _, err := os.Stat(studioPath); os.IsNotExist(err) {
		fmt.Println("Installing Roblox Studio...")
		
		cmd := exec.Command("wget", "-O", "/tmp/RobloxStudioLauncherBeta.exe", "https://setup.rbxcdn.com/RobloxStudioLauncherBeta.exe")
		cmd.Run()
		
		cmd = exec.Command("wine", "/tmp/RobloxStudioLauncherBeta.exe", "/S")
		cmd.Env = env
		cmd.Dir = prefixDir
		cmd.Run()
		
		os.Remove("/tmp/RobloxStudioLauncherBeta.exe")
	}
}

func func5(config Config) {
	homeDir, _ := os.UserHomeDir()
	prefixDir := filepath.Join(homeDir, ".qstudio", "wineprefix")
	
	env := os.Environ()
	env = append(env, "WINEPREFIX="+prefixDir)
	
	if config.Esync {
		env = append(env, "WINEESYNC=1")
	}
	
	if config.Fsync {
		env = append(env, "WINEFSYNC=1")
	}
	
	switch config.Renderer {
	case "vulkan":
		env = append(env, "DXVK_HUD=1")
		env = append(env, "__GL_SHADER_DISK_CACHE=1")
		env = append(env, "__GL_THREADED_OPTIMIZATIONS=1")
	case "dx11":
		env = append(env, "WINED3D_CONFIG=renderer=gl")
	case "opengl":
		env = append(env, "WINED3D_CONFIG=renderer=gl")
		env = append(env, "__GL_SHADER_DISK_CACHE=1")
	}
	
	if config.DpiFix {
		env = append(env, "WINEDLLOVERRIDES=winemenubuilder.exe=d")
	}
	
	studioPath := filepath.Join(prefixDir, "drive_c", "users", os.Getenv("USER"), "AppData", "Local", "Roblox", "Versions")
	
	var studioExe string
	if entries, err := ioutil.ReadDir(studioPath); err == nil {
		for _, entry := range entries {
			if entry.IsDir() && strings.HasPrefix(entry.Name(), "version-") {
				exePath := filepath.Join(studioPath, entry.Name(), "RobloxStudioBeta.exe")
				if _, err := os.Stat(exePath); err == nil {
					studioExe = exePath
					break
				}
			}
		}
	}
	
	if studioExe == "" {
		fmt.Println("Roblox Studio not found. Installing...")
		func4()
		return
	}
	
	fmt.Println("Launching Roblox Studio...")
	cmd := exec.Command("wine", studioExe)
	cmd.Env = env
	cmd.Dir = prefixDir
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Run()
}

func func6() {
	config := func1()
	
	fmt.Println("QStudio - Roblox Studio Launcher for Linux")
	fmt.Println("=========================================")
	fmt.Printf("Renderer: %s\n", config.Renderer)
	fmt.Printf("Esync: %t\n", config.Esync)
	fmt.Printf("Fsync: %t\n", config.Fsync)
	fmt.Printf("DPI Fix: %t\n", config.DpiFix)
	fmt.Println("=========================================")
	
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "install":
			func2()
			func3(config)
			func4()
			fmt.Println("Installation complete!")
		case "config":
			fmt.Println("Current configuration:")
			fmt.Printf("  renderer = \"%s\"\n", config.Renderer)
			fmt.Printf("  esync = %t\n", config.Esync)
			fmt.Printf("  fsync = %t\n", config.Fsync)
			fmt.Printf("  dpi_fix = %t\n", config.DpiFix)
		case "launch":
			func2()
			func3(config)
			func5(config)
		default:
			fmt.Println("Usage: qstudio [install|config|launch]")
		}
	} else {
		func2()
		func3(config)
		func5(config)
	}
}

func func7() {
	fmt.Println("Installing dependencies...")
	
	distroID := func8()
	
	switch distroID {
	case "ubuntu", "debian", "linuxmint", "pop":
		exec.Command("sudo", "apt", "update").Run()
		exec.Command("sudo", "apt", "install", "-y", "wine", "winetricks", "wget", "curl", "cabextract", "unzip").Run()
		exec.Command("sudo", "apt", "install", "-y", "wine-staging").Run()
		exec.Command("sudo", "apt", "install", "-y", "dxvk").Run()
	case "fedora", "rhel", "centos":
		exec.Command("sudo", "dnf", "install", "-y", "wine", "winetricks", "wget", "curl", "cabextract", "unzip").Run()
		exec.Command("sudo", "dnf", "install", "-y", "wine-staging").Run()
		exec.Command("sudo", "dnf", "copr", "enable", "kylegospo/dxvk").Run()
		exec.Command("sudo", "dnf", "install", "-y", "dxvk").Run()
	case "arch", "manjaro", "endeavouros":
		exec.Command("sudo", "pacman", "-S", "--noconfirm", "wine", "winetricks", "wget", "curl", "cabextract", "unzip").Run()
		exec.Command("sudo", "pacman", "-S", "--noconfirm", "wine-staging").Run()
		exec.Command("sudo", "pacman", "-S", "--noconfirm", "dxvk-bin").Run()
	case "opensuse", "sles":
		exec.Command("sudo", "zypper", "install", "-y", "wine", "winetricks", "wget", "curl", "cabextract", "unzip").Run()
		exec.Command("sudo", "zypper", "install", "-y", "wine-staging").Run()
	case "gentoo":
		exec.Command("sudo", "emerge", "-av", "wine-vanilla", "winetricks", "wget", "curl", "cabextract", "unzip").Run()
	default:
		fmt.Println("Unsupported distribution. Please install wine, winetricks, wget, curl, cabextract, unzip manually.")
		fmt.Println("Also install wine-staging and dxvk if available.")
	}
	
	homeDir, _ := os.UserHomeDir()
	dxvkPath := filepath.Join(homeDir, ".local", "share", "dxvk")
	if _, err := os.Stat(dxvkPath); os.IsNotExist(err) {
		func9()
	}
	
	fmt.Println("Dependencies installed successfully!")
}

func func8() string {
	releaseFile := "/etc/os-release"
	if content, err := ioutil.ReadFile(releaseFile); err == nil {
		lines := strings.Split(string(content), "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "ID=") {
				return strings.Trim(strings.TrimPrefix(line, "ID="), "\"")
			}
		}
	}
	return "unknown"
}

func func9() {
	fmt.Println("Installing DXVK...")
	homeDir, _ := os.UserHomeDir()
	dxvkDir := filepath.Join(homeDir, ".local", "share", "dxvk")
	os.MkdirAll(dxvkDir, 0755)
	
	cmd := exec.Command("wget", "-O", "/tmp/dxvk-release.tar.gz", "https://github.com/doitsujin/dxvk/releases/download/v2.3.1/dxvk-2.3.1.tar.gz")
	cmd.Run()
	
	cmd = exec.Command("tar", "-xzf", "/tmp/dxvk-release.tar.gz", "-C", "/tmp/")
	cmd.Run()
	
	cmd = exec.Command("cp", "-r", "/tmp/dxvk-2.3.1", dxvkDir)
	cmd.Run()
	
	os.Remove("/tmp/dxvk-release.tar.gz")
	exec.Command("rm", "-rf", "/tmp/dxvk-2.3.1").Run()
}
