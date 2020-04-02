## Modbus reader for 3 to 3 Phase AEG UPS

#### aeg3P3_modbus_rtu

Read modbus regestry from 1 to 86 address, and returns a json object.

Programm flags:

-serial - connected serial port (defaut value "/dev/ttyRS485-1")

-speed - connection speed (defaut value 9600)

-id - modbus slave ID (defaut value 1)
	
-q - quantity of janitza modbus arguments, value range 1 - 86 (defaut value 86)

**Note:** serial port static config "8E1" 




Example:

`aeg3P_modbus_rtu -serial=/dev/ttyRS485-1 -speed0=19200 -id=1 -q=10`
