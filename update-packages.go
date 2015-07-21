package main

import (
  "os"
  "fmt"
  "os/exec"
  "bufio"
  "strings"
)

var installedPackages []string

func check_package(cmd string) {
  _, err := exec.Command("sh", "-c", "which " + cmd).Output()

  if err != nil {
    fmt.Printf("Package manager: %s not found\n", cmd)
  } else {
    fmt.Printf("Package manager: %s installed\n", cmd)
    installedPackages = append(installedPackages, cmd)
  }
}

func ask_user() bool {
  consolereader := bufio.NewReader(os.Stdin)
  fmt.Print("Do you want to sudo (n): ")

  input, err := consolereader.ReadString('\n') // this will prompt the user for input

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  sudoMode := false
  if strings.TrimSpace(input) == "y" {
    sudoMode = true;
  }

  return sudoMode
}

func exec_cmd(cmd string) ([]byte, error) {
  out, err := exec.Command("sh", "-c", cmd).Output()
  return out, err
}

func process_cmd(cmd string) {
  outCmd, err := exec_cmd(cmd)

  if err != nil {
    println(err.Error())
    if ask_user() {
      outSudoCmd, err := exec_cmd("sudo " + cmd)
      if err != nil {
        println(err.Error())
      } else {
        fmt.Printf("%s", outSudoCmd)
      }
    } else {
      fmt.Println("User did not grant sudo permissions")
    }
  } else {
    fmt.Printf("%s", outCmd)
  }
}

func main() {
  packages := []string{"npm", "rvm", "brew", "gem"}

  for _, cmd := range packages {
    check_package(cmd)
  }

  for _, cmd := range installedPackages {
    switch cmd {
    case "npm":
      fmt.Println("npm: Looking for outdated packages...")
      process_cmd("npm outdated -g")

      fmt.Println("npm: Updating packages...")
      process_cmd("npm update -g")

    case "brew":
      fmt.Println("brew: Looking for new formulae...")
      process_cmd("brew update")

      fmt.Println("brew: Updating formulae...")
      process_cmd("brew upgrade --all")

    case "gem":
      fmt.Println("gem: Looking for outdated packages...")
      process_cmd("gem outdated")

      fmt.Println("gem: Updating packages...")
      process_cmd("gem update")

      fmt.Println("gem: Cleaning old packages...")
      process_cmd("gem clean")

    default:
      fmt.Println("Package manager not found")
    }
  }

  // show installed package managers
  // fmt.Println(installedPackages)

  os.Stdout.WriteString("done\n")
}
