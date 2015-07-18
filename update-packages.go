package main

import (
  "os"
  "fmt"
  "os/exec"
)

var installedPackages []string

func exec_cmd(cmd string) {
  out, err := exec.Command("sh","-c", "which " + cmd).Output()
  fmt.Printf("%s", out)

  if err != nil {
    println(err.Error())
  } else {
    installedPackages = append(installedPackages, cmd)
  }
}

func main() {
  packages := []string{"npm", "rvm", "brew", "gem"}

  for _, cmd := range packages {
    exec_cmd(cmd)
  }

  for _, cmd := range installedPackages {
    switch cmd {
    case "npm":
      fmt.Println("Looking for outdated packages...")
      out, err := exec.Command("sh","-c", "npm outdated -g").Output()
      if err != nil {
        println(err.Error())
      } else {
        fmt.Printf("%s", out)
      }
      fmt.Println("Updating packages...")
      out2, err := exec.Command("sh","-c", "npm update -g").Output()
      if err != nil {
        println(err.Error())
      } else {
        fmt.Printf("%s", out2)
      }

    case "brew":
      fmt.Println("Looking for new formulae...")
      out3, err := exec.Command("sh","-c", "brew update").Output()
      if err != nil {
        println(err.Error())
      } else {
        fmt.Printf("%s", out3)
      }
      fmt.Println("Updating formulae...")
      out4, err := exec.Command("sh","-c", "brew upgrade --all").Output()
      if err != nil {
        println(err.Error())
      } else {
        fmt.Printf("%s", out4)
      }

    case "gem":

    default:
      fmt.Println("not found")
    }
  }

  fmt.Println(installedPackages)

  os.Stdout.WriteString("done")
}
