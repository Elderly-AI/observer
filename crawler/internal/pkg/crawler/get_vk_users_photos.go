package crawler

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-vk-api/vk"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/Elderly-AI/observer/crawler/internal/pkg/model"
)

const MaxVkUsersCountRequest = 75
const MaxVkUserRequestBatchSize = 25
const MaxVkUsersRequestBatchCount = 3

// TODO обернуть в либку
const code = ` 
var id_cnt = 0;
var user_id_arr = [%s];
var users_photos = [];

while (id_cnt < user_id_arr.length) {
    var photos = API.photos.get({"owner_id": user_id_arr[id_cnt], "album_id": "profile"});
    var user_photos = {"user_id": user_id_arr[id_cnt], "photos":[]};
    var photo_cnt = 0;

    while (photo_cnt < photos.count) {
    	var index = (photos.items[photo_cnt].sizes@.type).indexOf("z");
		user_photos.photos.push(photos.items[photo_cnt].sizes[index].url);
    	photo_cnt = photo_cnt + 1;
    }

    users_photos.push(user_photos);
    id_cnt = id_cnt + 1;
}

return users_photos;
`

func (f *Facade) GetVkUsersPhotosHandler(ctx context.Context, users []uint64) (usersPhotos []model.UserPhotos, err error) {
	vkUsersPhotos, err := f.getVkUsersProfilePhotosBatch(users)
	if err != nil {
		return
	}
	usersPhotos, err = f.downloadUsersPhoto(vkUsersPhotos)
	return
}

func (f *Facade) downloadUsersPhoto(vkUsers []model.VkUser) (usersPhotos []model.UserPhotos, err error) {
	usersPhotos = make([]model.UserPhotos, 0, len(vkUsers))
	for _, vkUser := range vkUsers {
		var userPhotos model.UserPhotos
		userPhotos, err = f.downloadUserPhoto(vkUser)
		if err != nil {
			return
		}
		usersPhotos = append(usersPhotos, userPhotos)
	}
	return
}

func (f *Facade) downloadUserPhoto(vkUser model.VkUser) (userPhotos model.UserPhotos, err error) {
	userPhotos.UserID = vkUser.UserID
	userPhotos.Photos = make([][]byte, 0, len(vkUser.Photos))
	for _, photoUrl := range vkUser.Photos {
		var b []byte
		b, err = f.downloadPhoto(photoUrl)
		if err != nil {
			return
		}
		userPhotos.Photos = append(userPhotos.Photos, b)
	}
	return
}

func (f *Facade) downloadPhoto(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return nil, errors.New("error received non 200 response code")
	}
	b, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (f *Facade) getVkUsersProfilePhotosBatch(users []uint64) ([]model.VkUser, error) { // TODO async
	chunks, err := getVkUsersRequestChunks(users)
	if err != nil {
		return nil, err
	}
	usersPhotos := make([]model.VkUser, 0)
	for _, chunk := range chunks {
		var photos []model.VkUser
		photos, err = f.getVkUsersProfilePhotos(chunk)
		if err != nil {
			return nil, err
		}
		usersPhotos = append(usersPhotos, photos...)
	}
	return usersPhotos, nil
}

func (f *Facade) getVkUsersProfilePhotos(users []uint64) (usersProfilePhotos []model.VkUser, err error) {
	if len(users) > MaxVkUserRequestBatchSize {
		return nil, errors.New("error to many users in batch")
	}
	execCode := fmt.Sprintf(code, Uint64ArrToString(users))
	err = f.VkClient.CallMethod("execute", vk.RequestParams{
		"code": execCode,
	}, &usersProfilePhotos)
	return usersProfilePhotos, err
}

func getVkUsersRequestChunks(users []uint64) (chunks [][]uint64, err error) {
	if MaxVkUsersCountRequest < len(users) {
		err = errors.New("error too many users in request") // TODO var
		return
	}
	chunks = make([][]uint64, 0)
	for offset := 0; offset*MaxVkUserRequestBatchSize < len(users); offset += MaxVkUserRequestBatchSize {
		chunks = append(chunks, users[offset:min(offset+MaxVkUserRequestBatchSize, len(users))])
	}
	return
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
