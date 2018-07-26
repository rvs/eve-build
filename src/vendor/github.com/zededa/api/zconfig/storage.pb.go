// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage.proto

package zconfig

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DsType int32

const (
	DsType_DsUnknown DsType = 0
	DsType_DsHttp    DsType = 1
	DsType_DsHttps   DsType = 2
	DsType_DsS3      DsType = 3
	DsType_DsSFTP    DsType = 4
)

var DsType_name = map[int32]string{
	0: "DsUnknown",
	1: "DsHttp",
	2: "DsHttps",
	3: "DsS3",
	4: "DsSFTP",
}
var DsType_value = map[string]int32{
	"DsUnknown": 0,
	"DsHttp":    1,
	"DsHttps":   2,
	"DsS3":      3,
	"DsSFTP":    4,
}

func (x DsType) String() string {
	return proto.EnumName(DsType_name, int32(x))
}
func (DsType) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

type Format int32

const (
	Format_FmtUnknown Format = 0
	Format_RAW        Format = 1
	Format_QCOW       Format = 2
	Format_QCOW2      Format = 3
	Format_VHD        Format = 4
	Format_VMDK       Format = 5
	Format_OVA        Format = 6
	Format_VHDX       Format = 7
)

var Format_name = map[int32]string{
	0: "FmtUnknown",
	1: "RAW",
	2: "QCOW",
	3: "QCOW2",
	4: "VHD",
	5: "VMDK",
	6: "OVA",
	7: "VHDX",
}
var Format_value = map[string]int32{
	"FmtUnknown": 0,
	"RAW":        1,
	"QCOW":       2,
	"QCOW2":      3,
	"VHD":        4,
	"VMDK":       5,
	"OVA":        6,
	"VHDX":       7,
}

func (x Format) String() string {
	return proto.EnumName(Format_name, int32(x))
}
func (Format) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

type Target int32

const (
	Target_TgtUnknown Target = 0
	Target_Disk       Target = 1
	Target_Kernel     Target = 2
	Target_Initrd     Target = 3
	Target_RamDisk    Target = 4
)

var Target_name = map[int32]string{
	0: "TgtUnknown",
	1: "Disk",
	2: "Kernel",
	3: "Initrd",
	4: "RamDisk",
}
var Target_value = map[string]int32{
	"TgtUnknown": 0,
	"Disk":       1,
	"Kernel":     2,
	"Initrd":     3,
	"RamDisk":    4,
}

func (x Target) String() string {
	return proto.EnumName(Target_name, int32(x))
}
func (Target) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

type DriveType int32

const (
	DriveType_Unclassified DriveType = 0
	DriveType_CDROM        DriveType = 1
	DriveType_HDD          DriveType = 2
	DriveType_NET          DriveType = 3
)

var DriveType_name = map[int32]string{
	0: "Unclassified",
	1: "CDROM",
	2: "HDD",
	3: "NET",
}
var DriveType_value = map[string]int32{
	"Unclassified": 0,
	"CDROM":        1,
	"HDD":          2,
	"NET":          3,
}

func (x DriveType) String() string {
	return proto.EnumName(DriveType_name, int32(x))
}
func (DriveType) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

type SignatureInfo struct {
	Intercertsurl string `protobuf:"bytes,1,opt,name=intercertsurl" json:"intercertsurl,omitempty"`
	Signercerturl string `protobuf:"bytes,2,opt,name=signercerturl" json:"signercerturl,omitempty"`
	Signature     []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *SignatureInfo) Reset()                    { *m = SignatureInfo{} }
func (m *SignatureInfo) String() string            { return proto.CompactTextString(m) }
func (*SignatureInfo) ProtoMessage()               {}
func (*SignatureInfo) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *SignatureInfo) GetIntercertsurl() string {
	if m != nil {
		return m.Intercertsurl
	}
	return ""
}

func (m *SignatureInfo) GetSignercerturl() string {
	if m != nil {
		return m.Signercerturl
	}
	return ""
}

func (m *SignatureInfo) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type DatastoreConfig struct {
	Id       string `protobuf:"bytes,100,opt,name=id" json:"id,omitempty"`
	DType    DsType `protobuf:"varint,1,opt,name=dType,enum=DsType" json:"dType,omitempty"`
	Fqdn     string `protobuf:"bytes,2,opt,name=fqdn" json:"fqdn,omitempty"`
	ApiKey   string `protobuf:"bytes,3,opt,name=apiKey" json:"apiKey,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password" json:"password,omitempty"`
	// depending on datastore types, it could be bucket or path
	Dpath string `protobuf:"bytes,5,opt,name=dpath" json:"dpath,omitempty"`
}

func (m *DatastoreConfig) Reset()                    { *m = DatastoreConfig{} }
func (m *DatastoreConfig) String() string            { return proto.CompactTextString(m) }
func (*DatastoreConfig) ProtoMessage()               {}
func (*DatastoreConfig) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *DatastoreConfig) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DatastoreConfig) GetDType() DsType {
	if m != nil {
		return m.DType
	}
	return DsType_DsUnknown
}

func (m *DatastoreConfig) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *DatastoreConfig) GetApiKey() string {
	if m != nil {
		return m.ApiKey
	}
	return ""
}

func (m *DatastoreConfig) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *DatastoreConfig) GetDpath() string {
	if m != nil {
		return m.Dpath
	}
	return ""
}

type Image struct {
	Id             string          `protobuf:"bytes,100,opt,name=id" json:"id,omitempty"`
	Uuidandversion *UUIDandVersion `protobuf:"bytes,1,opt,name=uuidandversion" json:"uuidandversion,omitempty"`
	// it could be relative path/name as well
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Sha256  string `protobuf:"bytes,3,opt,name=sha256" json:"sha256,omitempty"`
	Size    int64  `protobuf:"varint,7,opt,name=size" json:"size,omitempty"`
	Iformat Format `protobuf:"varint,4,opt,name=iformat,enum=Format" json:"iformat,omitempty"`
	// if its signed image
	Siginfo   *SignatureInfo `protobuf:"bytes,5,opt,name=siginfo" json:"siginfo,omitempty"`
	DsId      string         `protobuf:"bytes,6,opt,name=dsId" json:"dsId,omitempty"`
	SizeBytes int64          `protobuf:"varint,8,opt,name=sizeBytes" json:"sizeBytes,omitempty"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()               {}
func (*Image) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

func (m *Image) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Image) GetUuidandversion() *UUIDandVersion {
	if m != nil {
		return m.Uuidandversion
	}
	return nil
}

func (m *Image) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Image) GetSha256() string {
	if m != nil {
		return m.Sha256
	}
	return ""
}

func (m *Image) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Image) GetIformat() Format {
	if m != nil {
		return m.Iformat
	}
	return Format_FmtUnknown
}

func (m *Image) GetSiginfo() *SignatureInfo {
	if m != nil {
		return m.Siginfo
	}
	return nil
}

func (m *Image) GetDsId() string {
	if m != nil {
		return m.DsId
	}
	return ""
}

func (m *Image) GetSizeBytes() int64 {
	if m != nil {
		return m.SizeBytes
	}
	return 0
}

type Drive struct {
	Image        *Image    `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
	Maxsize      int64     `protobuf:"varint,2,opt,name=maxsize" json:"maxsize,omitempty"`
	Readonly     bool      `protobuf:"varint,5,opt,name=readonly" json:"readonly,omitempty"`
	Preserve     bool      `protobuf:"varint,6,opt,name=preserve" json:"preserve,omitempty"`
	Drvtype      DriveType `protobuf:"varint,8,opt,name=drvtype,enum=DriveType" json:"drvtype,omitempty"`
	Target       Target    `protobuf:"varint,9,opt,name=target,enum=Target" json:"target,omitempty"`
	Maxsizebytes int64     `protobuf:"varint,10,opt,name=maxsizebytes" json:"maxsizebytes,omitempty"`
}

func (m *Drive) Reset()                    { *m = Drive{} }
func (m *Drive) String() string            { return proto.CompactTextString(m) }
func (*Drive) ProtoMessage()               {}
func (*Drive) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

func (m *Drive) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *Drive) GetMaxsize() int64 {
	if m != nil {
		return m.Maxsize
	}
	return 0
}

func (m *Drive) GetReadonly() bool {
	if m != nil {
		return m.Readonly
	}
	return false
}

func (m *Drive) GetPreserve() bool {
	if m != nil {
		return m.Preserve
	}
	return false
}

func (m *Drive) GetDrvtype() DriveType {
	if m != nil {
		return m.Drvtype
	}
	return DriveType_Unclassified
}

func (m *Drive) GetTarget() Target {
	if m != nil {
		return m.Target
	}
	return Target_TgtUnknown
}

func (m *Drive) GetMaxsizebytes() int64 {
	if m != nil {
		return m.Maxsizebytes
	}
	return 0
}

func init() {
	proto.RegisterType((*SignatureInfo)(nil), "SignatureInfo")
	proto.RegisterType((*DatastoreConfig)(nil), "DatastoreConfig")
	proto.RegisterType((*Image)(nil), "Image")
	proto.RegisterType((*Drive)(nil), "Drive")
	proto.RegisterEnum("DsType", DsType_name, DsType_value)
	proto.RegisterEnum("Format", Format_name, Format_value)
	proto.RegisterEnum("Target", Target_name, Target_value)
	proto.RegisterEnum("DriveType", DriveType_name, DriveType_value)
}

func init() { proto.RegisterFile("storage.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 709 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x94, 0x51, 0x6f, 0x22, 0x37,
	0x10, 0xc7, 0x6f, 0x17, 0x76, 0x17, 0x26, 0x81, 0x58, 0x56, 0x55, 0xad, 0x4e, 0x77, 0xba, 0x14,
	0xdd, 0x03, 0xe2, 0x61, 0x23, 0x71, 0x6a, 0x2b, 0xf5, 0xa9, 0x77, 0xd9, 0x52, 0x50, 0x74, 0x4d,
	0xbb, 0x01, 0xae, 0xaa, 0xfa, 0xe2, 0xac, 0xcd, 0xc6, 0x0a, 0x6b, 0x53, 0xdb, 0x90, 0xc2, 0x97,
	0xe8, 0x37, 0xe8, 0x87, 0xeb, 0x27, 0xa9, 0x6c, 0x03, 0x29, 0xbd, 0xb7, 0xf9, 0xff, 0x67, 0x18,
	0xff, 0xc6, 0x63, 0x16, 0x3a, 0xda, 0x48, 0x45, 0x2a, 0x96, 0xad, 0x94, 0x34, 0xf2, 0xe5, 0x05,
	0x65, 0x9b, 0x52, 0xd6, 0xb5, 0x14, 0xde, 0xe8, 0x6d, 0xa1, 0x73, 0xc7, 0x2b, 0x41, 0xcc, 0x5a,
	0xb1, 0x89, 0x58, 0x48, 0xfc, 0x16, 0x3a, 0x5c, 0x18, 0xa6, 0x4a, 0xa6, 0x8c, 0x5e, 0xab, 0x65,
	0x1a, 0x5c, 0x06, 0xfd, 0x76, 0x71, 0x6a, 0xda, 0x2a, 0xcd, 0x2b, 0xe1, 0x1d, 0x5b, 0x15, 0xfa,
	0xaa, 0x13, 0x13, 0xbf, 0x82, 0xb6, 0x3e, 0x34, 0x4f, 0x1b, 0x97, 0x41, 0xff, 0xbc, 0x78, 0x36,
	0x7a, 0x7f, 0x07, 0x70, 0x91, 0x13, 0x43, 0x2c, 0x21, 0xbb, 0x96, 0x62, 0xc1, 0x2b, 0xdc, 0x85,
	0x90, 0xd3, 0x94, 0xba, 0x66, 0x21, 0xa7, 0xf8, 0x35, 0x44, 0x74, 0xba, 0x5d, 0x31, 0x47, 0xd1,
	0x1d, 0x26, 0x59, 0xae, 0xad, 0x2c, 0xbc, 0x8b, 0x31, 0x34, 0x17, 0x7f, 0x50, 0xb1, 0x3f, 0xdd,
	0xc5, 0xf8, 0x4b, 0x88, 0xc9, 0x8a, 0xdf, 0xb0, 0xad, 0x3b, 0xb1, 0x5d, 0xec, 0x15, 0x7e, 0x09,
	0xad, 0x15, 0xd1, 0xfa, 0x49, 0x2a, 0x9a, 0x36, 0x5d, 0xe6, 0xa8, 0xf1, 0x17, 0x10, 0xd1, 0x15,
	0x31, 0x0f, 0x69, 0xe4, 0x12, 0x5e, 0xf4, 0xfe, 0x0a, 0x21, 0x9a, 0xd4, 0xa4, 0x62, 0x9f, 0x61,
	0x7d, 0x0b, 0xdd, 0xf5, 0x9a, 0x53, 0x22, 0xe8, 0x86, 0x29, 0xcd, 0xa5, 0x70, 0x7c, 0x67, 0xc3,
	0x8b, 0x6c, 0x36, 0x9b, 0xe4, 0x44, 0xd0, 0xb9, 0xb7, 0x8b, 0xff, 0x95, 0x59, 0x60, 0x41, 0x6a,
	0x76, 0x00, 0xb6, 0xb1, 0x05, 0xd6, 0x0f, 0x64, 0xf8, 0xf5, 0x37, 0x07, 0x60, 0xaf, 0x6c, 0xad,
	0xe6, 0x3b, 0x96, 0x26, 0x97, 0x41, 0xbf, 0x51, 0xb8, 0x18, 0x7f, 0x05, 0x09, 0x5f, 0x48, 0x55,
	0x13, 0xe3, 0x66, 0xb0, 0x37, 0x32, 0x72, 0xb2, 0x38, 0xf8, 0xb8, 0x0f, 0x89, 0xe6, 0x15, 0x17,
	0x0b, 0xe9, 0xa6, 0x39, 0x1b, 0x76, 0xb3, 0x93, 0x0d, 0x17, 0x87, 0xb4, 0x3d, 0x80, 0xea, 0x09,
	0x4d, 0x63, 0x0f, 0x63, 0x63, 0xbf, 0xb2, 0x1d, 0xfb, 0xb0, 0x35, 0x4c, 0xa7, 0x2d, 0x77, 0xf2,
	0xb3, 0xd1, 0xfb, 0x27, 0x80, 0x28, 0x57, 0x7c, 0xc3, 0xf0, 0x2b, 0x88, 0xb8, 0xbd, 0x9a, 0xfd,
	0xe0, 0x71, 0xe6, 0x2e, 0xaa, 0xf0, 0x26, 0x4e, 0x21, 0xa9, 0xc9, 0x9f, 0x8e, 0x3e, 0x74, 0x3d,
	0x0e, 0xd2, 0x6e, 0x41, 0x31, 0x42, 0xa5, 0x58, 0x6e, 0x1d, 0x5e, 0xab, 0x38, 0x6a, 0xb7, 0x21,
	0xc5, 0x34, 0x53, 0x1b, 0xe6, 0x98, 0x5a, 0xc5, 0x51, 0xe3, 0xb7, 0x90, 0x50, 0xb5, 0x31, 0xf6,
	0x29, 0xb4, 0xdc, 0xe0, 0x90, 0x39, 0x10, 0xf7, 0x1a, 0x0e, 0x29, 0xfc, 0x06, 0x62, 0x43, 0x54,
	0xc5, 0x4c, 0xda, 0xde, 0xdf, 0xce, 0xd4, 0xc9, 0x62, 0x6f, 0xe3, 0x1e, 0x9c, 0xef, 0x49, 0xee,
	0xdd, 0x84, 0xe0, 0xe8, 0x4e, 0xbc, 0xc1, 0x08, 0x62, 0xff, 0xca, 0x70, 0x07, 0xda, 0xb9, 0x9e,
	0x89, 0x47, 0x21, 0x9f, 0x04, 0x7a, 0x81, 0xc1, 0x26, 0xc6, 0xc6, 0xac, 0x50, 0x80, 0xcf, 0x20,
	0xf1, 0xb1, 0x46, 0x21, 0x6e, 0x41, 0x33, 0xd7, 0x77, 0xef, 0x50, 0xc3, 0x97, 0xdc, 0x8d, 0xa6,
	0x3f, 0xa3, 0xe6, 0xe0, 0x77, 0x88, 0xfd, 0x6e, 0x70, 0x17, 0x60, 0x54, 0x9b, 0xe7, 0x46, 0x09,
	0x34, 0x8a, 0xf7, 0x9f, 0x50, 0x60, 0x7f, 0xf8, 0xcb, 0xf5, 0xed, 0x27, 0x14, 0xe2, 0x36, 0x44,
	0x36, 0x1a, 0xa2, 0x86, 0xcd, 0xce, 0xc7, 0x39, 0x6a, 0xda, 0xec, 0xfc, 0x63, 0x7e, 0x83, 0x22,
	0x6b, 0xdd, 0xce, 0xdf, 0xa3, 0xd8, 0x59, 0xe3, 0xfc, 0x57, 0x94, 0x0c, 0x7e, 0x84, 0xd8, 0xcf,
	0x66, 0xbb, 0x4f, 0xab, 0xff, 0x74, 0xb7, 0x34, 0x5c, 0x3f, 0xa2, 0xc0, 0xd2, 0xdc, 0x30, 0x25,
	0xd8, 0x12, 0x85, 0x36, 0x9e, 0x08, 0x6e, 0x14, 0x45, 0x0d, 0x0b, 0x5f, 0x90, 0xda, 0x15, 0x35,
	0x07, 0xdf, 0x41, 0xfb, 0x78, 0x93, 0x18, 0xc1, 0xf9, 0x4c, 0x94, 0x4b, 0xa2, 0x35, 0x5f, 0x70,
	0x46, 0xd1, 0x0b, 0x0b, 0x76, 0x9d, 0x17, 0xb7, 0x1f, 0x51, 0x60, 0x29, 0xc6, 0x79, 0x8e, 0x42,
	0x1b, 0xfc, 0xf4, 0xc3, 0x14, 0x35, 0x3e, 0x7c, 0x0f, 0x6f, 0x4a, 0x59, 0x67, 0x3b, 0x46, 0x19,
	0x25, 0x59, 0xb9, 0x94, 0x6b, 0x9a, 0xad, 0xed, 0xc2, 0x78, 0xb9, 0xff, 0xe2, 0xfc, 0xf6, 0xba,
	0xe2, 0xe6, 0x61, 0x7d, 0x9f, 0x95, 0xb2, 0xbe, 0xf2, 0x75, 0x57, 0x64, 0xc5, 0xaf, 0x76, 0xa5,
	0xfb, 0xc3, 0xdf, 0xc7, 0xae, 0xea, 0xdd, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x33, 0xf9, 0x91,
	0xae, 0xa8, 0x04, 0x00, 0x00,
}
