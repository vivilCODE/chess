// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: proto/chess.proto

package pb

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

type Color int32

const (
	Color_white Color = 0
	Color_black Color = 1
)

// Enum value maps for Color.
var (
	Color_name = map[int32]string{
		0: "white",
		1: "black",
	}
	Color_value = map[string]int32{
		"white": 0,
		"black": 1,
	}
)

func (x Color) Enum() *Color {
	p := new(Color)
	*p = x
	return p
}

func (x Color) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Color) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_chess_proto_enumTypes[0].Descriptor()
}

func (Color) Type() protoreflect.EnumType {
	return &file_proto_chess_proto_enumTypes[0]
}

func (x Color) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Color.Descriptor instead.
func (Color) EnumDescriptor() ([]byte, []int) {
	return file_proto_chess_proto_rawDescGZIP(), []int{0}
}

type Piece int32

const (
	Piece_nil     Piece = 0
	Piece_wPawn   Piece = 1
	Piece_bPawn   Piece = 2
	Piece_wRook   Piece = 3
	Piece_bRook   Piece = 4
	Piece_wKnight Piece = 5
	Piece_bKnight Piece = 6
	Piece_wBishop Piece = 7
	Piece_bBishop Piece = 8
	Piece_wKing   Piece = 9
	Piece_bKing   Piece = 10
	Piece_wQueen  Piece = 11
	Piece_bQueen  Piece = 12
)

// Enum value maps for Piece.
var (
	Piece_name = map[int32]string{
		0:  "nil",
		1:  "wPawn",
		2:  "bPawn",
		3:  "wRook",
		4:  "bRook",
		5:  "wKnight",
		6:  "bKnight",
		7:  "wBishop",
		8:  "bBishop",
		9:  "wKing",
		10: "bKing",
		11: "wQueen",
		12: "bQueen",
	}
	Piece_value = map[string]int32{
		"nil":     0,
		"wPawn":   1,
		"bPawn":   2,
		"wRook":   3,
		"bRook":   4,
		"wKnight": 5,
		"bKnight": 6,
		"wBishop": 7,
		"bBishop": 8,
		"wKing":   9,
		"bKing":   10,
		"wQueen":  11,
		"bQueen":  12,
	}
)

func (x Piece) Enum() *Piece {
	p := new(Piece)
	*p = x
	return p
}

func (x Piece) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Piece) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_chess_proto_enumTypes[1].Descriptor()
}

func (Piece) Type() protoreflect.EnumType {
	return &file_proto_chess_proto_enumTypes[1]
}

func (x Piece) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Piece.Descriptor instead.
func (Piece) EnumDescriptor() ([]byte, []int) {
	return file_proto_chess_proto_rawDescGZIP(), []int{1}
}

type MakeMoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Game *Game `protobuf:"bytes,1,opt,name=game,proto3" json:"game,omitempty"`
	Move *Move `protobuf:"bytes,2,opt,name=move,proto3" json:"move,omitempty"`
}

func (x *MakeMoveRequest) Reset() {
	*x = MakeMoveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chess_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MakeMoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakeMoveRequest) ProtoMessage() {}

func (x *MakeMoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chess_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakeMoveRequest.ProtoReflect.Descriptor instead.
func (*MakeMoveRequest) Descriptor() ([]byte, []int) {
	return file_proto_chess_proto_rawDescGZIP(), []int{0}
}

func (x *MakeMoveRequest) GetGame() *Game {
	if x != nil {
		return x.Game
	}
	return nil
}

func (x *MakeMoveRequest) GetMove() *Move {
	if x != nil {
		return x.Move
	}
	return nil
}

type Move struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From *Square `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To   *Square `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *Move) Reset() {
	*x = Move{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chess_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Move) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Move) ProtoMessage() {}

func (x *Move) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chess_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Move.ProtoReflect.Descriptor instead.
func (*Move) Descriptor() ([]byte, []int) {
	return file_proto_chess_proto_rawDescGZIP(), []int{1}
}

func (x *Move) GetFrom() *Square {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *Move) GetTo() *Square {
	if x != nil {
		return x.To
	}
	return nil
}

type MakeMoveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *MakeMoveResponse) Reset() {
	*x = MakeMoveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chess_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MakeMoveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakeMoveResponse) ProtoMessage() {}

func (x *MakeMoveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chess_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakeMoveResponse.ProtoReflect.Descriptor instead.
func (*MakeMoveResponse) Descriptor() ([]byte, []int) {
	return file_proto_chess_proto_rawDescGZIP(), []int{2}
}

func (x *MakeMoveResponse) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type NewGameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *NewGameRequest) Reset() {
	*x = NewGameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chess_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewGameRequest) ProtoMessage() {}

func (x *NewGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chess_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewGameRequest.ProtoReflect.Descriptor instead.
func (*NewGameRequest) Descriptor() ([]byte, []int) {
	return file_proto_chess_proto_rawDescGZIP(), []int{3}
}

func (x *NewGameRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type NewGameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Game *Game `protobuf:"bytes,1,opt,name=game,proto3" json:"game,omitempty"`
}

func (x *NewGameResponse) Reset() {
	*x = NewGameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chess_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewGameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewGameResponse) ProtoMessage() {}

func (x *NewGameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chess_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewGameResponse.ProtoReflect.Descriptor instead.
func (*NewGameResponse) Descriptor() ([]byte, []int) {
	return file_proto_chess_proto_rawDescGZIP(), []int{4}
}

func (x *NewGameResponse) GetGame() *Game {
	if x != nil {
		return x.Game
	}
	return nil
}

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chess_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chess_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_proto_chess_proto_rawDescGZIP(), []int{5}
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chess_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chess_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_proto_chess_proto_rawDescGZIP(), []int{6}
}

func (x *PingResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chess_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chess_proto_msgTypes[7]
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
	return file_proto_chess_proto_rawDescGZIP(), []int{7}
}

func (x *Player) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Player) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Game struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        uint32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	PlayerOne *Player `protobuf:"bytes,2,opt,name=playerOne,proto3" json:"playerOne,omitempty"`
	PlayerTwo *Player `protobuf:"bytes,3,opt,name=playerTwo,proto3" json:"playerTwo,omitempty"`
	Board     *Board  `protobuf:"bytes,4,opt,name=board,proto3" json:"board,omitempty"`
}

func (x *Game) Reset() {
	*x = Game{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chess_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Game) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Game) ProtoMessage() {}

func (x *Game) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chess_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Game.ProtoReflect.Descriptor instead.
func (*Game) Descriptor() ([]byte, []int) {
	return file_proto_chess_proto_rawDescGZIP(), []int{8}
}

func (x *Game) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Game) GetPlayerOne() *Player {
	if x != nil {
		return x.PlayerOne
	}
	return nil
}

func (x *Game) GetPlayerTwo() *Player {
	if x != nil {
		return x.PlayerTwo
	}
	return nil
}

func (x *Game) GetBoard() *Board {
	if x != nil {
		return x.Board
	}
	return nil
}

type Board struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Squares []*Square `protobuf:"bytes,1,rep,name=squares,proto3" json:"squares,omitempty"`
}

func (x *Board) Reset() {
	*x = Board{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chess_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Board) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Board) ProtoMessage() {}

func (x *Board) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chess_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Board.ProtoReflect.Descriptor instead.
func (*Board) Descriptor() ([]byte, []int) {
	return file_proto_chess_proto_rawDescGZIP(), []int{9}
}

func (x *Board) GetSquares() []*Square {
	if x != nil {
		return x.Squares
	}
	return nil
}

type Square struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pos   *SquarePosition `protobuf:"bytes,1,opt,name=pos,proto3" json:"pos,omitempty"`
	Color Color           `protobuf:"varint,2,opt,name=color,proto3,enum=pb.Color" json:"color,omitempty"`
	Piece Piece           `protobuf:"varint,3,opt,name=piece,proto3,enum=pb.Piece" json:"piece,omitempty"`
}

func (x *Square) Reset() {
	*x = Square{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chess_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Square) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Square) ProtoMessage() {}

func (x *Square) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chess_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Square.ProtoReflect.Descriptor instead.
func (*Square) Descriptor() ([]byte, []int) {
	return file_proto_chess_proto_rawDescGZIP(), []int{10}
}

func (x *Square) GetPos() *SquarePosition {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (x *Square) GetColor() Color {
	if x != nil {
		return x.Color
	}
	return Color_white
}

func (x *Square) GetPiece() Piece {
	if x != nil {
		return x.Piece
	}
	return Piece_nil
}

type SquarePosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X uint32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y uint32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *SquarePosition) Reset() {
	*x = SquarePosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chess_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SquarePosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SquarePosition) ProtoMessage() {}

func (x *SquarePosition) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chess_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SquarePosition.ProtoReflect.Descriptor instead.
func (*SquarePosition) Descriptor() ([]byte, []int) {
	return file_proto_chess_proto_rawDescGZIP(), []int{11}
}

func (x *SquarePosition) GetX() uint32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *SquarePosition) GetY() uint32 {
	if x != nil {
		return x.Y
	}
	return 0
}

var File_proto_chess_proto protoreflect.FileDescriptor

var file_proto_chess_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x4d, 0x0a, 0x0f, 0x4d, 0x61, 0x6b, 0x65, 0x4d,
	0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x67, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x61,
	0x6d, 0x65, 0x52, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x6d, 0x6f, 0x76, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x6f, 0x76, 0x65,
	0x52, 0x04, 0x6d, 0x6f, 0x76, 0x65, 0x22, 0x42, 0x0a, 0x04, 0x4d, 0x6f, 0x76, 0x65, 0x12, 0x1e,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x1a,
	0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x52, 0x02, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x10, 0x4d, 0x61,
	0x6b, 0x65, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x22, 0x24, 0x0a, 0x0e, 0x4e, 0x65, 0x77, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x2f, 0x0a, 0x0f, 0x4e, 0x65, 0x77, 0x47,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x67,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x61, 0x6d, 0x65, 0x52, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x22, 0x0d, 0x0a, 0x0b, 0x50, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2a, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x8b, 0x01, 0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x09, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x4f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x4f, 0x6e, 0x65, 0x12, 0x28, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x54,
	0x77, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x77, 0x6f, 0x12,
	0x1f, 0x0a, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x70, 0x62, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x22, 0x2d, 0x0a, 0x05, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x24, 0x0a, 0x07, 0x73, 0x71, 0x75,
	0x61, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x52, 0x07, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x73, 0x22,
	0x70, 0x0a, 0x06, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x12, 0x24, 0x0a, 0x03, 0x70, 0x6f, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x71, 0x75, 0x61,
	0x72, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x12,
	0x1f, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72,
	0x12, 0x1f, 0x0a, 0x05, 0x70, 0x69, 0x65, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x09, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x69, 0x65, 0x63, 0x65, 0x52, 0x05, 0x70, 0x69, 0x65, 0x63,
	0x65, 0x22, 0x2c, 0x0a, 0x0e, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01,
	0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x79, 0x2a,
	0x1d, 0x0a, 0x05, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x09, 0x0a, 0x05, 0x77, 0x68, 0x69, 0x74,
	0x65, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x10, 0x01, 0x2a, 0x9e,
	0x01, 0x0a, 0x05, 0x50, 0x69, 0x65, 0x63, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x6e, 0x69, 0x6c, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x77, 0x50, 0x61, 0x77, 0x6e, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05,
	0x62, 0x50, 0x61, 0x77, 0x6e, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x77, 0x52, 0x6f, 0x6f, 0x6b,
	0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x62, 0x52, 0x6f, 0x6f, 0x6b, 0x10, 0x04, 0x12, 0x0b, 0x0a,
	0x07, 0x77, 0x4b, 0x6e, 0x69, 0x67, 0x68, 0x74, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x62, 0x4b,
	0x6e, 0x69, 0x67, 0x68, 0x74, 0x10, 0x06, 0x12, 0x0b, 0x0a, 0x07, 0x77, 0x42, 0x69, 0x73, 0x68,
	0x6f, 0x70, 0x10, 0x07, 0x12, 0x0b, 0x0a, 0x07, 0x62, 0x42, 0x69, 0x73, 0x68, 0x6f, 0x70, 0x10,
	0x08, 0x12, 0x09, 0x0a, 0x05, 0x77, 0x4b, 0x69, 0x6e, 0x67, 0x10, 0x09, 0x12, 0x09, 0x0a, 0x05,
	0x62, 0x4b, 0x69, 0x6e, 0x67, 0x10, 0x0a, 0x12, 0x0a, 0x0a, 0x06, 0x77, 0x51, 0x75, 0x65, 0x65,
	0x6e, 0x10, 0x0b, 0x12, 0x0a, 0x0a, 0x06, 0x62, 0x51, 0x75, 0x65, 0x65, 0x6e, 0x10, 0x0c, 0x32,
	0xa6, 0x01, 0x0a, 0x08, 0x43, 0x68, 0x65, 0x73, 0x73, 0x41, 0x70, 0x69, 0x12, 0x37, 0x0a, 0x08,
	0x4d, 0x61, 0x6b, 0x65, 0x4d, 0x6f, 0x76, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x61,
	0x6b, 0x65, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x70, 0x62, 0x2e, 0x4d, 0x61, 0x6b, 0x65, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x47, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x65, 0x77, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x65, 0x77, 0x47, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x04, 0x50,
	0x69, 0x6e, 0x67, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_chess_proto_rawDescOnce sync.Once
	file_proto_chess_proto_rawDescData = file_proto_chess_proto_rawDesc
)

func file_proto_chess_proto_rawDescGZIP() []byte {
	file_proto_chess_proto_rawDescOnce.Do(func() {
		file_proto_chess_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_chess_proto_rawDescData)
	})
	return file_proto_chess_proto_rawDescData
}

var file_proto_chess_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_chess_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_chess_proto_goTypes = []interface{}{
	(Color)(0),               // 0: pb.Color
	(Piece)(0),               // 1: pb.Piece
	(*MakeMoveRequest)(nil),  // 2: pb.MakeMoveRequest
	(*Move)(nil),             // 3: pb.Move
	(*MakeMoveResponse)(nil), // 4: pb.MakeMoveResponse
	(*NewGameRequest)(nil),   // 5: pb.NewGameRequest
	(*NewGameResponse)(nil),  // 6: pb.NewGameResponse
	(*PingRequest)(nil),      // 7: pb.PingRequest
	(*PingResponse)(nil),     // 8: pb.PingResponse
	(*Player)(nil),           // 9: pb.Player
	(*Game)(nil),             // 10: pb.Game
	(*Board)(nil),            // 11: pb.Board
	(*Square)(nil),           // 12: pb.Square
	(*SquarePosition)(nil),   // 13: pb.SquarePosition
}
var file_proto_chess_proto_depIdxs = []int32{
	10, // 0: pb.MakeMoveRequest.game:type_name -> pb.Game
	3,  // 1: pb.MakeMoveRequest.move:type_name -> pb.Move
	12, // 2: pb.Move.from:type_name -> pb.Square
	12, // 3: pb.Move.to:type_name -> pb.Square
	10, // 4: pb.NewGameResponse.game:type_name -> pb.Game
	9,  // 5: pb.Game.playerOne:type_name -> pb.Player
	9,  // 6: pb.Game.playerTwo:type_name -> pb.Player
	11, // 7: pb.Game.board:type_name -> pb.Board
	12, // 8: pb.Board.squares:type_name -> pb.Square
	13, // 9: pb.Square.pos:type_name -> pb.SquarePosition
	0,  // 10: pb.Square.color:type_name -> pb.Color
	1,  // 11: pb.Square.piece:type_name -> pb.Piece
	2,  // 12: pb.ChessApi.MakeMove:input_type -> pb.MakeMoveRequest
	5,  // 13: pb.ChessApi.NewGame:input_type -> pb.NewGameRequest
	7,  // 14: pb.ChessApi.Ping:input_type -> pb.PingRequest
	4,  // 15: pb.ChessApi.MakeMove:output_type -> pb.MakeMoveResponse
	6,  // 16: pb.ChessApi.NewGame:output_type -> pb.NewGameResponse
	8,  // 17: pb.ChessApi.Ping:output_type -> pb.PingResponse
	15, // [15:18] is the sub-list for method output_type
	12, // [12:15] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_proto_chess_proto_init() }
func file_proto_chess_proto_init() {
	if File_proto_chess_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_chess_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MakeMoveRequest); i {
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
		file_proto_chess_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Move); i {
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
		file_proto_chess_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MakeMoveResponse); i {
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
		file_proto_chess_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewGameRequest); i {
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
		file_proto_chess_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewGameResponse); i {
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
		file_proto_chess_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_proto_chess_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
		file_proto_chess_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_chess_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Game); i {
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
		file_proto_chess_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Board); i {
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
		file_proto_chess_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Square); i {
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
		file_proto_chess_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SquarePosition); i {
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
			RawDescriptor: file_proto_chess_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_chess_proto_goTypes,
		DependencyIndexes: file_proto_chess_proto_depIdxs,
		EnumInfos:         file_proto_chess_proto_enumTypes,
		MessageInfos:      file_proto_chess_proto_msgTypes,
	}.Build()
	File_proto_chess_proto = out.File
	file_proto_chess_proto_rawDesc = nil
	file_proto_chess_proto_goTypes = nil
	file_proto_chess_proto_depIdxs = nil
}
