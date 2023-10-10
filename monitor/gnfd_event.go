package monitor

import (
	"cosmossdk.io/math"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bnb-chain/greenfield/types/resource"
	permTypes "github.com/bnb-chain/greenfield/x/permission/types"
	storageTypes "github.com/bnb-chain/greenfield/x/storage/types"
	abciTypes "github.com/cometbft/cometbft/abci/types"
	"github.com/shopspring/decimal"
	"strconv"
)

const groupBucketRegex = "dm_b_.*"
const groupBucketPrefix = "dm_b_"
const groupObjectRegex = "dm_o_.*"
const groupObjectPrefix = "dm_o_" //TODO: how to parse object name

func parseEventCreateGroup(event abciTypes.Event) (*storageTypes.EventCreateGroup, error) {
	result := &storageTypes.EventCreateGroup{}
	for _, attr := range event.Attributes {
		switch attr.Key {
		case "group_id":
			groupId, err := strconv.Unquote(attr.Value)
			if err != nil {
				return nil, err
			}
			result.GroupId = math.NewUintFromString(groupId)
		case "group_name":
			groupName, err := strconv.Unquote(attr.Value)
			if err != nil {
				return nil, err
			}
			result.GroupName = groupName
		case "owner":
			owner, err := strconv.Unquote(attr.Value)
			if err != nil {
				return nil, err
			}
			result.Owner = owner
		case "extra":
			extra, err := strconv.Unquote(attr.Value)
			if err != nil {
				return nil, err
			}
			result.Extra = extra
		}
	}
	return result, nil
}

func parseEventDeleteGroup(event abciTypes.Event) (*storageTypes.EventDeleteGroup, error) {
	result := &storageTypes.EventDeleteGroup{}
	for _, attr := range event.Attributes {
		switch attr.Key {
		case "group_id":
			groupId, err := strconv.Unquote(attr.Value)
			if err != nil {
				return nil, err
			}
			result.GroupId = math.NewUintFromString(groupId)
		case "group_name":
			groupName, err := strconv.Unquote(attr.Value)
			if err != nil {
				return nil, err
			}
			result.GroupName = groupName
		case "owner":
			owner, err := strconv.Unquote(attr.Value)
			if err != nil {
				return nil, err
			}
			result.Owner = owner
		}
	}
	return result, nil
}

func parseEventUpdateGroupExtra(event abciTypes.Event) (*storageTypes.EventUpdateGroupExtra, error) {
	result := &storageTypes.EventUpdateGroupExtra{}
	for _, attr := range event.Attributes {
		switch attr.Key {
		case "group_id":
			groupId, err := strconv.Unquote(attr.Value)
			if err != nil {
				return nil, err
			}
			result.GroupId = math.NewUintFromString(groupId)
		case "group_name":
			groupName, err := strconv.Unquote(attr.Value)
			if err != nil {
				return nil, err
			}
			result.GroupName = groupName
		case "owner":
			owner, err := strconv.Unquote(attr.Value)
			if err != nil {
				return nil, err
			}
			result.Owner = owner
		case "extra":
			extra, err := strconv.Unquote(attr.Value)
			if err != nil {
				return nil, err
			}
			result.Extra = extra
		}
	}
	return result, nil
}

func parseEventPutPolicy(event abciTypes.Event) (*permTypes.EventPutPolicy, error) {
	result := &permTypes.EventPutPolicy{}
	for _, attr := range event.Attributes {
		switch attr.Key {
		case "resource_id":
			resourceId, err := strconv.Unquote(attr.Value)
			if err != nil {
				return nil, err
			}
			result.ResourceId = math.NewUintFromString(resourceId)
		case "resource_type":
			resourceType, err := strconv.Unquote(attr.Value)
			if err != nil {
				return nil, err
			}
			result.ResourceType = resource.ResourceType(resource.ResourceType_value[resourceType])
		case "principal":
			var res Principal
			err := json.Unmarshal([]byte(attr.Value), &res)
			fmt.Println(err)
			if err != nil {
				return nil, err
			}

			t, ok := permTypes.PrincipalType_value[res.Type]
			if !ok {
				return nil, errors.New("cannot parse principal")
			}

			result.Principal = &permTypes.Principal{
				Type:  permTypes.PrincipalType(t),
				Value: res.Value,
			}
		}
	}
	return result, nil
}

type Principal struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Extra struct {
	Desc  string          `json:"desc"`
	Url   string          `json:"url"`
	Price decimal.Decimal `json:"price"`
}

func parseExtra(str string) (*Extra, error) {
	var extra Extra
	err := json.Unmarshal([]byte(str), &extra)
	if err != nil {
		return nil, err
	}

	return &extra, nil
}
