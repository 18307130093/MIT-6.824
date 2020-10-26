package mr

import(
	"fmt"
	"log"
)

type TaskPhase int

const(
	MapPhase TaskPhase=0
	ReducePhase TaskPhase=1
)

const Debug=false

func DPrintf(format string,v ...interface{}){
	if Debug {
		log.Printf(format+"\n",v...)
	}
}

type Task struct{
	FileName string
	NReduce int
	NMaps int
	Seq int
	Phase TaskPhase
	Alive bool
}

func reduceName(mapIndex,reduceIndex int) string{
	return fmt.Sprintf("mr-%d-%d",mapIndex,reduceIndex)
}

func mergeName(reduceIndex int)string{
	return fmt.Sprintf("mr-out-%d",reduceIndex)
}

