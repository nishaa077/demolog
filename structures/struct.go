package structures

import (
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ZInfoMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Ztype         ZInfoTypes `protobuf:"varint,1,opt,name=ztype,proto3,enum=org.lfedge.eve.info.ZInfoTypes" json:"ztype,omitempty"`
	DevId         string     `protobuf:"bytes,2,opt,name=devId,proto3" json:"devId,omitempty"`
	// Types that are assignable to InfoContent:
	//
	//  *ZInfoMsg_Dinfo
	//  *ZInfoMsg_Ainfo
	//  *ZInfoMsg_Niinfo
	//  *ZInfoMsg_Vinfo
	//  *ZInfoMsg_Cinfo
	//  *ZInfoMsg_Binfo
	//  *ZInfoMsg_Amdinfo
	//  *ZInfoMsg_Evinfo
	//  *ZInfoMsg_Hwinfo
	//  *ZInfoMsg_Locinfo
	//  *ZInfoMsg_PatchInfo
	//  *ZInfoMsg_OpaqueAppInstStatus
	InfoContent isZInfoMsg_InfoContent `protobuf_oneof:"InfoContent"`
	AtTimeStamp *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=atTimeStamp,proto3" json:"atTimeStamp,omitempty"`
}

type isZInfoMsg_InfoContent interface {
	isZInfoMsg_InfoContent()
}

/*type ZInfoMsg_Dinfo struct {
	Dinfo *ZInfoDevice `protobuf:"bytes,3,opt,name=dinfo,proto3,oneof"`
}
func (*ZInfoMsg_Dinfo) isZInfoMsg_InfoContent() {}*/

type ZInfoTypes int32

const (
	ZInfoTypes_ZiNop             ZInfoTypes = 0
	ZInfoTypes_ZiDevice          ZInfoTypes = 1
	ZInfoTypes_ZiApp             ZInfoTypes = 3
	ZInfoTypes_ZiNetworkInstance ZInfoTypes = 6
	ZInfoTypes_ZiVolume          ZInfoTypes = 7
	ZInfoTypes_ZiContentTree     ZInfoTypes = 8
	ZInfoTypes_ZiBlobList        ZInfoTypes = 9
	ZInfoTypes_ZiAppInstMetaData ZInfoTypes = 10
	ZInfoTypes_ZiHardware        ZInfoTypes = 11
	ZInfoTypes_ZiEdgeview        ZInfoTypes = 12
	ZInfoTypes_ZiLocation        ZInfoTypes = 13
	ZInfoTypes_ZiPatchEnvelope   ZInfoTypes = 14
)
