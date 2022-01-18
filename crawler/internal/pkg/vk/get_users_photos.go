package vk

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"

	vkLib "github.com/go-vk-api/vk"
)

// TODO Причесать ассинхронность
func (v *Vk) GetUsersPhotos(users []uint64) ([]UserPhotos, error) {
	chunks, err := getVkUsersRequestChunks(users)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	wg.Add(len(chunks))

	buffer := make(chan []User, len(chunks))
	errs := make(chan error, len(chunks))

	for _, chunk := range chunks {
		chunkC := chunk
		go func() {
			v.getVkUsersProfilePhotos(chunkC, buffer, errs)
			wg.Done()
		}()
	}

	wg.Wait()
	usersPhotos := make([]User, 0, len(chunks))
	for i := 0; i < len(chunks); i++ {
		err = <-errs
		usersPhotos = append(usersPhotos, <-buffer...)
		if err != nil {
			return nil, err
		}
	}

	return filterUsersPhotos(usersPhotos), nil
}

func filterUsersPhotos(users []User) []UserPhotos {
	usersPhotos := make([]UserPhotos, len(users))
	for i, user := range users {
		usersPhotos[i].Photos = make([]string, len(user.Photos))
		usersPhotos[i].UserID = user.UserID
		for j, photo := range user.Photos {
			var max = ""
			for _, size := range photo.Sizes {
				if size.Type > max {
					usersPhotos[i].Photos[j] = size.Url
					max = size.Type
				}
			}
		}
	}
	return usersPhotos
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

func (v *Vk) getVkUsersProfilePhotos(users []uint64, usersPhotos chan []User, errs chan error) {
	var usersProfilePhotos []User
	var err error

	if len(users) > MaxVkUserRequestBatchSize {
		usersPhotos <- nil
		errs <- errors.New("error to many users in batch")
	}
	execCode := fmt.Sprintf(code, Uint64ArrToString(users))
	err = v.Client.CallMethod("execute", vkLib.RequestParams{
		"code": execCode,
	}, &usersProfilePhotos)

	usersPhotos <- usersProfilePhotos
	errs <- err
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
