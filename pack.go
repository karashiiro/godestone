// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package main generated by go-bindata.// sources:
// lodestone-css-selectors/meta.json
// lodestone-css-selectors/profile/achievements.json
// lodestone-css-selectors/profile/attributes.json
// lodestone-css-selectors/profile/character.json
// lodestone-css-selectors/profile/classjob.json
// lodestone-css-selectors/profile/gearset.json
// lodestone-css-selectors/profile/minion.json
// lodestone-css-selectors/profile/mount.json
package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _metaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xe6\xe5\x52\x50\x50\x50\x50\x2a\x4b\x2d\x2a\xce\xcc\xcf\x53\xb2\x52\x50\x32\xd0\x33\xd4\x33\x50\xe2\xe5\xaa\x05\x04\x00\x00\xff\xff\x56\xac\x6d\xc3\x1c\x00\x00\x00")

func metaJsonBytes() ([]byte, error) {
	return bindataRead(
		_metaJson,
		"meta.json",
	)
}

func metaJson() (*asset, error) {
	bytes, err := metaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "meta.json", size: 28, mode: os.FileMode(438), modTime: time.Unix(1608435847, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _profileAchievementsJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xe6\xe5\x52\x50\x50\x50\x50\xf2\xf1\x0c\x0e\x51\xb2\x52\x50\xd2\xcb\x49\x29\x2e\x89\x8f\x4f\x4c\xce\xc8\x4c\x2d\x4b\xcd\x4d\xcd\x2b\x51\xb0\x53\x28\xcd\xb1\xca\x2b\xc9\xd0\x4d\xce\xc8\xcc\x49\xd1\x30\xd6\x54\xd2\x81\x6a\xf2\x74\x01\x6b\x49\xcd\x2b\x29\xaa\x44\xd1\x03\x57\x01\x32\x36\xde\xcf\x35\x22\x24\xde\x29\x34\x24\xc4\xdf\x0f\xa4\xbe\x34\x47\x2f\xa9\x24\x2f\x3e\xbe\x20\x31\x3d\xb5\x08\xc9\x60\x23\x4d\x05\x3b\x85\x9c\x4c\x24\x11\x13\x90\x48\x22\x92\x80\xa1\xa6\x12\x2f\x57\x2d\x20\x00\x00\xff\xff\x6d\x82\xdf\x3c\xb4\x00\x00\x00")

func profileAchievementsJsonBytes() ([]byte, error) {
	return bindataRead(
		_profileAchievementsJson,
		"profile/achievements.json",
	)
}

func profileAchievementsJson() (*asset, error) {
	bytes, err := profileAchievementsJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "profile/achievements.json", size: 180, mode: os.FileMode(438), modTime: time.Unix(1608435847, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _profileAttributesJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\xd5\x4d\x4b\xf3\x40\x10\x07\xf0\x7b\xa1\xdf\x61\xc9\xe9\x79\x40\xc5\xd6\x2a\xd2\x43\x61\xd9\x8e\xc9\xd0\xed\x76\xd9\x0c\xbe\x9c\x96\xb4\x29\x34\x10\x6b\x49\xa3\x20\xe2\x77\x17\xb5\x88\xa6\x24\x4d\xca\xe6\x3a\x90\xfc\xf2\x9f\xcc\xec\xbe\x75\x3b\x8c\x31\xe6\x85\x64\x40\xf9\x14\x78\x43\xe6\xe5\xd1\x3c\x5d\x9e\x2d\x56\x51\x16\x2d\xf2\x65\x66\xed\x26\xca\xa2\x47\x6b\xd3\x64\x9b\x0f\xd7\xf9\xea\x74\xb1\x4a\xd2\xf8\x5f\xff\x3f\x1b\xb1\x7c\xfe\x14\xbf\xfe\x2a\xf6\xbe\x8a\xd9\x5e\x25\xfe\xf3\xa0\x77\xb2\x63\xc7\x70\x4f\x60\x90\x1e\xdc\xbb\xfd\x2a\xf7\x16\x89\xcb\x56\xd8\x8b\x2a\x16\x15\x81\x94\xe8\x83\x12\xe0\x9e\x1e\x54\xd1\x53\x54\x63\xf7\xe4\x65\x09\xb9\x43\x85\x41\x42\xc1\xa5\x0d\x90\xac\xe1\xd4\x24\xf4\xc0\xc9\x78\x11\x98\x29\x2a\x4e\x38\x53\xee\xed\xca\x11\x1b\xa3\x01\x41\x2d\x26\x2f\x9b\xb4\x9f\xec\x37\xa0\xc2\x26\xee\x95\x83\x8e\x4f\xb9\x8f\xc2\xb6\x66\x97\x75\x7c\xa7\x73\x22\x2e\x26\x56\xcf\xee\xc0\x34\xc0\xaf\x1d\x04\x0f\x27\x28\xa5\x0d\x35\x40\x93\x35\xab\x29\xd7\x8b\xfd\xdd\x7b\x3d\x23\x50\xa2\xc9\xc9\xd6\x3b\x77\x90\x3f\x00\x2e\x51\xf9\xad\x7f\x44\xe5\xce\x85\x1a\x8e\xf8\x09\x75\xe9\x03\x0b\x47\xa0\xb8\x68\x76\xa7\xf4\x5c\xdc\xa1\x1a\xa1\x15\xf4\xc0\xd0\x05\xfa\xd3\xdc\xe3\xd8\x88\x3d\xa7\xc5\xf7\xa6\x49\xb1\x12\x27\x2f\xc5\xd2\x76\x13\xad\x4b\x0e\x15\x6d\x7d\x6d\xc5\xb1\x60\xbf\x36\xd8\xed\xbc\x7f\x04\x00\x00\xff\xff\xc8\xae\x94\xe4\x11\x09\x00\x00")

func profileAttributesJsonBytes() ([]byte, error) {
	return bindataRead(
		_profileAttributesJson,
		"profile/attributes.json",
	)
}

func profileAttributesJson() (*asset, error) {
	bytes, err := profileAttributesJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "profile/attributes.json", size: 2321, mode: os.FileMode(438), modTime: time.Unix(1608435847, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _profileCharacterJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x94\x5b\x6f\x9b\x30\x14\xc7\xdf\x2b\xf5\x3b\x58\x3c\xad\xd2\x3a\xa9\x97\xa7\x3e\x44\x72\xc1\xab\x98\x00\x23\xc7\x62\xca\xd3\x91\x63\x9c\xe2\x8d\x8b\x05\x6e\xb6\x69\xca\x77\x9f\x80\x86\x06\x94\xa6\xec\xa2\xbd\x45\xe1\x9c\xdf\xf9\x73\xce\x4f\xfc\x3c\x3f\x43\x08\x21\x07\xbb\xdc\x4f\x08\xb8\x01\x5e\x2e\x3f\xd1\x7b\xe7\x0e\x39\x1f\x64\x26\x6a\x21\xad\xaa\x01\x64\x2e\x9a\x06\xb4\xac\x4a\xb4\x40\xba\x78\xbc\x2b\x6d\x76\x29\x33\x9d\xa7\xef\xae\x2e\x9c\xf7\xc7\x21\x10\x90\x84\x04\xc7\x51\x90\x0a\x2b\xd0\x02\x99\x57\x48\x09\xe6\x98\x75\xad\x9b\x5a\x14\x0a\xa0\x23\x00\x6c\x84\x54\xa7\x22\xdc\xfb\x74\x3a\xb0\x51\xf9\x46\x97\xb6\xae\xd2\x27\x69\x75\x55\x0e\xb5\x1f\x19\x21\xe0\xd2\x30\xc6\xd1\xca\xb9\x43\xcf\xab\xe8\x1e\x45\x38\x24\x53\xce\xa6\x56\x4a\x56\x85\x11\xe5\x0f\x80\x52\x14\x6d\x8e\xec\xf6\x20\xc6\xf5\x05\x5a\x20\x71\x34\x57\x07\xf5\x5d\x1a\x41\x80\x57\x84\x2d\xc7\xe3\xfa\xe4\x94\x73\x1a\xb6\x43\x53\xbd\x7d\x75\xae\xac\x55\x63\x4f\x2c\x60\xc0\x85\xbe\xe7\x05\xe4\x9f\xe1\x38\x8d\xff\x98\xf5\x82\xda\xf5\x3f\x77\xfb\x0b\x3c\x30\x1c\x79\x07\x27\x18\xf3\x2f\xd7\x79\x25\xbf\x1e\xb0\x6e\xdb\x05\xa7\x7a\x3b\xdd\xb9\x19\xfd\x31\xdc\x77\x7f\xc4\x16\x3a\xb6\x68\x5d\x7d\x9f\x32\x26\x15\xed\x7d\x47\x20\x0f\xaf\xc6\x42\xf4\xe9\x00\xd6\xba\xb6\xd9\x50\x1a\x53\xc6\x19\xf6\x79\x57\xfb\xa5\x01\xd0\x85\x78\x54\x60\x2a\xf3\x64\x4e\x79\x1b\x27\x31\x70\x82\xc3\x19\x1e\x9a\xad\xb1\x4a\x14\xff\xd5\xc1\x61\x66\x77\xe4\xe7\xb7\xfa\x2b\x0d\xe7\x11\xaf\xe7\x9a\x38\x0f\x77\x73\x4a\x46\x86\xdd\xee\xdb\x15\xc1\x03\x89\x3c\xc2\xde\xf6\xf1\xea\xf7\x7c\x5c\x12\x96\xf4\x58\x33\xb1\xed\x5b\x55\xe7\xe9\x38\xe7\xbe\x89\xfb\xbc\xdf\xdf\xa4\xc5\x6a\x9b\xbf\x18\xca\xe9\xe7\xe8\xed\xbc\x37\xf3\xf2\x9e\x9f\xed\x7e\x05\x00\x00\xff\xff\x76\xbd\x36\x59\x17\x06\x00\x00")

func profileCharacterJsonBytes() ([]byte, error) {
	return bindataRead(
		_profileCharacterJson,
		"profile/character.json",
	)
}

func profileCharacterJson() (*asset, error) {
	bytes, err := profileCharacterJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "profile/character.json", size: 1559, mode: os.FileMode(438), modTime: time.Unix(1608435847, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _profileClassjobJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x9a\x5f\x6f\xda\x30\x14\xc5\xdf\x2b\xf5\x3b\x20\x9e\x56\x69\x9b\x34\xcc\xfe\xf5\x61\x92\x09\x1e\xf1\xe2\xd8\x95\x49\xda\x69\x2f\xa8\xa3\x9d\xca\x84\x5a\xa9\xa3\xd5\xa4\x69\xdf\x7d\x1a\x8d\x89\x1b\xcc\x75\xb0\xb9\x79\x45\xe4\x9e\x1f\xe7\x24\xbe\xd7\x0e\x7f\x8e\x8f\x7a\xbd\x5e\xaf\x3f\x52\xdf\xbe\xd0\x99\x60\xe7\x4c\xf4\x4f\x7b\xfd\xab\xc5\xe3\xeb\xf9\xcd\xe5\xfd\xe5\x7c\x75\x7d\x3f\x9b\xfd\xbc\xfb\x3e\x9b\x2d\x17\xbf\x56\xa7\xb7\xab\x9b\x57\xf3\x9b\xc5\xf2\xea\xc5\xfb\x93\xde\xa7\xde\xd5\xe2\xd1\xfa\x68\x70\xd2\x7f\xf9\xac\x5e\xce\x8a\x42\xb0\xf0\x82\xc3\x66\x41\x49\xf3\x88\x72\xa4\x2e\xc7\x4a\xcd\x32\x3a\x63\x5f\xcf\x5a\x96\xfb\x08\xd2\x55\xe5\xf6\xf1\xcf\x51\x70\xb0\x55\x70\x8f\xdf\xeb\xa8\xf7\xf4\x7b\xab\x8a\x67\x54\xd0\x31\x97\x35\xa3\x5d\x70\x7e\x77\xbb\xba\xbe\x5d\x39\x88\xb6\x3e\x7a\xf3\xff\xa3\x87\x65\xf3\x4b\xcb\x45\xf3\x3b\x3b\x7f\x9b\x21\x29\xa5\x50\x49\x36\x2d\x68\xc1\xba\xe7\x21\xdb\x3c\xd5\xcd\xd0\x2d\x87\x75\x13\x5d\x50\xad\xb9\xd2\x9d\x24\xe4\xb8\x6c\xb0\x4d\xd2\x5d\x42\x8e\xcb\xc8\x36\x0f\x7e\x42\x8e\xcb\xac\x84\xc6\x54\x67\x99\xe4\x93\xb4\xe8\x24\x24\x02\x86\x64\xc1\x74\x97\x93\x03\x89\x38\x91\xf0\xa3\x72\xa0\x58\x51\x4d\x4a\x39\xd2\x8c\x66\xac\x9b\xe7\x69\x08\x46\x65\xc1\x74\x17\x95\x03\x89\x38\x91\xf0\xa3\x72\xa0\x0c\xed\xde\x94\x2b\x99\x45\xc7\x34\x88\x6e\x4c\x6b\x8c\xc3\x04\x14\x08\x43\x1a\x30\x71\xd1\x04\x42\xd8\x0b\x9e\xa6\x13\xa5\xe2\x87\x86\x36\x24\x70\x4b\x32\x24\xdd\xc5\x03\xb7\x24\xc3\x83\x9f\x10\xdc\x92\x24\x97\xf6\xe0\x8e\xc8\x01\x77\xa3\x27\x8e\xee\xd2\x81\x1b\xd1\x13\x0d\x7e\x36\x70\x0f\x9a\xd2\xbc\xd4\x94\x77\x92\x0e\xdc\x80\x0c\x49\x77\xf9\xc0\xdd\xc7\xf0\xe0\x27\xe4\x6b\x3d\x17\x29\x2f\x58\x4e\x27\xac\xb1\x79\x5b\x5e\x5f\xde\xff\x58\xfc\x86\xb3\x8e\xdf\x0b\xd5\xf2\x8d\x68\xd0\x21\x88\x0b\xc2\xde\x0f\x63\x8a\xdb\x0f\x49\x92\x2a\x41\x35\x96\xfd\x70\x57\x31\xe2\xa8\xe6\xc3\x8d\xc4\x20\xa0\x58\x0f\xf7\x0e\x3a\x2d\xb4\x12\x6a\xc2\xa9\xc4\xb2\x1f\x6e\x1a\x36\x00\x6a\x04\x70\xb7\xb0\x31\x50\x62\x00\xdb\x84\x39\xe0\xa2\x7a\x1c\x92\x42\xfc\xd8\xbb\x56\x0e\xb6\x3f\x7e\xd2\x5d\xeb\xef\xed\x7b\xfc\x70\x9b\xd3\x24\xe5\x92\x4f\x0b\x2c\xdb\xe1\xb5\xa7\x96\x47\xf5\x1e\x5e\x7d\x6a\x08\x94\x00\x7c\xc7\x29\x32\x61\x41\x2b\x7f\xfc\xbc\x5a\x69\xa3\x5a\xef\x3b\x2c\x59\x13\xa0\xf8\xbe\x73\xc1\x31\x8f\x9c\xa0\x49\xf6\x7c\xea\x79\x58\x36\x4e\x98\x9b\x43\xd4\x3e\x2b\xca\xa6\x7c\xc3\xdf\x68\x11\xe2\x12\xa9\x2c\x8c\x2e\x6e\x4f\x25\x65\x9e\x2b\x69\xdf\x9c\xfb\x55\xf7\x8c\x1d\xa6\x7a\x94\x3d\x9e\xb9\xc2\x68\x04\xb9\x03\x3f\xb8\x9a\x8d\x63\x6e\x1e\xf8\xc9\x34\xc5\xa3\xbc\x69\x27\x11\x64\x0d\x5c\x7a\x24\x4a\x16\xe3\x0d\xbc\x8f\xdb\x54\x8f\x32\xa7\xa5\x46\x90\x3b\x3b\x6b\x57\xd5\x13\xaa\xcf\x98\x2c\x5a\x2c\xfb\x6f\x31\xf6\x5b\xb5\x7c\xbb\x95\xff\x70\x10\xc4\x05\x01\x2f\xfe\x87\x13\x1f\x36\x96\xcd\x69\xce\x8b\x14\x2b\x01\x78\xed\xb3\xf4\x51\x23\x80\x57\x47\x8b\x02\x25\x03\xcf\xc6\x4b\xe7\x4a\xe3\x3d\x02\x9e\x4d\x57\x25\x8e\xea\xbe\x67\xc3\x55\x21\xa0\x58\xef\x79\x2f\xa4\xc4\x18\xf5\xee\xf7\xbc\x09\xda\xc8\xa3\xda\xef\x79\xf7\xb3\x81\x40\x09\x60\xe7\x91\xdb\x5a\x5c\x30\x5a\xa4\x4c\x5f\x28\x9d\xe1\x3d\x01\x8e\xcb\x06\xbb\x10\x50\x83\x70\x5c\x46\x76\x81\xa0\x84\xe1\xb8\xcc\xfe\xcb\x01\xa3\xe7\x78\x29\xbc\x03\x53\xa8\xb4\x51\xed\x77\x10\x90\x2d\x02\x14\xdf\x1d\xca\x76\x03\x10\x49\xca\xf2\x36\x47\x0f\x81\xf2\xf0\xbf\xb3\x6a\x79\x54\xf7\xe1\xbf\x60\xd5\x10\x28\x01\xc0\x7f\x27\x4b\x4a\xc1\x25\xd5\x6d\x4e\x3e\x03\xf5\x3f\xc0\x73\x68\xad\x8f\x1a\x81\x83\x82\x38\x29\x50\x32\x70\xa8\x3f\x7f\xef\xcf\x65\xd8\xea\x73\x80\x77\xfd\xdc\xb1\xff\x46\x07\x20\x4d\x80\xbd\x6d\x8f\x3f\xf6\x1c\xa9\x82\xb6\x3a\xf5\x0c\x54\xf7\x8c\xff\x46\x1d\xd5\x7a\xcf\xf0\x6f\x18\x50\xdc\x87\x47\xff\xcf\x7c\x9a\xe2\xdd\xf3\xf0\xe4\x5f\x69\xa3\x3a\x0f\x0f\xfe\x15\x01\x8a\xef\xee\xb9\xff\xf8\xe8\xef\xbf\x00\x00\x00\xff\xff\x6f\x01\x9c\x15\xe4\x2c\x00\x00")

func profileClassjobJsonBytes() ([]byte, error) {
	return bindataRead(
		_profileClassjobJson,
		"profile/classjob.json",
	)
}

func profileClassjobJson() (*asset, error) {
	bytes, err := profileClassjobJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "profile/classjob.json", size: 11492, mode: os.FileMode(438), modTime: time.Unix(1608435847, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _profileGearsetJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\xdb\xcf\x6b\x1a\x51\x14\xc5\xf1\x7d\xc0\xff\x61\x70\xd5\x42\x0c\xf5\x57\x6d\x5d\x04\x5e\xf5\xd5\x0c\x9d\x8c\x65\x66\x4a\xc8\x2a\xb4\xba\x50\x90\x98\x85\xed\xa6\xe4\x7f\x2f\xd3\xba\x88\xa0\xbc\xf4\xdc\xfb\xce\x78\x57\x81\x4b\x16\xdf\x18\xf2\xe1\x2d\x72\x7e\xb7\x2e\x92\x24\x49\xda\x77\xde\x7d\x9d\xe7\xed\x71\xb2\x3f\xfc\x3d\xe6\xee\xd6\xb7\xc7\x49\xfb\x6a\xbd\xd8\x3e\x76\x16\x9d\xce\xbb\xe4\x3a\xb9\x5a\xfe\xe8\xec\xb6\xdb\xcd\x6e\xfd\x94\x5c\x27\xcb\xf5\xaf\xf1\xe3\x6e\xd5\x59\xac\xd6\x9b\xe5\x9b\xee\x5b\xbd\x53\xaf\x3e\xad\x7a\x07\x97\xf6\xe5\x8b\xba\xdb\xb4\x70\x33\xff\xd0\x7c\xe4\xe1\xa9\x5f\x9f\x9e\x4e\x67\x97\x95\x4b\xf3\xc6\x82\x8f\x9e\x0e\x3f\x56\x57\xf9\x22\x75\x0f\xdd\x08\x8d\xc3\xfa\xf4\x73\xf3\xe2\x32\xaa\x2f\x9b\x75\xe8\x53\x3e\x5a\xd8\x6b\xa8\xf0\xc8\x2f\xfd\x44\x61\xbf\xa1\xc2\xfe\xab\x0b\x07\x0d\x15\x0e\x5e\x5d\x38\x6c\xa8\x70\x18\x2c\x9c\x14\xde\x55\xf3\x22\x96\x40\xa3\x13\xdf\xf5\xfd\xf0\x6f\xf7\x5f\xd0\xf3\x65\xeb\x62\x4f\x79\x79\x93\xfa\x6c\x1a\xa2\xbc\x6b\x81\x72\x7a\xa4\x94\x72\x66\x30\x4a\xb9\x4e\x63\x4c\xca\x79\x85\x28\xe5\xbc\x42\x94\x72\x5e\x21\x4a\x39\xaf\x50\x40\xb9\x4e\x24\x4a\xf9\x8d\x77\x41\xc8\x7b\x16\x20\xa7\x47\x4a\x21\x67\x06\xa3\x90\xeb\x34\xc6\x84\x9c\x57\x88\x42\xce\x2b\x44\x21\xe7\x15\xa2\x90\xf3\x0a\x05\x90\xeb\x44\xa2\x90\x7f\x9a\x4f\xef\x43\x90\xf7\x2d\x40\x4e\x8f\x94\x42\xce\x0c\x46\x21\xd7\x69\x8c\x09\x39\xaf\x10\x85\x9c\x57\x88\x42\xce\x2b\x44\x21\xe7\x15\x0a\x20\xd7\x89\x84\x5f\xe4\x2e\x9f\x96\x21\xc9\x07\x16\x24\xa7\x47\x4a\x25\x67\x06\xa3\x92\xeb\x34\xc6\x94\x9c\x57\x88\x4a\xce\x2b\x44\x25\xe7\x15\xa2\x92\xf3\x0a\x05\x92\xeb\x44\xa2\x92\xdf\xb9\xb4\xac\x42\x92\x0f\x2d\x48\x4e\x8f\x94\x4a\xce\x0c\x46\x25\xd7\x69\x8c\x29\x39\xaf\x10\x95\x9c\x57\x88\x4a\xce\x2b\x44\x25\xe7\x15\x0a\x24\xd7\x89\x44\x25\xcf\xfc\x2c\xf8\x24\x7f\x6f\x01\x72\x7a\xa4\x14\x72\x66\x30\x0a\xb9\x4e\x63\x4c\xc8\x79\x85\x28\xe4\xbc\x42\x14\x72\x5e\x21\x0a\x39\xaf\x50\x00\xb9\x4e\x24\x0a\xf9\x67\xef\x83\x2f\xf2\x91\x05\xc8\xe9\x91\x52\xc8\x99\xc1\x28\xe4\x3a\x8d\x31\x21\xe7\x15\xa2\x90\xf3\x0a\x51\xc8\x79\x85\x28\xe4\xbc\x42\x01\xe4\x3a\x91\xff\x0d\x79\xfd\x75\x8f\xb9\x77\x45\x91\xe6\xe1\x97\xf9\x07\x0b\xa0\xd3\x23\xa5\xa0\x33\x83\x51\xd0\x75\x1a\x63\x82\xce\x2b\x44\x41\xe7\x15\xa2\xa0\xf3\x0a\x51\xd0\x79\x85\x02\xd0\x75\x22\xd1\x97\x79\xee\x27\x5f\x32\x37\xf1\x21\xcc\x3f\x5a\xc0\x9c\x1e\x29\xc5\x9c\x19\x8c\x62\xae\xd3\x18\x13\x73\x5e\x21\x8a\x39\xaf\x10\xc5\x9c\x57\x88\x62\xce\x2b\x14\x60\xae\x13\x09\xff\x33\x62\xe1\x26\x3e\xf3\x55\xf0\x69\xde\x35\xb1\xf7\xe4\x57\x4a\x3d\xa7\x16\xa3\xa0\x2b\x45\x46\xdd\x09\x9d\xff\xe6\x93\x98\x08\x2f\x85\xce\x7f\xf5\x49\x4c\x94\x6c\x85\x9a\xdd\x7d\x16\x69\x3e\xeb\x06\x49\xb7\xb1\xfb\xb4\x37\xfc\x34\xb1\xfc\x34\x30\xfd\x34\xb0\xfd\x34\x30\xfe\x34\xb0\xfe\xb4\x31\xff\x6c\x76\xff\x59\x93\xde\x0b\x92\x6e\x62\x01\xca\xaf\x14\x93\x6e\x61\x03\xaa\x14\x19\x95\xf4\xf3\x5f\x81\x12\x13\x61\xd2\xcf\x7f\x07\x4a\x4c\x94\x90\xde\xec\x12\xb4\x9c\x7f\xcb\x26\xc5\x7d\x59\xb9\x2c\x08\x7b\xff\xb8\x62\x64\xdd\xf7\x3f\x42\xeb\xe2\xf9\x4f\x00\x00\x00\xff\xff\x19\x1d\xcd\xef\x2d\x4c\x00\x00")

func profileGearsetJsonBytes() ([]byte, error) {
	return bindataRead(
		_profileGearsetJson,
		"profile/gearset.json",
	)
}

func profileGearsetJson() (*asset, error) {
	bytes, err := profileGearsetJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "profile/gearset.json", size: 19501, mode: os.FileMode(438), modTime: time.Unix(1608435847, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _profileMinionJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xe6\xe5\x52\x50\x50\x50\x50\xf2\xf1\x0c\x0e\x51\xb2\x52\x50\xd2\xcb\xcd\xcc\xcb\xcc\xcf\x8b\x8f\xcf\xc9\x2c\x2e\x51\xd2\x81\xca\xfa\x39\xfa\xba\xa2\xc8\xe6\x25\xe6\xa6\x2a\xf1\x72\xd5\x02\x02\x00\x00\xff\xff\xeb\xdc\x01\x4f\x3f\x00\x00\x00")

func profileMinionJsonBytes() ([]byte, error) {
	return bindataRead(
		_profileMinionJson,
		"profile/minion.json",
	)
}

func profileMinionJson() (*asset, error) {
	bytes, err := profileMinionJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "profile/minion.json", size: 63, mode: os.FileMode(438), modTime: time.Unix(1608435847, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _profileMountJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xe6\xe5\x52\x50\x50\x50\x50\xf2\xf1\x0c\x0e\x51\xb2\x52\x50\xd2\xcb\xcd\x2f\xcd\x2b\x89\x8f\xcf\xc9\x2c\x2e\x51\xd2\x81\x4a\xfa\x39\xfa\xba\x22\x4b\xe6\x25\xe6\xa6\x2a\xf1\x72\xd5\x02\x02\x00\x00\xff\xff\xf1\xf9\x11\x1d\x3d\x00\x00\x00")

func profileMountJsonBytes() ([]byte, error) {
	return bindataRead(
		_profileMountJson,
		"profile/mount.json",
	)
}

func profileMountJson() (*asset, error) {
	bytes, err := profileMountJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "profile/mount.json", size: 61, mode: os.FileMode(438), modTime: time.Unix(1608435847, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"meta.json":                 metaJson,
	"profile/achievements.json": profileAchievementsJson,
	"profile/attributes.json":   profileAttributesJson,
	"profile/character.json":    profileCharacterJson,
	"profile/classjob.json":     profileClassjobJson,
	"profile/gearset.json":      profileGearsetJson,
	"profile/minion.json":       profileMinionJson,
	"profile/mount.json":        profileMountJson,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("nonexistent") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"meta.json": &bintree{metaJson, map[string]*bintree{}},
	"profile": &bintree{nil, map[string]*bintree{
		"achievements.json": &bintree{profileAchievementsJson, map[string]*bintree{}},
		"attributes.json":   &bintree{profileAttributesJson, map[string]*bintree{}},
		"character.json":    &bintree{profileCharacterJson, map[string]*bintree{}},
		"classjob.json":     &bintree{profileClassjobJson, map[string]*bintree{}},
		"gearset.json":      &bintree{profileGearsetJson, map[string]*bintree{}},
		"minion.json":       &bintree{profileMinionJson, map[string]*bintree{}},
		"mount.json":        &bintree{profileMountJson, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
