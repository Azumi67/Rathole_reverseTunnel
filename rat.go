//Author:github.com/Azumi67
//This script is for educational use and for my own learning, but I'd be happy if you find it useful too.
//This script simplifies the configuration of Rathole Reverse tunnel.
//You can send me feedback so I can use it to learn more.
//This script comes without any warranty
//Thank you.
package main

import (
    "time"
	"strconv"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"github.com/fatih/color"
	"log"
	"github.com/AlecAivazis/survey/v2"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"net"
	"io/ioutil"
	"io"
)
func getIPv4() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}

	for _, iface := range interfaces {
		name := iface.Name
		if strings.HasPrefix(name, "eth") || strings.HasPrefix(name, "en") {
			addresses, err := iface.Addrs()
			if err != nil {
				continue
			}

			for _, addr := range addresses {
				ipnet, ok := addr.(*net.IPNet)
				if ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
					return ipnet.IP.String()
				}
			}
		}
	}

	return ""
}
func displayProgress(total, current int) {
	width := 40
	percentage := current * 100 / total
	completed := width * current / total
	remaining := width - completed

	fmt.Printf("\r[%s>%s] %d%%", strings.Repeat("=", completed), strings.Repeat(" ", remaining), percentage)
}

func displayError(message string) {
	fmt.Printf("\u2718 Error: %s\n", message)
}

func displayNotification(message string) {
	fmt.Printf("\033[93m%s\033[0m\n", message)
}

func displayCheckmark(message string) {
	fmt.Printf("\033[92m\u2714 \033[0m%s\n", message)
}

func displayLoading() {
    frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
    delay := 100 * time.Millisecond
    duration := 5 * time.Second

    endTime := time.Now().Add(duration)

    for time.Now().Before(endTime) {
        for _, frame := range frames {
            fmt.Printf("\r[%s] Loading... ", frame)
            time.Sleep(delay)
        }
    }
    fmt.Println()
}
func displayLogo2() error {
	cmd := exec.Command("bash", "-c", "/etc/./logo.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
func displayLogo() {
	logo := `
   ______    _______    __      _______          __       _____  ___  
  /    " \  |   __ "\  |" \    /"      \        /""\      (\"   \|" \ 
 // ____  \ (. |__) :) ||  |  |:        |      /    \     |.\\   \   |
/  /    ) :)|:  ____/  |:  |  |_____/   )     /' /\  \    |: \.   \\ |
(: (____/ // (|  /     |.  |   //       /    //  __'  \   |.  \    \ |
\        // |__/ \     /\  |\  |:  __   \   /   /  \\  \  |    \    \|
 \"_____ / (_______)  (__\_|_) |__|  \___) (___/    \___) \___|\____\)
`
	
    cyan := color.New(color.FgCyan, color.Bold).SprintFunc()
    blue := color.New(color.FgBlue, color.Bold).SprintFunc()
	green := color.New(color.FgHiGreen, color.Bold).SprintFunc()      
    yellow := color.New(color.FgHiYellow, color.Bold).SprintFunc()   
    red := color.New(color.FgHiRed, color.Bold).SprintFunc()        


	

	    logo = cyan("  ______   ") + blue(" _______  ") + green("  __    ") + yellow("   ______   ") + red("     __      ") + cyan("  _____  ___  \n") +
		cyan(" /     \" \\  ") + blue("|   __ \" ") + green(" |\" \\  ") + yellow("   /\"     \\   ") + red("   /\"\"\\     ") + cyan(" (\\\"   \\|\"  \\ \n") +
		cyan("//  ____  \\ ")  + blue("(. |__) :)") + green("||  |  ") + yellow(" |:       |  ") + red("  /    \\   ") + cyan("  |.\\\\   \\   |\n") +
		cyan("/  /    ) :)") + blue("|:  ____/ ") + green("|:  |  ") + yellow(" |_____/  )  ") + red(" /' /\\  \\   ") + cyan(" |: \\.   \\\\ |\n") +
		cyan("(: (____/ / ") + blue("(|  /     ") + green("|.  | ") + yellow("  //      /  ") + red("//   __'  \\  ") + cyan(" |.  \\    \\ |\n") +
		cyan("\\        / ") + blue("/|__/ \\   ") + green(" /\\  |\\ ") + yellow(" |:  __  \\ ") + red(" /   /  \\\\   ") + cyan ("  |    \\    \\|\n") +
		cyan(" \"_____ / ") + blue("(_______)") + green("  (__\\_|_)") + yellow(" (__) \\___)") + red("(___/    \\___)") + cyan(" \\___|\\____\\)\n")


	fmt.Println(logo)
}
func main() {
	if os.Geteuid() != 0 {
		fmt.Println("\033[91mThis script must be run as root. Please use sudo -i.\033[0m")
		os.Exit(1)
	}

	mainMenu()
}
func readInput() {
	fmt.Print("Press Enter to continue..")
	fmt.Scanln()
	mainMenu()
}
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func mainMenu() {
	for {
		err := displayLogo2()
		if err != nil {
			log.Fatalf("failed to display logo: %v", err)
		}
		displayLogo()
		border := "\033[93m+" + strings.Repeat("=", 70) + "+\033[0m"
		content := "\033[93m║            ▌║█║▌│║▌│║▌║▌█║ \033[92mMain Menu\033[93m  ▌│║▌║▌│║║▌█║▌                  ║"
		footer := " \033[92m            Join Opiran Telegram \033[34m@https://t.me/OPIranClub\033[0m "

		borderLength := len(border) - 2
		centeredContent := fmt.Sprintf("%[1]*s", -borderLength, content)

		fmt.Println(border)
		fmt.Println(centeredContent)
		fmt.Println(border)

		fmt.Println(border)
		fmt.Println(footer)
		fmt.Println(border)
		prompt := &survey.Select{
			Message: "Enter your choice Please:",
			Options: []string{"0. \033[91mSTATUS Menu\033[0m", "1. \033[96mIPV4 \033[92mTCP \033[0m", "2. \033[93mIPV4 \033[92mUDP\033[0m", "3. \033[96mIPV6 \033[92mTCP\033[0m", "4. \033[93mIPV6 \033[92mUDP\033[0m", "5. \033[93mIPV4 \033[92mWs + TLS\033[0m", "6. \033[93mIPV6 \033[92mWs + TLS\033[0m", "7. \033[92mStop | Restart Service\033[0m", "8. \033[91mUninstall\033[0m", "q. Exit"},
		
		}
		fmt.Println("\033[93m╰─────────────────────────────────────────────────────────────────────╯\033[0m")

		var choice string
		err = survey.AskOne(prompt, &choice)
		if err != nil {
			log.Fatalf("\033[91muser input is wrong:\033[0m %v", err)
		}
		switch choice {
		case "0. \033[91mSTATUS Menu\033[0m":
			status()
		case "1. \033[96mIPV4 \033[92mTCP \033[0m":
			tcp4Menu()
		case "2. \033[93mIPV4 \033[92mUDP\033[0m":
			udp4Menu()
		case "3. \033[96mIPV6 \033[92mTCP\033[0m":
			tcp6Menu()
		case "4. \033[93mIPV6 \033[92mUDP\033[0m":
			udp6Menu()
		case "5. \033[93mIPV4 \033[92mWs + TLS\033[0m":
			ws4Menu()
		case "6. \033[93mIPV6 \033[92mWs + TLS\033[0m":
			ws6Menu()
		case "7. \033[92mStop | Restart Service\033[0m":
			startMain()
		case "8. \033[91mUninstall\033[0m":
			UniMenu()
		case "q. Exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice.")
		}

		
		readInput()
	}
}
func rmv() error {
	file := "/etc/rat.sh"
	if _, err := os.Stat(file); err == nil {
		err := os.Remove(file)
		if err != nil {
			return fmt.Errorf("\033[91mbash file doesn't exists:\033[0m %v", err)
		}
		fmt.Println("\033[91mbash file removed successfully!\033[0m")
	}
	return nil
}
func deleteCron() {
	entriesToDelete := []string{
		"0 */1 * * * /etc/rat.sh",
		"0 */2 * * * /etc/rat.sh",
		"0 */3 * * * /etc/rat.sh",
		"0 */4 * * * /etc/rat.sh",
		"0 */5 * * * /etc/rat.sh",
		"0 */6 * * * /etc/rat.sh",
		"0 */7 * * * /etc/rat.sh",
		"0 */8 * * * /etc/rat.sh",
		"0 */9 * * * /etc/rat.sh",
		"0 */10 * * * /etc/rat.sh",
		"0 */11 * * * /etc/rat.sh",
		"0 */12 * * * /etc/rat.sh",
		"0 */13 * * * /etc/rat.sh",
		"0 */14 * * * /etc/rat.sh",
		"0 */15 * * * /etc/rat.sh",
		"0 */16 * * * /etc/rat.sh",
		"0 */17 * * * /etc/rat.sh",
		"0 */18 * * * /etc/rat.sh",
		"0 */19 * * * /etc/rat.sh",
		"0 */20 * * * /etc/rat.sh",
		"0 */21 * * * /etc/rat.sh",
		"0 */22 * * * /etc/rat.sh",
		"0 */23 * * * /etc/rat.sh",
		"0 */24 * * * /etc/rat.sh",
	}

	existingCrontab, err := exec.Command("crontab", "-l").Output()
	if err != nil {
		fmt.Println("\033[91mNo existing cron found!\033[0m")
		return
	}

	newCrontab := string(existingCrontab)
	for _, entry := range entriesToDelete {
		if strings.Contains(newCrontab, entry) {
			newCrontab = strings.Replace(newCrontab, entry, "", -1)
		}
	}

	if newCrontab != string(existingCrontab) {
		cmd := exec.Command("crontab")
		cmd.Stdin = strings.NewReader(newCrontab)
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
		displayNotification("\033[92mDeleting Previous Crons..\033[0m")
	} else {
		fmt.Println("\033[91mCron doesn't exist, moving on..!\033[0m")
	}
}

const crontabFilePath = "/var/spool/cron/crontabs/root"

func resKharej() {
	deleteCron()
	if _, err := os.Stat("/etc/rat.sh"); err == nil {
		os.Remove("/etc/rat.sh")
	}

	file, err := os.Create("/etc/rat.sh")
	if err != nil {
		log.Fatalf("\033[91mbash creation error:\033[0m %v", err)
	}
	defer file.Close()

    file.WriteString("#!/bin/bash\n")
    file.WriteString("sudo systemctl daemon-reload\n")
    file.WriteString("sudo systemctl restart kharej-azumi\n")
	file.WriteString("sudo sync; echo 1 > /proc/sys/vm/drop_caches\n")
    file.WriteString("sudo journalctl --vacuum-size=1M\n")
	

	cmd := exec.Command("chmod", "+x", "/etc/rat.sh")
	if err := cmd.Run(); err != nil {
		log.Fatalf("\033[91mchmod cmd error:\033[0m %v", err)
	}

	fmt.Println("╭──────────────────────────────────────╮")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\033[93mEnter \033[92mReset timer\033[93m (hours):\033[0m ")
	hoursStr, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	hoursStr = strings.TrimSpace(hoursStr)
	fmt.Println("╰──────────────────────────────────────╯")

	hours, err := strconv.Atoi(hoursStr)
	if err != nil {
		log.Fatalf("\033[91mInvalid input for reset timer:\033[0m %v", err)
	}

	cronEntry := fmt.Sprintf("0 */%d * * * /etc/rat.sh", hours)

	crontabFile, err := os.OpenFile(crontabFilePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("\033[91mCouldn't open Cron:\033[0m %v", err)
	}
	defer crontabFile.Close()

	var crontabContent strings.Builder
	scanner := bufio.NewScanner(crontabFile)
	for scanner.Scan() {
		line := scanner.Text()
		if line == cronEntry {
			fmt.Println("\033[92mOh .. Cron entry already exists!\033[0m")
			return
		}
		crontabContent.WriteString(line)
		crontabContent.WriteString("\n")
	}

	crontabContent.WriteString(cronEntry)
	crontabContent.WriteString("\n")

	if err := scanner.Err(); err != nil {
		log.Fatalf("\033[91mcrontab Reading error:\033[0m %v", err)
	}

	if err := crontabFile.Truncate(0); err != nil {
		log.Fatalf("\033[91mcouldn't truncate cron file:\033[0m %v", err)
	}

	if _, err := crontabFile.Seek(0, 0); err != nil {
		log.Fatalf("\033[91mcouldn't find cron file: \033[0m%v", err)
	}

	if _, err := crontabFile.WriteString(crontabContent.String()); err != nil {
		log.Fatalf("\033[91mCouldn't write cron file:\033[0m %v", err)
	}

	fmt.Println("\033[92mCron entry added successfully!\033[0m")
}
func resIran() {
	deleteCron()
	if _, err := os.Stat("/etc/rat.sh"); err == nil {
		os.Remove("/etc/rat.sh")
	}

	file, err := os.Create("/etc/rat.sh")
	if err != nil {
		log.Fatalf("\033[91mbash creation error:\033[0m %v", err)
	}
	defer file.Close()

    file.WriteString("#!/bin/bash\n")
    file.WriteString("sudo systemctl daemon-reload\n")
    file.WriteString("sudo systemctl restart iran-azumi\n")
    file.WriteString("sudo sync; echo 1 > /proc/sys/vm/drop_caches\n")
    file.WriteString("sudo journalctl --vacuum-size=1M\n")

	cmd := exec.Command("chmod", "+x", "/etc/rat.sh")
	if err := cmd.Run(); err != nil {
		log.Fatalf("\033[91mchmod cmd error:\033[0m %v", err)
	}

	fmt.Println("╭──────────────────────────────────────╮")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\033[93mEnter \033[92mReset timer\033[93m (hours):\033[0m ")
	hoursStr, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	hoursStr = strings.TrimSpace(hoursStr)
	fmt.Println("╰──────────────────────────────────────╯")

	hours, err := strconv.Atoi(hoursStr)
	if err != nil {
		log.Fatalf("\033[91mInvalid input for reset timer:\033[0m %v", err)
	}

	cronEntry := fmt.Sprintf("0 */%d * * * /etc/rat.sh", hours)

	crontabFile, err := os.OpenFile(crontabFilePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("\033[91mCouldn't open Cron:\033[0m %v", err)
	}
	defer crontabFile.Close()

	var crontabContent strings.Builder
	scanner := bufio.NewScanner(crontabFile)
	for scanner.Scan() {
		line := scanner.Text()
		if line == cronEntry {
			fmt.Println("\033[92mOh .. Cron entry already exists!\033[0m")
			return
		}
		crontabContent.WriteString(line)
		crontabContent.WriteString("\n")
	}

	crontabContent.WriteString(cronEntry)
	crontabContent.WriteString("\n")

	if err := scanner.Err(); err != nil {
		log.Fatalf("\033[91mcrontab Reading error:\033[0m %v", err)
	}

	if err := crontabFile.Truncate(0); err != nil {
		log.Fatalf("\033[91mcouldn't truncate cron file:\033[0m %v", err)
	}

	if _, err := crontabFile.Seek(0, 0); err != nil {
		log.Fatalf("\033[91mcouldn't find cron file: \033[0m%v", err)
	}

	if _, err := crontabFile.WriteString(crontabContent.String()); err != nil {
		log.Fatalf("\033[91mCouldn't write cron file:\033[0m %v", err)
	}

	fmt.Println("\033[92mCron entry added successfully!\033[0m")
}
func startMain() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[92m Service \033[93mMenu\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. \033[92mRestart\033[0m", "2. \033[93mStop \033[0m", "0. \033[94mBack to the main menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. \033[92mRestart\033[0m":
		start()
	case "2. \033[93mStop \033[0m":
		stop()
	case "0. \033[94mBack to the main menu\033[0m":
	    clearScreen()
		mainMenu()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}
func start() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[92m Restart \033[93mMenu\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. \033[92mTCP + UDP\033[0m", "2. \033[93mWS +TLS \033[0m", "0. \033[94mBack to the previous menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. \033[92mTCP + UDP\033[0m":
		restarttcp()
	case "2. \033[93mWS +TLS \033[0m":
		restarttcp()
	case "0. \033[94mBack to the previous menu\033[0m":
	    clearScreen()
		startMain()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}
func restarttcp() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	displayNotification("\033[93mRestarting Reverse Tunnel \033[93m..\033[0m")
	fmt.Println("\033[93m╭─────────────────────────────────────────────╮\033[0m")

	cmd = exec.Command("systemctl", "restart", "kharej-azumi")
	cmd.Run()
	time.Sleep(1 * time.Second)

	cmd = exec.Command("systemctl", "restart", "iran-azumi")
	cmd.Run()
	time.Sleep(1 * time.Second)

	fmt.Print("Progress: ")

	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	delay := 0.1
	duration := 1.0
	endTime := time.Now().Add(time.Duration(duration) * time.Second)

	for time.Now().Before(endTime) {
		for _, frame := range frames {
			fmt.Printf("\r[%s] Loading...  ", frame)
			time.Sleep(time.Duration(delay * float64(time.Second)))
			fmt.Printf("\r[%s]             ", frame)
			time.Sleep(time.Duration(delay * float64(time.Second)))
		}
	}

	displayCheckmark("\033[92mRestart completed!\033[0m")
}
func stop() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[92m Stop \033[93mMenu\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. \033[92mTCP + UDP\033[0m", "2. \033[93mWS +TLS \033[0m", "0. \033[94mBack to the previous menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. \033[92mTCP + UDP\033[0m":
		stoptcp()
	case "2. \033[93mWS +TLS \033[0m":
		stoptcp()
	case "0. \033[94mBack to the previous menu\033[0m":
	    clearScreen()
		startMain()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}
func stoptcp() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	displayNotification("\033[93mStopping Reverse Tunnel \033[93m..\033[0m")
	fmt.Println("\033[93m╭─────────────────────────────────────────────╮\033[0m")

	cmd = exec.Command("systemctl", "stop", "kharej-azumi")
	cmd.Run()
	time.Sleep(1 * time.Second)

	cmd = exec.Command("systemctl", "stop", "iran-azumi")
	cmd.Run()
	time.Sleep(1 * time.Second)

	fmt.Print("Progress: ")

	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	delay := 0.1
	duration := 1.0
	endTime := time.Now().Add(time.Duration(duration) * time.Second)

	for time.Now().Before(endTime) {
		for _, frame := range frames {
			fmt.Printf("\r[%s] Loading...  ", frame)
			time.Sleep(time.Duration(delay * float64(time.Second)))
			fmt.Printf("\r[%s]             ", frame)
			time.Sleep(time.Duration(delay * float64(time.Second)))
		}
	}

	displayCheckmark("\033[92mService Stopped!\033[0m")
}
func status() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[92m Status \033[93mMenu\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. \033[92mTCP + UDP\033[0m", "2. \033[93mWS + TLS \033[0m", "0. \033[94mBack to the main menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. \033[92mTCP + UDP\033[0m":
		tcpStatus()
	case "2. \033[93mWS + TLS \033[0m":
		tcpStatus()
	case "0. \033[94mBack to the main menu\033[0m":
	    clearScreen()
		mainMenu()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}
func tcpStatus() {
	services := []string{"iran-azumi", "kharej-azumi"}

	fmt.Println("\033[93m            ╔════════════════════════════════════════════╗\033[0m")
	fmt.Println("\033[93m            ║               \033[92mReverse Status\033[93m               ║\033[0m")
	fmt.Println("\033[93m            ╠════════════════════════════════════════════╣\033[0m")

	for _, service := range services {
		cmd := exec.Command("systemctl", "is-active", "--quiet", service)
		err := cmd.Run()
		if err != nil {
			continue
		}

		status := "\033[92m✓ Active      \033[0m"
		displayName := ""
		switch service {
		case "iran-azumi":
			displayName = "\033[93mIRAN Server   \033[0m"
		case "kharej-azumi":
			displayName = "\033[93mKharej Server \033[0m"
		default:
			displayName = service
		}

		fmt.Printf("           \033[93m ║\033[0m    %s   |    %s\033[93m    ║\033[0m\n", displayName, status)
	}

	fmt.Println("\033[93m            ╚════════════════════════════════════════════╝\033[0m")
}
func UniMenu() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Uninstallation \033[96mMenu\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. \033[92mTCP + UDP\033[0m", "2. \033[93mWS + TLS \033[0m", "0. \033[94mBack to the main menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. \033[92mTCP + UDP\033[0m":
		removews()
	case "2. \033[93mWS + TLS \033[0m":
		removews()
	case "0. \033[94mBack to the main menu\033[0m":
	    clearScreen()
		mainMenu()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}
func removews() {
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("\033[93mRemoving Config ..\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	deleteCron()
	rmv()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\033[93mDo you want to remove Rathole project as well? (\033[92myes\033[93m/\033[91mno\033[93m):\033[0m ")
	scanner.Scan()
	removeRathole := strings.TrimSpace(scanner.Text())

	if removeRathole == "yes" || removeRathole == "y" {
		if _, err := os.Stat("/root/rathole"); err == nil {
			os.RemoveAll("/root/rathole")
		}
		fmt.Println("\033[92mRathole project removed.\033[0m")
	} else if removeRathole == "no" || removeRathole == "n" {
		fmt.Println("\033[91mSkipping..\033[0m")
	} else {
		fmt.Println("\033[91minvalid input.\033[0m")
	}

	azumiServices := []string{
		"iran-azumi", "kharej-azumi",
	}

	for _, serviceName := range azumiServices {
		hideCmd("systemctl", "disable", serviceName+".service")
		hideCmd("systemctl", "stop", serviceName+".service")
		hideCmd("rm", "/etc/systemd/system/"+serviceName+".service")
	}

	runCmd("systemctl", "daemon-reload")

	fmt.Print("Progress: ")

	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	delay := 100 * time.Millisecond
	duration := 1 * time.Second
	endTime := time.Now().Add(duration)

	for time.Now().Before(endTime) {
		for _, frame := range frames {
			fmt.Printf("\r[%s] Loading...  ", frame)
			time.Sleep(delay)
			fmt.Printf("\r[%s]             ", frame)
			time.Sleep(delay)
		}
	}

	displayCheckmark("\033[92m Uninstallation completed!\033[0m")
}
func runCmd(cmd string, args ...string) {
	command := exec.Command(cmd, args...)
	err := command.Run()
	if err != nil {
		fmt.Printf("\033[91mCouldn't run cmd: %s, %v\n\033[0m", cmd, err)
	}
}
func hideCmd(cmd string, args ...string) error {
	command := exec.Command(cmd, args...)

	nullDevice, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	command.Stdout = nullDevice
	command.Stderr = nullDevice

	return command.Run()
}
func tcp4Menu() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Reverse \033[92mTCP \033[96mIPV4 \033[93mMenu\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. \033[92mIRAN\033[0m", "2. \033[93mKHAREJ\033[92m[1]\033[0m", "3. \033[93mKHAREJ\033[92m[2]\033[0m", "4. \033[93mKHAREJ\033[92m[3]\033[0m", "5. \033[93mKHAREJ\033[92m[4]\033[0m", "6. \033[93mKHAREJ\033[92m[5]\033[0m", "7. \033[93mKHAREJ\033[92m[6]\033[0m", "8. \033[93mKHAREJ\033[92m[7]\033[0m", "9. \033[93mKHAREJ\033[92m[8]\033[0m", "10. \033[93mKHAREJ\033[92m[9]\033[0m", "11. \033[93mKHAREJ\033[92m[10]\033[0m", "0. \033[94mBack to the main menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. \033[92mIRAN\033[0m":
		iranTcp4()
	case "2. \033[93mKHAREJ\033[92m[1]\033[0m":
		kharejTcp4()
	case "3. \033[93mKHAREJ\033[92m[2]\033[0m":
		kharej2Tcp4()
    case "4. \033[93mKHAREJ\033[92m[3]\033[0m":
		kharej2Tcp4()
	case "5. \033[93mKHAREJ\033[92m[4]\033[0m":
		kharej2Tcp4()
	case "6. \033[93mKHAREJ\033[92m[5]\033[0m":
		kharej2Tcp4()
	case "7. \033[93mKHAREJ\033[92m[6]\033[0m":
		kharej2Tcp4()
	case "8. \033[93mKHAREJ\033[92m[7]\033[0m":
		kharej2Tcp4()
	case "9. \033[93mKHAREJ\033[92m[8]\033[0m":
		kharej2Tcp4()
	case "10. \033[93mKHAREJ\033[92m[9]\033[0m":
		kharej2Tcp4()
	case "11. \033[93mKHAREJ\033[92m[10]\033[0m":
		kharej2Tcp4()
	case "0. \033[94mBack to the main menu\033[0m":
	    clearScreen()
		mainMenu()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}

func iranTcp4() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Reverse \033[92mIPV4 \033[96mTCP\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring IRAN")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\033[93mHow many \033[92mconfigs\033[93m do you have \033[96m[All Servers Combined]\033[93m? \033[0m")
	scanner.Scan()
	numConfigsStr := scanner.Text()

	numConfigs, err := strconv.Atoi(numConfigsStr)
	if err != nil {
		fmt.Println("\033[91mPlease enter a valid number\033[0m")
		return
	}

	fmt.Print("\033[93mEnter \033[92mTunnel port:\033[0m ")
	scanner.Scan()
	tunnelPort := scanner.Text()

	kharejPorts := make([]string, numConfigs)
	for i := 0; i < numConfigs; i++ {
		fmt.Printf("\033[93mEnter \033[92mConfig %d\033[93m Port: \033[0m", i+1)
		scanner.Scan()
		kharejPorts[i] = scanner.Text()
	}

	fmt.Print("\033[93mDo you want nodelay enabled? (\033[92my/\033[91mn\033[93m): \033[0m")
	scanner.Scan()
	nodelayOp := scanner.Text()
	nodelay := "false"
	if strings.ToLower(nodelayOp) == "y" {
		nodelay = "true"
	}

	server := fmt.Sprintf(`[server]
bind_addr = "0.0.0.0:%s"
default_token = "azumiisinyourarea"

[server.transport]
type = "tcp"

[server.transport.tcp]
nodelay = %s
keepalive_secs = 20
keepalive_interval = 8

`, tunnelPort, nodelay)
	for i := 0; i < numConfigs; i++ {
		config := fmt.Sprintf(`[server.services.kharej%d]
type = "tcp"
bind_addr = "0.0.0.0:%s" 
`, i+1, kharejPorts[i])
		server += config
	}

	err = os.Remove("/root/rathole/server.toml")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting toml:\033[0m", err)
		return
	}

	file, err := os.Create("/root/rathole/server.toml")
	if err != nil {
		fmt.Println("\033[91merror creating toml:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(server)
	if err != nil {
		fmt.Println("\033[91merror putting configs into toml:\033[0m", err)
		return
	}
	service := `[Unit]
Description=Iran-Azumi Service
After=network.target

[Service]
Type=simple
Restart=on-failure
RestartSec=5s
LimitNOFILE=1048576
ExecStart=/root/rathole/target/debug/rathole /root/rathole/server.toml

[Install]
WantedBy=multi-user.target`

	err = os.Remove("/etc/systemd/system/iran-azumi.service")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting iran-azumi:\033[0m", err)
		return
	}

	file, err = os.Create("/etc/systemd/system/iran-azumi.service")
	if err != nil {
		fmt.Println("\033[91merror creating iran-azumi:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(service)
	if err != nil {
		fmt.Println("\033[91merror constructing iran-azumi:\033[0m", err)
		return
	}

	cmd := exec.Command("systemctl", "daemon-reload")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror reloading:\033[0m", err)
		return
	}

	cmd = exec.Command("sudo", "chmod", "u+x", "/etc/systemd/system/iran-azumi.service")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enablin da service:\033[0m", err)
		return
	}

	cmd = exec.Command("systemctl", "enable", "iran-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enabling da service:\033[0m", err)
		return
	}

	cmd = exec.Command("systemctl", "restart", "iran-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror restarting da service:\033[0m", err)
		return
	}
    resIran()
	displayCheckmark("\033[92mService created successfully!\033[0m")
}
func kharejTcp4() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Reverse \033[92mIPV4 \033[96mTCP\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring KHAREJ")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\033[93mEnter the \033[92mStarting number:\033[0m ")
	scanner.Scan()
	startingNumberStr := scanner.Text()

	startingNumber, err := strconv.Atoi(startingNumberStr)
	if err != nil {
		fmt.Println("\033[91mPlz enter a valid number\033[0m")
		return
	}

	fmt.Print("\033[93mEnter \033[92mIran IPV4:\033[0m ")
	scanner.Scan()
	iranIP := scanner.Text()

	fmt.Print("\033[93mEnter \033[92mTunnel port:\033[0m ")
	scanner.Scan()
	tunnelPort := scanner.Text()

	fmt.Print("\033[93mHow many \033[92mConfigs\033[93m do you have?\033[0m ")
	scanner.Scan()
	numConfigsStr := scanner.Text()

	numConfigs, err := strconv.Atoi(numConfigsStr)
	if err != nil {
		fmt.Println("\033[91mPlz enter a valid number\033[0m")
		return
	}

	kharejPorts := make([]string, numConfigs)
	for i := 0; i < numConfigs; i++ {
		fmt.Printf("\033[93mEnter \033[92mconfig %d\033[93m port:\033[0m ", i+1)
		scanner.Scan()
		kharejPorts[i] = scanner.Text()
	}

	fmt.Print("\033[93mDo you want nodelay enabled? (\033[92my/\033[91mn\033[93m): \033[0m")
	scanner.Scan()
	nodelayOp := scanner.Text()
	nodelay := "false"
	if strings.ToLower(nodelayOp) == "y" {
		nodelay = "true"
	}

	client := fmt.Sprintf(`[client]
remote_addr = "%s:%s"
default_token = "azumiisinyourarea"
retry_interval = 1

[client.transport]
type = "tcp"

[client.transport.tcp]
nodelay = %s
keepalive_secs = 20
keepalive_interval = 8
`, iranIP, tunnelPort, nodelay)

	for i := 0; i < numConfigs; i++ {
		config := fmt.Sprintf(`[client.services.kharej%d]
type = "tcp"
local_addr = "127.0.0.1:%s"
`, i+startingNumber, kharejPorts[i])
		client += config
	}

	err = os.Remove("/root/rathole/client.toml")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting toml:\033[0m", err)
		return
	}

	file, err := os.Create("/root/rathole/client.toml")
	if err != nil {
		fmt.Println("\033[91merror creating toml:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(client)
	if err != nil {
		fmt.Println("\033[91merror putting configs into toml:\033[0m", err)
		return
	}

	service := `[Unit]
Description=Kharej-Azumi Service
After=network.target

[Service]
Type=simple
Restart=on-failure
RestartSec=5s
LimitNOFILE=1048576
ExecStart=/root/rathole/target/debug/rathole /root/rathole/client.toml

[Install]
WantedBy=multi-user.target`

	err = os.Remove("/etc/systemd/system/kharej-azumi.service")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting iran-azumi:\033[0m", err)
		return
	}

	file, err = os.Create("/etc/systemd/system/kharej-azumi.service")
	if err != nil {
		fmt.Println("\033[91merror creating iran-azumi:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(service)
	if err != nil {
		fmt.Println("\033[91merror constructing iran-azumi:\033[0m", err)
		return
	}

	cmd := exec.Command("systemctl", "daemon-reload")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror reloading:\033[0m", err)
		return
	}

	cmd = exec.Command("sudo", "chmod", "u+x", "/etc/systemd/system/kharej-azumi.service")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enablin da service:\033[0m", err)
		return
	}

	cmd = exec.Command("systemctl", "enable", "kharej-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enabling da service:\033[0m", err)
		return
	}

	cmd = exec.Command("systemctl", "restart", "kharej-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror restarting da service:\033[0m", err)
		return
	}
	resKharej()
	displayCheckmark("\033[92mService created successfully!\033[0m")
	fmt.Println("╭─────────────────────────────────────────────╮")
	fmt.Printf("\033[92m Starting number for the next server : \033[96m%-9d\n\033[0m", numConfigs+1)
	fmt.Println("╰─────────────────────────────────────────────╯")
}
func kharej2Tcp4() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Reverse \033[92mIPV4 \033[96mTCP\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring KHAREJ")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\033[93mEnter the \033[92mStarting number:\033[0m ")
	scanner.Scan()
	startingNumberStr := scanner.Text()

	startingNumber, err := strconv.Atoi(startingNumberStr)
	if err != nil {
		fmt.Println("\033[91mPlz enter a valid number\033[0m")
		return
	}

	fmt.Print("\033[93mEnter \033[92mIran IPV4:\033[0m ")
	scanner.Scan()
	iranIP := scanner.Text()

	fmt.Print("\033[93mEnter \033[92mTunnel port:\033[0m ")
	scanner.Scan()
	tunnelPort := scanner.Text()

	fmt.Print("\033[93mHow many \033[92mConfigs\033[93m do you have?\033[0m ")
	scanner.Scan()
	numConfigsStr := scanner.Text()

	numConfigs, err := strconv.Atoi(numConfigsStr)
	if err != nil {
		fmt.Println("\033[91mPlz enter a valid number\033[0m")
		return
	}

	kharejPorts := make([]string, numConfigs)
	for i := 0; i < numConfigs; i++ {
		fmt.Printf("\033[93mEnter \033[92mconfig %d\033[93m port:\033[0m ", i+1)
		scanner.Scan()
		kharejPorts[i] = scanner.Text()
	}

	fmt.Print("\033[93mDo you want nodelay enabled? (\033[92my/\033[91mn\033[93m): \033[0m")
	scanner.Scan()
	nodelayOp := scanner.Text()
	nodelay := "false"
	if strings.ToLower(nodelayOp) == "y" {
		nodelay = "true"
	}

	client := fmt.Sprintf(`[client]
remote_addr = "%s:%s"
default_token = "azumiisinyourarea"
retry_interval = 1

[client.transport]
type = "tcp"

[client.transport.tcp]
nodelay = %s
keepalive_secs = 20
keepalive_interval = 8
`, iranIP, tunnelPort, nodelay)

	for i := 0; i < numConfigs; i++ {
		config := fmt.Sprintf(`[client.services.kharej%d]
type = "tcp"
local_addr = "127.0.0.1:%s"
`, i+startingNumber, kharejPorts[i])
		client += config
	}

	err = os.Remove("/root/rathole/client.toml")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting toml:\033[0m", err)
		return
	}

	file, err := os.Create("/root/rathole/client.toml")
	if err != nil {
		fmt.Println("\033[91merror creating toml:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(client)
	if err != nil {
		fmt.Println("\033[91merror putting configs into toml:\033[0m", err)
		return
	}

	service := `[Unit]
Description=Kharej-Azumi Service
After=network.target

[Service]
Type=simple
Restart=on-failure
RestartSec=5s
LimitNOFILE=1048576
ExecStart=/root/rathole/target/debug/rathole /root/rathole/client.toml

[Install]
WantedBy=multi-user.target`

	err = os.Remove("/etc/systemd/system/kharej-azumi.service")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting iran-azumi:\033[0m", err)
		return
	}

	file, err = os.Create("/etc/systemd/system/kharej-azumi.service")
	if err != nil {
		fmt.Println("\033[91merror creating iran-azumi:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(service)
	if err != nil {
		fmt.Println("\033[91merror constructing iran-azumi:\033[0m", err)
		return
	}

	cmd := exec.Command("systemctl", "daemon-reload")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror reloading:\033[0m", err)
		return
	}

	cmd = exec.Command("sudo", "chmod", "u+x", "/etc/systemd/system/kharej-azumi.service")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enablin da service:\033[0m", err)
		return
	}

	cmd = exec.Command("systemctl", "enable", "kharej-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enabling da service:\033[0m", err)
		return
	}

	cmd = exec.Command("systemctl", "restart", "kharej-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror restarting da service:\033[0m", err)
		return
	}
	resKharej()
	displayCheckmark("\033[92mService created successfully!\033[0m")
	if numConfigs == 1 {
	    fmt.Println("╭─────────────────────────────────────────────╮")
		fmt.Printf("\033[92m  Starting number for the next server:\033[96m %d\n\033[0m", startingNumber+1)
		fmt.Println("╰─────────────────────────────────────────────╯")
	} else {
	    fmt.Println("╭─────────────────────────────────────────────╮")
		fmt.Printf("\033[92m  Starting number for the next server:\033[96m %d\n\033[0m", numConfigs+startingNumber)
		fmt.Println("╰─────────────────────────────────────────────╯")
	}
}
func udp4Menu() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Reverse \033[92mUDP \033[96mIPV4 \033[93mMenu\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. \033[92mIRAN\033[0m", "2. \033[93mKHAREJ\033[92m[1]\033[0m", "3. \033[93mKHAREJ\033[92m[2]\033[0m", "4. \033[93mKHAREJ\033[92m[3]\033[0m", "5. \033[93mKHAREJ\033[92m[4]\033[0m", "6. \033[93mKHAREJ\033[92m[5]\033[0m", "7. \033[93mKHAREJ\033[92m[6]\033[0m", "8. \033[93mKHAREJ\033[92m[7]\033[0m", "9. \033[93mKHAREJ\033[92m[8]\033[0m", "10. \033[93mKHAREJ\033[92m[9]\033[0m", "11. \033[93mKHAREJ\033[92m[10]\033[0m", "0. \033[94mBack to the main menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. \033[92mIRAN\033[0m":
		iranUdp4()
	case "2. \033[93mKHAREJ\033[92m[1]\033[0m":
		kharejUdp4()
	case "3. \033[93mKHAREJ\033[92m[2]\033[0m":
		kharej2Udp4()
    case "4. \033[93mKHAREJ\033[92m[3]\033[0m":
		kharej2Udp4()
	case "5. \033[93mKHAREJ\033[92m[4]\033[0m":
		kharej2Udp4()
	case "6. \033[93mKHAREJ\033[92m[5]\033[0m":
		kharej2Udp4()
	case "7. \033[93mKHAREJ\033[92m[6]\033[0m":
		kharej2Udp4()
	case "8. \033[93mKHAREJ\033[92m[7]\033[0m":
		kharej2Udp4()
	case "9. \033[93mKHAREJ\033[92m[8]\033[0m":
		kharej2Udp4()
	case "10. \033[93mKHAREJ\033[92m[9]\033[0m":
		kharej2Udp4()
	case "11. \033[93mKHAREJ\033[92m[10]\033[0m":
		kharej2Udp4()
	case "0. \033[94mBack to the main menu\033[0m":
	    clearScreen()
		mainMenu()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}

func iranUdp4() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Reverse \033[92mIPV4 \033[96mUDP\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring IRAN")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\033[93mHow many \033[92mconfigs\033[93m do you have \033[96m[All Servers Combined]\033[93m? \033[0m")
	scanner.Scan()
	numConfigsStr := scanner.Text()

	numConfigs, err := strconv.Atoi(numConfigsStr)
	if err != nil {
		fmt.Println("\033[91mPlz enter a valid number\033[0m")
		return
	}

	fmt.Print("\033[93mEnter \033[92mTunnel port:\033[0m ")
	scanner.Scan()
	tunnelPort := scanner.Text()

	kharejPorts := make([]string, numConfigs)
	for i := 0; i < numConfigs; i++ {
		fmt.Printf("\033[93mEnter \033[92mConfig %d\033[93m Port: \033[0m", i+1)
		scanner.Scan()
		kharejPorts[i] = scanner.Text()
	}

	fmt.Print("\033[93mDo you want nodelay enabled? (\033[92my/\033[91mn\033[93m): \033[0m")
	scanner.Scan()
	nodelayOp := scanner.Text()
	nodelay := "false"
	if strings.ToLower(nodelayOp) == "y" {
		nodelay = "true"
	}

	server := fmt.Sprintf(`[server]
bind_addr = "0.0.0.0:%s"
default_token = "azumiisinyourarea"

[server.transport]
type = "tcp"

[server.transport.tcp]
nodelay = %s
keepalive_secs = 20
keepalive_interval = 8

`, tunnelPort, nodelay)
	for i := 0; i < numConfigs; i++ {
		config := fmt.Sprintf(`[server.services.kharej%d]
type = "udp"
bind_addr = "0.0.0.0:%s" 
`, i+1, kharejPorts[i])
		server += config
	}

	err = os.Remove("/root/rathole/server.toml")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting toml:\033[0m", err)
		return
	}

	file, err := os.Create("/root/rathole/server.toml")
	if err != nil {
		fmt.Println("\033[91merror creating toml:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(server)
	if err != nil {
		fmt.Println("\033[91merror putting configs into toml:\033[0m", err)
		return
	}
	service := `[Unit]
Description=Rathole Service
After=network.target

[Service]
Type=simple
Restart=on-failure
RestartSec=5s
LimitNOFILE=1048576
ExecStart=/root/rathole/target/debug/rathole /root/rathole/server.toml

[Install]
WantedBy=multi-user.target`

	err = os.Remove("/etc/systemd/system/iran-azumi.service")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting iran-azumi:\033[0m", err)
		return
	}

	file, err = os.Create("/etc/systemd/system/iran-azumi.service")
	if err != nil {
		fmt.Println("\033[91merror creating iran-azumi:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(service)
	if err != nil {
		fmt.Println("\033[91merror constructing iran-azumi:\033[0m", err)
		return
	}

	cmd := exec.Command("systemctl", "daemon-reload")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror reloading:\033[0m", err)
		return
	}

	cmd = exec.Command("sudo", "chmod", "u+x", "/etc/systemd/system/iran-azumi.service")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enablin da service:\033[0m", err)
		return
	}

	cmd = exec.Command("systemctl", "enable", "iran-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enabling da service:\033[0m", err)
		return
	}

	cmd = exec.Command("systemctl", "restart", "iran-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror restarting da service:\033[0m", err)
		return
	}
    resIran()
	displayCheckmark("\033[92mService created successfully!\033[0m")
}
func kharejUdp4() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Reverse \033[92mIPV4 \033[96mUDP\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring KHAREJ")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\033[93mEnter the \033[92mStarting number:\033[0m ")
	scanner.Scan()
	startingNumberStr := scanner.Text()

	startingNumber, err := strconv.Atoi(startingNumberStr)
	if err != nil {
		fmt.Println("\033[91mPlz enter a valid number\033[0m")
		return
	}

	fmt.Print("\033[93mEnter \033[92mIran IPV4:\033[0m ")
	scanner.Scan()
	iranIP := scanner.Text()

	fmt.Print("\033[93mEnter \033[92mTunnel port:\033[0m ")
	scanner.Scan()
	tunnelPort := scanner.Text()

	fmt.Print("\033[93mHow many \033[92mConfigs\033[93m do you have?\033[0m ")
	scanner.Scan()
	numConfigsStr := scanner.Text()

	numConfigs, err := strconv.Atoi(numConfigsStr)
	if err != nil {
		fmt.Println("\033[91mPlz enter a valid number\033[0m")
		return
	}

	kharejPorts := make([]string, numConfigs)
	for i := 0; i < numConfigs; i++ {
		fmt.Printf("\033[93mEnter \033[92mconfig %d\033[93m port:\033[0m ", i+1)
		scanner.Scan()
		kharejPorts[i] = scanner.Text()
	}

	fmt.Print("\033[93mDo you want nodelay enabled? (\033[92my/\033[91mn\033[93m): \033[0m")
	scanner.Scan()
	nodelayOp := scanner.Text()
	nodelay := "false"
	if strings.ToLower(nodelayOp) == "y" {
		nodelay = "true"
	}

	client := fmt.Sprintf(`[client]
remote_addr = "%s:%s"
default_token = "azumiisinyourarea"
retry_interval = 1

[client.transport]
type = "tcp"

[client.transport.tcp]
nodelay = %s
keepalive_secs = 20
keepalive_interval = 8
`, iranIP, tunnelPort, nodelay)

	for i := 0; i < numConfigs; i++ {
		config := fmt.Sprintf(`[client.services.kharej%d]
type = "udp"
local_addr = "127.0.0.1:%s"
`, i+startingNumber, kharejPorts[i])
		client += config
	}

	err = os.Remove("/root/rathole/client.toml")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting toml:\033[0m", err)
		return
	}

	file, err := os.Create("/root/rathole/client.toml")
	if err != nil {
		fmt.Println("\033[91merror creating toml:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(client)
	if err != nil {
		fmt.Println("\033[91merror putting configs into toml:\033[0m", err)
		return
	}

	service := `[Unit]
Description=Kharej-Azumi Service
After=network.target

[Service]
Type=simple
Restart=on-failure
RestartSec=5s
LimitNOFILE=1048576
ExecStart=/root/rathole/target/debug/rathole /root/rathole/client.toml

[Install]
WantedBy=multi-user.target`

	err = os.Remove("/etc/systemd/system/kharej-azumi.service")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting iran-azumi:\033[0m", err)
		return
	}

	file, err = os.Create("/etc/systemd/system/kharej-azumi.service")
	if err != nil {
		fmt.Println("\033[91merror creating iran-azumi:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(service)
	if err != nil {
		fmt.Println("\033[91merror constructing iran-azumi:\033[0m", err)
		return
	}

	cmd := exec.Command("systemctl", "daemon-reload")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror reloading:\033[0m", err)
		return
	}

	cmd = exec.Command("sudo", "chmod", "u+x", "/etc/systemd/system/kharej-azumi.service")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enablin da service:\033[0m", err)
		return
	}

	cmd = exec.Command("systemctl", "enable", "kharej-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enabling da service:\033[0m", err)
		return
	}

	cmd = exec.Command("systemctl", "restart", "kharej-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror restarting da service:\033[0m", err)
		return
	}
	resKharej()
	displayCheckmark("\033[92mService created successfully!\033[0m")
	fmt.Println("╭─────────────────────────────────────────────╮")
	fmt.Printf("\033[92m Starting number for the next server : \033[96m%-9d\n\033[0m", numConfigs+1)
	fmt.Println("╰─────────────────────────────────────────────╯")
}
func kharej2Udp4() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Reverse \033[92mIPV4 \033[96mUDP\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring KHAREJ")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\033[93mEnter the \033[92mStarting number:\033[0m ")
	scanner.Scan()
	startingNumberStr := scanner.Text()

	startingNumber, err := strconv.Atoi(startingNumberStr)
	if err != nil {
		fmt.Println("\033[91mPlz enter a valid number\033[0m")
		return
	}

	fmt.Print("\033[93mEnter \033[92mIran IPV4:\033[0m ")
	scanner.Scan()
	iranIP := scanner.Text()

	fmt.Print("\033[93mEnter \033[92mTunnel port:\033[0m ")
	scanner.Scan()
	tunnelPort := scanner.Text()

	fmt.Print("\033[93mHow many \033[92mConfigs\033[93m do you have?\033[0m ")
	scanner.Scan()
	numConfigsStr := scanner.Text()

	numConfigs, err := strconv.Atoi(numConfigsStr)
	if err != nil {
		fmt.Println("\033[91mPlz enter a valid number\033[0m")
		return
	}

	kharejPorts := make([]string, numConfigs)
	for i := 0; i < numConfigs; i++ {
		fmt.Printf("\033[93mEnter \033[92mconfig %d\033[93m port:\033[0m ", i+1)
		scanner.Scan()
		kharejPorts[i] = scanner.Text()
	}

	fmt.Print("\033[93mDo you want nodelay enabled? (\033[92my/\033[91mn\033[93m): \033[0m")
	scanner.Scan()
	nodelayOp := scanner.Text()
	nodelay := "false"
	if strings.ToLower(nodelayOp) == "y" {
		nodelay = "true"
	}

	client := fmt.Sprintf(`[client]
remote_addr = "%s:%s"
default_token = "azumiisinyourarea"
retry_interval = 1

[client.transport]
type = "tcp"

[client.transport.tcp]
nodelay = %s
keepalive_secs = 20
keepalive_interval = 8
`, iranIP, tunnelPort, nodelay)

	for i := 0; i < numConfigs; i++ {
		config := fmt.Sprintf(`[client.services.kharej%d]
type = "udp"
local_addr = "127.0.0.1:%s"
`, i+startingNumber, kharejPorts[i])
		client += config
	}

	err = os.Remove("/root/rathole/client.toml")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting toml:\033[0m", err)
		return
	}

	file, err := os.Create("/root/rathole/client.toml")
	if err != nil {
		fmt.Println("\033[91merror creating toml:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(client)
	if err != nil {
		fmt.Println("\033[91merror putting configs into toml:\033[0m", err)
		return
	}

	service := `[Unit]
Description=Kharej-Azumi Service
After=network.target

[Service]
Type=simple
Restart=on-failure
RestartSec=5s
LimitNOFILE=1048576
ExecStart=/root/rathole/target/debug/rathole /root/rathole/client.toml

[Install]
WantedBy=multi-user.target`

	err = os.Remove("/etc/systemd/system/kharej-azumi.service")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting iran-azumi:\033[0m", err)
		return
	}

	file, err = os.Create("/etc/systemd/system/kharej-azumi.service")
	if err != nil {
		fmt.Println("\033[91merror creating iran-azumi:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(service)
	if err != nil {
		fmt.Println("\033[91merror constructing iran-azumi:\033[0m", err)
		return
	}

	cmd := exec.Command("systemctl", "daemon-reload")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror reloading:\033[0m", err)
		return
	}

	cmd = exec.Command("sudo", "chmod", "u+x", "/etc/systemd/system/kharej-azumi.service")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enablin da service:\033[0m", err)
		return
	}

	cmd = exec.Command("systemctl", "enable", "kharej-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enabling da service:\033[0m", err)
		return
	}

	cmd = exec.Command("systemctl", "restart", "kharej-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror restarting da service:\033[0m", err)
		return
	}
	resKharej()
	displayCheckmark("\033[92mService created successfully!\033[0m")
	if numConfigs == 1 {
	    fmt.Println("╭─────────────────────────────────────────────╮")
		fmt.Printf("\033[92m  Starting number for the next server:\033[96m %d\n\033[0m", startingNumber+1)
		fmt.Println("╰─────────────────────────────────────────────╯")
	} else {
	    fmt.Println("╭─────────────────────────────────────────────╮")
		fmt.Printf("\033[92m  Starting number for the next server:\033[96m %d\n\033[0m", numConfigs+startingNumber)
		fmt.Println("╰─────────────────────────────────────────────╯")
	}
}
func tcp6Menu() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Reverse \033[92mTCP \033[96mIPV6 \033[93mMenu\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. \033[92mIRAN\033[0m", "2. \033[93mKHAREJ\033[92m[1]\033[0m", "3. \033[93mKHAREJ\033[92m[2]\033[0m", "4. \033[93mKHAREJ\033[92m[3]\033[0m", "5. \033[93mKHAREJ\033[92m[4]\033[0m", "6. \033[93mKHAREJ\033[92m[5]\033[0m", "7. \033[93mKHAREJ\033[92m[6]\033[0m", "8. \033[93mKHAREJ\033[92m[7]\033[0m", "9. \033[93mKHAREJ\033[92m[8]\033[0m", "10. \033[93mKHAREJ\033[92m[9]\033[0m", "11. \033[93mKHAREJ\033[92m[10]\033[0m", "0. \033[94mBack to the main menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. \033[92mIRAN\033[0m":
		iranTcp6()
	case "2. \033[93mKHAREJ\033[92m[1]\033[0m":
		kharejTcp6()
	case "3. \033[93mKHAREJ\033[92m[2]\033[0m":
		kharej2Tcp6()
    case "4. \033[93mKHAREJ\033[92m[3]\033[0m":
		kharej2Tcp6()
	case "5. \033[93mKHAREJ\033[92m[4]\033[0m":
		kharej2Tcp6()
	case "6. \033[93mKHAREJ\033[92m[5]\033[0m":
		kharej2Tcp6()
	case "7. \033[93mKHAREJ\033[92m[6]\033[0m":
		kharej2Tcp6()
	case "8. \033[93mKHAREJ\033[92m[7]\033[0m":
		kharej2Tcp6()
	case "9. \033[93mKHAREJ\033[92m[8]\033[0m":
		kharej2Tcp6()
	case "10. \033[93mKHAREJ\033[92m[9]\033[0m":
		kharej2Tcp6()
	case "11. \033[93mKHAREJ\033[92m[10]\033[0m":
		kharej2Tcp6()
	case "0. \033[94mBack to the main menu\033[0m":
	    clearScreen()
		mainMenu()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}

func iranTcp6() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Reverse \033[92mIPV6 \033[96mTCP\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring IRAN")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\033[93mHow many \033[92mconfigs\033[93m do you have \033[96m[All Servers Combined]\033[93m? \033[0m")
	scanner.Scan()
	numConfigsStr := scanner.Text()

	numConfigs, err := strconv.Atoi(numConfigsStr)
	if err != nil {
		fmt.Println("\033[91mPlz enter a valid number\033[0m")
		return
	}

	fmt.Print("\033[93mEnter \033[92mTunnel port:\033[0m ")
	scanner.Scan()
	tunnelPort := scanner.Text()

	kharejPorts := make([]string, numConfigs)
	for i := 0; i < numConfigs; i++ {
		fmt.Printf("\033[93mEnter \033[92mConfig %d\033[93m Port: \033[0m", i+1)
		scanner.Scan()
		kharejPorts[i] = scanner.Text()
	}

	fmt.Print("\033[93mDo you want nodelay enabled? (\033[92my/\033[91mn\033[93m): \033[0m")
	scanner.Scan()
	nodelayOp := scanner.Text()
	nodelay := "false"
	if strings.ToLower(nodelayOp) == "y" {
		nodelay = "true"
	}

	server := fmt.Sprintf(`[server]
bind_addr = "[::]:%s"
default_token = "azumiisinyourarea"

[server.transport]
type = "tcp"

[server.transport.tcp]
nodelay = %s
keepalive_secs = 20
keepalive_interval = 8

`, tunnelPort, nodelay)
	for i := 0; i < numConfigs; i++ {
		config := fmt.Sprintf(`[server.services.kharej%d]
type = "tcp"
bind_addr = "0.0.0.0:%s" 
`, i+1, kharejPorts[i])
		server += config
	}

	err = os.Remove("/root/rathole/server.toml")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting toml:\033[0m", err)
		return
	}

	file, err := os.Create("/root/rathole/server.toml")
	if err != nil {
		fmt.Println("\033[91merror creating toml:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(server)
	if err != nil {
		fmt.Println("\033[91merror putting configs into toml:\033[0m", err)
		return
	}
	service := `[Unit]
Description=Iran-Azumi Service
After=network.target

[Service]
Type=simple
Restart=on-failure
RestartSec=5s
LimitNOFILE=1048576
ExecStart=/root/rathole/target/debug/rathole /root/rathole/server.toml

[Install]
WantedBy=multi-user.target`

	err = os.Remove("/etc/systemd/system/iran-azumi.service")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting iran-azumi:\033[0m", err)
		return
	}

	file, err = os.Create("/etc/systemd/system/iran-azumi.service")
	if err != nil {
		fmt.Println("\033[91merror creating iran-azumi:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(service)
	if err != nil {
		fmt.Println("\033[91merror constructing iran-azumi:\033[0m", err)
		return
	}

	cmd := exec.Command("systemctl", "daemon-reload")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror reloading:\033[0m", err)
		return
	}

	cmd = exec.Command("sudo", "chmod", "u+x", "/etc/systemd/system/iran-azumi.service")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enablin da service:\033[0m", err)
		return
	}

	cmd = exec.Command("systemctl", "enable", "iran-azumi.")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enabling da service:\033[0m", err)
		return
	}

	cmd = exec.Command("systemctl", "restart", "iran-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror restarting da service:\033[0m", err)
		return
	}
    resIran()
	displayCheckmark("\033[92mService created successfully!\033[0m")
}
func kharejTcp6() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Reverse \033[92mIPV6 \033[96mTCP\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring KHAREJ")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\033[93mEnter the \033[92mStarting number:\033[0m ")
	scanner.Scan()
	startingNumberStr := scanner.Text()

	startingNumber, err := strconv.Atoi(startingNumberStr)
	if err != nil {
		fmt.Println("\033[91mPlz enter a valid number\033[0m")
		return
	}

	fmt.Print("\033[93mEnter \033[92mIran IPV6:\033[0m ")
	scanner.Scan()
	iranIP := scanner.Text()

	fmt.Print("\033[93mEnter \033[92mTunnel port:\033[0m ")
	scanner.Scan()
	tunnelPort := scanner.Text()

	fmt.Print("\033[93mHow many \033[92mConfigs\033[93m do you have?\033[0m ")
	scanner.Scan()
	numConfigsStr := scanner.Text()

	numConfigs, err := strconv.Atoi(numConfigsStr)
	if err != nil {
		fmt.Println("\033[91mPlz enter a valid number\033[0m")
		return
	}

	kharejPorts := make([]string, numConfigs)
	for i := 0; i < numConfigs; i++ {
		fmt.Printf("\033[93mEnter \033[92mconfig %d\033[93m port:\033[0m ", i+1)
		scanner.Scan()
		kharejPorts[i] = scanner.Text()
	}

	fmt.Print("\033[93mDo you want nodelay enabled? (\033[92my/\033[91mn\033[93m): \033[0m")
	scanner.Scan()
	nodelayOp := scanner.Text()
	nodelay := "false"
	if strings.ToLower(nodelayOp) == "y" {
		nodelay = "true"
	}

	client := fmt.Sprintf(`[client]
remote_addr = "[%s]:%s"
default_token = "azumiisinyourarea"
retry_interval = 1

[client.transport]
type = "tcp"

[client.transport.tcp]
nodelay = %s
keepalive_secs = 20
keepalive_interval = 8
`, iranIP, tunnelPort, nodelay)

	for i := 0; i < numConfigs; i++ {
		config := fmt.Sprintf(`[client.services.kharej%d]
type = "tcp"
local_addr = "127.0.0.1:%s"
`, i+startingNumber, kharejPorts[i])
		client += config
	}

	err = os.Remove("/root/rathole/client.toml")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting toml:\033[0m", err)
		return
	}

	file, err := os.Create("/root/rathole/client.toml")
	if err != nil {
		fmt.Println("\033[91merror creating toml:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(client)
	if err != nil {
		fmt.Println("\033[91merror putting configs into toml:\033[0m", err)
		return
	}

	service := `[Unit]
Description=Kharej-Azumi Service
After=network.target

[Service]
Type=simple
Restart=on-failure
RestartSec=5s
LimitNOFILE=1048576
ExecStart=/root/rathole/target/debug/rathole /root/rathole/client.toml

[Install]
WantedBy=multi-user.target`

	err = os.Remove("/etc/systemd/system/kharej-azumi.service")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting iran-azumi:\033[0m", err)
		return
	}

	file, err = os.Create("/etc/systemd/system/kharej-azumi.service")
	if err != nil {
		fmt.Println("\033[91merror creating iran-azumi:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(service)
	if err != nil {
		fmt.Println("\033[91merror constructing iran-azumi:\033[0m", err)
		return
	}

	cmd := exec.Command("systemctl", "daemon-reload")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror reloading:\033[0m", err)
		return
	}

	cmd = exec.Command("sudo", "chmod", "u+x", "/etc/systemd/system/kharej-azumi.service")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enablin da service:\033[0m", err)
		return
	}

	cmd = exec.Command("systemctl", "enable", "kharej-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enabling da service:\033[0m", err)
		return
	}

	cmd = exec.Command("systemctl", "restart", "kharej-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror restarting da service:\033[0m", err)
		return
	}
	resKharej()
	displayCheckmark("\033[92mService created successfully!\033[0m")
	fmt.Println("╭─────────────────────────────────────────────╮")
	fmt.Printf("\033[92m Starting number for the next server : \033[96m%-9d\n\033[0m", numConfigs+1)
	fmt.Println("╰─────────────────────────────────────────────╯")
}
func kharej2Tcp6() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Reverse \033[92mIPV6 \033[96mTCP\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring KHAREJ")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\033[93mEnter the \033[92mStarting number:\033[0m ")
	scanner.Scan()
	startingNumberStr := scanner.Text()

	startingNumber, err := strconv.Atoi(startingNumberStr)
	if err != nil {
		fmt.Println("\033[91mPlz enter a valid number\033[0m")
		return
	}

	fmt.Print("\033[93mEnter \033[92mIran IPV6:\033[0m ")
	scanner.Scan()
	iranIP := scanner.Text()

	fmt.Print("\033[93mEnter \033[92mTunnel port:\033[0m ")
	scanner.Scan()
	tunnelPort := scanner.Text()

	fmt.Print("\033[93mHow many \033[92mConfigs\033[93m do you have?\033[0m ")
	scanner.Scan()
	numConfigsStr := scanner.Text()

	numConfigs, err := strconv.Atoi(numConfigsStr)
	if err != nil {
		fmt.Println("\033[91mPlz enter a valid number\033[0m")
		return
	}

	kharejPorts := make([]string, numConfigs)
	for i := 0; i < numConfigs; i++ {
		fmt.Printf("\033[93mEnter \033[92mconfig %d\033[93m port:\033[0m ", i+1)
		scanner.Scan()
		kharejPorts[i] = scanner.Text()
	}

	fmt.Print("\033[93mDo you want nodelay enabled? (\033[92my/\033[91mn\033[93m): \033[0m")
	scanner.Scan()
	nodelayOp := scanner.Text()
	nodelay := "false"
	if strings.ToLower(nodelayOp) == "y" {
		nodelay = "true"
	}

	client := fmt.Sprintf(`[client]
remote_addr = "[%s]:%s"
default_token = "azumiisinyourarea"
retry_interval = 1

[client.transport]
type = "tcp"

[client.transport.tcp]
nodelay = %s
keepalive_secs = 20
keepalive_interval = 8
`, iranIP, tunnelPort, nodelay)

	for i := 0; i < numConfigs; i++ {
		config := fmt.Sprintf(`[client.services.kharej%d]
type = "tcp"
local_addr = "127.0.0.1:%s"
`, i+startingNumber, kharejPorts[i])
		client += config
	}

	err = os.Remove("/root/rathole/client.toml")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting toml:\033[0m", err)
		return
	}

	file, err := os.Create("/root/rathole/client.toml")
	if err != nil {
		fmt.Println("\033[91merror creating toml:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(client)
	if err != nil {
		fmt.Println("\033[91merror putting configs into toml:\033[0m", err)
		return
	}

	service := `[Unit]
Description=Kharej-Azumi Service
After=network.target

[Service]
Type=simple
Restart=on-failure
RestartSec=5s
LimitNOFILE=1048576
ExecStart=/root/rathole/target/debug/rathole /root/rathole/client.toml

[Install]
WantedBy=multi-user.target`

	err = os.Remove("/etc/systemd/system/kharej-azumi.service")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting iran-azumi:\033[0m", err)
		return
	}

	file, err = os.Create("/etc/systemd/system/kharej-azumi.service")
	if err != nil {
		fmt.Println("\033[91merror creating iran-azumi:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(service)
	if err != nil {
		fmt.Println("\033[91merror constructing iran-azumi:\033[0m", err)
		return
	}

	cmd := exec.Command("systemctl", "daemon-reload")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror reloading:\033[0m", err)
		return
	}

	cmd = exec.Command("sudo", "chmod", "u+x", "/etc/systemd/system/kharej-azumi.service")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enablin da service:\033[0m", err)
		return
	}
	cmd = exec.Command("systemctl", "enable", "kharej-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enabling da service:\033[0m", err)
		return
	}

	cmd = exec.Command("systemctl", "restart", "kharej-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror restarting da service:\033[0m", err)
		return
	}
	resKharej()
	displayCheckmark("\033[92mService created successfully!\033[0m")
	if numConfigs == 1 {
	    fmt.Println("╭─────────────────────────────────────────────╮")
		fmt.Printf("\033[92m  Starting number for the next server:\033[96m %d\n\033[0m", startingNumber+1)
		fmt.Println("╰─────────────────────────────────────────────╯")
	} else {
	    fmt.Println("╭─────────────────────────────────────────────╮")
		fmt.Printf("\033[92m  Starting number for the next server:\033[96m %d\n\033[0m", numConfigs+startingNumber)
		fmt.Println("╰─────────────────────────────────────────────╯")
	}
}
func udp6Menu() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Reverse \033[92mUDP \033[96mIPV6 \033[93mMenu\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. \033[92mIRAN\033[0m", "2. \033[93mKHAREJ\033[92m[1]\033[0m", "3. \033[93mKHAREJ\033[92m[2]\033[0m", "4. \033[93mKHAREJ\033[92m[3]\033[0m", "5. \033[93mKHAREJ\033[92m[4]\033[0m", "6. \033[93mKHAREJ\033[92m[5]\033[0m", "7. \033[93mKHAREJ\033[92m[6]\033[0m", "8. \033[93mKHAREJ\033[92m[7]\033[0m", "9. \033[93mKHAREJ\033[92m[8]\033[0m", "10. \033[93mKHAREJ\033[92m[9]\033[0m", "11. \033[93mKHAREJ\033[92m[10]\033[0m", "0. \033[94mBack to the main menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. \033[92mIRAN\033[0m":
		iranUdp6()
	case "2. \033[93mKHAREJ\033[92m[1]\033[0m":
		kharejUdp6()
	case "3. \033[93mKHAREJ\033[92m[2]\033[0m":
		kharej2Udp6()
    case "4. \033[93mKHAREJ\033[92m[3]\033[0m":
		kharej2Udp6()
	case "5. \033[93mKHAREJ\033[92m[4]\033[0m":
		kharej2Udp6()
	case "6. \033[93mKHAREJ\033[92m[5]\033[0m":
		kharej2Udp6()
	case "7. \033[93mKHAREJ\033[92m[6]\033[0m":
		kharej2Udp6()
	case "8. \033[93mKHAREJ\033[92m[7]\033[0m":
		kharej2Udp6()
	case "9. \033[93mKHAREJ\033[92m[8]\033[0m":
		kharej2Udp6()
	case "10. \033[93mKHAREJ\033[92m[9]\033[0m":
		kharej2Udp6()
	case "11. \033[93mKHAREJ\033[92m[10]\033[0m":
		kharej2Udp6()
	case "0. \033[94mBack to the main menu\033[0m":
	    clearScreen()
		mainMenu()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}

func iranUdp6() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Reverse \033[92mIPV6 \033[96mUDP\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring IRAN")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\033[93mHow many \033[92mconfigs\033[93m do you have \033[96m[All Servers Combined]\033[93m? \033[0m")
	scanner.Scan()
	numConfigsStr := scanner.Text()

	numConfigs, err := strconv.Atoi(numConfigsStr)
	if err != nil {
		fmt.Println("\033[91mPlz enter a valid number\033[0m")
		return
	}

	fmt.Print("\033[93mEnter \033[92mTunnel port:\033[0m ")
	scanner.Scan()
	tunnelPort := scanner.Text()

	kharejPorts := make([]string, numConfigs)
	for i := 0; i < numConfigs; i++ {
		fmt.Printf("\033[93mEnter \033[92mConfig %d\033[93m Port: \033[0m", i+1)
		scanner.Scan()
		kharejPorts[i] = scanner.Text()
	}

	fmt.Print("\033[93mDo you want nodelay enabled? (\033[92my/\033[91mn\033[93m): \033[0m")
	scanner.Scan()
	nodelayOp := scanner.Text()
	nodelay := "false"
	if strings.ToLower(nodelayOp) == "y" {
		nodelay = "true"
	}

	server := fmt.Sprintf(`[server]
bind_addr = "[::]:%s"
default_token = "azumiisinyourarea"

[server.transport]
type = "tcp"

[server.transport.tcp]
nodelay = %s
keepalive_secs = 20
keepalive_interval = 8

`, tunnelPort, nodelay)
	for i := 0; i < numConfigs; i++ {
		config := fmt.Sprintf(`[server.services.kharej%d]
type = "udp"
bind_addr = "0.0.0.0:%s" 
`, i+1, kharejPorts[i])
		server += config
	}

	err = os.Remove("/root/rathole/server.toml")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting toml:\033[0m", err)
		return
	}

	file, err := os.Create("/root/rathole/server.toml")
	if err != nil {
		fmt.Println("\033[91merror creating toml:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(server)
	if err != nil {
		fmt.Println("\033[91merror putting configs into toml:\033[0m", err)
		return
	}
	service := `[Unit]
Description=Iran-Azumi Service
After=network.target

[Service]
Type=simple
Restart=on-failure
RestartSec=5s
LimitNOFILE=1048576
ExecStart=/root/rathole/target/debug/rathole /root/rathole/server.toml

[Install]
WantedBy=multi-user.target`

	err = os.Remove("/etc/systemd/system/iran-azumi.service")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting iran-azumi:\033[0m", err)
		return
	}

	file, err = os.Create("/etc/systemd/system/iran-azumi.service")
	if err != nil {
		fmt.Println("\033[91merror creating iran-azumi:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(service)
	if err != nil {
		fmt.Println("\033[91merror constructing iran-azumi:\033[0m", err)
		return
	}

	cmd := exec.Command("systemctl", "daemon-reload")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror reloading:\033[0m", err)
		return
	}

	cmd = exec.Command("sudo", "chmod", "u+x", "/etc/systemd/system/iran-azumi.service")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enablin da service:\033[0m", err)
		return
	}
	
	cmd = exec.Command("systemctl", "enable", "iran-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enabling da service:\033[0m", err)
		return
	}

	cmd = exec.Command("systemctl", "restart", "iran-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror restarting da service:\033[0m", err)
		return
	}
    resIran()
	displayCheckmark("\033[92mService created successfully!\033[0m")
}
func kharejUdp6() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Reverse \033[92mIPV6 \033[96mUDP\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring KHAREJ")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\033[93mEnter the \033[92mStarting number:\033[0m ")
	scanner.Scan()
	startingNumberStr := scanner.Text()

	startingNumber, err := strconv.Atoi(startingNumberStr)
	if err != nil {
		fmt.Println("\033[91mPlz enter a valid number\033[0m")
		return
	}

	fmt.Print("\033[93mEnter \033[92mIran IPV6:\033[0m ")
	scanner.Scan()
	iranIP := scanner.Text()

	fmt.Print("\033[93mEnter \033[92mTunnel port:\033[0m ")
	scanner.Scan()
	tunnelPort := scanner.Text()

	fmt.Print("\033[93mHow many \033[92mConfigs\033[93m do you have?\033[0m ")
	scanner.Scan()
	numConfigsStr := scanner.Text()

	numConfigs, err := strconv.Atoi(numConfigsStr)
	if err != nil {
		fmt.Println("\033[91mPlz enter a valid number\033[0m")
		return
	}

	kharejPorts := make([]string, numConfigs)
	for i := 0; i < numConfigs; i++ {
		fmt.Printf("\033[93mEnter \033[92mconfig %d\033[93m port:\033[0m ", i+1)
		scanner.Scan()
		kharejPorts[i] = scanner.Text()
	}

	fmt.Print("\033[93mDo you want nodelay enabled? (\033[92my/\033[91mn\033[93m): \033[0m")
	scanner.Scan()
	nodelayOp := scanner.Text()
	nodelay := "false"
	if strings.ToLower(nodelayOp) == "y" {
		nodelay = "true"
	}

	client := fmt.Sprintf(`[client]
remote_addr = "[%s]:%s"
default_token = "azumiisinyourarea"
retry_interval = 1

[client.transport]
type = "tcp"

[client.transport.tcp]
nodelay = %s
keepalive_secs = 20
keepalive_interval = 8
`, iranIP, tunnelPort, nodelay)

	for i := 0; i < numConfigs; i++ {
		config := fmt.Sprintf(`[client.services.kharej%d]
type = "udp"
local_addr = "127.0.0.1:%s"
`, i+startingNumber, kharejPorts[i])
		client += config
	}

	err = os.Remove("/root/rathole/client.toml")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting toml:\033[0m", err)
		return
	}

	file, err := os.Create("/root/rathole/client.toml")
	if err != nil {
		fmt.Println("\033[91merror creating toml:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(client)
	if err != nil {
		fmt.Println("\033[91merror putting configs into toml:\033[0m", err)
		return
	}

	service := `[Unit]
Description=Kharej-Azumi Service
After=network.target

[Service]
Type=simple
Restart=on-failure
RestartSec=5s
LimitNOFILE=1048576
ExecStart=/root/rathole/target/debug/rathole /root/rathole/client.toml

[Install]
WantedBy=multi-user.target`

	err = os.Remove("/etc/systemd/system/kharej-azumi.service")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting iran-azumi:\033[0m", err)
		return
	}

	file, err = os.Create("/etc/systemd/system/kharej-azumi.service")
	if err != nil {
		fmt.Println("\033[91merror creating iran-azumi:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(service)
	if err != nil {
		fmt.Println("\033[91merror constructing iran-azumi:\033[0m", err)
		return
	}

	cmd := exec.Command("systemctl", "daemon-reload")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror reloading:\033[0m", err)
		return
	}

	cmd = exec.Command("sudo", "chmod", "u+x", "/etc/systemd/system/kharej-azumi.service")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enablin da service:\033[0m", err)
		return
	}
	
	cmd = exec.Command("systemctl", "enable", "kharej-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enabling da service:\033[0m", err)
		return
	}

	cmd = exec.Command("systemctl", "restart", "kharej-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror restarting da service:\033[0m", err)
		return
	}
	resKharej()
	displayCheckmark("\033[92mService created successfully!\033[0m")
	fmt.Println("╭─────────────────────────────────────────────╮")
	fmt.Printf("\033[92m Starting number for the next server : \033[96m%-9d\n\033[0m", numConfigs+1)
	fmt.Println("╰─────────────────────────────────────────────╯")
}
func kharej2Udp6() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Reverse \033[92mIPV6 \033[96mUDP\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring KHAREJ")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\033[93mEnter the \033[92mStarting number:\033[0m ")
	scanner.Scan()
	startingNumberStr := scanner.Text()

	startingNumber, err := strconv.Atoi(startingNumberStr)
	if err != nil {
		fmt.Println("\033[91mPlz enter a valid number\033[0m")
		return
	}

	fmt.Print("\033[93mEnter \033[92mIran IPV6:\033[0m ")
	scanner.Scan()
	iranIP := scanner.Text()

	fmt.Print("\033[93mEnter \033[92mTunnel port:\033[0m ")
	scanner.Scan()
	tunnelPort := scanner.Text()

	fmt.Print("\033[93mHow many \033[92mConfigs\033[93m do you have?\033[0m ")
	scanner.Scan()
	numConfigsStr := scanner.Text()

	numConfigs, err := strconv.Atoi(numConfigsStr)
	if err != nil {
		fmt.Println("\033[91mPlz enter a valid number\033[0m")
		return
	}

	kharejPorts := make([]string, numConfigs)
	for i := 0; i < numConfigs; i++ {
		fmt.Printf("\033[93mEnter \033[92mconfig %d\033[93m port:\033[0m ", i+1)
		scanner.Scan()
		kharejPorts[i] = scanner.Text()
	}

	fmt.Print("\033[93mDo you want nodelay enabled? (\033[92my/\033[91mn\033[93m): \033[0m")
	scanner.Scan()
	nodelayOp := scanner.Text()
	nodelay := "false"
	if strings.ToLower(nodelayOp) == "y" {
		nodelay = "true"
	}

	client := fmt.Sprintf(`[client]
remote_addr = "[%s]:%s"
default_token = "azumiisinyourarea"
retry_interval = 1

[client.transport]
type = "tcp"

[client.transport.tcp]
nodelay = %s
keepalive_secs = 20
keepalive_interval = 8
`, iranIP, tunnelPort, nodelay)

	for i := 0; i < numConfigs; i++ {
		config := fmt.Sprintf(`[client.services.kharej%d]
type = "udp"
local_addr = "127.0.0.1:%s"
`, i+startingNumber, kharejPorts[i])
		client += config
	}

	err = os.Remove("/root/rathole/client.toml")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting toml:\033[0m", err)
		return
	}

	file, err := os.Create("/root/rathole/client.toml")
	if err != nil {
		fmt.Println("\033[91merror creating toml:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(client)
	if err != nil {
		fmt.Println("\033[91merror putting configs into toml:\033[0m", err)
		return
	}

	service := `[Unit]
Description=Kharej-Azumi Service
After=network.target

[Service]
Type=simple
Restart=on-failure
RestartSec=5s
LimitNOFILE=1048576
ExecStart=/root/rathole/target/debug/rathole /root/rathole/client.toml

[Install]
WantedBy=multi-user.target`

	err = os.Remove("/etc/systemd/system/kharej-azumi.service")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting iran-azumi:\033[0m", err)
		return
	}

	file, err = os.Create("/etc/systemd/system/kharej-azumi.service")
	if err != nil {
		fmt.Println("\033[91merror creating iran-azumi:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(service)
	if err != nil {
		fmt.Println("\033[91merror constructing iran-azumi:\033[0m", err)
		return
	}

	cmd := exec.Command("systemctl", "daemon-reload")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror reloading:\033[0m", err)
		return
	}

	
	cmd = exec.Command("sudo", "chmod", "u+x", "/etc/systemd/system/kharej-azumi.service")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enablin da service:\033[0m", err)
		return
	}
	
	cmd = exec.Command("systemctl", "enable", "kharej-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enabling da service:\033[0m", err)
		return
	}

	cmd = exec.Command("systemctl", "restart", "kharej-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror restarting da service:\033[0m", err)
		return
	}
	resKharej()
	displayCheckmark("\033[92mService created successfully!\033[0m")
	if numConfigs == 1 {
	    fmt.Println("╭─────────────────────────────────────────────╮")
		fmt.Printf("\033[92m  Starting number for the next server:\033[96m %d\n\033[0m", startingNumber+1)
		fmt.Println("╰─────────────────────────────────────────────╯")
	} else {
	    fmt.Println("╭─────────────────────────────────────────────╮")
		fmt.Printf("\033[92m  Starting number for the next server:\033[96m %d\n\033[0m", numConfigs+startingNumber)
		fmt.Println("╰─────────────────────────────────────────────╯")
	}
}
func scm(cmd *exec.Cmd) error {
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("\033[91mcmd failed: %v, output: %s\033[0m", err, output)
	}
	return nil
}

func ssl2() {
	caKeyCmd := exec.Command("openssl", "genrsa", "-out", "/root/rathole/rootCA.key", "2048")
	if err := scm(caKeyCmd); err != nil {
		log.Fatal(err)
	}

	caCertCmd := exec.Command("openssl", "req", "-x509", "-sha256", "-days", "356", "-new", "-nodes",
		"-key", "/root/rathole/rootCA.key", "-subj", "/CN=MyOwnCA/C=US/L=San Francisco", "-out", "/root/rathole/rootCA.crt")
	if err := scm(caCertCmd); err != nil {
		log.Fatal(err)
	}

	serverKeyCmd := exec.Command("openssl", "genrsa", "-out", "/root/rathole/server.key", "2048")
	if err := scm(serverKeyCmd); err != nil {
		log.Fatal(err)
	}

	csrConf := `# OpenSSL configuration file for generating server certificate signing request (CSR)
[ req ]
default_bits = 2048
prompt = no
default_md = sha256
req_extensions = req_ext
distinguished_name = dn

[ dn ]
C = US
ST = California
L = San Francisco
O = Someone
OU = Someone
CN = localhost

[ req_ext ]
subjectAltName = @alt_names

[ alt_names ]
DNS.1 = localhost
`
	csrConfFile := "/root/rathole/csr.conf"
	if err := ioutil.WriteFile(csrConfFile, []byte(csrConf), 0644); err != nil {
		log.Fatal(err)
	}
	defer os.Remove(csrConfFile)

	serverCertCmd := exec.Command("openssl", "req", "-new", "-key", "/root/rathole/server.key", "-out", "/root/rathole/server.csr",
		"-config", csrConfFile)
	if err := scm(serverCertCmd); err != nil {
		log.Fatal(err)
	}

	serverCertCmd = exec.Command("openssl", "x509", "-req", "-in", "/root/rathole/server.csr", "-CA", "/root/rathole/rootCA.crt",
		"-CAkey", "/root/rathole/rootCA.key", "-CAcreateserial", "-out", "/root/rathole/server.crt", "-days", "365", "-sha256",
		"-extfile", csrConfFile)
	if err := scm(serverCertCmd); err != nil {
		log.Fatal(err)
	}

	pkcs12Cmd := exec.Command("openssl", "pkcs12", "-export", "-out", "/root/rathole/identity.pfx",
		"-inkey", "/root/rathole/server.key", "-in", "/root/rathole/server.crt", "-certfile", "/root/rathole/rootCA.crt", "-passout", "pass:azumi1234")
	if err := scm(pkcs12Cmd); err != nil {
		log.Fatal(err)
	}

	displayCheckmark("\033[92mCertificate process completed!\033[0m")
}
func ssl() {
	displayNotification("\033[93mGetting Certs..\033[0m")

	dest := "/root/rathole"
	err := os.Chdir(dest)
	if err != nil {
		log.Fatalf("\033[91mThis dir doesn't Exist:\033[0m %v", err)
	}

	caKey := []string{
		"req", "-x509",
		"-sha256", "-days", "365",
		"-nodes",
		"-newkey", "rsa:2048",
		"-subj", "/CN=MyOwnCA/C=US/L=San Francisco",
		"-keyout", "rootCA.key", "-out", "rootCA.crt",
	}

	err = OpenSSL(caKey)
	if err != nil {
		log.Fatalf("\033[91mCouldn't create CA:\033[0m %v", err)
	}

	serverKey := []string{
		"genrsa", "-out", "server.key", "2048",
	}

	err = OpenSSL(serverKey)
	if err != nil {
		log.Fatalf("\033[91mCouldn't create server's private key:\033[0m %v", err)
	}

	csrKey := `[ req ]
	default_bits = 2048
	prompt = no
	default_md = sha256
	req_extensions = req_ext
	distinguished_name = dn

[ dn ]
C = US
ST = California
L = San Francisco
O = Someone
OU = Someone
CN = localhost

[ req_ext ]
subjectAltName = @alt_names

[ alt_names ]
DNS.1 = localhost
`

	err = ioutil.WriteFile("csr.conf", []byte(csrKey), 0644)
	if err != nil {
		log.Fatalf("\033[91mCoudln't create CSR Key:\033[0m %v", err)
	}

	serverCSR := []string{
		"req", "-new", "-key", "server.key", "-out", "server.csr",
		"-config", "csr.conf",
	}

	err = OpenSSL(serverCSR)
	if err != nil {
		log.Fatalf("\033[91mCouldn't create server's CSR Key:\033[0m %v", err)
	}

	certKey := `authorityKeyIdentifier=keyid,issuer
basicConstraints=CA:FALSE
keyUsage = digitalSignature, nonRepudiation, keyEncipherment, dataEncipherment
subjectAltName = @alt_names

[alt_names]
DNS.1 = localhost
`

	err = ioutil.WriteFile("cert.conf", []byte(certKey), 0644)
	if err != nil {
		log.Fatalf("\033[91mCouldn't create cert key:\033[0m %v", err)
	}

	serverCert := []string{
		"x509", "-req",
		"-in", "server.csr",
		"-CA", "rootCA.crt", "-CAkey", "rootCA.key",
		"-out", "server.crt",
		"-days", "365",
		"-sha256", "-extfile", "cert.conf",
	}

	err = OpenSSL(serverCert)
	if err != nil {
		log.Fatalf("\033[91mCouldn't create server's cert:\033[0m %v", err)
	}

	serverPfx := []string{
		"pkcs12", "-export",
		"-out", "identity.pfx",
		"-inkey", "server.key", "-in", "server.crt", "-certfile", "rootCA.crt",
		"-passout", "pass:azumi1234",
	}

	err = OpenSSL(serverPfx)
	if err != nil {
		log.Fatalf("\033[91mCouldn't create identity cert:\033[0m %v", err)
	}

	os.Remove("server.csr")
	os.Remove("csr.conf")
	os.Remove("cert.conf")

	displayCheckmark("\033[92mCertificate process completed!\033[0m")
}

func OpenSSL(args []string) error {
	cmd := exec.Command("openssl", args...)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("\033[91mopenssl command has failed:\033[0m %w", err)
	}
	return nil
}
func copy() {
	fmt.Print("\033[93mEnter \033[92mKharej IPv4\033[93m address:\033[0m ")
	var ipAddress string
	fmt.Scanln(&ipAddress)

	fmt.Print("\033[93mEnter \033[92mSSH port\033[93m:\033[0m ")
	var port string
	fmt.Scanln(&port)

	fmt.Print("\033[93mEnter the \033[92musername\033[93m:\033[0m ")
	var username string
	fmt.Scanln(&username)

	fmt.Print("\033[93mEnter the \033[92mpassword\033[93m:\033[0m ")
	var password string
	fmt.Scanln(&password)

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", ipAddress+":"+port, config)
	if err != nil {
		log.Fatalf("\033[91mCouldn't connect to Kharej server:\033[0m %v", err)
	}
	defer conn.Close()

	sftpClient, err := sftp.NewClient(conn)
	if err != nil {
		log.Fatalf("\033[91mCouldn't establish an SFTP session:\033[0m %v", err)
	}
	defer sftpClient.Close()

	file, err := os.Open("/root/rathole/rootCA.crt") 
	if err != nil {
		log.Fatalf("\033[91mThere is no rootCA key in your dir:\033[0m %v", err)
	}
	defer file.Close()

	remoteFilePath := "/root/rathole/rootCA.crt"
	remoteFile, err := sftpClient.Create(remoteFilePath)
	if err != nil {
		log.Fatalf("\033[91mCouldn't copy rootCA key into kharej's dir:\033[0m %v", err)
	}
	defer remoteFile.Close()

	_, err = io.Copy(remoteFile, file)
	if err != nil {
		log.Fatalf("\033[91mCopying process was not successful:\033[0m %v", err)
	}

	displayCheckmark("\033[92mCert copied successfully!\033[0m")
}
func ws4Menu() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Reverse \033[92mWs + TLS \033[96mIPV4 \033[93mMenu\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. \033[92mIRAN\033[0m", "2. \033[93mKHAREJ\033[92m[1]\033[0m", "3. \033[93mKHAREJ\033[92m[2]\033[0m", "4. \033[93mKHAREJ\033[92m[3]\033[0m", "5. \033[93mKHAREJ\033[92m[4]\033[0m", "6. \033[93mKHAREJ\033[92m[5]\033[0m", "0. \033[94mBack to the main menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. \033[92mIRAN\033[0m":
		iranWs4M()
	case "2. \033[93mKHAREJ\033[92m[1]\033[0m":
		kharejWs4()
	case "3. \033[93mKHAREJ\033[92m[2]\033[0m":
		kharej2Ws4()
    case "4. \033[93mKHAREJ\033[92m[3]\033[0m":
		kharej2Ws4()
	case "5. \033[93mKHAREJ\033[92m[4]\033[0m":
		kharej2Ws4()
	case "6. \033[93mKHAREJ\033[92m[5]\033[0m":
		kharej2Ws4()
	case "0. \033[94mBack to the main menu\033[0m":
	    clearScreen()
		mainMenu()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}
func iranWs4M() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m IRAN \033[92mWs + TLS \033[93mMenu\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. \033[92mIRAN Config \033[96m[Method 1]\033[0m", "2. \033[92mIRAN Config \033[96m[Method 2]\033[0m", "3. \033[93mCopy Cert\033[0m", "0. \033[94mBack to the previous menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. \033[92mIRAN Config \033[96m[Method 1]\033[0m":
		iranWs4()
	case "2. \033[92mIRAN Config \033[96m[Method 2]\033[0m":
		iranWs42()	
	case "3. \033[93mCopy Cert\033[0m":
		copy()
	case "0. \033[94mBack to the previous menu\033[0m":
	    clearScreen()
		ws4Menu()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}
func iranWs4() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Reverse \033[92mIPV4 \033[96mWS + TLS\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring IRAN")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	rootCA := "/root/rathole/rootCA.crt"

	if _, err := os.Stat(rootCA); os.IsNotExist(err) {
		ssl()
	} else {
		fmt.Println("\033[93mSkip getting Cert..\033[0m")
	}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\033[93mHow many \033[92mconfigs\033[93m do you have \033[96m[All Servers Combined]\033[93m? \033[0m")
	scanner.Scan()
	numConfigsStr := scanner.Text()

	numConfigs, err := strconv.Atoi(numConfigsStr)
	if err != nil {
		fmt.Println("\033[91mPlz enter a valid number\033[0m")
		return
	}

	fmt.Print("\033[93mEnter \033[92mTunnel port:\033[0m ")
	scanner.Scan()
	tunnelPort := scanner.Text()

	kharejPorts := make([]string, numConfigs)
	for i := 0; i < numConfigs; i++ {
		fmt.Printf("\033[93mEnter \033[92mConfig %d\033[93m Port: \033[0m", i+1)
		scanner.Scan()
		kharejPorts[i] = scanner.Text()
	}

	server := fmt.Sprintf(`[server]
bind_addr = "0.0.0.0:%s"
default_token = "azumiisinyourarea"

[server.transport]
type = "tls"

[server.transport.tls]
pkcs12 = "/root/rathole/identity.pfx"
pkcs12_password = "azumi1234"

[server.transport.websocket] 
tls = true 

`, tunnelPort)
	for i := 0; i < numConfigs; i++ {
		config := fmt.Sprintf(`[server.services.kharej%d]
bind_addr = "0.0.0.0:%s" 
`, i+1, kharejPorts[i])
		server += config
	}

	err = os.Remove("/root/rathole/server.toml")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting toml:\033[0m", err)
		return
	}

	file, err := os.Create("/root/rathole/server.toml")
	if err != nil {
		fmt.Println("\033[91merror creating toml:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(server)
	if err != nil {
		fmt.Println("\033[91merror putting configs into toml:\033[0m", err)
		return
	}
	service := `[Unit]
Description=Iran-Azumi Service
After=network.target

[Service]
Type=simple
Restart=on-failure
RestartSec=5s
LimitNOFILE=1048576
ExecStart=/root/rathole/target/debug/rathole /root/rathole/server.toml

[Install]
WantedBy=multi-user.target`

	err = os.Remove("/etc/systemd/system/iran-azumi.service")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting iran-azumi:\033[0m", err)
		return
	}

	file, err = os.Create("/etc/systemd/system/iran-azumi.service")
	if err != nil {
		fmt.Println("\033[91merror creating iran-azumi:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(service)
	if err != nil {
		fmt.Println("\033[91merror constructing iran-azumi:\033[0m", err)
		return
	}

	cmd := exec.Command("systemctl", "daemon-reload")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror reloading:\033[0m", err)
		return
	}

	
	cmd = exec.Command("sudo", "chmod", "u+x", "/etc/systemd/system/iran-azumi.service")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enablin da service:\033[0m", err)
		return
	}
	
	cmd = exec.Command("systemctl", "enable", "iran-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enabling da service:\033[0m", err)
		return
	}

	cmd = exec.Command("systemctl", "restart", "iran-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror restarting da service:\033[0m", err)
		return
	}
    resIran()
	displayCheckmark("\033[92mService created successfully!\033[0m")
}
func iranWs42() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Reverse \033[92mIPV4 \033[96mWS + TLS\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring IRAN")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	rootCA := "/root/rathole/rootCA.crt"

	if _, err := os.Stat(rootCA); os.IsNotExist(err) {
		ssl2()
	} else {
		fmt.Println("\033[93mSkip getting Cert..\033[0m")
	}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\033[93mHow many \033[92mconfigs\033[93m do you have \033[96m[All Servers Combined]\033[93m? \033[0m")
	scanner.Scan()
	numConfigsStr := scanner.Text()

	numConfigs, err := strconv.Atoi(numConfigsStr)
	if err != nil {
		fmt.Println("\033[91mPlz enter a valid number\033[0m")
		return
	}

	fmt.Print("\033[93mEnter \033[92mTunnel port:\033[0m ")
	scanner.Scan()
	tunnelPort := scanner.Text()

	kharejPorts := make([]string, numConfigs)
	for i := 0; i < numConfigs; i++ {
		fmt.Printf("\033[93mEnter \033[92mConfig %d\033[93m Port: \033[0m", i+1)
		scanner.Scan()
		kharejPorts[i] = scanner.Text()
	}

	server := fmt.Sprintf(`[server]
bind_addr = "0.0.0.0:%s"
default_token = "azumiisinyourarea"

[server.transport]
type = "tls"

[server.transport.tls]
pkcs12 = "/root/rathole/identity.pfx"
pkcs12_password = "azumi1234"

[server.transport.websocket] 
tls = true 

`, tunnelPort)
	for i := 0; i < numConfigs; i++ {
		config := fmt.Sprintf(`[server.services.kharej%d]
bind_addr = "0.0.0.0:%s" 
`, i+1, kharejPorts[i])
		server += config
	}

	err = os.Remove("/root/rathole/server.toml")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting toml:\033[0m", err)
		return
	}

	file, err := os.Create("/root/rathole/server.toml")
	if err != nil {
		fmt.Println("\033[91merror creating toml:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(server)
	if err != nil {
		fmt.Println("\033[91merror putting configs into toml:\033[0m", err)
		return
	}
	service := `[Unit]
Description=Iran-Azumi Service
After=network.target

[Service]
Type=simple
Restart=on-failure
RestartSec=5s
LimitNOFILE=1048576
ExecStart=/root/rathole/target/debug/rathole /root/rathole/server.toml

[Install]
WantedBy=multi-user.target`

	err = os.Remove("/etc/systemd/system/iran-azumi.service")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting iran-azumi:\033[0m", err)
		return
	}

	file, err = os.Create("/etc/systemd/system/iran-azumi.service")
	if err != nil {
		fmt.Println("\033[91merror creating iran-azumi:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(service)
	if err != nil {
		fmt.Println("\033[91merror constructing iran-azumi:\033[0m", err)
		return
	}

	cmd := exec.Command("systemctl", "daemon-reload")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror reloading:\033[0m", err)
		return
	}

	cmd = exec.Command("sudo", "chmod", "u+x", "/etc/systemd/system/iran-azumi.service")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enablin da service:\033[0m", err)
		return
	}
	cmd = exec.Command("systemctl", "enable", "iran-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enabling da service:\033[0m", err)
		return
	}

	cmd = exec.Command("systemctl", "restart", "iran-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror restarting da service:\033[0m", err)
		return
	}
    resIran()
	displayCheckmark("\033[92mService created successfully!\033[0m")
}
func kharejWs4() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Reverse \033[92mIPV4 \033[96mTCP\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring KHAREJ")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\033[93mEnter the \033[92mStarting number:\033[0m ")
	scanner.Scan()
	startingNumberStr := scanner.Text()

	startingNumber, err := strconv.Atoi(startingNumberStr)
	if err != nil {
		fmt.Println("\033[91mPlz enter a valid number\033[0m")
		return
	}

	fmt.Print("\033[93mEnter \033[92mIran IPV4:\033[0m ")
	scanner.Scan()
	iranIP := scanner.Text()

	fmt.Print("\033[93mEnter \033[92mTunnel port:\033[0m ")
	scanner.Scan()
	tunnelPort := scanner.Text()

	fmt.Print("\033[93mHow many \033[92mConfigs\033[93m do you have?\033[0m ")
	scanner.Scan()
	numConfigsStr := scanner.Text()

	numConfigs, err := strconv.Atoi(numConfigsStr)
	if err != nil {
		fmt.Println("\033[91mPlz enter a valid number\033[0m")
		return
	}

	kharejPorts := make([]string, numConfigs)
	for i := 0; i < numConfigs; i++ {
		fmt.Printf("\033[93mEnter \033[92mconfig %d\033[93m port:\033[0m ", i+1)
		scanner.Scan()
		kharejPorts[i] = scanner.Text()
	}

	client := fmt.Sprintf(`[client]
remote_addr = "%s:%s"
default_token = "azumiisinyourarea"

[client.transport]
type = "tls"

[client.transport.tls]
trusted_root = "/root/rathole/rootCA.crt"
hostname = "localhost"

[client.transport.websocket] 
tls = true 

`, iranIP, tunnelPort)

	for i := 0; i < numConfigs; i++ {
		config := fmt.Sprintf(`[client.services.kharej%d]
local_addr = "127.0.0.1:%s"
`, i+startingNumber, kharejPorts[i])
		client += config
	}

	err = os.Remove("/root/rathole/client.toml")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting toml:\033[0m", err)
		return
	}

	file, err := os.Create("/root/rathole/client.toml")
	if err != nil {
		fmt.Println("\033[91merror creating toml:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(client)
	if err != nil {
		fmt.Println("\033[91merror putting configs into toml:\033[0m", err)
		return
	}

	service := `[Unit]
Description=Kharej-Azumi Service
After=network.target

[Service]
Type=simple
Restart=on-failure
RestartSec=5s
LimitNOFILE=1048576
ExecStart=/root/rathole/target/debug/rathole /root/rathole/client.toml

[Install]
WantedBy=multi-user.target`

	err = os.Remove("/etc/systemd/system/kharej-azumi.service")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting iran-azumi:\033[0m", err)
		return
	}

	file, err = os.Create("/etc/systemd/system/kharej-azumi.service")
	if err != nil {
		fmt.Println("\033[91merror creating iran-azumi:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(service)
	if err != nil {
		fmt.Println("\033[91merror constructing iran-azumi:\033[0m", err)
		return
	}

	cmd := exec.Command("systemctl", "daemon-reload")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror reloading:\033[0m", err)
		return
	}

	cmd = exec.Command("sudo", "chmod", "u+x", "/etc/systemd/system/kharej-azumi.service")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enablin da service:\033[0m", err)
		return
	}
	
	cmd = exec.Command("systemctl", "enable", "kharej-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enabling da service:\033[0m", err)
		return
	}

	cmd = exec.Command("systemctl", "restart", "kharej-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror restarting da service:\033[0m", err)
		return
	}
	resKharej()
	displayCheckmark("\033[92mService created successfully!\033[0m")
	fmt.Println("╭─────────────────────────────────────────────╮")
	fmt.Printf("\033[92m Starting number for the next server : \033[96m%-9d\n\033[0m", numConfigs+1)
	fmt.Println("╰─────────────────────────────────────────────╯")
}
func kharej2Ws4() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Reverse \033[92mIPV4 \033[96mTCP\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring KHAREJ")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\033[93mEnter the \033[92mStarting number:\033[0m ")
	scanner.Scan()
	startingNumberStr := scanner.Text()

	startingNumber, err := strconv.Atoi(startingNumberStr)
	if err != nil {
		fmt.Println("\033[91mPlz enter a valid number\033[0m")
		return
	}

	fmt.Print("\033[93mEnter \033[92mIran IPV4:\033[0m ")
	scanner.Scan()
	iranIP := scanner.Text()

	fmt.Print("\033[93mEnter \033[92mTunnel port:\033[0m ")
	scanner.Scan()
	tunnelPort := scanner.Text()

	fmt.Print("\033[93mHow many \033[92mConfigs\033[93m do you have?\033[0m ")
	scanner.Scan()
	numConfigsStr := scanner.Text()

	numConfigs, err := strconv.Atoi(numConfigsStr)
	if err != nil {
		fmt.Println("\033[91mPlz enter a valid number\033[0m")
		return
	}

	kharejPorts := make([]string, numConfigs)
	for i := 0; i < numConfigs; i++ {
		fmt.Printf("\033[93mEnter \033[92mconfig %d\033[93m port:\033[0m ", i+1)
		scanner.Scan()
		kharejPorts[i] = scanner.Text()
	}

	client := fmt.Sprintf(`[client]
remote_addr = "%s:%s"
default_token = "azumiisinyourarea"

[client.transport]
type = "tls"

[client.transport.tls]
trusted_root = "/root/rathole/rootCA.crt"
hostname = "localhost"

[client.transport.websocket] 
tls = true 

`, iranIP, tunnelPort)

	for i := 0; i < numConfigs; i++ {
		config := fmt.Sprintf(`[client.services.kharej%d]
local_addr = "127.0.0.1:%s"
`, i+startingNumber, kharejPorts[i])
		client += config
	}

	err = os.Remove("/root/rathole/client.toml")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting toml:\033[0m", err)
		return
	}

	file, err := os.Create("/root/rathole/client.toml")
	if err != nil {
		fmt.Println("\033[91merror creating toml:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(client)
	if err != nil {
		fmt.Println("\033[91merror putting configs into toml:\033[0m", err)
		return
	}

	service := `[Unit]
Description=Kharej-Azumi Service
After=network.target

[Service]
Type=simple
Restart=on-failure
RestartSec=5s
LimitNOFILE=1048576
ExecStart=/root/rathole/target/debug/rathole /root/rathole/client.toml

[Install]
WantedBy=multi-user.target`

	err = os.Remove("/etc/systemd/system/kharej-azumi.service")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting iran-azumi:\033[0m", err)
		return
	}

	file, err = os.Create("/etc/systemd/system/kharej-azumi.service")
	if err != nil {
		fmt.Println("\033[91merror creating iran-azumi:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(service)
	if err != nil {
		fmt.Println("\033[91merror constructing iran-azumi:\033[0m", err)
		return
	}

	cmd := exec.Command("systemctl", "daemon-reload")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror reloading:\033[0m", err)
		return
	}

	cmd = exec.Command("sudo", "chmod", "u+x", "/etc/systemd/system/kharej-azumi.service")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enablin da service:\033[0m", err)
		return
	}
	
	cmd = exec.Command("systemctl", "enable", "kharej-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enabling da service:\033[0m", err)
		return
	}

	cmd = exec.Command("systemctl", "restart", "kharej-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror restarting da service:\033[0m", err)
		return
	}
	resKharej()
	displayCheckmark("\033[92mService created successfully!\033[0m")
	if numConfigs == 1 {
	    fmt.Println("╭─────────────────────────────────────────────╮")
		fmt.Printf("\033[92m  Starting number for the next server:\033[96m %d\n\033[0m", startingNumber+1)
		fmt.Println("╰─────────────────────────────────────────────╯")
	} else {
	    fmt.Println("╭─────────────────────────────────────────────╮")
		fmt.Printf("\033[92m  Starting number for the next server:\033[96m %d\n\033[0m", numConfigs+startingNumber)
		fmt.Println("╰─────────────────────────────────────────────╯")
	}
}
func ws6Menu() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Reverse \033[92mWs + TLS \033[96mIPV6 \033[93mMenu\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. \033[92mIRAN\033[0m", "2. \033[93mKHAREJ\033[92m[1]\033[0m", "3. \033[93mKHAREJ\033[92m[2]\033[0m", "4. \033[93mKHAREJ\033[92m[3]\033[0m", "5. \033[93mKHAREJ\033[92m[4]\033[0m", "6. \033[93mKHAREJ\033[92m[5]\033[0m", "0. \033[94mBack to the main menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. \033[92mIRAN\033[0m":
		iranWs6M()
	case "2. \033[93mKHAREJ\033[92m[1]\033[0m":
		kharejWs6()
	case "3. \033[93mKHAREJ\033[92m[2]\033[0m":
		kharej2Ws6()
    case "4. \033[93mKHAREJ\033[92m[3]\033[0m":
		kharej2Ws6()
	case "5. \033[93mKHAREJ\033[92m[4]\033[0m":
		kharej2Ws6()
	case "6. \033[93mKHAREJ\033[92m[5]\033[0m":
		kharej2Ws6()
	case "0. \033[94mBack to the main menu\033[0m":
	    clearScreen()
		mainMenu()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}
func iranWs6M() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m IRAN \033[92mWs + TLS \033[93mMenu\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. \033[92mIRAN Config \033[96m[Method 1]\033[0m", "2. \033[92mIRAN Config \033[96m[Method 2]\033[0m", "3. \033[93mCopy Cert\033[0m", "0. \033[94mBack to the previous menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. \033[92mIRAN Config \033[96m[Method 1]\033[0m":
		iranWs6()
	case "2. \033[92mIRAN Config \033[96m[Method 2]\033[0m":
		iranWs62()
	case "3. \033[93mCopy Cert\033[0m":
		copy()
	case "0. \033[94mBack to the previous menu\033[0m":
	    clearScreen()
		ws4Menu()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}
func iranWs6() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Reverse \033[92mIPV6 \033[96mWS + TLS\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring IRAN")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	rootCA := "/root/rathole/rootCA.crt"

	if _, err := os.Stat(rootCA); os.IsNotExist(err) {
		ssl()
	} else {
		fmt.Println("\033[93mSkip getting Cert..\033[0m")
	}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\033[93mHow many \033[92mconfigs\033[93m do you have \033[96m[All Servers Combined]\033[93m? \033[0m")
	scanner.Scan()
	numConfigsStr := scanner.Text()

	numConfigs, err := strconv.Atoi(numConfigsStr)
	if err != nil {
		fmt.Println("\033[91mPlz enter a valid number\033[0m")
		return
	}

	fmt.Print("\033[93mEnter \033[92mTunnel port:\033[0m ")
	scanner.Scan()
	tunnelPort := scanner.Text()

	kharejPorts := make([]string, numConfigs)
	for i := 0; i < numConfigs; i++ {
		fmt.Printf("\033[93mEnter \033[92mConfig %d\033[93m Port: \033[0m", i+1)
		scanner.Scan()
		kharejPorts[i] = scanner.Text()
	}

	server := fmt.Sprintf(`[server]
bind_addr = "[::]:%s"
default_token = "azumiisinyourarea"

[server.transport]
type = "tls"

[server.transport.tls]
pkcs12 = "/root/rathole/identity.pfx"
pkcs12_password = "azumi1234"

[server.transport.websocket] 
tls = true 

`, tunnelPort)
	for i := 0; i < numConfigs; i++ {
		config := fmt.Sprintf(`[server.services.kharej%d]
bind_addr = "0.0.0.0:%s" 
`, i+1, kharejPorts[i])
		server += config
	}

	err = os.Remove("/root/rathole/server.toml")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting toml:\033[0m", err)
		return
	}

	file, err := os.Create("/root/rathole/server.toml")
	if err != nil {
		fmt.Println("\033[91merror creating toml:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(server)
	if err != nil {
		fmt.Println("\033[91merror putting configs into toml:\033[0m", err)
		return
	}
	service := `[Unit]
Description=Iran-Azumi Service
After=network.target

[Service]
Type=simple
Restart=on-failure
RestartSec=5s
LimitNOFILE=1048576
ExecStart=/root/rathole/target/debug/rathole /root/rathole/server.toml

[Install]
WantedBy=multi-user.target`

	err = os.Remove("/etc/systemd/system/iran-azumi.service")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting iran-azumi:\033[0m", err)
		return
	}

	file, err = os.Create("/etc/systemd/system/iran-azumi.service")
	if err != nil {
		fmt.Println("\033[91merror creating iran-azumi:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(service)
	if err != nil {
		fmt.Println("\033[91merror constructing iran-azumi:\033[0m", err)
		return
	}

	cmd := exec.Command("systemctl", "daemon-reload")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror reloading:\033[0m", err)
		return
	}

	cmd = exec.Command("sudo", "chmod", "u+x", "/etc/systemd/system/iran-azumi.service")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enablin da service:\033[0m", err)
		return
	}
	
	cmd = exec.Command("systemctl", "enable", "iran-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enabling da service:\033[0m", err)
		return
	}

	cmd = exec.Command("systemctl", "restart", "iran-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror restarting da service:\033[0m", err)
		return
	}
    resIran()
	displayCheckmark("\033[92mService created successfully!\033[0m")
}
func iranWs62() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Reverse \033[92mIPV6 \033[96mWS + TLS\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring IRAN")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	rootCA := "/root/rathole/rootCA.crt"

	if _, err := os.Stat(rootCA); os.IsNotExist(err) {
		ssl2()
	} else {
		fmt.Println("\033[93mSkip getting Cert..\033[0m")
	}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\033[93mHow many \033[92mconfigs\033[93m do you have \033[96m[All Servers Combined]\033[93m? \033[0m")
	scanner.Scan()
	numConfigsStr := scanner.Text()

	numConfigs, err := strconv.Atoi(numConfigsStr)
	if err != nil {
		fmt.Println("\033[91mPlz enter a valid number\033[0m")
		return
	}

	fmt.Print("\033[93mEnter \033[92mTunnel port:\033[0m ")
	scanner.Scan()
	tunnelPort := scanner.Text()

	kharejPorts := make([]string, numConfigs)
	for i := 0; i < numConfigs; i++ {
		fmt.Printf("\033[93mEnter \033[92mConfig %d\033[93m Port: \033[0m", i+1)
		scanner.Scan()
		kharejPorts[i] = scanner.Text()
	}

	server := fmt.Sprintf(`[server]
bind_addr = "[::]:%s"
default_token = "azumiisinyourarea"

[server.transport]
type = "tls"

[server.transport.tls]
pkcs12 = "/root/rathole/identity.pfx"
pkcs12_password = "azumi1234"

[server.transport.websocket] 
tls = true 

`, tunnelPort)
	for i := 0; i < numConfigs; i++ {
		config := fmt.Sprintf(`[server.services.kharej%d]
bind_addr = "0.0.0.0:%s" 
`, i+1, kharejPorts[i])
		server += config
	}

	err = os.Remove("/root/rathole/server.toml")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting toml:\033[0m", err)
		return
	}

	file, err := os.Create("/root/rathole/server.toml")
	if err != nil {
		fmt.Println("\033[91merror creating toml:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(server)
	if err != nil {
		fmt.Println("\033[91merror putting configs into toml:\033[0m", err)
		return
	}
	service := `[Unit]
Description=Iran-Azumi Service
After=network.target

[Service]
Type=simple
Restart=on-failure
RestartSec=5s
LimitNOFILE=1048576
ExecStart=/root/rathole/target/debug/rathole /root/rathole/server.toml

[Install]
WantedBy=multi-user.target`

	err = os.Remove("/etc/systemd/system/iran-azumi.service")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting iran-azumi:\033[0m", err)
		return
	}

	file, err = os.Create("/etc/systemd/system/iran-azumi.service")
	if err != nil {
		fmt.Println("\033[91merror creating iran-azumi:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(service)
	if err != nil {
		fmt.Println("\033[91merror constructing iran-azumi:\033[0m", err)
		return
	}

	cmd := exec.Command("systemctl", "daemon-reload")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror reloading:\033[0m", err)
		return
	}

	cmd = exec.Command("sudo", "chmod", "u+x", "/etc/systemd/system/iran-azumi.service")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enablin da service:\033[0m", err)
		return
	}
	
	cmd = exec.Command("systemctl", "enable", "iran-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enabling da service:\033[0m", err)
		return
	}

	cmd = exec.Command("systemctl", "restart", "iran-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror restarting da service:\033[0m", err)
		return
	}
    resIran()
	displayCheckmark("\033[92mService created successfully!\033[0m")
}
func kharejWs6() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Reverse \033[92mIPV6 \033[96mTCP\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring KHAREJ")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\033[93mEnter the \033[92mStarting number:\033[0m ")
	scanner.Scan()
	startingNumberStr := scanner.Text()

	startingNumber, err := strconv.Atoi(startingNumberStr)
	if err != nil {
		fmt.Println("\033[91mPlz enter a valid number\033[0m")
		return
	}

	fmt.Print("\033[93mEnter \033[92mIran IPV6:\033[0m ")
	scanner.Scan()
	iranIP := scanner.Text()

	fmt.Print("\033[93mEnter \033[92mTunnel port:\033[0m ")
	scanner.Scan()
	tunnelPort := scanner.Text()

	fmt.Print("\033[93mHow many \033[92mConfigs\033[93m do you have?\033[0m ")
	scanner.Scan()
	numConfigsStr := scanner.Text()

	numConfigs, err := strconv.Atoi(numConfigsStr)
	if err != nil {
		fmt.Println("\033[91mPlz enter a valid number\033[0m")
		return
	}

	kharejPorts := make([]string, numConfigs)
	for i := 0; i < numConfigs; i++ {
		fmt.Printf("\033[93mEnter \033[92mconfig %d\033[93m port:\033[0m ", i+1)
		scanner.Scan()
		kharejPorts[i] = scanner.Text()
	}

	client := fmt.Sprintf(`[client]
remote_addr = "[%s]:%s"
default_token = "azumiisinyourarea"

[client.transport]
type = "tls"

[client.transport.tls]
trusted_root = "/root/rathole/rootCA.crt"
hostname = "localhost"

[client.transport.websocket] 
tls = true 

`, iranIP, tunnelPort)

	for i := 0; i < numConfigs; i++ {
		config := fmt.Sprintf(`[client.services.kharej%d]
local_addr = "127.0.0.1:%s"
`, i+startingNumber, kharejPorts[i])
		client += config
	}

	err = os.Remove("/root/rathole/client.toml")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting toml:\033[0m", err)
		return
	}

	file, err := os.Create("/root/rathole/client.toml")
	if err != nil {
		fmt.Println("\033[91merror creating toml:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(client)
	if err != nil {
		fmt.Println("\033[91merror putting configs into toml:\033[0m", err)
		return
	}

	service := `[Unit]
Description=Kharej-Azumi Service
After=network.target

[Service]
Type=simple
Restart=on-failure
RestartSec=5s
LimitNOFILE=1048576
ExecStart=/root/rathole/target/debug/rathole /root/rathole/client.toml

[Install]
WantedBy=multi-user.target`

	err = os.Remove("/etc/systemd/system/kharej-azumi.service")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting iran-azumi:\033[0m", err)
		return
	}

	file, err = os.Create("/etc/systemd/system/kharej-azumi.service")
	if err != nil {
		fmt.Println("\033[91merror creating iran-azumi:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(service)
	if err != nil {
		fmt.Println("\033[91merror constructing iran-azumi:\033[0m", err)
		return
	}

	cmd := exec.Command("systemctl", "daemon-reload")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror reloading:\033[0m", err)
		return
	}

	cmd = exec.Command("sudo", "chmod", "u+x", "/etc/systemd/system/kharej-azumi.service")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enablin da service:\033[0m", err)
		return
	}
	
	cmd = exec.Command("systemctl", "enable", "kharej-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enabling da service:\033[0m", err)
		return
	}

	cmd = exec.Command("systemctl", "restart", "kharej-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror restarting da service:\033[0m", err)
		return
	}
	resKharej()
	displayCheckmark("\033[92mService created successfully!\033[0m")
	fmt.Println("╭─────────────────────────────────────────────╮")
	fmt.Printf("\033[92m Starting number for the next server : \033[96m%-9d\n\033[0m", numConfigs+1)
	fmt.Println("╰─────────────────────────────────────────────╯")
}
func kharej2Ws6() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Reverse \033[92mIPV6 \033[96mTCP\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring KHAREJ")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\033[93mEnter the \033[92mStarting number:\033[0m ")
	scanner.Scan()
	startingNumberStr := scanner.Text()

	startingNumber, err := strconv.Atoi(startingNumberStr)
	if err != nil {
		fmt.Println("\033[91mPlz enter a valid number\033[0m")
		return
	}

	fmt.Print("\033[93mEnter \033[92mIran IPV6:\033[0m ")
	scanner.Scan()
	iranIP := scanner.Text()

	fmt.Print("\033[93mEnter \033[92mTunnel port:\033[0m ")
	scanner.Scan()
	tunnelPort := scanner.Text()

	fmt.Print("\033[93mHow many \033[92mConfigs\033[93m do you have?\033[0m ")
	scanner.Scan()
	numConfigsStr := scanner.Text()

	numConfigs, err := strconv.Atoi(numConfigsStr)
	if err != nil {
		fmt.Println("\033[91mPlz enter a valid number\033[0m")
		return
	}

	kharejPorts := make([]string, numConfigs)
	for i := 0; i < numConfigs; i++ {
		fmt.Printf("\033[93mEnter \033[92mconfig %d\033[93m port:\033[0m ", i+1)
		scanner.Scan()
		kharejPorts[i] = scanner.Text()
	}

	client := fmt.Sprintf(`[client]
remote_addr = "[%s]:%s"
default_token = "azumiisinyourarea"

[client.transport]
type = "tls"

[client.transport.tls]
trusted_root = "/root/rathole/rootCA.crt"
hostname = "localhost"

[client.transport.websocket] 
tls = true 

`, iranIP, tunnelPort)

	for i := 0; i < numConfigs; i++ {
		config := fmt.Sprintf(`[client.services.kharej%d]
local_addr = "127.0.0.1:%s"
`, i+startingNumber, kharejPorts[i])
		client += config
	}

	err = os.Remove("/root/rathole/client.toml")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting toml:\033[0m", err)
		return
	}

	file, err := os.Create("/root/rathole/client.toml")
	if err != nil {
		fmt.Println("\033[91merror creating toml:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(client)
	if err != nil {
		fmt.Println("\033[91merror putting configs into toml:\033[0m", err)
		return
	}

	service := `[Unit]
Description=Kharej-Azumi Service
After=network.target

[Service]
Type=simple
Restart=on-failure
RestartSec=5s
LimitNOFILE=1048576
ExecStart=/root/rathole/target/debug/rathole /root/rathole/client.toml

[Install]
WantedBy=multi-user.target`

	err = os.Remove("/etc/systemd/system/kharej-azumi.service")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("\033[91merror deleting iran-azumi:\033[0m", err)
		return
	}

	file, err = os.Create("/etc/systemd/system/kharej-azumi.service")
	if err != nil {
		fmt.Println("\033[91merror creating iran-azumi:\033[0m", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(service)
	if err != nil {
		fmt.Println("\033[91merror constructing iran-azumi:\033[0m", err)
		return
	}

	cmd := exec.Command("systemctl", "daemon-reload")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror reloading:\033[0m", err)
		return
	}

	cmd = exec.Command("sudo", "chmod", "u+x", "/etc/systemd/system/kharej-azumi.service")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enablin da service:\033[0m", err)
		return
	}
	
	cmd = exec.Command("systemctl", "enable", "kharej-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror enabling da service:\033[0m", err)
		return
	}

	cmd = exec.Command("systemctl", "restart", "kharej-azumi")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91merror restarting da service:\033[0m", err)
		return
	}
	resKharej()
	displayCheckmark("\033[92mService created successfully!\033[0m")
	if numConfigs == 1 {
	    fmt.Println("╭─────────────────────────────────────────────╮")
		fmt.Printf("\033[92m  Starting number for the next server:\033[96m %d\n\033[0m", startingNumber+1)
	    fmt.Println("╰─────────────────────────────────────────────╯")
	} else {
	    fmt.Println("╭─────────────────────────────────────────────╮")
		fmt.Printf("\033[92m  Starting number for the next server:\033[96m %d\n\033[0m", numConfigs+startingNumber)
	    fmt.Println("╰─────────────────────────────────────────────╯")
	}
}

    
	
