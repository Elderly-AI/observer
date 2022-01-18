package vk

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	vkLib "github.com/go-vk-api/vk"
)

func (v *Vk) GetUsersPhotos(users []uint64) ([]User, error) {
	chunks, err := getVkUsersRequestChunks(users)
	if err != nil {
		return nil, err
	}
	usersPhotos := make([]User, 0)
	for _, chunk := range chunks {
		var photos []User
		photos, err = v.getVkUsersProfilePhotos(chunk)
		if err != nil {
			return nil, err
		}
		usersPhotos = append(usersPhotos, photos...)
	}
	return usersPhotos, nil
}

func getVkUsersRequestChunks(users []uint64) (chunks [][]uint64, err error) {
	if MaxVkUsersCountRequest < len(users) {
		err = errors.New("error too many users in request") // TODO var
		return
	}
	chunks = make([][]uint64, 0)
	for offset := 0; offset < len(users); offset += MaxVkUserRequestBatchSize {
		chunks = append(chunks, users[offset:min(offset+MaxVkUserRequestBatchSize, len(users))])
	}
	return
}

func (v *Vk) getVkUsersProfilePhotos(users []uint64) (usersProfilePhotos []User, err error) {
	if len(users) > MaxVkUserRequestBatchSize {
		return nil, errors.New("error to many users in batch")
	}
	execCode := fmt.Sprintf(code, Uint64ArrToString(users))
	err = v.Client.CallMethod("execute", vkLib.RequestParams{
		"code": execCode,
	}, &usersProfilePhotos)
	return usersProfilePhotos, err
}

func Uint64ArrToString(arr []uint64) string {
	b := make([]string, len(arr))
	for i, v := range arr {
		b[i] = strconv.FormatUint(v, 10)
	}
	return strings.Join(b, ",")
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
