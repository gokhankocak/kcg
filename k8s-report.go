package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"kcg"
	bs4 "kcg/bs4"
)

// ProgArgs holds program arguments
type ProgArgs struct {
	InputDir string
}

var progArgs ProgArgs

func initProgArgs() {
	flag.StringVar(&progArgs.InputDir, "InputDir", ".", "is where json input files are stored")
}

/*
func displayPodInfo(pod *kcg.Pod) (jsonAsStr string) {
	jsonAsStr += fmt.Sprintf("\n\t\t\"Pod\": \"%s\", \"Namespace\": \"%s\"\n", pod.Metadata.Name, pod.Metadata.Namespace)
	if len(pod.Metadata.OwnerReferences) != 0 {
		jsonAsStr += fmt.Sprintf("\t\t\t,\"OwnerReferences\":\n\t\t\t[\n")
	}
	for ow := 0; ow < len(pod.Metadata.OwnerReferences); ow++ {
		owner := &pod.Metadata.OwnerReferences[ow]
		if ow != 0 {
			jsonAsStr += fmt.Sprintf(",\n")
		}
		jsonAsStr += fmt.Sprintf("\t\t\t\t{\"CreatedBy\": \"%s/%s\"}", owner.Kind, owner.Name)
	}
	if len(pod.Metadata.OwnerReferences) != 0 {
		jsonAsStr += fmt.Sprintf("\t\t\t]\n")
	}
	jsonAsStr += fmt.Sprintf("\n")

	return jsonAsStr

	for c := 0; c < len(pod.Spec.Containers); c++ {
		cntr := &pod.Spec.Containers[c]
		byColon := strings.Split(cntr.Image, ":")
		if len(byColon) >= 2 {
			fmt.Printf("\t\tContainerImage: %s\tVersion: %s\n", byColon[0], byColon[1])
		} else {
			fmt.Printf("\t\tContainerImage: %s\n", byColon[0])
		}
		for cp := 0; cp < len(cntr.Ports); cp++ {
			fmt.Printf("\t\t\tContainerPort: %s\t%s\n", cntr.Ports[cp].CntrPort, cntr.Ports[cp].Protocol)
		}
		for vm := 0; vm < len(cntr.VolumeMounts); vm++ {
			fmt.Printf("\t\t\tVolumeName: %s\tMountPath: %s\n", cntr.VolumeMounts[vm].Name, cntr.VolumeMounts[vm].MountPath)
		}
	}

	for c := 0; c < len(pod.Spec.InitContainers); c++ {
		cntr := &pod.Spec.InitContainers[c]
		byColon := strings.Split(cntr.Image, ":")
		if len(byColon) >= 2 {
			fmt.Printf("\t\tInitContainerImage: %s\tVersion: %s\n", byColon[0], byColon[1])
		} else {
			fmt.Printf("\t\tInitContainerImage: %s\n", byColon[0])
		}
	}

		for v := 0; v < len(pod.Spec.Volumes); v++ {
			vol := &pod.Spec.Volumes[v]
			fmt.Printf("\t\tVolume: %s\n", vol.Name)
		}

	return jsonAsStr
}

*/

func containerInfoToString(pod *kcg.Pod) (cntrStr string) {
	for c := 0; c < len(pod.Spec.Containers); c++ {
		cntr := &pod.Spec.Containers[c]
		if c != 0 {
			cntrStr += "<br>"
		}
		byColon := strings.Split(cntr.Image, ":")
		image := byColon[0]
		version := ""
		if len(byColon) > 1 {
			version = byColon[1]
		}
		portStr := ""
		for p := 0; p < len(cntr.Ports); p++ {
			port := &cntr.Ports[p]
			portStr += fmt.Sprintf("<%d %s %s> ", port.CntrPort, port.Protocol, port.Name)
		}
		cntrStr += fmt.Sprintf("{container: %s, image: %s, version: %s, ports: {%s}}", cntr.Name, image, version, portStr)
	}
	return cntrStr
}

func initContainerInfoToString(pod *kcg.Pod) (cntrStr string) {
	for c := 0; c < len(pod.Spec.InitContainers); c++ {
		cntr := &pod.Spec.InitContainers[c]
		if c != 0 {
			cntrStr += "<br>"
		}
		byColon := strings.Split(cntr.Image, ":")
		image := byColon[0]
		version := ""
		if len(byColon) > 1 {
			version = byColon[1]
		}
		cntrStr += fmt.Sprintf("{initcontainer: %s, image: %s, version: %s}", cntr.Name, image, version)
	}
	return cntrStr
}

func main() {

	initProgArgs()
	flag.Parse()

	htm := bs4.NewBs4Html()
	htm.New("Kubernetes Report")

	nav := bs4.NewBs4Nav()
	nav.Start()
	nav.SetLogo("Gokhan Kocak")
	nav.AddDropDownMenu(
		bs4.Bs4NameHref{Name: "Nodes", Href: "Nodes"},
		[]bs4.Bs4NameHref{
			{Name: "Basics", Href: "Node_Basics"},
			{Name: "Network", Href: "Node_Network"},
			{Name: "CPU & Memory", Href: "Node_CPU_Memory"},
		},
	)
	nav.AddItem(bs4.Bs4NameHref{Name: "Namespaces", Href: "Namespaces"})
	nav.AddItem(bs4.Bs4NameHref{Name: "Services", Href: "Services"})
	nav.AddItem(bs4.Bs4NameHref{Name: "Pods", Href: "Pods"})
	nav.AddItem(bs4.Bs4NameHref{Name: "Images", Href: "Images"})

	htm.Insert(nav.End())

	// nodes
	var nds kcg.Nodes
	content, err := ioutil.ReadFile(filepath.Join(filepath.Dir(progArgs.InputDir), "/", filepath.Base(progArgs.InputDir), "/", "nodes.json"))
	if err == nil {
		err = json.Unmarshal(content, &nds)
		if err != nil {
			fmt.Println("nodes.json", err.Error())
		}
	}
	fmt.Printf("Loaded %d items from nodes.json\n", len(nds.Items))

	tbl := bs4.NewBs4Table()
	tbl.Start("Nodes", "Nodes")
	tbl.SetHeader([]string{"Name", "Resource Version", "UID", "Architecture", "Operating System", "OS Image", "Container Runtime Version", "Kernel Version", "Kube Proxy Version", "Kubelet Version"})
	tbl.NewBody()

	for k := 0; k < len(nds.Items); k++ {
		nd := &nds.Items[k]
		tbl.AddRow([]string{"<b>" + nd.Metadata.Name + "</b>", nd.Metadata.ResourceVersion, nd.Metadata.UID, nd.Status.Info.Architecture, nd.Status.Info.OperatingSystem, nd.Status.Info.OsImage, nd.Status.Info.ContainerRuntimeVersion, nd.Status.Info.KernelVersion, nd.Status.Info.KubeProxyVersion, nd.Status.Info.KubeletVersion})
	}

	tbl.EndBody()
	htm.Insert(tbl.End())

	// node network
	tbl = bs4.NewBs4Table()
	tbl.Start("Node Network", "Node_Network")
	tbl.SetHeader([]string{"Node", "PodCIDR", "PodCIDRs", "Provider ID", "Addresses"})
	tbl.NewBody()

	for k := 0; k < len(nds.Items); k++ {
		nd := &nds.Items[k]
		cidrs := ""
		for t := 0; t < len(nd.Spec.PodCIDRs); t++ {
			cidrs += nd.Spec.PodCIDRs[t] + "<br>"
		}
		addr := ""
		for t := 0; t < len(nd.Status.Addresses); t++ {
			addr += nd.Status.Addresses[t].Type + ": " + nd.Status.Addresses[t].Address + "<br>"
		}
		tbl.AddRow([]string{"<b>" + nd.Metadata.Name + "</b>", nd.Spec.PodCIDR, cidrs, nd.Spec.ProviderID, addr})
	}
	tbl.EndBody()
	htm.Insert(tbl.End())

	// node CPU & Memory
	tbl = bs4.NewBs4Table()
	tbl.Start("Node CPU & Memory (Actual : Capacity)", "Node_CPU_Memory")
	tbl.SetHeader([]string{"Node", "cpu", "ephemeral-storage", "hugepages-1Gi", "hugepages-2Mi", "memory", "pods"})
	tbl.NewBody()
	for k := 0; k < len(nds.Items); k++ {
		nd := &nds.Items[k]
		a := &nds.Items[k].Status.Allocatable
		c := &nds.Items[k].Status.Capacity
		tbl.AddRow([]string{"<b>" + nd.Metadata.Name + "</b>", a.CPU + " : " + c.CPU, a.EphemeralStorage + " : " + c.EphemeralStorage, fmt.Sprintf("%v : %v", a.HugePages1Gi, c.HugePages1Gi), fmt.Sprintf("%v : %v", a.HugePages2Mi, c.HugePages2Mi), a.Memory + " : " + c.Memory, a.Pods + " : " + c.Pods})
	}
	tbl.EndBody()
	htm.Insert(tbl.End())

	// node conditions
	tbl = bs4.NewBs4Table()
	tbl.Start("Node Conditions", "Node_Conditions")
	tbl.SetHeader([]string{"node", "lastHeartbeatTime", "lastTransitionTime", "message", "reason", "status", "type"})
	tbl.NewBody()
	for k := 0; k < len(nds.Items); k++ {
		nd := &nds.Items[k]
		nc := kcg.Condition{}
		for t := 0; t < len(nd.Status.Conditions); t++ {
			cs := &nd.Status.Conditions[t]
			nc.LastHeartbeatTime += cs.LastHeartbeatTime + "<br>"
			nc.LastTransitionTime += cs.LastTransitionTime + "<br>"
			nc.Message += cs.Message + "<br>"
			nc.Reason += cs.Reason + "<br>"
			nc.Status += cs.Status + "<br>"
			nc.Type += cs.Type + "<br>"
		}
		tbl.AddRow([]string{"<b>" + nd.Metadata.Name + "</b>", nc.LastHeartbeatTime, nc.LastTransitionTime, nc.Message, nc.Reason, nc.Status, nc.Type})
	}
	tbl.EndBody()
	htm.Insert(tbl.End())

	// namespaces
	var nss kcg.Namespaces
	content, err = ioutil.ReadFile(filepath.Join(filepath.Dir(progArgs.InputDir), "/", filepath.Base(progArgs.InputDir), "/", "namespaces.json"))
	if err == nil {
		err = json.Unmarshal(content, &nss)
		if err != nil {
			fmt.Println("namespaces.json", err.Error())
		}
	}
	fmt.Printf("Loaded %d items from namespaces.json\n", len(nss.Items))

	nstbl := bs4.NewBs4Table()
	nstbl.Start("Namespaces", "Namespaces")
	nstbl.SetHeader([]string{"Name", "Resource Version", "UID"})
	nstbl.NewBody()

	for k := 0; k < len(nss.Items); k++ {
		ns := &nss.Items[k]
		nstbl.AddRow([]string{"<b>" + ns.Metadata.Name + "</b>", ns.Metadata.ResourceVersion, ns.Metadata.UID})
	}

	nstbl.EndBody()
	htm.Insert(nstbl.End())

	// services
	var svcs kcg.Services
	content, err = ioutil.ReadFile(filepath.Join(filepath.Dir(progArgs.InputDir), "/", filepath.Base(progArgs.InputDir), "/", "services.json"))
	if err == nil {
		err = json.Unmarshal(content, &svcs)
		if err != nil {
			fmt.Println("services.json", err.Error())
		}
	}
	fmt.Printf("Loaded %d items from services.json\n", len(svcs.Items))

	svctbl := bs4.NewBs4Table()
	svctbl.Start("Services", "Services")
	svctbl.SetHeader([]string{"Name", "Namespace", "Resource Version", "UID"})
	svctbl.NewBody()

	for k := 0; k < len(svcs.Items); k++ {
		svc := &svcs.Items[k]
		svctbl.AddRow([]string{"<b>" + svc.Metadata.Name + "</b>", svc.Metadata.Namespace, svc.Metadata.ResourceVersion, svc.Metadata.UID})
	}

	svctbl.EndBody()
	htm.Insert(svctbl.End())

	// pods
	var pods kcg.Pods
	content, err = ioutil.ReadFile(filepath.Join(filepath.Dir(progArgs.InputDir), "/", filepath.Base(progArgs.InputDir), "/", "pods.json"))
	if err == nil {
		err = json.Unmarshal(content, &pods)
		if err == nil {
			fmt.Println("pods.json", err.Error())
		}
	}
	fmt.Printf("Loaded %d items from pods.json\n", len(pods.Items))

	podtbl := bs4.NewBs4Table()
	podtbl.Start("Pods", "Pods")
	podtbl.SetHeader([]string{"Name", "Namespace", "Containers", "Resource Version", "UID"})
	podtbl.NewBody()

	for k := 0; k < len(pods.Items); k++ {
		pod := &pods.Items[k]
		cntrStr := containerInfoToString(pod)
		initCntrStr := initContainerInfoToString(pod)
		if cntrStr != "" && initCntrStr != "" {
			cntrStr += "<br>"
		}
		cntrStr += initCntrStr
		podtbl.AddRow([]string{"<b>" + pod.Metadata.Name + "</b>", pod.Metadata.Namespace, cntrStr, pod.Metadata.ResourceVersion, pod.Metadata.UID})
	}

	matchedPods := make(map[string]int)

	// Link Services and Pods and vice versa
	for k := 0; k < len(svcs.Items); k++ {
		svc := &svcs.Items[k]
		for t := 0; t < len(pods.Items); t++ {
			pod := &pods.Items[t]
			MatchCount := 0
			for SelKey, SelVal := range svc.Spec.Selector {
				if pod.Metadata.Labels[SelKey] == SelVal {
					MatchCount++
				}
			}
			if MatchCount != 0 && MatchCount == len(svc.Spec.Selector) {
				matchedPods[pod.Metadata.Name] = 1
			}
		}
		if false {
			fmt.Printf("\tClusterIP: %s\n", svc.Spec.ClusterIP)
			for i := 0; i < len(svc.Spec.Ports); i++ {
				port := &svc.Spec.Ports[i]
				fmt.Printf("\tname: %s port: %d protocol: %s targetPort: %v nodePort: %d\n", port.Name, port.Port, port.Protocol, port.TargetPort, port.NodePort)
			}
		}
	}

	// display unmatched pod (pods without services)
	for t := 0; t < len(pods.Items); t++ {
		pod := &pods.Items[t]
		if matchedPods[pod.Metadata.Name] != 1 {
			cntrStr := containerInfoToString(pod)
			cntrStr += initContainerInfoToString(pod)
			podtbl.AddRow([]string{"<b>" + pod.Metadata.Name + "</b>", pod.Metadata.Namespace, cntrStr, pod.Metadata.ResourceVersion, pod.Metadata.UID})
		}
	}

	podtbl.EndBody()
	htm.Insert(podtbl.End())

	var dps kcg.Deployments
	content, err = ioutil.ReadFile(filepath.Join(filepath.Dir(progArgs.InputDir), "/", filepath.Base(progArgs.InputDir), "/", "deployments.json"))
	if err == nil {
		err = json.Unmarshal(content, &dps)
		if err != nil {
			fmt.Println("deployments.json", err.Error())
		}
	}
	fmt.Printf("Loaded %d items from deployments.json\n", len(dps.Items))

	if false {
		for k := 0; k < len(dps.Items); k++ {
			dp := &dps.Items[k]
			fmt.Println(dp.Metadata.Name, dp.Metadata.Namespace, dp.Metadata.Generation)
		}
	}

	var dmnss kcg.DaemonSets
	content, err = ioutil.ReadFile(filepath.Join(filepath.Dir(progArgs.InputDir), "/", filepath.Base(progArgs.InputDir), "/", "daemonsets.json"))
	if err == nil {
		err = json.Unmarshal(content, &dmnss)
		if err != nil {
			fmt.Println("daemonsets.json", err.Error())
		}
	}
	fmt.Printf("Loaded %d items from daemonsets.json\n", len(dmnss.Items))

	if false {
		for k := 0; k < len(dmnss.Items); k++ {
			dmns := &dps.Items[k]
			fmt.Println(dmns.Metadata.Name, dmns.Metadata.Namespace, dmns.Metadata.Generation)
		}
	}

	var rpss kcg.ReplicaSets
	content, err = ioutil.ReadFile(filepath.Join(filepath.Dir(progArgs.InputDir), "/", filepath.Base(progArgs.InputDir), "/", "replicasets.json"))
	if err == nil {
		err = json.Unmarshal(content, &rpss)
		if err != nil {
			fmt.Println("replicasets.json", err.Error())
		}
	}
	fmt.Printf("Loaded %d items from replicasets.json\n", len(rpss.Items))

	if false {
		for k := 0; k < len(rpss.Items); k++ {
			rps := &rpss.Items[k]
			fmt.Println(rps.Metadata.Name, rps.Metadata.Namespace, rps.Metadata.Generation)
		}
	}

	var sfss kcg.StatefulSets
	content, err = ioutil.ReadFile(filepath.Join(filepath.Dir(progArgs.InputDir), "/", filepath.Base(progArgs.InputDir), "/", "statefulsets.json"))
	if err == nil {
		err = json.Unmarshal(content, &sfss)
		if err != nil {
			fmt.Println("statefulsets.json", err.Error())
		}
	}
	fmt.Printf("Loaded %d items from statefulsets.json\n", len(sfss.Items))

	if false {
		for k := 0; k < len(sfss.Items); k++ {
			sfs := &sfss.Items[k]
			fmt.Println(sfs.Metadata.Name, sfs.Metadata.Namespace, sfs.Metadata.Generation)
		}
	}

	// images

	imageMap := make(map[string]int64)
	for k := 0; k < len(nds.Items); k++ {
		nd := &nds.Items[k]
		for i := 0; i < len(nd.Status.Images); i++ {
			for n := 0; n < len(nd.Status.Images[i].Names); n++ {
				name := nd.Status.Images[i].Names[n]
				if strings.Contains(name, "@") == false {
					imageMap[name] = nd.Status.Images[i].SizeBytes
				}
			}
		}
	}

	type tempImageSort struct {
		Name      string
		Version   string
		SizeBytes int64
	}
	var sortedImages []tempImageSort
	for key, value := range imageMap {
		byColon := strings.Split(key, ":")
		sortedImages = append(sortedImages, tempImageSort{Name: byColon[0], Version: byColon[1], SizeBytes: value})
	}
	sort.Slice(sortedImages, func(i, j int) bool { return sortedImages[i].Name < sortedImages[j].Name })

	imgtbl := bs4.NewBs4Table()
	imgtbl.Start("Images", "Images")
	imgtbl.SetHeader([]string{"Name", "Version", "Bytes"})
	imgtbl.NewBody()

	for k := 0; k < len(sortedImages); k++ {
		imgtbl.AddRow([]string{"<b>" + sortedImages[k].Name + "</b>", sortedImages[k].Version, strconv.FormatInt(sortedImages[k].SizeBytes, 10)})
	}
	imgtbl.EndBody()
	htm.Insert(imgtbl.End())

	nowStr := time.Now().String()
	filename := strings.ReplaceAll(strings.ReplaceAll(nowStr, " ", "_"), ":", "-")[:19] + ".html"
	err = ioutil.WriteFile(filename, []byte(htm.End()), 0644)
	if err != nil {
		log.Println(err.Error())
	}
}
