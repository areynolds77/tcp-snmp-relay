# tcp-snmp-relay
A lightweight Go program to relay TCP strings as SNMP Traps.


## Usage
From a command prompt run ".\tcp-snmp-relay.exe". This will open a TCP socket on port 8162. TCP messages must be formatted as follows:
TrapDestinationIP,OID,SNMP Community. For example:
```
192.168.1.1,.1..3.6.1.4.1.42463.12.3.0.1,public
```
## Triggering alarms on a Burk ArcPlus

Download the custom Telos MIB and SMI files and place them in C:\ProgramData\Burk Technology\AutoLoad Plus\MIBs (or wherever your default MIB directory is). 

In AutoLoad, connect to the ArcPlus you wish to send traps to and create a new SNMP device. The IP address of the device must match the IP address of the computer that is running the tcp-snmp-relay. Select the custom TelosTraps.mib file. 

Create a new SNMP trap to monitor for by right clicking on the ArcPlus>Settings>SNMP Traps. Add an SNMP Trap to monitor and use the following settings: 
  Version: v2Trap
  Type: specific
  Selected Trap/Inform: [Alarm1-10]
  
You can then customize the Alarm message to your liking. (You can also do this in the MIB file if you prefer). Choose the alert settings as desired and then save the configuration.

# Notes
Only SNMPv2c supported right now. No input sanitation/validation. No error handling. 
