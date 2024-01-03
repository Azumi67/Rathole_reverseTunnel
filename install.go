package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func swapFile() error {
	cmd := exec.Command("sudo", "fallocate", "-l", "2G", "/swapfile")
	err := cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("sudo", "chmod", "600", "/swapfile")
	err = cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("sudo", "mkswap", "/swapfile")
	err = cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("sudo", "swapon", "/swapfile")
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func install() error {
	log.Println("\033[93mInstalling Git..\033[0m")
	cmd := exec.Command("sudo", "apt-get", "install", "git", "-y")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("\033[91merror installing Git: %s\033[0m", err.Error())
	}
	log.Println("\033[92mGit installation completed!\033[0m")

	log.Println("\033[93mCloning repo..\033[0m")
	err = os.RemoveAll("/root/rathole")
	if err != nil {
		return fmt.Errorf("\033[91merror @removing existing rathole dir: %s\033[0m", err.Error())
	}

	cmd = exec.Command("git", "clone", "https://github.com/miyugundam/rathole.git")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("\033[91merror cloning repository: %s\033[0m", err.Error())
	}
	log.Println("\033[92mRepo's clone was successful\033[0m")

	return nil
}

func cargoToml() error {
	cargoFile := "/root/rathole/Cargo.toml"
	_, err := os.Stat(cargoFile)
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}

	file, err := os.OpenFile(cargoFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("\n[workspace]\n")
	if err != nil {
		return err
	}

	return nil
}

func buildRat() error {
	log.Println("\033[92mBuilding rathole..\033[0m")
	cmd := exec.Command("cargo", "build")
	cmd.Dir = "/root/rathole"

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}
	log.Println("\033[92mRathole built successfully!\033[0m")

	return nil
}

func swap() (bool, error) {
	cmd := exec.Command("free", "-h")
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}

	fmt.Println(string(output))

	cmd = exec.Command("swapon", "--show")
	output, err = cmd.Output()
	if err != nil {
		return false, err
	}

	return strings.Contains(string(output), "/swapfile"), nil
}


func main() {
	swapEnabled, err := swap()
	if err != nil {
		log.Fatal("\033[91merror checking swap status:\033[0m", err)
	}

	if !swapEnabled {
		err := swapFile()
		if err != nil {
			log.Fatal("\033[91merror swapping:\033[0m", err)
		}
	}
	err = install()
	if err != nil {
		log.Fatal("\033[91merror installing Git & cloning repo:\033[0m", err)
	}

	err = os.Chdir("/root/rathole")
	if err != nil {
		log.Fatal("\033[91merror changing dir to rathole:\033[0m", err)
	}

	err = cargoToml()
	if err != nil {
		log.Fatal("\033[91merror modifying Cargo:\033[0m", err)
	}

	err = buildRat()
	if err != nil {
		log.Fatal("\033[91merror building rat holz!:\033[0m", err)
	}
}