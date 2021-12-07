package util

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/urfave/cli"
	apinet "k8s.io/apimachinery/pkg/util/net"
)

// JoinIPs stringifies and joins a list of IP addresses with commas.
func JoinIPs(elems []net.IP) string {
	var strs []string
	for _, elem := range elems {
		strs = append(strs, elem.String())
	}
	return strings.Join(strs, ",")
}

// JoinIPNets stringifies and joins a list of IP networks with commas.
func JoinIPNets(elems []*net.IPNet) string {
	var strs []string
	for _, elem := range elems {
		strs = append(strs, elem.String())
	}
	return strings.Join(strs, ",")
}

// GetFirst4Net returns the first IPv4 network from the list of IP networks.
// If no IPv4 addresses are found, an error is raised.
func GetFirst4Net(elems []*net.IPNet) (*net.IPNet, error) {
	for _, elem := range elems {
		if elem == nil || elem.IP.To4() == nil {
			continue
		}
		return elem, nil
	}
	return nil, errors.New("no IPv4 CIDRs found")
}

// GetFirst4 returns the first IPv4 address from the list of IP addresses.
// If no IPv4 addresses are found, an error is raised.
func GetFirst4(elems []net.IP) (net.IP, error) {
	for _, elem := range elems {
		if elem == nil || elem.To4() == nil {
			continue
		}
		return elem, nil
	}
	return nil, errors.New("no IPv4 address found")
}

// GetFirst4String returns the first IPv4 address from a list of IP address strings.
// If no IPv4 addresses are found, an error is raised.
func GetFirst4String(elems []string) (string, error) {
	ips := []net.IP{}
	for _, elem := range elems {
		for _, v := range strings.Split(elem, ",") {
			ips = append(ips, net.ParseIP(v))
		}
	}
	ip, err := GetFirst4(ips)
	if err != nil {
		return "", err
	}
	return ip.String(), nil
}

// GetHostnameAndIPs takes a node name and list of IPs, usually from CLI args.
// If set, these are used to return the node's name and addresses. If not set,
// the system hostname and primary interface address are returned instead.
func GetHostnameAndIPs(name string, nodeIPs cli.StringSlice) (string, []net.IP, error) {
	ips := []net.IP{}
	if len(nodeIPs) == 0 {
		hostIP, err := apinet.ChooseHostInterface()
		if err != nil {
			return "", nil, err
		}
		ips = append(ips, hostIP)
	} else {
		var err error
		ips, err = ParseStringSliceToIPs(nodeIPs)
		if err != nil {
			return "", nil, fmt.Errorf("invalid node-ip: %w", err)
		}
	}

	if name == "" {
		hostname, err := os.Hostname()
		if err != nil {
			return "", nil, err
		}
		name = hostname
	}

	// Use lower case hostname to comply with kubernetes constraint:
	// https://github.com/kubernetes/kubernetes/issues/71140
	name = strings.ToLower(name)

	return name, ips, nil
}

// ParseStringSliceToIPs converts slice of strings that in turn can be lists of comma separated unparsed IP addresses
// into a single slice of net.IP, it returns error if at any point parsing failed
func ParseStringSliceToIPs(s cli.StringSlice) ([]net.IP, error) {
	var ips []net.IP
	for _, unparsedIP := range s {
		for _, v := range strings.Split(unparsedIP, ",") {
			ip := net.ParseIP(v)
			if ip == nil {
				return nil, fmt.Errorf("invalid ip format '%s'", v)
			}
			ips = append(ips, ip)
		}
	}

	return ips, nil
}
