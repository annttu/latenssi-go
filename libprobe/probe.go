package libprobe


type ProbeInitiator func(name string, config map[string]interface{}) ProbeFunction

type ProbeFunction func(host string, interval uint64) ProbeRunner
