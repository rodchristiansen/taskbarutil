# TaskbarUtil

**TaskbarUtil** is a command-line utility for managing Windows taskbar items. It allows you to pin, unpin, and list applications and shortcuts on the taskbar. This tool is useful for system administrators, DevOps teams, and power users who need to automate taskbar configuration across single or multi-user environments. Written in GoLang, it offers cross-platform possibilities and can be integrated with version control for easy management.

- Compatible with Windows 10 and later
- Pin, unpin, and list Taskbar items
- Supports Applications and shortcuts
- Can manage taskbar layouts for all users or specific user profiles

---

## **CHANGELOG**

### Version 2024.09.20
- Initial release: Pin, unpin, and list taskbar items
- Basic taskbar layout management for individual and all users
- Written in GoLang for cross-platform potential

---

## **USAGE**

### Pin an Application to the Taskbar
```bash
taskbarutil pin "C:\Program Files\Microsoft Office\root\Office16\OUTLOOK.EXE"
```

### Unpin an Application from the Taskbar
```bash
taskbarutil unpin "OUTLOOK.EXE"
```

### List All Pinned Applications
```bash
taskbarutil list
```

## **OPTIONS**

### Command Options:
- **pin**: Pin a new item to the taskbar
- **unpin**: Unpin an existing item from the taskbar
- **list**: List all pinned items on the taskbar

### Examples:
1. Pin **Outlook** to the taskbar:
   ```bash
   taskbarutil pin "C:\Program Files\Microsoft Office\root\Office16\OUTLOOK.EXE"
   ```

2. Unpin **Outlook** from the taskbar:
   ```bash
   taskbarutil unpin "OUTLOOK.EXE"
   ```

3. List all currently pinned items:
   ```bash
   taskbarutil list
   ```

## **ADVANCED USAGE**

TaskbarUtil can work in multi-user environments. Administrators can apply taskbar configurations for all user profiles on a system or for specific profiles by specifying the **TaskBar** directories.

### Examples:
1. Pin **Outlook** to all user profiles:
   ```bash
   taskbarutil pin "C:\Program Files\Microsoft Office\root\Office16\OUTLOOK.EXE" --allusers
   ```

2. Unpin an item from a specific user’s taskbar:
   ```bash
   taskbarutil unpin "OUTLOOK.EXE" --user "C:\Users\jsmith"
   ```

---

## **LIMITATIONS AND DEPENDENCIES**
- Works only on **Windows 10** and **Windows 11**
- Requires access to taskbar pinned items folder for multi-user tasks

---

## **CONTRIBUTING**

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/my-feature`)
3. Commit your changes (`git commit -am 'Add my feature'`)
4. Push to the branch (`git push origin feature/my-feature`)
5. Create a new pull request

## **LICENSE**
Licensed under the [Apache 2.0 License](LICENSE).
