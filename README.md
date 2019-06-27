# tcp-snmp-relay
A lightweight Go program to relay TCP strings as SNMP Traps.


## Usage
From a command prompt run ".\tcp-snmp-relay.exe". This will open a TCP socket on port 8162. TCP messages must be formatted as follows:
TrapDestinationIP,OID,SNMP Community. For example:
```
192.168.1.1,.1.3.6.1.4.1.42463.12.3.0.1,public
```

You can send a trap for any OID, but I have created a custom MIB file (based off the existing Telos xNode MIB) that includes 10 Generic alarms you can trigger. OIDs are as follows:
```
  .1.3.6.1.4.1.42463.12.3.0.1 | ALARM_01_ON
  .1.3.6.1.4.1.42463.12.3.0.2 | ALARM_02_ON
  .1.3.6.1.4.1.42463.12.3.0.3 | ALARM_03_ON
  .1.3.6.1.4.1.42463.12.3.0.4 | ALARM_04_ON
  .1.3.6.1.4.1.42463.12.3.0.5 | ALARM_05_ON
  .1.3.6.1.4.1.42463.12.3.0.6 | ALARM_06_ON
  .1.3.6.1.4.1.42463.12.3.0.7 | ALARM_07_ON
  .1.3.6.1.4.1.42463.12.3.0.8 | ALARM_08_ON
  .1.3.6.1.4.1.42463.12.3.0.9 | ALARM_09_ON
  .1.3.6.1.4.1.42463.12.3.0.10 | ALARM_10_ON
  
````

You can customize the MIB file to add more alarms or to customize the names. 
## Triggering alarms on a Burk ArcPlus

Download the custom Telos MIB and SMI files and place them in C:\ProgramData\Burk Technology\AutoLoad Plus\MIBs (or wherever your default MIB directory is). 

In AutoLoad, connect to the ArcPlus you wish to send traps to and create a new SNMP device. The IP address of the device must match the IP address of the computer that is running the tcp-snmp-relay. Select the custom TelosTraps.mib file. 

Create a new SNMP trap to monitor for by right clicking on the ArcPlus>Settings>SNMP Traps. Add an SNMP Trap to monitor and use the following settings: 
  Version: v2Trap
  Type: specific
  Selected Trap/Inform: [Alarm1-10]
  
You can then customize the Alarm message to your liking. (You can also do this in the MIB file if you prefer). Choose the alert settings as desired and then save the configuration.

# Notes
Only SNMPv2c supported right now. No input sanitation/validation. No error handling. As of right now I can't find a way to clear an alarm via an SNMP trap. 
