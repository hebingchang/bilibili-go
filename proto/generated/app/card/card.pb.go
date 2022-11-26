// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: bilibili/app/card/v1/card.proto

package card

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 卡片信息
type Card struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Item:
	//
	//	*Card_SmallCoverV5
	//	*Card_LargeCoverV1
	//	*Card_ThreeItemAllV2
	//	*Card_ThreeItemV1
	//	*Card_HotTopic
	//	*Card_ThreeItemHV5
	//	*Card_MiddleCoverV3
	//	*Card_LargeCoverV4
	//	*Card_PopularTopEntrance
	//	*Card_RcmdOneItem
	//	*Card_SmallCoverV5Ad
	Item isCard_Item `protobuf_oneof:"item"`
}

func (x *Card) Reset() {
	*x = Card{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_card_v1_card_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Card) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Card) ProtoMessage() {}

func (x *Card) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_card_v1_card_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Card.ProtoReflect.Descriptor instead.
func (*Card) Descriptor() ([]byte, []int) {
	return file_bilibili_app_card_v1_card_proto_rawDescGZIP(), []int{0}
}

func (m *Card) GetItem() isCard_Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (x *Card) GetSmallCoverV5() *SmallCoverV5 {
	if x, ok := x.GetItem().(*Card_SmallCoverV5); ok {
		return x.SmallCoverV5
	}
	return nil
}

func (x *Card) GetLargeCoverV1() *LargeCoverV1 {
	if x, ok := x.GetItem().(*Card_LargeCoverV1); ok {
		return x.LargeCoverV1
	}
	return nil
}

func (x *Card) GetThreeItemAllV2() *ThreeItemAllV2 {
	if x, ok := x.GetItem().(*Card_ThreeItemAllV2); ok {
		return x.ThreeItemAllV2
	}
	return nil
}

func (x *Card) GetThreeItemV1() *ThreeItemV1 {
	if x, ok := x.GetItem().(*Card_ThreeItemV1); ok {
		return x.ThreeItemV1
	}
	return nil
}

func (x *Card) GetHotTopic() *HotTopic {
	if x, ok := x.GetItem().(*Card_HotTopic); ok {
		return x.HotTopic
	}
	return nil
}

func (x *Card) GetThreeItemHV5() *DynamicHot {
	if x, ok := x.GetItem().(*Card_ThreeItemHV5); ok {
		return x.ThreeItemHV5
	}
	return nil
}

func (x *Card) GetMiddleCoverV3() *MiddleCoverV3 {
	if x, ok := x.GetItem().(*Card_MiddleCoverV3); ok {
		return x.MiddleCoverV3
	}
	return nil
}

func (x *Card) GetLargeCoverV4() *LargeCoverV4 {
	if x, ok := x.GetItem().(*Card_LargeCoverV4); ok {
		return x.LargeCoverV4
	}
	return nil
}

func (x *Card) GetPopularTopEntrance() *PopularTopEntrance {
	if x, ok := x.GetItem().(*Card_PopularTopEntrance); ok {
		return x.PopularTopEntrance
	}
	return nil
}

func (x *Card) GetRcmdOneItem() *RcmdOneItem {
	if x, ok := x.GetItem().(*Card_RcmdOneItem); ok {
		return x.RcmdOneItem
	}
	return nil
}

func (x *Card) GetSmallCoverV5Ad() *SmallCoverV5Ad {
	if x, ok := x.GetItem().(*Card_SmallCoverV5Ad); ok {
		return x.SmallCoverV5Ad
	}
	return nil
}

type isCard_Item interface {
	isCard_Item()
}

type Card_SmallCoverV5 struct {
	// 小封面条目
	SmallCoverV5 *SmallCoverV5 `protobuf:"bytes,1,opt,name=small_cover_v5,json=smallCoverV5,proto3,oneof"`
}

type Card_LargeCoverV1 struct {
	LargeCoverV1 *LargeCoverV1 `protobuf:"bytes,2,opt,name=large_cover_v1,json=largeCoverV1,proto3,oneof"`
}

type Card_ThreeItemAllV2 struct {
	ThreeItemAllV2 *ThreeItemAllV2 `protobuf:"bytes,3,opt,name=three_item_all_v2,json=threeItemAllV2,proto3,oneof"`
}

type Card_ThreeItemV1 struct {
	ThreeItemV1 *ThreeItemV1 `protobuf:"bytes,4,opt,name=three_item_v1,json=threeItemV1,proto3,oneof"`
}

type Card_HotTopic struct {
	HotTopic *HotTopic `protobuf:"bytes,5,opt,name=hot_topic,json=hotTopic,proto3,oneof"`
}

type Card_ThreeItemHV5 struct {
	ThreeItemHV5 *DynamicHot `protobuf:"bytes,6,opt,name=three_item_h_v5,json=threeItemHV5,proto3,oneof"`
}

type Card_MiddleCoverV3 struct {
	MiddleCoverV3 *MiddleCoverV3 `protobuf:"bytes,7,opt,name=middle_cover_v3,json=middleCoverV3,proto3,oneof"`
}

type Card_LargeCoverV4 struct {
	LargeCoverV4 *LargeCoverV4 `protobuf:"bytes,8,opt,name=large_cover_v4,json=largeCoverV4,proto3,oneof"`
}

type Card_PopularTopEntrance struct {
	// 热门列表顶部按钮
	PopularTopEntrance *PopularTopEntrance `protobuf:"bytes,9,opt,name=popular_top_entrance,json=popularTopEntrance,proto3,oneof"`
}

type Card_RcmdOneItem struct {
	RcmdOneItem *RcmdOneItem `protobuf:"bytes,10,opt,name=rcmd_one_item,json=rcmdOneItem,proto3,oneof"`
}

type Card_SmallCoverV5Ad struct {
	SmallCoverV5Ad *SmallCoverV5Ad `protobuf:"bytes,11,opt,name=small_cover_v5_ad,json=smallCoverV5Ad,proto3,oneof"`
}

func (*Card_SmallCoverV5) isCard_Item() {}

func (*Card_LargeCoverV1) isCard_Item() {}

func (*Card_ThreeItemAllV2) isCard_Item() {}

func (*Card_ThreeItemV1) isCard_Item() {}

func (*Card_HotTopic) isCard_Item() {}

func (*Card_ThreeItemHV5) isCard_Item() {}

func (*Card_MiddleCoverV3) isCard_Item() {}

func (*Card_LargeCoverV4) isCard_Item() {}

func (*Card_PopularTopEntrance) isCard_Item() {}

func (*Card_RcmdOneItem) isCard_Item() {}

func (*Card_SmallCoverV5Ad) isCard_Item() {}

var File_bilibili_app_card_v1_card_proto protoreflect.FileDescriptor

var file_bilibili_app_card_v1_card_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x63,
	0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x14, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x63, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x21, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x06, 0x0a, 0x04, 0x43,
	0x61, 0x72, 0x64, 0x12, 0x4a, 0x0a, 0x0e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x5f, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x5f, 0x76, 0x35, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62, 0x69,
	0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x6d, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x56, 0x35, 0x48,
	0x00, 0x52, 0x0c, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x56, 0x35, 0x12,
	0x4a, 0x0a, 0x0e, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x76,
	0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x61, 0x72, 0x67, 0x65, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x56, 0x31, 0x48, 0x00, 0x52, 0x0c, 0x6c,
	0x61, 0x72, 0x67, 0x65, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x56, 0x31, 0x12, 0x51, 0x0a, 0x11, 0x74,
	0x68, 0x72, 0x65, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x76, 0x32,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x68,
	0x72, 0x65, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x6c, 0x6c, 0x56, 0x32, 0x48, 0x00, 0x52, 0x0e,
	0x74, 0x68, 0x72, 0x65, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x6c, 0x6c, 0x56, 0x32, 0x12, 0x47,
	0x0a, 0x0d, 0x74, 0x68, 0x72, 0x65, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x76, 0x31, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x68, 0x72,
	0x65, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x56, 0x31, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x68, 0x72, 0x65,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x56, 0x31, 0x12, 0x3d, 0x0a, 0x09, 0x68, 0x6f, 0x74, 0x5f, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x69, 0x6c,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x48, 0x6f, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x48, 0x00, 0x52, 0x08, 0x68, 0x6f,
	0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x49, 0x0a, 0x0f, 0x74, 0x68, 0x72, 0x65, 0x65, 0x5f,
	0x69, 0x74, 0x65, 0x6d, 0x5f, 0x68, 0x5f, 0x76, 0x35, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x63,
	0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x48, 0x6f,
	0x74, 0x48, 0x00, 0x52, 0x0c, 0x74, 0x68, 0x72, 0x65, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x48, 0x56,
	0x35, 0x12, 0x4d, 0x0a, 0x0f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x5f, 0x76, 0x33, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x62, 0x69, 0x6c,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x56, 0x33, 0x48,
	0x00, 0x52, 0x0d, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x56, 0x33,
	0x12, 0x4a, 0x0a, 0x0e, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f,
	0x76, 0x34, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x61, 0x72, 0x67, 0x65, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x56, 0x34, 0x48, 0x00, 0x52, 0x0c,
	0x6c, 0x61, 0x72, 0x67, 0x65, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x56, 0x34, 0x12, 0x5c, 0x0a, 0x14,
	0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x74, 0x6f, 0x70, 0x5f, 0x65, 0x6e, 0x74, 0x72,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x62, 0x69, 0x6c,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x54, 0x6f, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x61, 0x6e, 0x63, 0x65, 0x48, 0x00, 0x52, 0x12, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x54,
	0x6f, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0d, 0x72, 0x63,
	0x6d, 0x64, 0x5f, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x63, 0x6d, 0x64, 0x4f, 0x6e, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x48, 0x00, 0x52, 0x0b, 0x72, 0x63, 0x6d, 0x64, 0x4f, 0x6e, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x51, 0x0a, 0x11, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x5f, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x5f, 0x76, 0x35, 0x5f, 0x61, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x61,
	0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6d, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x76, 0x65, 0x72,
	0x56, 0x35, 0x41, 0x64, 0x48, 0x00, 0x52, 0x0e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x76,
	0x65, 0x72, 0x56, 0x35, 0x41, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x42, 0x3d,
	0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x65, 0x62,
	0x69, 0x6e, 0x67, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_app_card_v1_card_proto_rawDescOnce sync.Once
	file_bilibili_app_card_v1_card_proto_rawDescData = file_bilibili_app_card_v1_card_proto_rawDesc
)

func file_bilibili_app_card_v1_card_proto_rawDescGZIP() []byte {
	file_bilibili_app_card_v1_card_proto_rawDescOnce.Do(func() {
		file_bilibili_app_card_v1_card_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_app_card_v1_card_proto_rawDescData)
	})
	return file_bilibili_app_card_v1_card_proto_rawDescData
}

var file_bilibili_app_card_v1_card_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_bilibili_app_card_v1_card_proto_goTypes = []interface{}{
	(*Card)(nil),               // 0: bilibili.app.card.v1.Card
	(*SmallCoverV5)(nil),       // 1: bilibili.app.card.v1.SmallCoverV5
	(*LargeCoverV1)(nil),       // 2: bilibili.app.card.v1.LargeCoverV1
	(*ThreeItemAllV2)(nil),     // 3: bilibili.app.card.v1.ThreeItemAllV2
	(*ThreeItemV1)(nil),        // 4: bilibili.app.card.v1.ThreeItemV1
	(*HotTopic)(nil),           // 5: bilibili.app.card.v1.HotTopic
	(*DynamicHot)(nil),         // 6: bilibili.app.card.v1.DynamicHot
	(*MiddleCoverV3)(nil),      // 7: bilibili.app.card.v1.MiddleCoverV3
	(*LargeCoverV4)(nil),       // 8: bilibili.app.card.v1.LargeCoverV4
	(*PopularTopEntrance)(nil), // 9: bilibili.app.card.v1.PopularTopEntrance
	(*RcmdOneItem)(nil),        // 10: bilibili.app.card.v1.RcmdOneItem
	(*SmallCoverV5Ad)(nil),     // 11: bilibili.app.card.v1.SmallCoverV5Ad
}
var file_bilibili_app_card_v1_card_proto_depIdxs = []int32{
	1,  // 0: bilibili.app.card.v1.Card.small_cover_v5:type_name -> bilibili.app.card.v1.SmallCoverV5
	2,  // 1: bilibili.app.card.v1.Card.large_cover_v1:type_name -> bilibili.app.card.v1.LargeCoverV1
	3,  // 2: bilibili.app.card.v1.Card.three_item_all_v2:type_name -> bilibili.app.card.v1.ThreeItemAllV2
	4,  // 3: bilibili.app.card.v1.Card.three_item_v1:type_name -> bilibili.app.card.v1.ThreeItemV1
	5,  // 4: bilibili.app.card.v1.Card.hot_topic:type_name -> bilibili.app.card.v1.HotTopic
	6,  // 5: bilibili.app.card.v1.Card.three_item_h_v5:type_name -> bilibili.app.card.v1.DynamicHot
	7,  // 6: bilibili.app.card.v1.Card.middle_cover_v3:type_name -> bilibili.app.card.v1.MiddleCoverV3
	8,  // 7: bilibili.app.card.v1.Card.large_cover_v4:type_name -> bilibili.app.card.v1.LargeCoverV4
	9,  // 8: bilibili.app.card.v1.Card.popular_top_entrance:type_name -> bilibili.app.card.v1.PopularTopEntrance
	10, // 9: bilibili.app.card.v1.Card.rcmd_one_item:type_name -> bilibili.app.card.v1.RcmdOneItem
	11, // 10: bilibili.app.card.v1.Card.small_cover_v5_ad:type_name -> bilibili.app.card.v1.SmallCoverV5Ad
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_bilibili_app_card_v1_card_proto_init() }
func file_bilibili_app_card_v1_card_proto_init() {
	if File_bilibili_app_card_v1_card_proto != nil {
		return
	}
	file_bilibili_app_card_v1_single_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_bilibili_app_card_v1_card_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Card); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_bilibili_app_card_v1_card_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Card_SmallCoverV5)(nil),
		(*Card_LargeCoverV1)(nil),
		(*Card_ThreeItemAllV2)(nil),
		(*Card_ThreeItemV1)(nil),
		(*Card_HotTopic)(nil),
		(*Card_ThreeItemHV5)(nil),
		(*Card_MiddleCoverV3)(nil),
		(*Card_LargeCoverV4)(nil),
		(*Card_PopularTopEntrance)(nil),
		(*Card_RcmdOneItem)(nil),
		(*Card_SmallCoverV5Ad)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bilibili_app_card_v1_card_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bilibili_app_card_v1_card_proto_goTypes,
		DependencyIndexes: file_bilibili_app_card_v1_card_proto_depIdxs,
		MessageInfos:      file_bilibili_app_card_v1_card_proto_msgTypes,
	}.Build()
	File_bilibili_app_card_v1_card_proto = out.File
	file_bilibili_app_card_v1_card_proto_rawDesc = nil
	file_bilibili_app_card_v1_card_proto_goTypes = nil
	file_bilibili_app_card_v1_card_proto_depIdxs = nil
}
