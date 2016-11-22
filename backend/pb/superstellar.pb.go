// Code generated by protoc-gen-go.
// source: superstellar.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	superstellar.proto

It has these top-level messages:
	Message
	Point
	Vector
	Spaceship
	ProjectileFired
	ProjectileHit
	PlayerLeft
	PlayerJoined
	PlayerDied
	Ping
	Pong
	Space
	Hello
	JoinGame
	JoinGameAck
	UserAction
	TargetAngle
	UserMessage
	Rank
	Leaderboard
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UserInput int32

const (
	UserInput_CENTER     UserInput = 0
	UserInput_LEFT       UserInput = 1
	UserInput_RIGHT      UserInput = 2
	UserInput_THRUST_ON  UserInput = 3
	UserInput_THRUST_OFF UserInput = 4
	UserInput_FIRE_START UserInput = 5
	UserInput_FIRE_STOP  UserInput = 6
)

var UserInput_name = map[int32]string{
	0: "CENTER",
	1: "LEFT",
	2: "RIGHT",
	3: "THRUST_ON",
	4: "THRUST_OFF",
	5: "FIRE_START",
	6: "FIRE_STOP",
}
var UserInput_value = map[string]int32{
	"CENTER":     0,
	"LEFT":       1,
	"RIGHT":      2,
	"THRUST_ON":  3,
	"THRUST_OFF": 4,
	"FIRE_START": 5,
	"FIRE_STOP":  6,
}

func (x UserInput) String() string {
	return proto.EnumName(UserInput_name, int32(x))
}
func (UserInput) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Message struct {
	// Types that are valid to be assigned to Content:
	//	*Message_Space
	//	*Message_Hello
	//	*Message_PlayerLeft
	//	*Message_ProjectileFired
	//	*Message_ProjectileHit
	//	*Message_PlayerJoined
	//	*Message_JoinGameAck
	//	*Message_Leaderboard
	//	*Message_PlayerDied
	//	*Message_Pong
	Content isMessage_Content `protobuf_oneof:"content"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isMessage_Content interface {
	isMessage_Content()
}

type Message_Space struct {
	Space *Space `protobuf:"bytes,1,opt,name=space,oneof"`
}
type Message_Hello struct {
	Hello *Hello `protobuf:"bytes,2,opt,name=hello,oneof"`
}
type Message_PlayerLeft struct {
	PlayerLeft *PlayerLeft `protobuf:"bytes,3,opt,name=playerLeft,oneof"`
}
type Message_ProjectileFired struct {
	ProjectileFired *ProjectileFired `protobuf:"bytes,4,opt,name=projectileFired,oneof"`
}
type Message_ProjectileHit struct {
	ProjectileHit *ProjectileHit `protobuf:"bytes,5,opt,name=projectileHit,oneof"`
}
type Message_PlayerJoined struct {
	PlayerJoined *PlayerJoined `protobuf:"bytes,6,opt,name=playerJoined,oneof"`
}
type Message_JoinGameAck struct {
	JoinGameAck *JoinGameAck `protobuf:"bytes,7,opt,name=joinGameAck,oneof"`
}
type Message_Leaderboard struct {
	Leaderboard *Leaderboard `protobuf:"bytes,8,opt,name=leaderboard,oneof"`
}
type Message_PlayerDied struct {
	PlayerDied *PlayerDied `protobuf:"bytes,9,opt,name=playerDied,oneof"`
}
type Message_Pong struct {
	Pong *Pong `protobuf:"bytes,10,opt,name=pong,oneof"`
}

func (*Message_Space) isMessage_Content()           {}
func (*Message_Hello) isMessage_Content()           {}
func (*Message_PlayerLeft) isMessage_Content()      {}
func (*Message_ProjectileFired) isMessage_Content() {}
func (*Message_ProjectileHit) isMessage_Content()   {}
func (*Message_PlayerJoined) isMessage_Content()    {}
func (*Message_JoinGameAck) isMessage_Content()     {}
func (*Message_Leaderboard) isMessage_Content()     {}
func (*Message_PlayerDied) isMessage_Content()      {}
func (*Message_Pong) isMessage_Content()            {}

func (m *Message) GetContent() isMessage_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Message) GetSpace() *Space {
	if x, ok := m.GetContent().(*Message_Space); ok {
		return x.Space
	}
	return nil
}

func (m *Message) GetHello() *Hello {
	if x, ok := m.GetContent().(*Message_Hello); ok {
		return x.Hello
	}
	return nil
}

func (m *Message) GetPlayerLeft() *PlayerLeft {
	if x, ok := m.GetContent().(*Message_PlayerLeft); ok {
		return x.PlayerLeft
	}
	return nil
}

func (m *Message) GetProjectileFired() *ProjectileFired {
	if x, ok := m.GetContent().(*Message_ProjectileFired); ok {
		return x.ProjectileFired
	}
	return nil
}

func (m *Message) GetProjectileHit() *ProjectileHit {
	if x, ok := m.GetContent().(*Message_ProjectileHit); ok {
		return x.ProjectileHit
	}
	return nil
}

func (m *Message) GetPlayerJoined() *PlayerJoined {
	if x, ok := m.GetContent().(*Message_PlayerJoined); ok {
		return x.PlayerJoined
	}
	return nil
}

func (m *Message) GetJoinGameAck() *JoinGameAck {
	if x, ok := m.GetContent().(*Message_JoinGameAck); ok {
		return x.JoinGameAck
	}
	return nil
}

func (m *Message) GetLeaderboard() *Leaderboard {
	if x, ok := m.GetContent().(*Message_Leaderboard); ok {
		return x.Leaderboard
	}
	return nil
}

func (m *Message) GetPlayerDied() *PlayerDied {
	if x, ok := m.GetContent().(*Message_PlayerDied); ok {
		return x.PlayerDied
	}
	return nil
}

func (m *Message) GetPong() *Pong {
	if x, ok := m.GetContent().(*Message_Pong); ok {
		return x.Pong
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Message) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Message_OneofMarshaler, _Message_OneofUnmarshaler, _Message_OneofSizer, []interface{}{
		(*Message_Space)(nil),
		(*Message_Hello)(nil),
		(*Message_PlayerLeft)(nil),
		(*Message_ProjectileFired)(nil),
		(*Message_ProjectileHit)(nil),
		(*Message_PlayerJoined)(nil),
		(*Message_JoinGameAck)(nil),
		(*Message_Leaderboard)(nil),
		(*Message_PlayerDied)(nil),
		(*Message_Pong)(nil),
	}
}

func _Message_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Message)
	// content
	switch x := m.Content.(type) {
	case *Message_Space:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Space); err != nil {
			return err
		}
	case *Message_Hello:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Hello); err != nil {
			return err
		}
	case *Message_PlayerLeft:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PlayerLeft); err != nil {
			return err
		}
	case *Message_ProjectileFired:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ProjectileFired); err != nil {
			return err
		}
	case *Message_ProjectileHit:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ProjectileHit); err != nil {
			return err
		}
	case *Message_PlayerJoined:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PlayerJoined); err != nil {
			return err
		}
	case *Message_JoinGameAck:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.JoinGameAck); err != nil {
			return err
		}
	case *Message_Leaderboard:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Leaderboard); err != nil {
			return err
		}
	case *Message_PlayerDied:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PlayerDied); err != nil {
			return err
		}
	case *Message_Pong:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Pong); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Message.Content has unexpected type %T", x)
	}
	return nil
}

func _Message_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Message)
	switch tag {
	case 1: // content.space
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Space)
		err := b.DecodeMessage(msg)
		m.Content = &Message_Space{msg}
		return true, err
	case 2: // content.hello
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Hello)
		err := b.DecodeMessage(msg)
		m.Content = &Message_Hello{msg}
		return true, err
	case 3: // content.playerLeft
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PlayerLeft)
		err := b.DecodeMessage(msg)
		m.Content = &Message_PlayerLeft{msg}
		return true, err
	case 4: // content.projectileFired
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ProjectileFired)
		err := b.DecodeMessage(msg)
		m.Content = &Message_ProjectileFired{msg}
		return true, err
	case 5: // content.projectileHit
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ProjectileHit)
		err := b.DecodeMessage(msg)
		m.Content = &Message_ProjectileHit{msg}
		return true, err
	case 6: // content.playerJoined
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PlayerJoined)
		err := b.DecodeMessage(msg)
		m.Content = &Message_PlayerJoined{msg}
		return true, err
	case 7: // content.joinGameAck
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(JoinGameAck)
		err := b.DecodeMessage(msg)
		m.Content = &Message_JoinGameAck{msg}
		return true, err
	case 8: // content.leaderboard
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Leaderboard)
		err := b.DecodeMessage(msg)
		m.Content = &Message_Leaderboard{msg}
		return true, err
	case 9: // content.playerDied
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PlayerDied)
		err := b.DecodeMessage(msg)
		m.Content = &Message_PlayerDied{msg}
		return true, err
	case 10: // content.pong
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Pong)
		err := b.DecodeMessage(msg)
		m.Content = &Message_Pong{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Message_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Message)
	// content
	switch x := m.Content.(type) {
	case *Message_Space:
		s := proto.Size(x.Space)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_Hello:
		s := proto.Size(x.Hello)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_PlayerLeft:
		s := proto.Size(x.PlayerLeft)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_ProjectileFired:
		s := proto.Size(x.ProjectileFired)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_ProjectileHit:
		s := proto.Size(x.ProjectileHit)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_PlayerJoined:
		s := proto.Size(x.PlayerJoined)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_JoinGameAck:
		s := proto.Size(x.JoinGameAck)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_Leaderboard:
		s := proto.Size(x.Leaderboard)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_PlayerDied:
		s := proto.Size(x.PlayerDied)
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_Pong:
		s := proto.Size(x.Pong)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Point struct {
	X int32 `protobuf:"zigzag32,1,opt,name=x" json:"x,omitempty"`
	Y int32 `protobuf:"zigzag32,2,opt,name=y" json:"y,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Vector struct {
	X float32 `protobuf:"fixed32,1,opt,name=x" json:"x,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=y" json:"y,omitempty"`
}

func (m *Vector) Reset()                    { *m = Vector{} }
func (m *Vector) String() string            { return proto.CompactTextString(m) }
func (*Vector) ProtoMessage()               {}
func (*Vector) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Spaceship struct {
	Id          uint32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Position    *Point  `protobuf:"bytes,2,opt,name=position" json:"position,omitempty"`
	Velocity    *Vector `protobuf:"bytes,3,opt,name=velocity" json:"velocity,omitempty"`
	Facing      float32 `protobuf:"fixed32,4,opt,name=facing" json:"facing,omitempty"`
	InputThrust bool    `protobuf:"varint,5,opt,name=inputThrust" json:"inputThrust,omitempty"`
	MaxHp       uint32  `protobuf:"varint,6,opt,name=maxHp" json:"maxHp,omitempty"`
	Hp          uint32  `protobuf:"varint,7,opt,name=hp" json:"hp,omitempty"`
}

func (m *Spaceship) Reset()                    { *m = Spaceship{} }
func (m *Spaceship) String() string            { return proto.CompactTextString(m) }
func (*Spaceship) ProtoMessage()               {}
func (*Spaceship) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Spaceship) GetPosition() *Point {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *Spaceship) GetVelocity() *Vector {
	if m != nil {
		return m.Velocity
	}
	return nil
}

type ProjectileFired struct {
	Id       uint32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	FrameId  uint32  `protobuf:"varint,2,opt,name=frameId" json:"frameId,omitempty"`
	Origin   *Point  `protobuf:"bytes,3,opt,name=origin" json:"origin,omitempty"`
	Velocity *Vector `protobuf:"bytes,4,opt,name=velocity" json:"velocity,omitempty"`
	Facing   float32 `protobuf:"fixed32,5,opt,name=facing" json:"facing,omitempty"`
	Ttl      uint32  `protobuf:"varint,6,opt,name=ttl" json:"ttl,omitempty"`
}

func (m *ProjectileFired) Reset()                    { *m = ProjectileFired{} }
func (m *ProjectileFired) String() string            { return proto.CompactTextString(m) }
func (*ProjectileFired) ProtoMessage()               {}
func (*ProjectileFired) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ProjectileFired) GetOrigin() *Point {
	if m != nil {
		return m.Origin
	}
	return nil
}

func (m *ProjectileFired) GetVelocity() *Vector {
	if m != nil {
		return m.Velocity
	}
	return nil
}

type ProjectileHit struct {
	Id uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *ProjectileHit) Reset()                    { *m = ProjectileHit{} }
func (m *ProjectileHit) String() string            { return proto.CompactTextString(m) }
func (*ProjectileHit) ProtoMessage()               {}
func (*ProjectileHit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type PlayerLeft struct {
	Id uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *PlayerLeft) Reset()                    { *m = PlayerLeft{} }
func (m *PlayerLeft) String() string            { return proto.CompactTextString(m) }
func (*PlayerLeft) ProtoMessage()               {}
func (*PlayerLeft) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type PlayerJoined struct {
	Id       uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
}

func (m *PlayerJoined) Reset()                    { *m = PlayerJoined{} }
func (m *PlayerJoined) String() string            { return proto.CompactTextString(m) }
func (*PlayerJoined) ProtoMessage()               {}
func (*PlayerJoined) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type PlayerDied struct {
	Id       uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	KilledBy uint32 `protobuf:"varint,2,opt,name=killedBy" json:"killedBy,omitempty"`
}

func (m *PlayerDied) Reset()                    { *m = PlayerDied{} }
func (m *PlayerDied) String() string            { return proto.CompactTextString(m) }
func (*PlayerDied) ProtoMessage()               {}
func (*PlayerDied) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type Ping struct {
	Id uint32 `protobuf:"varint,1,opt,name=Id,json=id" json:"Id,omitempty"`
}

func (m *Ping) Reset()                    { *m = Ping{} }
func (m *Ping) String() string            { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()               {}
func (*Ping) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type Pong struct {
	Id uint32 `protobuf:"varint,1,opt,name=Id,json=id" json:"Id,omitempty"`
}

func (m *Pong) Reset()                    { *m = Pong{} }
func (m *Pong) String() string            { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()               {}
func (*Pong) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type Space struct {
	PhysicsFrameID uint32       `protobuf:"varint,1,opt,name=physicsFrameID" json:"physicsFrameID,omitempty"`
	Spaceships     []*Spaceship `protobuf:"bytes,2,rep,name=spaceships" json:"spaceships,omitempty"`
}

func (m *Space) Reset()                    { *m = Space{} }
func (m *Space) String() string            { return proto.CompactTextString(m) }
func (*Space) ProtoMessage()               {}
func (*Space) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Space) GetSpaceships() []*Spaceship {
	if m != nil {
		return m.Spaceships
	}
	return nil
}

type Hello struct {
	MyId                 uint32            `protobuf:"varint,1,opt,name=myId" json:"myId,omitempty"`
	IdToUsername         map[uint32]string `protobuf:"bytes,2,rep,name=idToUsername" json:"idToUsername,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	WorldRadius          float32           `protobuf:"fixed32,3,opt,name=worldRadius" json:"worldRadius,omitempty"`
	BoundaryAnnulusWidth float32           `protobuf:"fixed32,4,opt,name=boundaryAnnulusWidth" json:"boundaryAnnulusWidth,omitempty"`
}

func (m *Hello) Reset()                    { *m = Hello{} }
func (m *Hello) String() string            { return proto.CompactTextString(m) }
func (*Hello) ProtoMessage()               {}
func (*Hello) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *Hello) GetIdToUsername() map[uint32]string {
	if m != nil {
		return m.IdToUsername
	}
	return nil
}

type JoinGame struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
}

func (m *JoinGame) Reset()                    { *m = JoinGame{} }
func (m *JoinGame) String() string            { return proto.CompactTextString(m) }
func (*JoinGame) ProtoMessage()               {}
func (*JoinGame) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

type JoinGameAck struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *JoinGameAck) Reset()                    { *m = JoinGameAck{} }
func (m *JoinGameAck) String() string            { return proto.CompactTextString(m) }
func (*JoinGameAck) ProtoMessage()               {}
func (*JoinGameAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

type UserAction struct {
	UserInput UserInput `protobuf:"varint,1,opt,name=userInput,enum=superstellar.UserInput" json:"userInput,omitempty"`
}

func (m *UserAction) Reset()                    { *m = UserAction{} }
func (m *UserAction) String() string            { return proto.CompactTextString(m) }
func (*UserAction) ProtoMessage()               {}
func (*UserAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

type TargetAngle struct {
	Angle float32 `protobuf:"fixed32,1,opt,name=angle" json:"angle,omitempty"`
}

func (m *TargetAngle) Reset()                    { *m = TargetAngle{} }
func (m *TargetAngle) String() string            { return proto.CompactTextString(m) }
func (*TargetAngle) ProtoMessage()               {}
func (*TargetAngle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

type UserMessage struct {
	// Types that are valid to be assigned to Content:
	//	*UserMessage_UserAction
	//	*UserMessage_JoinGame
	//	*UserMessage_Ping
	//	*UserMessage_TargetAngle
	Content isUserMessage_Content `protobuf_oneof:"content"`
}

func (m *UserMessage) Reset()                    { *m = UserMessage{} }
func (m *UserMessage) String() string            { return proto.CompactTextString(m) }
func (*UserMessage) ProtoMessage()               {}
func (*UserMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

type isUserMessage_Content interface {
	isUserMessage_Content()
}

type UserMessage_UserAction struct {
	UserAction *UserAction `protobuf:"bytes,1,opt,name=userAction,oneof"`
}
type UserMessage_JoinGame struct {
	JoinGame *JoinGame `protobuf:"bytes,2,opt,name=joinGame,oneof"`
}
type UserMessage_Ping struct {
	Ping *Ping `protobuf:"bytes,3,opt,name=ping,oneof"`
}
type UserMessage_TargetAngle struct {
	TargetAngle *TargetAngle `protobuf:"bytes,4,opt,name=targetAngle,oneof"`
}

func (*UserMessage_UserAction) isUserMessage_Content()  {}
func (*UserMessage_JoinGame) isUserMessage_Content()    {}
func (*UserMessage_Ping) isUserMessage_Content()        {}
func (*UserMessage_TargetAngle) isUserMessage_Content() {}

func (m *UserMessage) GetContent() isUserMessage_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *UserMessage) GetUserAction() *UserAction {
	if x, ok := m.GetContent().(*UserMessage_UserAction); ok {
		return x.UserAction
	}
	return nil
}

func (m *UserMessage) GetJoinGame() *JoinGame {
	if x, ok := m.GetContent().(*UserMessage_JoinGame); ok {
		return x.JoinGame
	}
	return nil
}

func (m *UserMessage) GetPing() *Ping {
	if x, ok := m.GetContent().(*UserMessage_Ping); ok {
		return x.Ping
	}
	return nil
}

func (m *UserMessage) GetTargetAngle() *TargetAngle {
	if x, ok := m.GetContent().(*UserMessage_TargetAngle); ok {
		return x.TargetAngle
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*UserMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _UserMessage_OneofMarshaler, _UserMessage_OneofUnmarshaler, _UserMessage_OneofSizer, []interface{}{
		(*UserMessage_UserAction)(nil),
		(*UserMessage_JoinGame)(nil),
		(*UserMessage_Ping)(nil),
		(*UserMessage_TargetAngle)(nil),
	}
}

func _UserMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*UserMessage)
	// content
	switch x := m.Content.(type) {
	case *UserMessage_UserAction:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UserAction); err != nil {
			return err
		}
	case *UserMessage_JoinGame:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.JoinGame); err != nil {
			return err
		}
	case *UserMessage_Ping:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Ping); err != nil {
			return err
		}
	case *UserMessage_TargetAngle:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TargetAngle); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("UserMessage.Content has unexpected type %T", x)
	}
	return nil
}

func _UserMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*UserMessage)
	switch tag {
	case 1: // content.userAction
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UserAction)
		err := b.DecodeMessage(msg)
		m.Content = &UserMessage_UserAction{msg}
		return true, err
	case 2: // content.joinGame
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(JoinGame)
		err := b.DecodeMessage(msg)
		m.Content = &UserMessage_JoinGame{msg}
		return true, err
	case 3: // content.ping
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Ping)
		err := b.DecodeMessage(msg)
		m.Content = &UserMessage_Ping{msg}
		return true, err
	case 4: // content.targetAngle
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TargetAngle)
		err := b.DecodeMessage(msg)
		m.Content = &UserMessage_TargetAngle{msg}
		return true, err
	default:
		return false, nil
	}
}

func _UserMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*UserMessage)
	// content
	switch x := m.Content.(type) {
	case *UserMessage_UserAction:
		s := proto.Size(x.UserAction)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UserMessage_JoinGame:
		s := proto.Size(x.JoinGame)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UserMessage_Ping:
		s := proto.Size(x.Ping)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UserMessage_TargetAngle:
		s := proto.Size(x.TargetAngle)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Rank struct {
	Id    uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Score uint32 `protobuf:"varint,2,opt,name=score" json:"score,omitempty"`
}

func (m *Rank) Reset()                    { *m = Rank{} }
func (m *Rank) String() string            { return proto.CompactTextString(m) }
func (*Rank) ProtoMessage()               {}
func (*Rank) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

type Leaderboard struct {
	Ranks        []*Rank `protobuf:"bytes,1,rep,name=ranks" json:"ranks,omitempty"`
	ClientId     uint32  `protobuf:"varint,2,opt,name=clientId" json:"clientId,omitempty"`
	UserScore    uint32  `protobuf:"varint,3,opt,name=userScore" json:"userScore,omitempty"`
	UserPosition uint32  `protobuf:"varint,4,opt,name=userPosition" json:"userPosition,omitempty"`
}

func (m *Leaderboard) Reset()                    { *m = Leaderboard{} }
func (m *Leaderboard) String() string            { return proto.CompactTextString(m) }
func (*Leaderboard) ProtoMessage()               {}
func (*Leaderboard) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *Leaderboard) GetRanks() []*Rank {
	if m != nil {
		return m.Ranks
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "superstellar.Message")
	proto.RegisterType((*Point)(nil), "superstellar.Point")
	proto.RegisterType((*Vector)(nil), "superstellar.Vector")
	proto.RegisterType((*Spaceship)(nil), "superstellar.Spaceship")
	proto.RegisterType((*ProjectileFired)(nil), "superstellar.ProjectileFired")
	proto.RegisterType((*ProjectileHit)(nil), "superstellar.ProjectileHit")
	proto.RegisterType((*PlayerLeft)(nil), "superstellar.PlayerLeft")
	proto.RegisterType((*PlayerJoined)(nil), "superstellar.PlayerJoined")
	proto.RegisterType((*PlayerDied)(nil), "superstellar.PlayerDied")
	proto.RegisterType((*Ping)(nil), "superstellar.Ping")
	proto.RegisterType((*Pong)(nil), "superstellar.Pong")
	proto.RegisterType((*Space)(nil), "superstellar.Space")
	proto.RegisterType((*Hello)(nil), "superstellar.Hello")
	proto.RegisterType((*JoinGame)(nil), "superstellar.JoinGame")
	proto.RegisterType((*JoinGameAck)(nil), "superstellar.JoinGameAck")
	proto.RegisterType((*UserAction)(nil), "superstellar.UserAction")
	proto.RegisterType((*TargetAngle)(nil), "superstellar.TargetAngle")
	proto.RegisterType((*UserMessage)(nil), "superstellar.UserMessage")
	proto.RegisterType((*Rank)(nil), "superstellar.Rank")
	proto.RegisterType((*Leaderboard)(nil), "superstellar.Leaderboard")
	proto.RegisterEnum("superstellar.UserInput", UserInput_name, UserInput_value)
}

func init() { proto.RegisterFile("superstellar.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1039 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x56, 0xdd, 0x6e, 0xdc, 0x44,
	0x14, 0xae, 0x77, 0xed, 0xcd, 0xee, 0xc9, 0x6e, 0xba, 0x1d, 0xa2, 0x60, 0x42, 0x11, 0x95, 0x81,
	0xaa, 0x02, 0x14, 0x50, 0x01, 0x51, 0x45, 0x42, 0x90, 0xa4, 0xd9, 0xee, 0x56, 0xa1, 0x8d, 0x26,
	0x0e, 0x48, 0xdc, 0x54, 0x8e, 0x3d, 0xc9, 0x4e, 0xe2, 0xd8, 0xd6, 0xd8, 0x2e, 0xf1, 0x83, 0xf0,
	0x46, 0x3c, 0x07, 0xe2, 0x1d, 0xb8, 0x46, 0x62, 0xce, 0xf8, 0x7f, 0xd7, 0xbd, 0xe0, 0xce, 0x67,
	0xce, 0xf7, 0x9d, 0x39, 0xff, 0x63, 0x20, 0x71, 0x1a, 0x31, 0x11, 0x27, 0xcc, 0xf7, 0x1d, 0xb1,
	0x17, 0x89, 0x30, 0x09, 0xc9, 0xb8, 0x79, 0x66, 0xfd, 0xad, 0xc3, 0xc6, 0xcf, 0x2c, 0x8e, 0x9d,
	0x2b, 0x46, 0xbe, 0x00, 0x23, 0x8e, 0x1c, 0x97, 0x99, 0xda, 0x23, 0xed, 0xc9, 0xe6, 0xd3, 0xf7,
	0xf6, 0x5a, 0xec, 0x33, 0x54, 0xcd, 0xef, 0xd1, 0x1c, 0x83, 0xe0, 0xa5, 0x54, 0x84, 0x66, 0xaf,
	0x0b, 0x3c, 0x47, 0x15, 0x82, 0x15, 0x86, 0xec, 0x03, 0x44, 0xbe, 0x93, 0x31, 0x71, 0xc2, 0x2e,
	0x13, 0xb3, 0xaf, 0x18, 0x66, 0x9b, 0x71, 0x5a, 0xe9, 0x25, 0xad, 0x81, 0x26, 0x0b, 0xb8, 0x2f,
	0x1d, 0xbf, 0x66, 0x6e, 0xc2, 0x7d, 0x36, 0xe3, 0x82, 0x79, 0xa6, 0xae, 0x0c, 0x7c, 0xb4, 0x62,
	0xa0, 0x0d, 0x92, 0x56, 0x56, 0x79, 0xe4, 0x08, 0x26, 0xf5, 0xd1, 0x9c, 0x27, 0xa6, 0xa1, 0x0c,
	0x7d, 0xf8, 0x2e, 0x43, 0x12, 0x22, 0xcd, 0xb4, 0x39, 0xe4, 0x27, 0x18, 0xe7, 0xde, 0xbd, 0x0c,
	0x79, 0x20, 0x9d, 0x19, 0x28, 0x1b, 0xbb, 0x5d, 0xd1, 0xe4, 0x08, 0x69, 0xa2, 0xc5, 0x20, 0x3f,
	0xc0, 0xe6, 0xb5, 0xfc, 0x7a, 0xe1, 0xdc, 0xb2, 0x03, 0xf7, 0xc6, 0xdc, 0x50, 0x06, 0x3e, 0x68,
	0x1b, 0x78, 0x59, 0x03, 0x24, 0xbf, 0x89, 0x47, 0xba, 0xcf, 0x1c, 0x8f, 0x89, 0x8b, 0xd0, 0x11,
	0x9e, 0x39, 0xec, 0xa2, 0x9f, 0xd4, 0x00, 0xa4, 0x37, 0xf0, 0x75, 0x2d, 0x9e, 0x73, 0xe9, 0xfd,
	0xe8, 0xdd, 0xb5, 0x40, 0x7d, 0x5d, 0x0b, 0x94, 0xc8, 0x13, 0xd0, 0xa3, 0x30, 0xb8, 0x32, 0x41,
	0xb1, 0xc8, 0x0a, 0x4b, 0x6a, 0x24, 0x5e, 0x21, 0x0e, 0x47, 0xb0, 0xe1, 0x86, 0x41, 0xc2, 0x82,
	0xc4, 0xfa, 0x04, 0x8c, 0x53, 0xe9, 0x7e, 0x42, 0xc6, 0xa0, 0xdd, 0xa9, 0xde, 0x7a, 0x40, 0xb5,
	0x3b, 0x94, 0x32, 0xd5, 0x3c, 0x52, 0xca, 0xac, 0x4f, 0x61, 0xf0, 0x8b, 0xcc, 0x71, 0x28, 0x6a,
	0x54, 0xaf, 0x85, 0xea, 0x21, 0xea, 0x2f, 0x0d, 0x46, 0xaa, 0x0f, 0xe3, 0x25, 0x8f, 0xc8, 0x16,
	0xf4, 0xb8, 0xa7, 0xa0, 0x13, 0x2a, 0xbf, 0xc8, 0x57, 0x30, 0x8c, 0xc2, 0x98, 0x27, 0x3c, 0x0c,
	0xba, 0xbb, 0x52, 0xb9, 0x41, 0x2b, 0x10, 0xf9, 0x1a, 0x86, 0x6f, 0x99, 0x1f, 0xba, 0x3c, 0xc9,
	0x8a, 0xa6, 0xdc, 0x6e, 0x13, 0x72, 0x97, 0x68, 0x85, 0x22, 0x3b, 0x30, 0xb8, 0x74, 0x5c, 0x2e,
	0x53, 0xa0, 0x2b, 0x9f, 0x0a, 0x89, 0x3c, 0x82, 0x4d, 0x1e, 0x44, 0x69, 0x62, 0x2f, 0x45, 0x1a,
	0xe7, 0x7d, 0x35, 0xa4, 0xcd, 0x23, 0xb2, 0x0d, 0xc6, 0xad, 0x73, 0x37, 0x8f, 0x54, 0xbf, 0x4c,
	0x68, 0x2e, 0x60, 0x08, 0xcb, 0x48, 0x75, 0x80, 0x0c, 0x61, 0x19, 0x59, 0x7f, 0x6a, 0x70, 0x7f,
	0xa5, 0x91, 0xd7, 0xc2, 0x34, 0x61, 0xe3, 0x52, 0xc8, 0x5e, 0x58, 0x78, 0x2a, 0xca, 0x09, 0x2d,
	0x45, 0x39, 0x93, 0x83, 0x50, 0xf0, 0x2b, 0x1e, 0x14, 0xd1, 0x74, 0x86, 0x5f, 0x40, 0x5a, 0xc1,
	0xeb, 0xff, 0x33, 0x78, 0xa3, 0x15, 0xfc, 0x14, 0xfa, 0x49, 0xe2, 0x17, 0x81, 0xe1, 0xa7, 0xf5,
	0x31, 0x4c, 0x5a, 0x53, 0xb4, 0x1a, 0x83, 0xf5, 0x10, 0xa0, 0x1e, 0xf8, 0x35, 0xed, 0x3e, 0x8c,
	0x9b, 0x03, 0xb4, 0x96, 0x81, 0x5d, 0x18, 0xa6, 0x31, 0x13, 0x81, 0x8c, 0x5a, 0xa5, 0x60, 0x44,
	0x2b, 0xd9, 0x7a, 0x56, 0x5a, 0x56, 0x0d, 0xdb, 0xc1, 0xbc, 0xe1, 0xbe, 0xcf, 0xbc, 0xc3, 0xac,
	0x48, 0x5e, 0x25, 0x5b, 0x3b, 0xa0, 0x9f, 0x62, 0x38, 0x92, 0xb3, 0x68, 0x7a, 0x83, 0xe7, 0x61,
	0xc7, 0xf9, 0x12, 0x0c, 0xd5, 0x8b, 0xe4, 0x31, 0x6c, 0x45, 0xcb, 0x2c, 0xe6, 0x6e, 0x3c, 0x53,
	0x85, 0x78, 0x5e, 0x80, 0x56, 0x4e, 0xc9, 0xf7, 0x00, 0x71, 0xd9, 0xbc, 0xb1, 0xbc, 0xbe, 0x2f,
	0x73, 0xfe, 0x7e, 0xc7, 0x92, 0x45, 0x3d, 0x6d, 0x40, 0xad, 0x7f, 0x35, 0x30, 0xd4, 0x46, 0x25,
	0x04, 0xf4, 0xdb, 0xac, 0xf2, 0x42, 0x7d, 0xcb, 0x05, 0x39, 0xe6, 0x9e, 0x1d, 0x9e, 0xd7, 0x19,
	0x41, 0xc3, 0x9f, 0x75, 0x2c, 0xe4, 0xbd, 0x45, 0x03, 0x77, 0x1c, 0x24, 0x22, 0xa3, 0x2d, 0x2a,
	0xb6, 0xf1, 0xef, 0xa1, 0xf0, 0x3d, 0xea, 0x78, 0x3c, 0x8d, 0x55, 0x17, 0xf5, 0x68, 0xf3, 0x88,
	0x3c, 0x85, 0xed, 0x8b, 0x30, 0x0d, 0x3c, 0x47, 0x64, 0x07, 0x41, 0x90, 0xfa, 0x69, 0xfc, 0x2b,
	0xf7, 0x92, 0x65, 0x31, 0x0e, 0x9d, 0xba, 0xdd, 0x1f, 0xe1, 0xc1, 0xda, 0xc5, 0xd8, 0x34, 0x37,
	0x2c, 0x2b, 0x02, 0xc1, 0x4f, 0x9c, 0x90, 0xb7, 0x8e, 0x9f, 0x96, 0x25, 0xcd, 0x85, 0xfd, 0xde,
	0x33, 0xcd, 0x7a, 0x0c, 0xc3, 0x72, 0x1f, 0xb6, 0x6a, 0xaf, 0xad, 0xd4, 0x5e, 0x6e, 0xc6, 0xc6,
	0xde, 0xc4, 0x41, 0x89, 0x53, 0x57, 0x26, 0x31, 0x56, 0xc8, 0x21, 0x2d, 0x45, 0xbc, 0x8a, 0x09,
	0x11, 0x8a, 0xf2, 0x2a, 0x25, 0x58, 0x47, 0x00, 0xe8, 0xe3, 0x81, 0xab, 0x96, 0xc3, 0x77, 0x30,
	0x42, 0xc3, 0x0b, 0x9c, 0x61, 0xc5, 0xdf, 0x5a, 0x2d, 0xd6, 0x79, 0xa9, 0xa6, 0x35, 0x52, 0x6e,
	0xbb, 0x4d, 0xdb, 0x11, 0x57, 0x2c, 0x39, 0x08, 0xae, 0x7c, 0x86, 0x37, 0x39, 0xf8, 0x51, 0x6c,
	0xb4, 0x5c, 0xb0, 0xfe, 0xd1, 0x60, 0x13, 0xd9, 0xe5, 0xcb, 0x2b, 0x77, 0x72, 0x5a, 0xdd, 0x5c,
	0x3c, 0xbf, 0xe6, 0xfa, 0x65, 0xb9, 0x1e, 0x77, 0x72, 0x8d, 0x26, 0xdf, 0xc2, 0xb0, 0x7c, 0x1d,
	0x8a, 0xad, 0xb7, 0xd3, 0xfd, 0x94, 0x48, 0x5e, 0x85, 0x54, 0x9b, 0x1c, 0x27, 0xb9, 0xdf, 0xb9,
	0xc9, 0x79, 0xb1, 0xc9, 0x71, 0x1c, 0x64, 0x52, 0x93, 0x3a, 0xa0, 0x62, 0x55, 0xac, 0x3c, 0x37,
	0x8d, 0x88, 0xf1, 0xb9, 0x69, 0xe0, 0x9b, 0x0f, 0xc1, 0x97, 0xa0, 0x53, 0x27, 0xb8, 0x59, 0x1b,
	0x4a, 0x99, 0xa3, 0xd8, 0x0d, 0x05, 0x2b, 0x26, 0x32, 0x17, 0xac, 0x3f, 0x64, 0x8e, 0x1a, 0xcf,
	0x98, 0xf4, 0xd8, 0x10, 0x92, 0x8d, 0xb5, 0xec, 0xaf, 0xbb, 0x8c, 0x86, 0x69, 0x0e, 0xc0, 0x16,
	0x71, 0x7d, 0x2e, 0x6f, 0xac, 0x36, 0x64, 0x25, 0x93, 0x87, 0x79, 0x55, 0xcf, 0xd4, 0x7d, 0x7d,
	0xa5, 0xac, 0x0f, 0x88, 0x05, 0x63, 0x14, 0x4e, 0xcb, 0x57, 0x44, 0x57, 0x80, 0xd6, 0xd9, 0xe7,
	0xd7, 0x30, 0xaa, 0x0a, 0x4f, 0x00, 0x06, 0x47, 0xc7, 0xaf, 0xec, 0x63, 0x3a, 0xbd, 0x47, 0x86,
	0xa0, 0x9f, 0x1c, 0xcf, 0xec, 0xa9, 0x46, 0x46, 0x60, 0xd0, 0xc5, 0x8b, 0xb9, 0x3d, 0xed, 0x91,
	0x09, 0x8c, 0xec, 0x39, 0x3d, 0x3f, 0xb3, 0xdf, 0xbc, 0x7e, 0x35, 0xed, 0xcb, 0xd0, 0xa1, 0x14,
	0x67, 0xb3, 0xa9, 0x8e, 0xf2, 0x6c, 0x41, 0x8f, 0xdf, 0x9c, 0xd9, 0x07, 0xd4, 0x9e, 0x1a, 0x08,
	0x2f, 0xe4, 0xd7, 0xa7, 0xd3, 0xc1, 0xa1, 0xfe, 0x5b, 0x2f, 0xba, 0xb8, 0x18, 0xa8, 0x1f, 0xb7,
	0x6f, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x5f, 0x3b, 0x82, 0xce, 0x09, 0x00, 0x00,
}
