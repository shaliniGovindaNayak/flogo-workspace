package hardware

import (
	"fmt"

	"github.com/project-flogo/core/activity"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	nett "github.com/shirou/gopsutil/net"
	"runtime"
	"strconv"
	"encoding/json"
)

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Output{})

//New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {

	act := &Activity{} //add aSetting to instance

	return act, nil
}

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

func dealwithErr(err error) {
	if err != nil {
			fmt.Println(err)
			//os.Exit(-1)
	}
}

type Details struct {
	Total_memory string
	Free_memory string
	Percentage_used_memory string
	Total_disk_space string
	Used_disk_space string
	Free_disk_space string
	Percentage_disk_space_usage string
	CPU_index_number string
	VendorID string
	Family string
	Speed string
	Uptime string
	Number_of_processes_running string
	Host_ID string
 }	

func GetHardwareData() string{
	runtimeOS := runtime.GOOS
		 
		 fmt.Println("operation system:",runtimeOS)
         // memory
		 vmStat, err := mem.VirtualMemory()
		 //fmt.Println(strconv.FormatUint(vmStat.Total, 10))

		//fmt.Println(vmStat)
        // dealwithErr(err)

         // disk - start from "/" mount point for Linux
         // might have to change for Windows!!
         // don't have a Window to test this out, if detect OS == windows
         // then use "\" instead of "/"

		 diskStat, err := disk.Usage("/")
	
		 dealwithErr(err)
		 //fmt.Println(diskStat)

         // cpu - get CPU number of cores and speed
         cpuStat, err := cpu.Info()
		 dealwithErr(err)
		 //fmt.Println(cpuStat)
		 //percentage, err := cpu.Percent(0, true)
		 //fmt.Println(percentage)
		 //dealwithErr(err)

         // host or machine kernel, uptime, platform Info
		 hostStat, err := host.Info()
		 //fmt.Println(hostStat)
         dealwithErr(err)

         // get interfaces MAC/hardware address
		 //interfStat, err := nett.Interfaces()
		 //fmt.Println(interfStat)
		// dealwithErr(err)
		 
		 //serial := disk.GetDiskSerialNumber("/dev/sda")


		 fmt.Println( "Total memory:",strconv.FormatUint(diskStat.Total, 10))
		 fmt.Println("Free memory:",strconv.FormatUint(vmStat.Free, 10))
		 fmt.Println("Percentage used memory: " ,strconv.FormatFloat(vmStat.UsedPercent, 'f', 2, 64))
		 //fmt.Println( "Disk serial number: ", serial)
		 fmt.Println( "Total disk space: " , strconv.FormatUint(diskStat.Total, 10))
		 fmt.Println( "Used disk space: " , strconv.FormatUint(diskStat.Used, 10))
		 fmt.Println( "Free disk space: " , strconv.FormatUint(diskStat.Free, 10))
		 fmt.Println( "Percentage disk space usage: " , strconv.FormatFloat(diskStat.UsedPercent, 'f', 2, 64))
		 fmt.Println( "CPU index number: " , strconv.FormatInt(int64(cpuStat[0].CPU), 10))
		 fmt.Println( "VendorID: " , cpuStat[0].VendorID)
		 fmt.Println( "Family: " , cpuStat[0].Family)
		 fmt.Println( "Speed: " , strconv.FormatFloat(cpuStat[0].Mhz, 'f', 2, 64))
		 fmt.Println( "Uptime: " + strconv.FormatUint(hostStat.Uptime, 10))
		 fmt.Println( "Number of processes running: " + strconv.FormatUint(hostStat.Procs, 10))
		 fmt.Println( "Host ID(uuid): " + hostStat.HostID)
	

		
		for idx, cpupercent := range percentage {
			fmt.Println("Current CPU utilization: [" + strconv.Itoa(idx) + "] " + strconv.FormatFloat(cpupercent, 'f', 2, 64) )

		}
		
		/*for _, interf := range interfStat {
		 	fmt.Println("Interface Name: " + interf.Name) 

			if interf.HardwareAddr != "" {
					fmt.Println("Hardware(MAC) Address: " + interf.HardwareAddr)
			}

			for _, flag := range interf.Flags {
					fmt.Println("Interface behavior or flags: " + flag)
			}

			for _, addr := range interf.Addrs {
					fmt.Println("IPv6 or IPv4 addresses: " + addr.String())

			}

	}*/

	jsondata := Details{
		Total_memory: strconv.FormatUint(diskStat.Total, 10),
		Free_memory : strconv.FormatUint(vmStat.Free, 10),
		Percentage_used_memory:strconv.FormatFloat(vmStat.UsedPercent, 'f', 2, 64),
		Total_disk_space:strconv.FormatUint(diskStat.Total, 10),
		Used_disk_space:strconv.FormatUint(diskStat.Used, 10),
		Free_disk_space:strconv.FormatUint(diskStat.Free, 10),
		Percentage_disk_space_usage:strconv.FormatFloat(diskStat.UsedPercent, 'f', 2, 64),
		CPU_index_number:strconv.FormatInt(int64(cpuStat[0].CPU), 10),
		VendorID:cpuStat[0].VendorID,
		Family:cpuStat[0].Family,
		Speed:strconv.FormatFloat(cpuStat[0].Mhz, 'f', 2, 64),
		Uptime:strconv.FormatUint(hostStat.Uptime, 10),
		Number_of_processes_running:strconv.FormatUint(hostStat.Procs, 10),
		Host_ID:hostStat.HostID,
	}

	b, err := json.Marshal(jsondata)
	if err != nil {
    fmt.Println("error:", err)
	}
	//os.Stdout.Write(b)
	out := string(b)
	return(out)
}


// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {


	out := GetHardwareData()

    //for _, a := range as {	//log.Println("setting:", settings.ASetting)
	//ctx.Logger().Debug("Output: %s", settings.ASetting)
	//ctx.Logger().Debugf("Input: %s", input.AnInput)

	
	output := &Output{Output: out}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}
