# windows-battery-remainder

[![Go Report Card](https://goreportcard.com/badge/github.com/sam0392in/windows-battery-remainder)](https://goreportcard.com/report/github.com/sam0392in/windows-battery-remainder)

![alt text](./img/alert.png?raw=true)

This tool when installed on windows machine, reminds user to plug the battery when battery % is lower than 20% and unplug battery once battery is fully charged.

## OS Supported
Windows Version >= 10

## Setup

1. Download batter-alert.exe binary from releases section of this repo.

2. Open Task Schedular and click on create task.

3. Enter the name : batter-alert and specified details in general section
![alt text](./img/ts1.png?raw=true)

4. Setup trigger to run the binary for every 5 minutes.
![alt text](./img/ts2.png?raw=true)

5. Specify the location of binary in double quotes in Action setion.
![alt text](./img/ts3.png?raw=true)

6. Click save and enjoy. Now you save your battery from over-charging.