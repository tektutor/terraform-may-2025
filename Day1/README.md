# Day 1

## Info - Dual/multi booting
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
    - Microsoft Hyper-V
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
</pre>
