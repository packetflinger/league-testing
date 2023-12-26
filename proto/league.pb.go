// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: proto/league.proto

package proto

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

type TeamType int32

const (
	TeamType_TeamSingle   TeamType = 0
	TeamType_TeamMultiple TeamType = 1
)

// Enum value maps for TeamType.
var (
	TeamType_name = map[int32]string{
		0: "TeamSingle",
		1: "TeamMultiple",
	}
	TeamType_value = map[string]int32{
		"TeamSingle":   0,
		"TeamMultiple": 1,
	}
)

func (x TeamType) Enum() *TeamType {
	p := new(TeamType)
	*p = x
	return p
}

func (x TeamType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TeamType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_league_proto_enumTypes[0].Descriptor()
}

func (TeamType) Type() protoreflect.EnumType {
	return &file_proto_league_proto_enumTypes[0]
}

func (x TeamType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TeamType.Descriptor instead.
func (TeamType) EnumDescriptor() ([]byte, []int) {
	return file_proto_league_proto_rawDescGZIP(), []int{0}
}

type League struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid      string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name      string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	StartTime int64    `protobuf:"varint,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   int64    `protobuf:"varint,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Sort      int32    `protobuf:"varint,5,opt,name=sort,proto3" json:"sort,omitempty"`
	TeamType  TeamType `protobuf:"varint,6,opt,name=team_type,json=teamType,proto3,enum=proto.TeamType" json:"team_type,omitempty"`
	Team      []*Team  `protobuf:"bytes,7,rep,name=team,proto3" json:"team,omitempty"`
}

func (x *League) Reset() {
	*x = League{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_league_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *League) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*League) ProtoMessage() {}

func (x *League) ProtoReflect() protoreflect.Message {
	mi := &file_proto_league_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use League.ProtoReflect.Descriptor instead.
func (*League) Descriptor() ([]byte, []int) {
	return file_proto_league_proto_rawDescGZIP(), []int{0}
}

func (x *League) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *League) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *League) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *League) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *League) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *League) GetTeamType() TeamType {
	if x != nil {
		return x.TeamType
	}
	return TeamType_TeamSingle
}

func (x *League) GetTeam() []*Team {
	if x != nil {
		return x.Team
	}
	return nil
}

type Team struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid   string    `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name   string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Player []*Player `protobuf:"bytes,3,rep,name=player,proto3" json:"player,omitempty"`
}

func (x *Team) Reset() {
	*x = Team{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_league_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Team) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Team) ProtoMessage() {}

func (x *Team) ProtoReflect() protoreflect.Message {
	mi := &file_proto_league_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Team.ProtoReflect.Descriptor instead.
func (*Team) Descriptor() ([]byte, []int) {
	return file_proto_league_proto_rawDescGZIP(), []int{1}
}

func (x *Team) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Team) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Team) GetPlayer() []*Player {
	if x != nil {
		return x.Player
	}
	return nil
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_league_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_proto_league_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_proto_league_proto_rawDescGZIP(), []int{2}
}

func (x *Player) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Match struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid          string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	TeamHome      *Team  `protobuf:"bytes,2,opt,name=team_home,json=teamHome,proto3" json:"team_home,omitempty"`
	TeamAway      *Team  `protobuf:"bytes,3,opt,name=team_away,json=teamAway,proto3" json:"team_away,omitempty"`
	Played        bool   `protobuf:"varint,4,opt,name=played,proto3" json:"played,omitempty"`
	TimeScheduled int64  `protobuf:"varint,5,opt,name=time_scheduled,json=timeScheduled,proto3" json:"time_scheduled,omitempty"`
	TimePlayed    int64  `protobuf:"varint,6,opt,name=time_played,json=timePlayed,proto3" json:"time_played,omitempty"`
	Forfeit       bool   `protobuf:"varint,7,opt,name=forfeit,proto3" json:"forfeit,omitempty"`
	Winner        *Team  `protobuf:"bytes,8,opt,name=winner,proto3" json:"winner,omitempty"`
}

func (x *Match) Reset() {
	*x = Match{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_league_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Match) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Match) ProtoMessage() {}

func (x *Match) ProtoReflect() protoreflect.Message {
	mi := &file_proto_league_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Match.ProtoReflect.Descriptor instead.
func (*Match) Descriptor() ([]byte, []int) {
	return file_proto_league_proto_rawDescGZIP(), []int{3}
}

func (x *Match) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Match) GetTeamHome() *Team {
	if x != nil {
		return x.TeamHome
	}
	return nil
}

func (x *Match) GetTeamAway() *Team {
	if x != nil {
		return x.TeamAway
	}
	return nil
}

func (x *Match) GetPlayed() bool {
	if x != nil {
		return x.Played
	}
	return false
}

func (x *Match) GetTimeScheduled() int64 {
	if x != nil {
		return x.TimeScheduled
	}
	return 0
}

func (x *Match) GetTimePlayed() int64 {
	if x != nil {
		return x.TimePlayed
	}
	return 0
}

func (x *Match) GetForfeit() bool {
	if x != nil {
		return x.Forfeit
	}
	return false
}

func (x *Match) GetWinner() *Team {
	if x != nil {
		return x.Winner
	}
	return nil
}

var File_proto_league_proto protoreflect.FileDescriptor

var file_proto_league_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x01, 0x0a, 0x06,
	0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x2c, 0x0a, 0x09,
	0x74, 0x65, 0x61, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x74, 0x65,
	0x61, 0x6d, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x22, 0x55, 0x0a, 0x04, 0x54,
	0x65, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x22, 0x1c, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x8e, 0x02, 0x0a, 0x05, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x28,
	0x0a, 0x09, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x68, 0x6f, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x08,
	0x74, 0x65, 0x61, 0x6d, 0x48, 0x6f, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x09, 0x74, 0x65, 0x61, 0x6d,
	0x5f, 0x61, 0x77, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x41, 0x77,
	0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x6f, 0x72, 0x66, 0x65, 0x69, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x66, 0x6f, 0x72, 0x66, 0x65, 0x69, 0x74, 0x12, 0x23, 0x0a, 0x06,
	0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65,
	0x72, 0x2a, 0x2c, 0x0a, 0x08, 0x54, 0x65, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a,
	0x0a, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x10, 0x00, 0x12, 0x10, 0x0a,
	0x0c, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x10, 0x01, 0x42,
	0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x66, 0x6c, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2f, 0x6c, 0x65, 0x61, 0x67,
	0x75, 0x65, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_league_proto_rawDescOnce sync.Once
	file_proto_league_proto_rawDescData = file_proto_league_proto_rawDesc
)

func file_proto_league_proto_rawDescGZIP() []byte {
	file_proto_league_proto_rawDescOnce.Do(func() {
		file_proto_league_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_league_proto_rawDescData)
	})
	return file_proto_league_proto_rawDescData
}

var file_proto_league_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_league_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_league_proto_goTypes = []interface{}{
	(TeamType)(0),  // 0: proto.TeamType
	(*League)(nil), // 1: proto.League
	(*Team)(nil),   // 2: proto.Team
	(*Player)(nil), // 3: proto.Player
	(*Match)(nil),  // 4: proto.Match
}
var file_proto_league_proto_depIdxs = []int32{
	0, // 0: proto.League.team_type:type_name -> proto.TeamType
	2, // 1: proto.League.team:type_name -> proto.Team
	3, // 2: proto.Team.player:type_name -> proto.Player
	2, // 3: proto.Match.team_home:type_name -> proto.Team
	2, // 4: proto.Match.team_away:type_name -> proto.Team
	2, // 5: proto.Match.winner:type_name -> proto.Team
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_league_proto_init() }
func file_proto_league_proto_init() {
	if File_proto_league_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_league_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*League); i {
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
		file_proto_league_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Team); i {
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
		file_proto_league_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
		file_proto_league_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Match); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_league_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_league_proto_goTypes,
		DependencyIndexes: file_proto_league_proto_depIdxs,
		EnumInfos:         file_proto_league_proto_enumTypes,
		MessageInfos:      file_proto_league_proto_msgTypes,
	}.Build()
	File_proto_league_proto = out.File
	file_proto_league_proto_rawDesc = nil
	file_proto_league_proto_goTypes = nil
	file_proto_league_proto_depIdxs = nil
}
