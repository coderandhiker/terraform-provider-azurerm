package parse

// NOTE: this file is generated via 'go:generate' - manual changes will be overwritten

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
)

type BgpConnectionId struct {
	SubscriptionId string
	ResourceGroup  string
	VirtualHubName string
	Name           string
}

func NewBgpConnectionID(subscriptionId, resourceGroup, virtualHubName, name string) BgpConnectionId {
	return BgpConnectionId{
		SubscriptionId: subscriptionId,
		ResourceGroup:  resourceGroup,
		VirtualHubName: virtualHubName,
		Name:           name,
	}
}

func (id BgpConnectionId) ID(_ string) string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/virtualHubs/%s/bgpConnections/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroup, id.VirtualHubName, id.Name)
}

// BgpConnectionID parses a BgpConnection ID into an BgpConnectionId struct
func BgpConnectionID(input string) (*BgpConnectionId, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, err
	}

	resourceId := BgpConnectionId{
		SubscriptionId: id.SubscriptionID,
		ResourceGroup:  id.ResourceGroup,
	}

	if resourceId.SubscriptionId == "" {
		return nil, fmt.Errorf("ID was missing the 'subscriptions' element")
	}

	if resourceId.ResourceGroup == "" {
		return nil, fmt.Errorf("ID was missing the 'resourceGroups' element")
	}

	if resourceId.VirtualHubName, err = id.PopSegment("virtualHubs"); err != nil {
		return nil, err
	}
	if resourceId.Name, err = id.PopSegment("bgpConnections"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}
