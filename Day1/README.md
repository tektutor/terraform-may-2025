# Day 1

## Info - Dual/multi booting  - dynamic inventory ( python script )

<pre>
- let's say you already have windows OS installed on your laptop, for R&D purpose you need some Linux OS
- you could use Boot Loader utilities like LILO(deprecated/outdated), GRUB, BootCamp(Mac)
- Master Boot Record(MBR - Hard Disk - Byte 0 Sector 0) - 512 bytes
  - Boot Loader software is installed in the MBR 
- When we power on a machine, BIOS will perform POST (Power On Self Test)
  - BIOS will instruct the CPU to launch the Boot Loader software residing in MBR
  - Boot Loader then scans all your hard disks looking for Operating Systems, if it finds more than one OS, it gives a menu for you to choose between those OS
- though more than one OS is installed on your system, only one can be actively running at any point of time
</pre>  

## Info - Hypervisor Overview
<pre>
- a virtualization technology
- running more than one Operating System on the same machine ( laptop/desktop/workstation/server )
- more than one OS can actively running
- examples
  - Type 1 ( is used in workstations/servers )
    - VMWare vSphere/Vcenter
    - KVM ( Opensource )
  - Type 2 ( are used in desktops/laptops/workstations )
    - Oracle Virtualbox ( Free )
    - Parallels (Mac OS-X)
    - VMWare Fusion(Mac OS-X)
    - VMWare Workstation ( Windows & Linux )
    - Microsoft Hyper-V  - dynamic inventory ( python script )

- this type of virtualization is called heavy-weight
  - for each Virtual Machine(VM) we need to allocated dedicated hardware resources
    - CPU Cores
    - RAM
    - Storage ( HDD/SSD )
- Hyperthreading
  - each physical cpu core can support two logical(virtual) core
  - i.e if you have a machine with 4 Cores, we have total 4x2 = 8 virtual cores
- the max number of VMs that can be supported by a machine is based on
  - how many virtual cores your machine has ( Primary factor )
  - how much RAM your machine has
  - how much storage your machine has
- each VM represents one Operating System
</pre>

## Info - Containerization
<pre>
- light-weight application virtualization technology
- all containers running on the same OS(machine) they share the hardware resources on the underlying os
- each container represents on one application
- containers are not OS
- containers will never be able to replace OS/VM
- containers don't have their own OS Kernel
- similarities between containers and virtual machines(OS)
  - just like each VM get's it own IP address, each container gets its own IP address 
  - just like each VM has its own Network card, each container has its own network card
  - just like each VM has its own file system, each container has its own file system
  - just like each VM has its own port range, each container has its own port range 0-65535 ports
- examples
  - Docker, Podman, containerd, etc
</pre>
Commands
## Info - Container Engine
<pre>
- is a high-level software that helps managing container images and containers
- is very user-friendly
- container engines internally depends on Container Runtimes to manage images and containers
- examples
  - Docker Container Engine depends on containerd which in turn depends on runC container runtime
  - Podman Container Engine depends on CRI-O container runtime
</pre>  

## Info - Container Runtime
<pre>
- is a low-level software that helps managing container images and containers
- is not so user-friendly, hence normally no end-user use this
- examples
  - runC is a container runtime
  - CRI-O is a container runtime
</pre>

## Info - Configuration Management
<pre>
- helps in automating all the system administration activities
- if you already have machine with OS pre-installed, you could use configuration management tools
  - to install/uninstall/upgrade softwares into that machine
  - the machine where the automation is done is called nodes ( Unix, Linux,Mac, Windows, Network Switches/Routers, etc., )
- examples
  - Ansible
  - Puppet
  - Chef
  - Salt/Saltstack ( almost dead )
</pre>

## Info - Provisioners
<pre>
- these tools helps us create a virtual machine in your data center, private/public cloud ( AWS, Azure, GCP, Digital Ocean etc., )
- examples
  - Docker
  - AWS Cloudformation ( only supports automating infrastructure in AWS environment )
  - Terraform ( Cloud newtral, works locally or in public/private/hybrid cloud )
</pre>

## Info - Ansible Configuration Management Tool
<pre>
- opensource
- developed by Michael Deehan ( former employee of Red Hat )
- is developed in Python language
- is agent-less
- follows PUSH based architecture
- Domain Specific Language (DSL)
  - the language in which automation code is written
  - YAML in case of Ansible
- Michael Deehan left Red Hat and founded a company called Ansible Inc, developed Ansible Core as an opensource project
- comes in 3 flavours
  - Ansible Core
    - open source
    - command-line linux tool
    - though Ansible core can only be installed in Linux Distributions it can manage servers with any OS (  Windows, Mac, Unix, Linux, etc )
    - the machine where Ansible is installed is called Ansible Controller Machine
    - the machine where the automation is performed is called Ansible Node
    - Ansible Node ( Remote Servers )
      - can be a container
      - a local virtual machine
      - an ec2 instance running on AWS
      - an Azure VM 
      - a physical server with Windows/Unix/Linux/Mac
      - Network Switch/Router
  - AWX
    - opensource
    - developed on top of Ansible Core
    - supports web-interface
    - you don't get any support from any organization in case you are struck with some issues in AWX
  - Ansible Tower ( these days called as Ansible Automation Platform )
    - an enterprise product from Red Hat
    - is developed on top of opensource AWX
    - hence supports all the features of AWX
    - you get world-wide support from Red Hat ( an IBM company )
</pre>

## Info - Installing Ansible core in Ubuntu ( Just for your reference - you don't have to do this in lab machine )
```
sudo apt update
sudo apt install software-properties-common
sudo add-apt-repository --yes --update ppa:ansible/ansible
sudo apt install ansible
```

## Info - Ansible Inventory
<pre>
- is a text file with no extension
- technically, the file name can be anything but as a best practice name it either hosts or inventory
- any plain text editor can be used to create the inventory file
- it captures the ansible node(remote servers) connectivity details
- there are 2 types of inventory
  - static inventory ( text file )
  - dynamic inventory ( python script )
</pre>

## Info - Ansible Module
<pre>
- ansible comes out of the box with many ansible modules
- each ansible modules automates one functionality
- for instance, 
  - copy module - helps in copying a file from local machine to the ansible node or vice versa
  - file module - helps in creating a file/directory on the remote machine with specific permissions
  - service module - helps in managing service 
    - start a service
    - enable/disable a service
    - restart/reload a service
- ansible modules are implemented either as Python scripts or as Powershell scripts
  - to automate configuration management on  windows ansible nodes, ansible supports modules written in Powershell script
  - to perform configuration management on unix/linux/mac ansible nodes, ansible supports modules written in Python
</pre>

## Lab - Cloning the training repository on your RPS Cloud Lab machine ( From the Ubuntu terminal )
```
cd ~
git clone https://github.com/tektutor/terraform-may-2025.git
cd terraform-may-2025
ls -l | grep terraform-may-2025
```

Expected output
![image](https://github.com/user-attachments/assets/d3187df3-dac5-4604-8846-609f6e02ac65)
![image](https://github.com/user-attachments/assets/280df926-e57e-42a3-bdf5-7fd30032f16c)
