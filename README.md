## Modbus reader for 3 to 3 Phase AEG UPS

#### aeg3P3_modbus_rtu

Read modbus tcp/rtu regestry from 1 to 86 address, and returns a json object.

Programm flags:

-serial - connected serial port (defaut value "/dev/ttyRS485-1")

-speed  - connection speed (defaut value 9600)

-id -  modbus slave ID (defaut value 1)

-t - timeout milisec (defaut value 3000)

-q - quantity of janitza modbus arguments, value range 1 - 86 (defaut value 86)

-rtype - request type tcp/rtu - true/false (default value true)

-ip - modbus slave ip adress (defaut value localhost)

-port - modbus tcp port (defaut value 502)

=======================================

**Note:** serial port static config "8E1" 


###### Example:

For modbus-rtu:

`aeg3P3_modbus_rtu -serial=/dev/ttyRS485-1 -rtype=false -speed0=19200 -id=1 -q=10`

For modbus-tcp:

`aeg3P3_modbus_rtu -ip=192.168.10.10 -rtype=true -port=503 -id=1 -q=10`