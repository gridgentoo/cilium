// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Hubble

package api

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"k8s.io/utils/strings/slices"

	pb "github.com/cilium/cilium/api/v1/flow"
	"github.com/cilium/cilium/pkg/labels"
)

// ContextIdentifier describes the identification method of a transmission or
// receiving context
type ContextIdentifier int

const (
	// ContextDisabled disables context identification
	ContextDisabled ContextIdentifier = iota
	// ContextIdentity uses the full set of identity labels for identification purposes
	ContextIdentity
	// ContextNamespace uses the namespace name for identification purposes
	ContextNamespace
	// ContextPod uses the pod name for identification purposes
	ContextPod
	// ContextPodShort uses a short version of the pod name. It should
	// typically map to the deployment/replicaset name
	ContextPodShort
	// ContextDNS uses the DNS name for identification purposes
	ContextDNS
	// ContextIP uses the IP address for identification purposes
	ContextIP
	// ContextReservedIdentity uses reserved labels in the identity label list for identification
	// purpose. It uses "reserved:kube-apiserver" label if it's present in the identity label list.
	// Otherwise, it uses the first label in the identity label list with "reserved:" prefix.
	ContextReservedIdentity
)

// ContextOptionsHelp is the help text for context options
const ContextOptionsHelp = `
 sourceContext          := identifier , { "|", identifier }
 destinationContext     := identifier , { "|", identifier }
 labels                 := label , { ",", label }
 identifier             := identity | namespace | pod | pod-short | dns | ip | reserved-identity
 label                  := source_pod | source_namespace | source_workload | destination_pod | destination_namespace | destination_workload | reporter
`

var (
	shortPodPattern    = regexp.MustCompile("^(.+?)(-[a-z0-9]+){1,2}$")
	kubeAPIServerLabel = labels.LabelKubeAPIServer.String()
	// contextLabelsList defines available labels for the ContextLabels
	// ContextIdentifier and the order of those labels for GetLabelNames and GetLabelValues.
	contextLabelsList = []string{
		"source_pod",
		"source_namespace",
		"source_workload",
		"destination_pod",
		"destination_namespace",
		"destination_workload",
		"reporter",
	}
	allowedContextLabels = newLabelsSet(contextLabelsList)
)

// String return the context identifier as string
func (c ContextIdentifier) String() string {
	switch c {
	case ContextDisabled:
		return "disabled"
	case ContextIdentity:
		return "identity"
	case ContextNamespace:
		return "namespace"
	case ContextPod:
		return "pod"
	case ContextPodShort:
		return "pod-short"
	case ContextDNS:
		return "dns"
	case ContextIP:
		return "ip"
	case ContextReservedIdentity:
		return "reserved-identity"
	}
	return fmt.Sprintf("%d", c)
}

type ContextIdentifierList []ContextIdentifier

func (cs ContextIdentifierList) String() string {
	s := make([]string, 0, len(cs))
	for _, c := range cs {
		s = append(s, c.String())
	}
	return strings.Join(s, "|")
}

// ContextOptions is the set of options to define whether and how to include
// sending and/or receiving context information
type ContextOptions struct {
	// Destination is the destination context to include in metrics
	Destination ContextIdentifierList
	// Source is the source context to include in metrics
	Source ContextIdentifierList

	// Labels is the full set of labels that have been allowlisted when using the
	// ContextLabels ContextIdentifier.
	Labels labelsSet
}

func parseContextIdentifier(s string) (ContextIdentifier, error) {
	switch strings.ToLower(s) {
	case "identity":
		return ContextIdentity, nil
	case "namespace":
		return ContextNamespace, nil
	case "pod":
		return ContextPod, nil
	case "pod-short":
		return ContextPodShort, nil
	case "dns":
		return ContextDNS, nil
	case "ip":
		return ContextIP, nil
	case "reserved-identity":
		return ContextReservedIdentity, nil
	default:
		return ContextDisabled, fmt.Errorf("unknown context '%s'", s)
	}
}

func parseContext(s string) (cs ContextIdentifierList, err error) {
	for _, v := range strings.Split(s, "|") {
		c, err := parseContextIdentifier(v)
		if err != nil {
			return nil, err
		}
		cs = append(cs, c)
	}

	return cs, nil
}

func parseLabels(s string) (labelsSet, error) {
	labels := strings.Split(s, ",")
	for _, label := range labels {
		if !allowedContextLabels.HasLabel(label) {
			return labelsSet{}, fmt.Errorf("invalid labelsContext value: %s", label)
		}
	}
	ls := newLabelsSet(labels)
	return ls, nil
}

// ParseContextOptions parses a set of options and extracts the context
// relevant options
func ParseContextOptions(options Options) (*ContextOptions, error) {
	o := &ContextOptions{}
	var err error
	for key, value := range options {
		switch strings.ToLower(key) {
		case "destinationcontext":
			o.Destination, err = parseContext(value)
			if err != nil {
				return nil, err
			}
		case "sourcecontext":
			o.Source, err = parseContext(value)
			if err != nil {
				return nil, err
			}
		case "labelscontext":
			o.Labels, err = parseLabels(value)
			if err != nil {
				return nil, err
			}
		}
	}

	return o, nil
}

type labelsSet map[string]struct{}

func newLabelsSet(labels []string) labelsSet {
	m := make(map[string]struct{}, len(labels))
	for _, label := range labels {
		m[label] = struct{}{}
	}
	return labelsSet(m)
}

func (ls labelsSet) HasLabel(label string) bool {
	_, exists := ls[label]
	return exists
}

func labelsContext(wantedLabels labelsSet, flow *pb.Flow) (outputLabels []string) {
	// Iterate over contextLabelsList so that the label order is stable,
	// otherwise GetLabelNames and GetLabelValues might be mismatched
	for _, label := range contextLabelsList {
		if wantedLabels.HasLabel(label) {
			var labelValue string
			switch label {
			case "source_pod":
				if flow.GetSource() != nil {
					labelValue = flow.GetSource().PodName
				}
			case "source_namespace":
				if flow.GetSource() != nil {
					labelValue = flow.GetSource().Namespace
				}
			case "source_workload":
				if flow.GetSource() != nil {
					if workloads := flow.GetSource().GetWorkloads(); len(workloads) != 0 {
						labelValue = workloads[0].Name
					}
				}
			case "destination_pod":
				if flow.GetDestination() != nil {
					labelValue = flow.GetDestination().PodName
				}
			case "destination_namespace":
				if flow.GetDestination() != nil {
					labelValue = flow.GetDestination().Namespace
				}
			case "destination_workload":
				if flow.GetDestination() != nil {
					if workloads := flow.GetDestination().GetWorkloads(); len(workloads) != 0 {
						labelValue = workloads[0].Name
					}
				}
			case "reporter":
				switch flow.GetTrafficDirection() {
				case pb.TrafficDirection_EGRESS:
					labelValue = "client"
				case pb.TrafficDirection_INGRESS:
					labelValue = "server"
				default:
					labelValue = "unknown"
				}
			default:
				// Label is in contextLabelsList but isn't handled the switch
				// statement. Programmer error.
				panic(fmt.Sprintf("label %s not mapped in labelsContext, please update labelsContext", label))
			}
			outputLabels = append(outputLabels, labelValue)
		}
	}
	return outputLabels
}

func sourceNamespaceContext(flow *pb.Flow) (context string) {
	if flow.GetSource() != nil {
		context = flow.GetSource().Namespace
	}
	return
}

func sourceIdentityContext(flow *pb.Flow) (context string) {
	if flow.GetSource() != nil {
		context = strings.Join(flow.GetSource().Labels, ",")
	}
	return
}

func sourceReservedIdentityContext(flow *pb.Flow) (context string) {
	if flow.GetSource() != nil {
		context = handleReservedIdentityLabels(flow.GetSource().Labels)
	}
	return
}

func sourcePodContext(flow *pb.Flow) (context string) {
	if flow.GetSource() != nil {
		context = flow.GetSource().PodName
		if flow.GetSource().Namespace != "" {
			context = flow.GetSource().Namespace + "/" + context
		}
	}
	return
}

func shortenPodName(name string) string {
	return shortPodPattern.ReplaceAllString(name, "${1}")
}

func sourcePodShortContext(flow *pb.Flow) (context string) {
	if flow.GetSource() != nil {
		context = shortenPodName(flow.GetSource().PodName)
		if flow.GetSource().Namespace != "" {
			context = flow.GetSource().Namespace + "/" + context
		}
	}
	return
}

func sourceDNSContext(flow *pb.Flow) (context string) {
	if flow.GetSourceNames() != nil {
		context = strings.Join(flow.GetSourceNames(), ",")
	}
	return
}

func sourceIPContext(flow *pb.Flow) (context string) {
	if flow.GetIP() != nil {
		context = flow.GetIP().GetSource()
	}
	return
}

func destinationNamespaceContext(flow *pb.Flow) (context string) {
	if flow.GetDestination() != nil {
		context = flow.GetDestination().Namespace
	}
	return
}

func handleReservedIdentityLabels(lbls []string) string {
	// if reserved:kube-apiserver label is present, return it (instead of reserved:world, etc..)
	if slices.Contains(lbls, kubeAPIServerLabel) {
		return kubeAPIServerLabel
	}
	// else return the first reserved label.
	for _, label := range lbls {
		if strings.HasPrefix(label, labels.LabelSourceReserved+":") {
			return label
		}
	}
	return ""
}

func destinationIdentityContext(flow *pb.Flow) (context string) {
	if flow.GetDestination() != nil {
		context = strings.Join(flow.GetDestination().Labels, ",")
	}
	return
}

func destinationReservedIdentityContext(flow *pb.Flow) (context string) {
	if flow.GetDestination() != nil {
		context = handleReservedIdentityLabels(flow.GetDestination().Labels)
	}
	return
}

func destinationPodContext(flow *pb.Flow) (context string) {
	if flow.GetDestination() != nil {
		context = flow.GetDestination().PodName
		if flow.GetDestination().Namespace != "" {
			context = flow.GetDestination().Namespace + "/" + context
		}
	}
	return
}

func destinationPodShortContext(flow *pb.Flow) (context string) {
	if flow.GetDestination() != nil {
		context = shortenPodName(flow.GetDestination().PodName)
		if flow.GetDestination().Namespace != "" {
			context = flow.GetDestination().Namespace + "/" + context
		}
	}
	return
}

func destinationDNSContext(flow *pb.Flow) (context string) {
	if flow.GetDestinationNames() != nil {
		context = strings.Join(flow.GetDestinationNames(), ",")
	}
	return
}

func destinationIPContext(flow *pb.Flow) (context string) {
	if flow.GetIP() != nil {
		context = flow.GetIP().GetDestination()
	}
	return
}

// GetLabelValues returns the values of the context relevant labels according
// to the configured options. The order of the values is the same as the order
// of the label names returned by GetLabelNames()
func (o *ContextOptions) GetLabelValues(flow *pb.Flow) (labels []string) {
	if len(o.Labels) != 0 {
		labels = append(labels, labelsContext(o.Labels, flow)...)
	}

	if len(o.Source) != 0 {
		var context string
		for _, source := range o.Source {
			switch source {
			case ContextNamespace:
				context = sourceNamespaceContext(flow)
			case ContextIdentity:
				context = sourceIdentityContext(flow)
			case ContextPod:
				context = sourcePodContext(flow)
			case ContextPodShort:
				context = sourcePodShortContext(flow)
			case ContextDNS:
				context = sourceDNSContext(flow)
			case ContextIP:
				context = sourceIPContext(flow)
			case ContextReservedIdentity:
				context = sourceReservedIdentityContext(flow)
			}
			// always use first non-empty context
			if context != "" {
				break
			}
		}
		labels = append(labels, context)
	}

	if len(o.Destination) != 0 {
		var context string
		for _, destination := range o.Destination {
			switch destination {
			case ContextNamespace:
				context = destinationNamespaceContext(flow)
			case ContextIdentity:
				context = destinationIdentityContext(flow)
			case ContextPod:
				context = destinationPodContext(flow)
			case ContextPodShort:
				context = destinationPodShortContext(flow)
			case ContextDNS:
				context = destinationDNSContext(flow)
			case ContextIP:
				context = destinationIPContext(flow)
			case ContextReservedIdentity:
				context = destinationReservedIdentityContext(flow)
			}
			// always use first non-empty context
			if context != "" {
				break
			}
		}
		labels = append(labels, context)
	}

	return
}

// GetLabelNames returns a slice of label names required to fulfil the
// configured context description requirements
func (o *ContextOptions) GetLabelNames() (labels []string) {
	if len(o.Labels) != 0 {
		// We must iterate over contextLabelsList to ensure the order of the label
		// names the same order as label values in GetLabelValues
		for _, label := range contextLabelsList {
			if o.Labels.HasLabel(label) {
				labels = append(labels, label)
			}
		}
	}

	if len(o.Source) != 0 {
		labels = append(labels, "source")
	}

	if len(o.Destination) != 0 {
		labels = append(labels, "destination")
	}

	return
}

// Status returns the configuration status of context options suitable for use
// with Handler.Status
func (o *ContextOptions) Status() string {
	var status []string
	if len(o.Source) != 0 {
		status = append(status, "source="+o.Source.String())
	}

	if len(o.Destination) != 0 {
		status = append(status, "destination="+o.Destination.String())
	}

	sort.Strings(status)

	return strings.Join(status, ",")
}
