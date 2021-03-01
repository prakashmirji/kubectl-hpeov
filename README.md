This repo hosts the code and binary for kubectl extension for hpe oneview

This is Go CLi application based on Cobra. It is still in development phase.

### Pre-reqs for test
    - Go 1.14 or higher
    - k8s ( can be simple Kind based k8s cluster)
    - Access t0 HPE Oneview
        Depends on https://github.com/HewlettPackard/oneview-golang


### How to build

    git clone https://github.com/prakashmirji/kubectl-hpeov
    cd to kubectl-hpeov
    go install <path> or
    go build -o <name>

    Copy the binary to /usr/local/bin/ as kubectl-<name>
    For example
    >cp hpeov /usr/local/bin/kubectl-hpeov
    >chmod +x /usr/local/bin/kubectl-hpeov

### How to test using kubectl
    If you copied the binary using above steps, then you can 
    >kubectl hpeov serverhardware get --all

    You will see a result that looks this
    ```
    Server Name		Power State	Model		Memory	Status	iLO Address
    0000A66102, bay 3	On		SY 660 Gen9	294912	OK	172.18.6.10
    0000A66103, bay 3	On	SY 660 Gen9	294912	OK	172.18.6.17
    0000A66103, bay 5	On	SY 480 Gen9	147456	OK	172.18.6.19
    0000A66102, bay 4	On	SY 660 Gen9	294912	OK	172.18.6.11
    0000A66103, bay 4	On	SY 660 Gen9	294912	OK	172.18.6.18
    0000A66102, bay 7	Off	SY 480 Gen9	147456	OK	172.18.6.14
    0000A66102, bay 5	On	SY 480 Gen9	147456	OK	172.18.6.12
    0000A66102, bay 8	Off	SY 480 Gen9	147456	OK	172.18.6.15
    0000A66103, bay 6	On	SY 660 Gen10	65536	OK	172.18.31.6
    0000A66102, bay 6	Off	SY 660 Gen10	65536	OK	172.18.31.4
    0000A66101, bay 4	On	SY 660 Gen9	294912	OK	172.18.6.5
    0000A66101, bay 3	On	SY 660 Gen9	294912	OK	172.18.6.4
    0000A66103, bay 7	On	SY 480 Gen9	147456	OK	172.18.6.21
    0000A66101, bay 5	On	SY 480 Gen9	147456	OK	172.18.6.2
    0000A66102, bay 11	Off	SY 480 Gen10	32768	OK	172.18.31.3
    0000A66101, bay 6	On	SY 660 Gen10	65536	OK	172.18.31.2
    0000A66103, bay 8	On	SY 480 Gen9	147456	OK	172.18.6.22
    0000A66103, bay 11	On	SY 480 Gen10	32768	OK	172.18.31.5
    0000A66101, bay 8	On	SY 480 Gen9	147456	OK	172.18.6.7
    0000A66101, bay 7	On	SY 480 Gen9	147456	OK	172.18.6.6
    0000A66101, bay 11	On	SY 480 Gen10	32768	OK	172.18.31.1
    ````
    You can get the CLI help message
    >./kubectl-hpeov -h  

        ```
        A kubectl extension for hpe oneview product. For example:

        kubectl hpeov serverhardware get --all
        kubectl hpeov serverhardware get --name=<name of server hardware> 
        kubectl hpeov serverhardware power --name=<server name> --powerstate=On
        kubectl hpeov serverprofile create --file=<json payload to create profile>.
        kubectl hpeov serverprofile get --all
        kubectl hpeov serverprofile get --name=<name of the profile>

        Usage:
        hpeov [command]

        Available Commands:
        help           Help about any command
        serverhardware A subcommand of hpeov cli for getting server hardware details

        Flags:
            --config string   config file (default is $HOME/.kubectl-hpeov.yaml)
        -h, --help            help for hpeov
        -t, --toggle          Help message for toggle

        Use "hpeov [command] --help" for more information about a command.
        ```

### How to test during development without kubectl
Just build the go CLI binary and test it with commands, args and flags.
Use help command to know the args and flags currently supported

### Contributions
Looking for contributors who can help in adding more features
