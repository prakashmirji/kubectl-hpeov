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
    If you copied the binary using above steps, edit env file
    >vi env.sh
    Update the OneView details and source it like
    >. ./env.sh
    >kubectl hpeov serverhardware get --all

    You will see a result that looks this
    ```
    Server Name		Power State	Model		Memory	Status	iLO Address	Profle State
    0000A66102, bay 3	On		SY 660 Gen9	294912	OK	    172.18.6.10	NoProfileApplied
    0000A66103, bay 3	On	    SY 660 Gen9	294912	OK	    172.18.6.17	NoProfileApplied
    0000A66103, bay 5	On	    SY 480 Gen9	147456	OK	    172.18.6.19	NoProfileApplied
    0000A66102, bay 4	On	    SY 660 Gen9	294912	OK	    172.18.6.11	NoProfileApplied
    0000A66103, bay 4	On	    SY 660 Gen9	294912	OK	    172.18.6.18	NoProfileApplied
    0000A66102, bay 7	Off	    SY 480 Gen9	147456	OK	    172.18.6.14	NoProfileApplied
    0000A66102, bay 5	On	    SY 480 Gen9	147456	OK	    172.18.6.12	NoProfileApplied
    0000A66102, bay 8	Off	    SY 480 Gen9	147456	OK	    172.18.6.15	NoProfileApplied
    ````
    You can get the CLI help message
    >./kubectl-hpeov -h  

        ```
        A kubectl extension for hpe oneview product. For example:

        kubectl hpeov serverhardware get --all
        kubectl hpeov serverhardware get --name=<name of server hardware> 
        kubectl hpeov serverhardware power --name=<server name> --powerstate=On
        kubectl hpeov serverprofile get --all
        kubectl hpeov serverprofile get --profilename=<name of server profile> 
        kubectl hpeov serverprofile create --profilename=<name of server profile> --templatename=<name of server template>
        kubectl hpeov serverprofile delete --profilename=<name of server profile>
        kubectl hpeov servertemplate get --all
        kubectl hpeov servertemplate get --name=<templa name>


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
