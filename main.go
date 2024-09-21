package main

import (
    "fmt"
    "os"
    "path/filepath"
    "os/exec"
    "log"
)

// Pin an application to the taskbar
func pinToTaskbar(appPath string) error {
    userProfile := os.Getenv("USERPROFILE")
    taskbarPinnedFolder := filepath.Join(userProfile, "AppData", "Roaming", "Microsoft", "Internet Explorer", "Quick Launch", "User Pinned", "TaskBar")

    appName := filepath.Base(appPath)
    shortcutPath := filepath.Join(taskbarPinnedFolder, appName+".lnk")

    cmd := exec.Command("powershell", "-Command", fmt.Sprintf(`$WScriptShell = New-Object -ComObject WScript.Shell; $Shortcut = $WScriptShell.CreateShortcut("%s"); $Shortcut.TargetPath = "%s"; $Shortcut.Save()`, shortcutPath, appPath))
    
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to pin app: %v", err)
    }

    fmt.Printf("Application pinned to taskbar: %s\n", appName)
    return nil
}

// Unpin an application from the taskbar
func unpinFromTaskbar(appName string) error {
    userProfile := os.Getenv("USERPROFILE")
    taskbarPinnedFolder := filepath.Join(userProfile, "AppData", "Roaming", "Microsoft", "Internet Explorer", "Quick Launch", "User Pinned", "TaskBar")

    shortcutPath := filepath.Join(taskbarPinnedFolder, appName+".lnk")

    if _, err := os.Stat(shortcutPath); err == nil {
        if err := os.Remove(shortcutPath); err != nil {
            return fmt.Errorf("failed to unpin app: %v", err)
        }
        fmt.Printf("Application unpinned from taskbar: %s\n", appName)
    } else {
        return fmt.Errorf("shortcut not found for %s", appName)
    }

    return nil
}

// List all pinned applications in the taskbar
func listPinnedItems() ([]string, error) {
    userProfile := os.Getenv("USERPROFILE")
    taskbarPinnedFolder := filepath.Join(userProfile, "AppData", "Roaming", "Microsoft", "Internet Explorer", "Quick Launch", "User Pinned", "TaskBar")

    var pinnedApps []string
    files, err := os.ReadDir(taskbarPinnedFolder)
    if err != nil {
        return nil, fmt.Errorf("failed to read pinned items: %v", err)
    }

    for _, file := range files {
        pinnedApps = append(pinnedApps, file.Name())
    }

    return pinnedApps, nil
}

// Main handler for command-line interface
func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: taskbarutil <command> [appPath/appName]")
        fmt.Println("Commands: pin, unpin, list")
        os.Exit(1)
    }

    command := os.Args[1]

    switch command {
    case "pin":
        if len(os.Args) < 3 {
            log.Fatal("Please provide the path of the application to pin")
        }
        appPath := os.Args[2]
        if err := pinToTaskbar(appPath); err != nil {
            log.Fatal(err)
        }

    case "unpin":
        if len(os.Args) < 3 {
            log.Fatal("Please provide the name of the application to unpin")
        }
        appName := os.Args[2]
        if err := unpinFromTaskbar(appName); err != nil {
            log.Fatal(err)
        }

    case "list":
        pinnedItems, err := listPinnedItems()
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println("Pinned Applications:")
        for _, app := range pinnedItems {
            fmt.Println(app)
        }

    default:
        fmt.Println("Invalid command. Use pin, unpin, or list.")
    }
}
