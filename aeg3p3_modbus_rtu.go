package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"github.com/goburrow/modbus"
	"os"
	"time"
)

/*
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!!!!!!!!!!! VERSION !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
*/
const version = "0.0.1"

/*
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!!!!!!!!!!! VERSION !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
*/

type par struct {
	reg     uint16
	regtype uint
	regName string
}
type param []par

var aegPar = param{
	{1, 0, "UPSStatus"},
	{2, 1, "NonCriticalFault"},
	{3, 1, "CriticalFault"},
	{4, 1, "Userdef_Digln1Аctive"},
	{5, 1, "Userdef_Digln2Аctive"},
	{6, 1, "Userdef_Digln3Аctive"},
	{7, 1, "EmergencySwitchOff"},
	{8, 1, "DouCanFault"},
	{9, 1, "MainsFault"},
	{10, 1, "RectifierWarning"},
	{11, 1, "RectifierFault"},
	{12, 1, "BatteryАlarm"},
	{13, 1, "lnverterWarning"},
	{14, 1, "lnverterFault"},
	{15, 1, "SBSWarning"},
	{16, 1, "SBSFault"},
	{17, 1, "EqualisingCharge"},
	{18, 1, "Charge"},
	{19, 1, "TrickleCharge"},
	{20, 1, "GenSetOperation"},
	{21, 1, "BattTemp.SensFault"},
	{22, 1, "BatteryTemp.ToHigh"},
	{23, 1, "CircuitBreaker"},
	{24, 1, "BatteryWarning"},
	{25, 1, "Batterylow"},
	{26, 1, "Warning_InverterOverload"},
	{27, 1, "FanFault"},
	{28, 1, "Alarm_InverterOverload"},
	{29, 1, "ShortCircuit"},
	{30, 1, "DcUnderVoltage"},
	{31, 1, "DcOverVoltage"},
	{32, 1, "PowerStackOverТemp"},
	{33, 1, "SBSReady"},
	{34, 1, "SBSMainsFault"},
	{35, 1, "SBSBlocked"},
	{36, 1, "RectifierOn"},
	{37, 1, "1nverterOn"},
	{38, 1, "SBSOn"},
	{39, 2, "RectMainsFreq."},
	{40, 1, "RectMainsVoltL1"},
	{41, 1, "RectMainsVoltL2"},
	{42, 1, "RectMainsVoltLЗ"},
	{43, 2, "SBSMainsFreq."},
	{44, 1, "SBSMainsVoltL1"},
	{45, 1, "SBSMainsVoltL2"},
	{46, 1, "SBSMainsVoltLЗ"},
	{47, 1, "BatteryVoltage"},
	{48, 3, "BatteryCurrent"}, //!!!!!!!!!!!!!! Надо проверить
	{49, 2, "АutonomyTime"},
	{50, 1, "BatteryCapacity"},
	{51, 3, "BatteryTemperature"}, //!!!!!!!!!!!!!! Надо проверить
	{52, 2, "OutputFreq."},
	{53, 1, "OutputVoltageL1"},
	{54, 1, "OutputLoadL1"},
	{55, 1, "OutputCurrentL1"},
	{56, 2, "OutputPowerL1"},
	{57, 1, "OutputVoltageL2"},
	{58, 1, "OutputLoadL2"},
	{59, 1, "OutputCurrentL2"},
	{60, 2, "OutputPowerL2"},
	{61, 1, "OutputVoltageLЗ"},
	{62, 1, "OutputLoadLЗ"},
	{63, 1, "OutputCurrentLЗ"},
	{64, 2, "OutputPowerLЗ"},
	{65, 1, "LifeCheck"},
	{66, 1, "Userdef_AUX1-Rect."},
	{67, 1, "Userdef_AUX2-Rect."},
	{68, 1, "Userdef_AUXЗ-Rect."},
	{69, 1, "Userdef_AUX4-Rect."},
	{70, 1, "Userdef_AUX5-Rect."},
	{71, 1, "Userdef_AUX6-Rect."},
	{72, 1, "Userdef_AUX7-Rect."},
	{73, 1, "Userdef_AUX1-lnv."},
	{74, 1, "Userdef_AUX2-lnv."},
	{75, 1, "Userdef_AUXЗ-lnv."},
	{76, 1, "Userdef_AUX4-lnv."},
	{77, 1, "Userdef_AUX5-lnv."},
	{78, 1, "Userdef_AUX6-lnv."},
	{79, 1, "Userdef_AUX7-lnv."},
	{80, 1, "Userdef_AUX1-SBS"},
	{81, 1, "Userdef_AUX2-SBS"},
	{82, 1, "Userdef_AUXЗ-SBS"},
	{83, 1, "Userdef_AUX4-SBS"},
	{84, 1, "Userdef_AUX5-SBS"},
	{85, 1, "Userdef_AUX6-SB"},
	{86, 1, "Userdef_AUX7-SBS"},
}

type respStruct struct {
	namePar  string
	valuePar string
}

var respons []respStruct

func main() {

	serialPort := flag.String("serial", "/dev/ttyRS485-1", "a string")
	serialSpeed := flag.Int("speed", 9600, "a int")
	slaveID := flag.Int("id", 1, "an int")
	timeout := flag.Int("t", 3000, "an int")
	typOfReg := flag.Int("type", 1, "an int")
	regQuantity := flag.Uint("q", 86, "an uint")
	/*dataBits := flag.Int("dbits", 1, "a int")
	parity := flag.String("parity", "E", "a string")
	stopBits := flag.Int("sbits", 1, "an int")*/
	flag.Parse()

	handler := modbus.NewRTUClientHandler(fmt.Sprint(*serialPort))
	handler.BaudRate = *serialSpeed
	handler.SlaveId = byte(*slaveID)
	handler.Timeout = time.Duration(*timeout) * time.Millisecond
	/*handler.DataBits = int(*dataBits)
	handler.Parity = fmt.Sprint(*parity)
	handler.StopBits = *stopBits*/

	defer handler.Close()
	client := modbus.NewClient(handler)

	/**
	Конфигурация для чтения параметров из json файла ./config.json
	*/
	/*		data, err := ioutil.ReadFile("./config.json")
			if err != nil {
				fmt.Print(err)
			}
			type Config struct {
				DEV       string
				SPEED     int
				DATABITS  int
				PARITY    string
				STOPBITS  int
				TIMEOUT   time.Duration
				ID        int
				STARTREG  int
				QUANTITY  int
				TYPEOFREG int
			}
			var conf Config
			err = json.Unmarshal(data, &conf)
			if err != nil {
				printError(err)
			}

			handler := modbus.NewRTUClientHandler(conf.DEV)
			handler.BaudRate = conf.SPEED
			handler.DataBits = conf.DATABITS
			handler.Parity = conf.PARITY
			handler.StopBits = conf.STOPBITS
			handler.SlaveId = byte(conf.ID)
			handler.Timeout = conf.TIMEOUT * time.Millisecond
	*/
	if *typOfReg == 1 {
		results, err := client.ReadHoldingRegisters(uint16(1), uint16(*regQuantity))
		if err != nil {
			printError(err)
		}
		printResult(results)
		/*fmt.Println(len(results))
		fmt.Println(results)
		os.Exit(0)*/
	}

	if *typOfReg == 2 {
		results, err := client.ReadInputRegisters(uint16(1), uint16(*regQuantity))
		if err != nil {
			printError(err)
		}
		printResult(results)
		/*fmt.Println(len(results))
		fmt.Println(results)
		os.Exit(0)*/
	}
	fmt.Printf("{\"status\":\"error\", \"error\":\"typeofreg not 1 or 2 \"} \n")

	/*var ttt = []byte{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 1, 244, 0, 226, 0, 225, 0, 225, 1, 244, 0, 230, 0, 229, 0, 229, 0, 243, 0, 0, 5, 220, 0, 100, 4, 86, 1, 244, 0, 219, 0, 6, 0, 1, 0, 1, 0, 219, 0, 19, 0, 2, 0, 1, 0, 219, 0, 43, 0, 4, 0, 4, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	var ttt = []byte{}
	printResult(ttt)*/
}

func printError(err error) {
	fmt.Printf("{\"status\":\"error\", \"error\":\"%s\", \"version\":\"%s\"} \n", err, version)
	os.Exit(1)
}

func printResult(data []byte) {
	if len(data) != 0 {
		for i := 0; i < (len(data) / 2); i++ {
			if aegPar[i].regtype == 0 {
				temp1 := aegPar[i].regName
				var temp2 string
				a := binary.BigEndian.Uint16(data[i*2 : (i*2)+2])
				switch a {
				case 1:
					temp2 = "\"Normal_mode\""
				case 2:
					temp2 = "\"Baypass_mode\""
				case 3:
					temp2 = "\"Battery_mode\""
				case 4:
					temp2 = "\"Eco_mode\""
				case 6:
					temp2 = "\"Shutdown_run\""
				case 7:
					temp2 = "\"Shutdown\""
				}
				r := respStruct{namePar: temp1, valuePar: temp2}
				respons = append(respons, r)
			} else if aegPar[i].regtype == 2 {
				temp1 := aegPar[i].regName
				data1 := float64(binary.BigEndian.Uint16(data[i*2:(i*2)+2])) / 10
				temp2 := fmt.Sprintf("%.2f", data1)
				r := respStruct{namePar: temp1, valuePar: temp2}
				respons = append(respons, r)
			} else if aegPar[i].regtype == 3 {
				temp1 := aegPar[i].regName
				i := binary.BigEndian.Uint16(data[i*2 : (i*2)+2])
				b := float64(int16(i)) / 10
				temp2 := fmt.Sprintf("%.2f", b)
				r := respStruct{namePar: temp1, valuePar: temp2}
				respons = append(respons, r)
			} else {
				temp1 := aegPar[i].regName
				data1 := binary.BigEndian.Uint16(data[i*2 : (i*2)+2])
				temp2 := fmt.Sprintf("%d", data1)
				r := respStruct{namePar: temp1, valuePar: temp2}
				respons = append(respons, r)
			}
		}
		printJson(respons)
	} else {
		fmt.Printf("{\"status\":\"error\", \"error\":\"lengt of data is 0\", \"version\":\"%s\"} \n", version)
		os.Exit(100)
	}

}

func printJson(data []respStruct) {
	for l := 0; l < len(data); l++ {
		if l == 0 {
			fmt.Printf("{")
		}
		fmt.Printf("\"%s\":", data[l].namePar)
		fmt.Printf("%s,", data[l].valuePar)
		if l == len(data)-1 {
			fmt.Printf("\"version\":\"%s\"}\n", version)
		}
	}
	os.Exit(0)
}

/* build for rapberry
env GOOS=linux GOARCH=arm GOARM=5 go build
*/
