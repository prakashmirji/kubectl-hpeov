### Overview

This repo hosts the code and binary for kubectl extension for hpe oneview


The goal of the project is to allow K8s user to use their regular kubectl commands to interact
with underlying bare metal servers. 

It allows K8s users with following use cases

- List server details ( name, status, model, state...etc), you can below kubectl command
    ```
    kubectl hpeov serverhardware get --all
    ```
- Power On/Off server, you can use below kubectl command
    ```
    kubectl hpeov serverhardware power --name=<server name> --powerstate=On
    ```
- Provisioning of bare metal (applying configuration)
    ```
    kubectl hpeov serverprofile create --profilename=<name of server profile> --templatename=<name of server template>
    ```
- Deprovisioning of bare metal
    ```
    kubectl hpeov serverprofile delete --profilename=<name of server profile>
    ```
- List available templates
    ```
    kubectl hpeov servertemplate get --all
    ```
- List running servers having profiles
    ```
    kubectl hpeov serverprofile get --profilename=<name of server profile>
    ```

Note: You can query the resources using name flag. For example


    kubectl hpeov serverhardware get --name=<name of server hardware>


Here 'hpeov' is the plugin name.


This is Go CLi application based on Cobra. It is still in development phase.


### Pre-reqs for test
    - Go 1.14 or higher
    - k8s ( can be simple Kind based k8s cluster)
    - Access to HPE Oneview
        Depends on https://github.com/HewlettPackard/oneview-golang


### How to build
Clone the repo

    git clone https://github.com/prakashmirji/kubectl-hpeov
    cd to kubectl-hpeov
    go install <path> or
    go build -o <name>

    For example

    go build -o hpeov

Copy the binary you built to /usr/local/bin/ directory

For example

    >cp hpeov /usr/local/bin/kubectl-hpeov
    >chmod +x /usr/local/bin/kubectl-hpeov

The format and name when you copy to /usr/local/bin should be in the form 'kubectl-(plugin name)'

### How to test using kubectl
If you copied the binary using above steps, edit env file


    >vi env.sh


Update the OneView details and source it like


    >. ./env.sh


Try sample command like this


    >kubectl hpeov serverhardware get --all


You will see a result that looks this

    
    Server Name		Power State	Model		Memory	Status	iLO Address	Profle State
    0000A66102, bay 3	On		SY 660 Gen9	294912	OK	    172.18.6.10	NoProfileApplied
    0000A66103, bay 3	On	    SY 660 Gen9	294912	OK	    172.18.6.17	NoProfileApplied
    0000A66103, bay 5	On	    SY 480 Gen9	147456	OK	    172.18.6.19	NoProfileApplied
    0000A66102, bay 4	On	    SY 660 Gen9	294912	OK	    172.18.6.11	NoProfileApplied
    0000A66103, bay 4	On	    SY 660 Gen9	294912	OK	    172.18.6.18	NoProfileApplied
    0000A66102, bay 7	Off	    SY 480 Gen9	147456	OK	    172.18.6.14	NoProfileApplied
    0000A66102, bay 5	On	    SY 480 Gen9	147456	OK	    172.18.6.12	NoProfileApplied
    0000A66102, bay 8	Off	    SY 480 Gen9	147456	OK	    172.18.6.15	NoProfileApplied
    
You can get the CLI help message using below command
    
    >./kubectl-hpeov -h  

It prints help like this

        
    A kubectl extension for hpe oneview product. For example:

        kubectl hpeov serverhardware get --all
        kubectl hpeov serverhardware get --name=<name of server hardware> 
        kubectl hpeov serverhardware power --name=<server name> --powerstate=On
        kubectl hpeov serverprofile create --file=<json payload to create profile>
        kubectl hpeov serverprofile get --all
        kubectl hpeov serverprofile get --name=<name of the profile>

    Usage:
    hpeov [command]

    Available Commands:
    help           Help about any command
    serverhardware A subcommand of hpeov cli for getting server hardware details
    serverprofile  A subcommand of hpeov cli for operating with server profile
    servertemplate A subcommand of hpeov cli for getting server template details

    Flags:
        --config string   config file (default is $HOME/.kubectl-hpeov.yaml)
    -h, --help            help for hpeov
    -t, --toggle          Help message for toggle

    Use "hpeov [command] --help" for more information about a command.

### How to test during development without kubectl
Just build the go CLI binary and test it with commands, args and flags.
Use help command to know the args and flags currently supported

### Contributions
Looking for contributors who can help in adding more features
