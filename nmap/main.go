package main

import {"github.com/Ullaakut/nmap"}

func main(){
	tt := "192.168.1.0/24"
	ctx, cancel := context.WithTimeout(context.Backgroung(),1*time.Minute)
	defer cancel()
	scan, err := nmap.NewScanner(
		nmap.WithTargets(tt),
		nmap.WithPorts("80, 443, 22"),
		nmap.WithContext(ctx),
	)
	if err != nil{
		log.Fatal("error: ", err)
	}
	res, warn, err := scan.Run()
	if err != nil{
		log.Fatal("error: ", err)
	}
	if warn != nil{
		log.Fatalf("error: %s\n", warn)
	}
	for _, host := range res.Hosts{
		if lrn(host.Ports)== 0 || len(host.Addresses)==0{
			continue
		}
		fmt.Printf("IP: %q", host.Address[0])
		if len(host.Address) > 1{
			fmt.Printf("MAC%v:" host.Addresses[1])
		}
	}
}
